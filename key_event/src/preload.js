const { contextBridge, ipcRenderer} = require('electron')
const fs = require('fs')
const path = require("path")

contextBridge.exposeInMainWorld('hzfui', {
    openAutoOpen(){
        ipcRenderer.send("openAutoStart")
    },
    closeAutoOpen(){
        ipcRenderer.send("closeAutoStart")
    },
    getAppDataPath(){
        let dir = process.env.LOCALAPPDATA;
        if(!dir){
            dir = path.join("C:", "Users",process.env.USERNAME, "AppData", "Local")
        }
        return dir
    },
    readKeyDetail(dir, key){
        return new Promise((resolve, reject) => {
            const keyDir = path.join(dir, "keyEvent");
            const result = {};
            fs.readdir(keyDir, (err, files) => {
                if(err){
                    return reject(err);
                }
                files.forEach(file => {
                    const keyBuf = fs.readFileSync(path.join(keyDir, file), {encoding: "utf-8"})
                    const keyObj = JSON.parse(keyBuf);
                    Object.keys(keyObj).forEach(c_key => {
                        if(c_key + "" === key + ""){
                           const date = file.split("-key.json")[0];
                            console.log(keyObj[c_key].count, date)
                            result[date] = keyObj[c_key].count;
                        }
                    })
                })
                resolve(result);
            })
        })
    },
    readKeys(dir){
        //make(map[uint16]*KeyCount)
        return new Promise((resolve, reject) => {
            const keyDir = path.join(dir, "keyEvent");
            const keyCounts = {};
            fs.readdir(keyDir, (err, files) => {
                // 判断为文件并且以-keys.json结尾
                if (err) {
                    reject(err)
                    return;
                }
                files.forEach(file => {
                    if (file.endsWith("-key.json")) {
                        const keyBuf = fs.readFileSync(path.join(keyDir, file), {encoding: "utf-8"});
                        let keyData = JSON.parse(keyBuf);
                        for (let key in Object.keys(keyData)) {
                            if(!keyData.hasOwnProperty(Object.keys(keyData)[key])){
                                continue;
                            }
                            if (keyCounts[Object.keys(keyData)[key]]) {
                                keyCounts[Object.keys(keyData)[key]].count += keyData[Object.keys(keyData)[key]].count;
                            } else {
                                keyCounts[Object.keys(keyData)[key]] = keyData[Object.keys(keyData)[key]];
                            }
                        }
                    }

                })
                resolve(keyCounts);
            })
        })
    },
    getShortcutKey(dir){
        return new Promise((resolve, reject) => {
            const keyDir = path.join(dir, "keyEvent_ShortcutKey");
            const keyCounts = {};
            fs.readdir(keyDir, (err, files) => {
                // 判断为文件并且以-shortcutKey.json结尾
                if (err) {
                    reject(err)
                    return;
                }
                files.forEach(file => {
                    if (file.endsWith("-shortcutKey.json")) {
                        const keyBuf = fs.readFileSync
                        (path.join(keyDir, file), {encoding: "utf-8"});
                        let keyData = JSON.parse(keyBuf);
                        const keys = Object.keys(keyData);
                        for (let key in keys) {
                            if(!keyData.hasOwnProperty(keys[key])){
                                continue;
                            }
                            if (keyCounts[keys[key]]) {
                                keyCounts[keys[key]].count += keyData[keys[key]].count;
                            } else {
                                keyCounts[keys[key]] = keyData[keys[key]];
                            }
                        }
                    }
                })
                resolve(keyCounts);
            })
        })
    }
})