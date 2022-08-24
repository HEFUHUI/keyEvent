package controllers

import (
	"encoding/json"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

type KeyCount struct {
	RawCode uint16 `json:"raw_code"`
	Name    string `json:"name"`
	Count   uint64 `json:"count"`
}

func (c *MainController) Get() {
	// c.TplName = "index.html"
	file, err := os.ReadFile("./key.json")
	if err != nil {
		println(err.Error())
		return
	}

	keys := make(map[uint16]*KeyCount)

	c.Data["keys"] = keys
	json.Unmarshal(file, &keys)
	c.JSONResp(keys)
}

// @router /key_count [get]
func (c *MainController) GetKeyCount() {
	// 读取keys下的所有文件
	files, err := os.ReadDir("./keys")
	if err != nil {
		println(err.Error())
		return
	}
	keyCounts := make(map[uint16]*KeyCount)
	for _, file := range files {
		kcs := make(map[uint16]*KeyCount)
		fileName := file.Name()
		if fileName[len(fileName)-5:] == "-key.json" {
			OFile, err := os.ReadFile("./keys/" + fileName)
			if err != nil {
				println(err.Error())
				continue
			}
			json.Unmarshal(OFile, &kcs)
		}
	}
	c.JSONResp(keyCounts)

	// file, err := os.ReadFile("./key.json")
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }

	// keys := make(map[uint16]*KeyCount)

	// c.Data["keys"] = keys
	// json.Unmarshal(file, &keys)
	// c.JSONResp(keys)
}
