# 介绍

ERod 是基于 [energy](https://github.com/energye/energy) 和 [rod](https://github.com/go-rod/rod) 封装的 DevTools 协议。 
它专为 Web 自动化和抓取而设计，适用于高级和低级使用，
高级开发人员可以轻松使用低级包和功能 自定义或构建自己的 Rod 版本，高级函数只是构建默认版本 Rod 的示例。

ERod 底层直接使用 CEF API 操作 DevTools 协议方法, 而不是使用浏览器DevTools协议

- energy: 基于 LCL 和 CEF 的跨平台GUI框架
- rod: 基于 DevTools 协议的高级驱动程序，在ERod中移除了浏览器DevTools协议消息，底层消息处理直接使用 CEF API 回调消息函数
