package controllers

import (
	"encoding/json"
	"os"
	"path"

	"github.com/beego/beego/v2/server/web"
)

type StatisticsController struct {
	Dir string
	web.Controller
}

func (c *StatisticsController) Get() {
	key := c.Ctx.Request.URL.Query().Get("key")
	files, err := os.ReadDir(c.Dir)
	if err != nil {
		println(err.Error())
		return
	}
	keyCounts := make(map[string]uint64)
	for _, file := range files {
		kcs := make(map[string]*KeyCount)
		fileName := file.Name()
		if fileName[len(fileName)-9:] == "-key.json" {
			OFile, err := os.ReadFile(path.Join(c.Dir, fileName))
			if err != nil {
				continue
			}
			json.Unmarshal(OFile, &kcs)
		}
		if _, ok := kcs[key]; ok {
			keyCounts[fileName[:len(fileName)-9]] += kcs[key].Count
		}
	}
	c.JSONResp(keyCounts)
}
