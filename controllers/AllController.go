package controllers

import (
	"encoding/json"
	"os"
	"path"

	"github.com/beego/beego/v2/server/web"
)

type AllController struct {
	Dir string
	web.Controller
}

type KeyCount struct {
	RawCode uint16 `json:"raw_code"`
	Name    string `json:"name"`
	Count   uint64 `json:"count"`
}

// func (c *MainController) Get() {
// 	// c.TplName = "index.html"
// 	file, err := os.ReadFile("./key.json")
// 	if err != nil {
// 		println(err.Error())
// 		return
// 	}

// 	keys := make(map[uint16]*KeyCount)

// 	c.Data["keys"] = keys
// 	json.Unmarshal(file, &keys)
// 	c.JSONResp(keys)
// }

func (c *AllController) Get() {
	files, err := os.ReadDir(c.Dir)
	if err != nil {
		println(err.Error())
		return
	}
	keyCounts := make(map[uint16]*KeyCount)
	for _, file := range files {
		kcs := make(map[uint16]*KeyCount)
		fileName := file.Name()
		if fileName[len(fileName)-9:] == "-key.json" {
			OFile, err := os.ReadFile(path.Join(c.Dir, fileName))
			if err != nil {
				continue
			}
			json.Unmarshal(OFile, &kcs)
		}
		// 合并
		for k, v := range kcs {
			if _, ok := keyCounts[k]; ok {
				keyCounts[k].Count += v.Count
			} else {
				keyCounts[k] = v
			}
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
