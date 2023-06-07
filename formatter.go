package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
)

type FormatResponse struct {
	Success bool
	Data    string
}

func fileFormatterController() func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		log.Println(c.RemoteAddr(), "Connection opened for format request!")
		for {
			mt, msg, err := c.ReadMessage()
			log.Println(c.RemoteAddr(), "Received format request!")

			if err != nil {
				logAndSendFormatErrorToClient(c, mt, "Format controller :: Error in Reading message from websocket:", err)
				return
			}

			input, err := base64Decoder(string(msg))
			if err != nil {
				logAndSendFormatErrorToClient(c, mt, "Format controller :: Error in decoding input base64 code:", err)
				continue
			}

			result, err := runFormatter(string(input))
			if err != nil {
				logAndSendFormatErrorToClient(c, mt, "Format controller :: Error in formatting input code:", err)
				continue
			}

			response := FormatResponse{
				Success: true,
				Data:    base64Encoder(string(result)),
			}
			responseJSON, err := json.Marshal(response)
			if err != nil {
				logAndSendFormatErrorToClient(c, mt, "Format controller :: Error in marshaling response:", err)
				continue
			}

			base64EncodedResponse := base64Encoder(string(responseJSON))
			if err = c.WriteMessage(mt, []byte(base64EncodedResponse)); err != nil {
				logAndSendFormatErrorToClient(c, mt, "Format controller :: Error in writing back to the client:", err)
				continue
			}
			log.Println(c.RemoteAddr(), "Sent format response!")
		}
	}
}

func logAndSendFormatErrorToClient(c *websocket.Conn, mt int, appErrorMessage string, sysError error) {
	response := FormatResponse{
		Success: false,
		Data:    appErrorMessage + sysError.Error(),
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
