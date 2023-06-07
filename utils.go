package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

func ParseInputMessageAsProjectDirectory(message []byte) (InputProjectDir, error) {
	var projDir InputProjectDir
	err := json.Unmarshal([]byte(message), &projDir)
	if err != nil {
		return projDir, err
	}

	decodedCargoToml, err := base64Decoder(projDir.CargoToml)
	if err != nil {
		return projDir, err
	}
	projDir.CargoToml = string(decodedCargoToml)

	decodedMainRs, err := base64Decoder(projDir.MainRs)
	if err != nil {
		return projDir, err
	}
	projDir.MainRs = string(decodedMainRs)

	return projDir, nil
}

func GetMd5StringOfInput(input string) string {
	hashAsByteArray := md5.Sum([]byte(input))
	hashAsString := hex.EncodeToString(hashAsByteArray[:])

	return hashAsString
}

func base64Decoder(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}

func base64Encoder(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}
