package main

import (
	"encoding/json"
	"fmt"
	"ketEvent/controllers"
	"os"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	hook "github.com/robotn/gohook"
)

func main() {
	go low()
	web.SetStaticPath("/key_event/dist", "static")
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	web.Router("/", &controllers.MainController{})
	web.Run()
}

func low() {
	// 判断keys 文件夹是否存在，如果不存在则创建
	_, err := os.Stat("keys")
	if os.IsNotExist(err) {
		os.Mkdir("keys", os.ModePerm)
	}
	today := time.Now().Format("2006-01-02")
	keysFile := "./keys/" + today + "-key.json"
	file, err := os.ReadFile(keysFile)
	var keyCounts map[uint16]*controllers.KeyCount
	if err == nil {
		_ = json.Unmarshal(file, &keyCounts)
		if keyCounts == nil {
			keyCounts = make(map[uint16]*controllers.KeyCount)
		}
	} else {
		keyCounts = make(map[uint16]*controllers.KeyCount)
	}

	evChan := hook.Start()
	defer hook.End()
	for ev := range evChan {
		if ev.Kind == 5 {
			str := string(ev.Rawcode)
			if ev.Rawcode == 112 {
				str = "F1"
			} else if ev.Rawcode == 113 {
				str = "F2"
			} else if ev.Rawcode == 114 {
				str = "F3"
			} else if ev.Rawcode == 115 {
				str = "F4"
			} else if ev.Rawcode == 116 {
				str = "F5"
			} else if ev.Rawcode == 117 {
				str = "F6"
			} else if ev.Rawcode == 118 {
				str = "F7"
			} else if ev.Rawcode == 119 {
				str = "F8"
			} else if ev.Rawcode == 120 {
				str = "F9"
			} else if ev.Rawcode == 121 {
				str = "F10"
			} else if ev.Rawcode == 122 {
				str = "F11"
			} else if ev.Rawcode == 123 {
				str = "F12"
			} else if ev.Rawcode == 36 {
				str = "Home"
			} else if ev.Rawcode == 35 {
				str = "End"
			} else if ev.Rawcode == 33 {
				str = "PageUp"
			} else if ev.Rawcode == 34 {
				str = "PageDown"
			} else if ev.Rawcode == 45 {
				str = "Insert"
			} else if ev.Rawcode == 46 {
				str = "Delete"
			} else if ev.Rawcode == 27 {
				str = "Escape"
			} else if ev.Rawcode == 9 {
				str = "Tab"
			} else if ev.Rawcode == 8 {
				str = "Backspace"
			} else if ev.Rawcode == 13 {
				str = "Enter"
			} else if ev.Rawcode == 32 {
				str = "Space"
			} else if ev.Rawcode == 37 {
				str = "Left"
			} else if ev.Rawcode == 162 {
				str = "Ctrl"
			} else if ev.Rawcode == 164 {
				str = "Alt"
			} else if ev.Rawcode == 160 {
				str = "Shift"
			} else if ev.Rawcode == 20 {
				str = "CapsLock"
			} else if ev.Rawcode == 44 {
				str = "PrintScreen"
			} else if ev.Rawcode == 145 {
				str = "ScrollLock"
			} else if ev.Rawcode == 19 {
				str = "Pause"
			} else if ev.Rawcode == 91 {
				str = "Windowns"
			} else if ev.Rawcode == 39 {
				str = "Right"
			} else if ev.Rawcode == 38 {
				str = "Up"
			} else if ev.Rawcode == 40 {
				str = "Down"
			} else if ev.Rawcode == 144 {
				str = "NumLock"
			} else if ev.Rawcode == 188 {
				str = "<"
			} else if ev.Rawcode == 190 {
				str = ">"
			} else if ev.Rawcode == 191 {
				str = "?"
			} else if ev.Rawcode == 192 {
				str = "`"
			} else if ev.Rawcode == 219 {
				str = "{"
			} else if ev.Rawcode == 220 {
				str = "|"
			} else if ev.Rawcode == 221 {
				str = "}"
			} else if ev.Rawcode == 222 {
				str = "\""
			} else if ev.Rawcode == 186 {
				str = ";"
			} else if ev.Rawcode == 187 {
				str = "+"
			} else if ev.Rawcode == 189 {
				str = "-"
			}

			fmt.Println(str, ev.Rawcode)

			if keyCounts[ev.Rawcode] != nil {
				keyCounts[ev.Rawcode].Count++
			} else {
				keyCounts[ev.Rawcode] = &controllers.KeyCount{
					RawCode: ev.Rawcode,
					Name:    str,
					Count:   1,
				}
			}
			marshal, err := json.Marshal(keyCounts)
			if err != nil {
				fmt.Println(err)
			}
			err = os.WriteFile(keysFile, marshal, 0644)
			if err != nil {
				println(err)
			}

		}
	}
}
