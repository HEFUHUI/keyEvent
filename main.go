package main

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	hook "github.com/robotn/gohook"
	"io"
	"keyEvent/controllers"
	"log"
	"os"
	"os/user"
	"path"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	today := time.Now().Format("2006-01-02")

	logDir := path.Join(LocalDataDir, "keyEvent_log")
	err := os.MkdirAll(logDir, os.ModePerm)
	// 设置日志输出到文件
	LocalDataDir = path.Join(logDir, today+"-ke.log")
	logFile, err := os.OpenFile(LocalDataDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	writer := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(writer)
}

var LocalDataDir string

func main() {
	getLocalDataDir()
	go low()
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	web.Router("/", &controllers.AllController{
		Dir: LocalDataDir,
	})
	web.Router("/statistics", &controllers.StatisticsController{
		Dir: LocalDataDir,
	})
	web.Run(":8888")
}

func getLocalDataDir() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("获取当前用户失败", err.Error())
	}
	LocalDataDir = path.Join(currentUser.HomeDir, "AppData", "Local")
	//LocalDataDir = "./keys"
	err = os.MkdirAll(LocalDataDir, os.ModePerm)
	if err != nil {
		log.Fatal("创建本地数据目录失败", err.Error())
	}
}

func low() {
	var err error
	today := time.Now().Format("2006-01-02")
	keysDir := path.Join(LocalDataDir, "keyEvent")
	shortcutKeysDir := path.Join(LocalDataDir, "keyEvent_ShortcutKey")
	err = os.MkdirAll(keysDir, os.ModePerm)
	if err != nil {
		log.Fatal("创建本地数据目录失败，按键统计", err.Error())
	}
	err = os.MkdirAll(shortcutKeysDir, os.ModePerm)
	if err != nil {
		log.Fatal("创建本地数据目录失败，快捷键统计", err.Error())
	}
	keysFile := path.Join(keysDir, today+"-key.json")
	shortcutKeysFile := path.Join(shortcutKeysDir, today+"-shortcutKey.json")
	file, err := os.ReadFile(keysFile)
	shortcutKeyFile, err := os.ReadFile(shortcutKeysFile)
	var keyCounts map[uint16]*controllers.KeyCount
	var shortcutKeys map[string]*controllers.ShortcutKey
	if err == nil {
		_ = json.Unmarshal(file, &keyCounts)
		_ = json.Unmarshal(shortcutKeyFile, &shortcutKeys)
		if keyCounts == nil {
			keyCounts = make(map[uint16]*controllers.KeyCount)
		}
		if shortcutKeys == nil {
			shortcutKeys = make(map[string]*controllers.ShortcutKey)
		}
	} else {
		keyCounts = make(map[uint16]*controllers.KeyCount)
		shortcutKeys = make(map[string]*controllers.ShortcutKey)
	}

	currentKey := make(map[uint16]int64)

	evChan := hook.Start()
	defer hook.End()
	shortcutKey := ""
	for ev := range evChan {
		if ev.Kind == hook.KeyHold {
			// 当按下列表不为空并且当前按键不为ctrl、alt、shift、win时
			if len(currentKey) > 0 && ev.Rawcode != 162 && ev.Rawcode != 164 && ev.Rawcode != 160 && ev.Rawcode != 91 {
				for u := range currentKey {
					shortcutKey += fmt.Sprintf("%d+", u)
				}
				shortcutKey += fmt.Sprintf("%d", ev.Rawcode)
				if shortcutKeys[shortcutKey] == nil {
					shortcutKeys[shortcutKey] = &controllers.ShortcutKey{
						Count: 1,
					}
				} else {
					shortcutKeys[shortcutKey].Count++
				}

				shortcutKey = ""
			}
			// 如果current_key中没有这个键，就累计按下时间 ( 毫秒 )
			if _, ok := currentKey[ev.Rawcode]; !ok {
				currentKey[ev.Rawcode] = time.Now().UnixNano() / 1e6
			}
		}
		if ev.Kind == hook.KeyUp {
			// 获取按键持续时长
			using := (time.Now().UnixNano() / 1e6) - currentKey[ev.Rawcode]
			delete(currentKey, ev.Rawcode)
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

			if keyCounts[ev.Rawcode] != nil {
				keyCounts[ev.Rawcode].Count++
				keyCounts[ev.Rawcode].Using += using
			} else {
				keyCounts[ev.Rawcode] = &controllers.KeyCount{
					RawCode: ev.Rawcode,
					Name:    str,
					Count:   1,
					Using:   using,
				}
			}
			marshal, err := json.Marshal(keyCounts)
			shortcutKeyMarshal, err := json.Marshal(shortcutKeys)
			if err != nil {
				fmt.Println(err)
			}
			err = os.WriteFile(keysFile, marshal, 0644)
			err = os.WriteFile(shortcutKeysFile, shortcutKeyMarshal, 0644)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
