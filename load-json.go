package amba

import (
	"encoding/json"
	"io"
	"os"
)

func LoadJson(targetStruct interface{},targetFile string) error {
	file, err := os.Open(targetFile)
	if err != nil {
		panic(err)
	}
	dataByte, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(dataByte, &targetStruct)

	return nil
}