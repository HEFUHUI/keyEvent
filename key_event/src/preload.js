const { contextBridge, ipcRenderer} = require('electron')

contextBridge.exposeInMainWorld('hzfui', {
    getKeys(){

    },
    openAutoOpen(){
        console.log('openAutoOpen')
        ipcRenderer.send("openAutoStart")
    },
    closeAutoOpen(){
        ipcRenderer.send("closeAutoStart")
    }
})