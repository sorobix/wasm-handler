package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
)

type RedisObject struct {
	Success bool
	Message string
	Wasm    string
}

type CompilerResponse struct {
	Id      string
	Success bool
	Message string
}

type InputProjectDir struct {
	CargoToml string
	MainRs    string
}

func compilerController() func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		log.Println(c.RemoteAddr(), "Connection opened for compile request!")
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				logAndSendFormatErrorToClient(c, mt, "Compile controller :: Error in Reading message from websocket:", err)
				return
			}
			log.Println(c.RemoteAddr(), "Received compile request!")

			projDir, err := ParseInputMessageAsProjectDirectory(msg)
			if err != nil {
				logAndSendCompileErrorToClient(c, mt, "no-id", "Compile controller :: Error in parsing input as proj dir:", err)
				continue
			}

			redisKey := FormRedisKey(projDir.CargoToml, projDir.MainRs)

			valueFromRedis := GetFromRedis(redisKey)
			foundInCache := len(valueFromRedis) != 0

			if err := c.WriteMessage(mt, []byte("Compiling")); err != nil {
				logAndSendCompileErrorToClient(c, mt, redisKey, "Compile controller :: Error in writing back to the client:", err)
				break
			}

			compileResponse := CompilerResponse{
				Id: redisKey,
			}

			if foundInCache {
				var redisValue RedisObject
				if err := json.Unmarshal([]byte(valueFromRedis), &redisValue); err != nil {
					logAndSendCompileErrorToClient(c, mt, redisKey, "Compile controller :: Error in unmarshaling redis object from cache:", err)
					continue
				}

				decodedMessage, err := base64Decoder(redisValue.Message)
				if err != nil {
					logAndSendCompileErrorToClient(c, mt, redisKey, "Compile controller :: Error in decoding message from cache:", err)
					continue
				}

				compileResponse.Success = redisValue.Success
				compileResponse.Message = string(decodedMessage)

			} else {
				compilationOutput := BuildWasmFromMainRS(projDir.CargoToml, projDir.MainRs)

				decodedMessage, err := base64Decoder(compilationOutput.Message)
				if err != nil {
					logAndSendCompileErrorToClient(c, mt, redisKey, "Compile controller :: Error in decoding input message from cache:", err)
					continue
				}

				compileResponse.Success = compilationOutput.Success
				compileResponse.Message = string(decodedMessage)

				redisValue, err := json.Marshal(compilationOutput)
				if err != nil {
					logAndSendCompileErrorToClient(c, mt, redisKey, "Compile controller :: Error in marshaling redis value:", err)
					continue
				}
				SetInRedis(redisKey, string(redisValue))
			}

			responseInJson, err := json.Marshal(compileResponse)
			if err != nil {
				logAndSendCompileErrorToClient(c, mt, redisKey, "Compile controller :: Error in marshaling output:", err)
				continue
			}

			base64EncodedResponse := base64Encoder(string(responseInJson))
			if err := c.WriteMessage(mt, []byte(base64EncodedResponse)); err != nil {
				logAndSendCompileErrorToClient(c, mt, redisKey, "Compile controller :: Error in writing back to the clients:", err)
				continue
			}
			log.Println(c.RemoteAddr(), "Sent compile response!")
		}
	}
}

func logAndSendCompileErrorToClient(c *websocket.Conn, mt int, id string, appErrorMessage string, sysError error) {
	response := CompilerResponse{
		Id:      id,
		Success: false,
		Message: appErrorMessage + sysError.Error(),
	}
	log.Println(c.RemoteAddr(), appErrorMessage, sysError.Error())
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Println(c.RemoteAddr(), "For ", appErrorMessage, " due to: ", sysError.Error(), "couldn't marshal response JSON due to: ", err.Error())
	}
	if err := c.WriteMessage(mt, responseJSON); err != nil {
		log.Println(c.RemoteAddr(), "For ", appErrorMessage, " due to: ", sysError.Error(), "couldn't write back to client: ", err.Error())
	}
}
