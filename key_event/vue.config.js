const {
  defineConfig
} = require('@vue/cli-service')

module.exports = defineConfig({
  transpileDependencies: true,
  pluginOptions: {
    electronBuilder: {
      chainWebpackMainProcess: (config) => {
        config.output.filename((file) => {
          if (file.chunk.name === 'index') {
            return 'background.js';
          } else {
            return '[name].js';
          }
        });
      },
      preload: './src/preload.js',
      customFileProtocol: './',
      builderOptions: {
        "productName": "键盘监听器", //项目名 这也是生成的exe文件的前缀名
        "appId": "top.hzfui.key", //包名
        "copyright": "HFH Copyright © 2022", //版权  信息
        "compression": "maximum", // "store" | "normal"| "maximum" 打包压缩情况(store 相对较快)，store 39749kb, maximum 39186kb
        "directories": {
          "output": "dist_electron" // 输出文件夹
        },
        extraFiles: [
          {
            from: "./src/build/keyEvent.exe",
            to: "./resources/app/keyEvent.exe"
          },
          {
            from: "./src/assets",
            to: "./resources/app/assets"
          },
        ],
        "asar": false, // asar打包
        "win": {
          "icon": "public/favicon.ico", //图标路径
          "target": [{
            "target": "nsis",
            "arch": ["x64"]
          }]
        },
        "nsis": {
          "oneClick": false, // 一键安装
          "perMachine": true, // 是否开启安装时权限限制（此电脑或当前用户）
          "allowElevation": true, // 允许请求提升。 如果为false，则用户必须使用提升的权限重新启动安装程序。
          "allowToChangeInstallationDirectory": true, // 允许修改安装目录
          "installerIcon": "public/favicon.ico", // 安装图标
          "uninstallerIcon": "public/favicon.ico", //卸载图标
          "installerHeaderIcon": "public/favicon.ico", // 安装时头部图标
          "createDesktopShortcut": true, // 创建桌面图标
          "createStartMenuShortcut": true, // 创建开始菜单图标
          "shortcutName": "键盘监听器" // 图标名称，
        }
      }
    }
  }
})