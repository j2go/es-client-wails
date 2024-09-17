# Elasticsearch查询客户端

原仓库 https://gitee.com/qiaoshengda/es-client 打包的桌面板在我 M2 Mac 上用不了，仔细看了一下是直接请求被跨域拦截了

本仓库基于 es-client 的前端，用 wails 代理 http 调用解决跨域问题

## 开发

```shell
wails dev
```

## 发布包下载

### macos m1/m2
https://gitee.com/tiangao/es-client-wails/releases/download/v3.1.7-alpha/es-client-wails_darwin_arm64.zip

### macos intel
https://gitee.com/tiangao/es-client-wails/releases/download/v3.1.7-alpha/es-client-wails_darwin_amd64.zip
