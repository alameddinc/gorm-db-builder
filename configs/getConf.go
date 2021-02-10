package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func GetConfigs(s string) map[string]interface{} {
	arg := strings.Split(s, ".")
	jsonFile, err := os.Open("configs/" + arg[0] + ".json")
	if err != nil {
		panic(err.Error() + " configs can not be read!")
	}
	byteVal, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err.Error() + " configs can not be read!")
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(byteVal), &result)
	return result
}
