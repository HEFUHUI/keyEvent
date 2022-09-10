'use strict'

import {app, BrowserWindow, dialog, Menu, protocol, Tray} from 'electron'
import {createProtocol} from 'vue-cli-plugin-electron-builder/lib'

const isDevelopment = process.env.NODE_ENV !== 'production'
const path = require('path')
const {spawn} = require('child_process')

protocol.registerSchemesAsPrivileged([
    {scheme: 'app', privileges: {secure: true, standard: true}}
])

let appTray = null;

async function createAboutWindow(){
    const aboutWin = new BrowserWindow({
        width: 400,
        height: 350,
        webPreferences:{
        }
    })
    await aboutWin.loadURL('https://bafybeiffaktrgsjkbere23fkq43jjq5ao2rdpqromcdkdgvaeiwp57kkka.ipfs.4everland.io')
}

async function createWindow() {
    Menu.setApplicationMenu(null)
    const win = new BrowserWindow({
        width: 1200,
        height: 900,
        webPreferences: {
            nodeIntegration: true,
            contextIsolation: !process.env.ELECTRON_NODE_INTEGRATION
        }
    })

    if (process.env.WEBPACK_DEV_SERVER_URL) {
        // Load the url of the dev server if in development mode
        await win.loadURL(process.env.WEBPACK_DEV_SERVER_URL)
        // if (!process.env.IS_TEST) win.webContents.openDevTools()
    } else {
        createProtocol('app')
        await win.loadURL('app://./index.html')
    }
    if (appTray == null) {
        tray(win);
    }

    const execPath = isDevelopment ? path.join(__dirname, "..", "src", "build") :
        path.join(__dirname)
    spawn(path.join(execPath, "keyEvent.exe"), {
        cwd: execPath
    }).on("error", (err) => {
        dialog.showErrorBox("发生错误", err.message)
    })
}

function tray() {
    const trayMenuTemplate = [
        {
            label: "关于", click: () => {
                createAboutWindow().catch(err=>{
                    dialog.showErrorBox("发生错误", err.message)
                })
            }
        },
        {
            label: '退出程序',
            click: function () {
                // 判断是否确定
                dialog.showMessageBox({
                    type: 'info',
                    title: '退出程序',
                    message: '确定要退出程序吗？',
                    buttons: ['确定', '取消']
                }).then(r => {
                    if (r.response === 0) {
                        app.quit();
                    }
                })
            }
        }
    ];

    const iconPath = isDevelopment ? path.join(__dirname, "..", "src", "assets", "2.png") : path.join(__dirname, "assets", "2.png");
    appTray = new Tray(iconPath);

    const contextMenu = Menu.buildFromTemplate(trayMenuTemplate);
    Menu.setApplicationMenu(contextMenu);
    appTray.setToolTip('键盘按键统计.');
    appTray.setContextMenu(contextMenu);
    appTray.on("double-click", async () => {
        if (BrowserWindow.getAllWindows().length === 0) {
            await createWindow()
        }else{
            BrowserWindow.getAllWindows()[0].show()
        }
    })
}

app.on('window-all-closed', () => {
})

app.on("activate", () => {
    if(!app.requestSingleInstanceLock()){
        app.quit()
    }
    if (BrowserWindow.getAllWindows().length === 0){
        createWindow().catch(err => {
            dialog.showErrorBox("发生错误", err.message)
        })
    }else{
        BrowserWindow.getAllWindows()[0].show()
    }
})

app.on('ready', async () => {
    await createWindow()
    if(!app.requestSingleInstanceLock()){
        dialog.showMessageBox({
            type: 'info',
            title: '程序已经在运行',
            message: '程序已经在运行',
            buttons: ['确定']
        }).then(r => {
            if (r.response === 0) {
                app.quit();
            }
        })
    }
})
