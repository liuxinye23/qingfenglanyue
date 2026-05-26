---
name: wasm-reverse-triage
description: 面向 CTF 中 WASM、前端校验与浏览器端逆向题目的首轮分流技能包。
metadata:
  version: "1.0.0"
  tags:
    - ctf
    - wasm
    - reverse
    - frontend
    - browser
  triggers:
    - wasm
    - webassembly
    - frontend reverse
    - browser challenge
    - client-side check
    - js reverse
  recommended_tools:
    - strings
    - http-framework-test
    - ghidra
  role_hints:
    - CTF
    - 二进制分析
    - Web应用扫描
  stages:
    - triage
    - verify
    - solve
  autoload_priority: 17
---

# WASM 与前端逆向题首轮分流

## 何时使用

当题目包含 `.wasm`、前端大包、浏览器端校验、客户端 flag 逻辑、JS 混淆或前端解谜流程时使用。

## 快速流程

1. 先确认校验在浏览器端还是后端。
2. 找出关键资源：JS bundle、WASM、Source Map、配置 JSON、动态加载脚本。
3. 识别题目是“还原算法”“提取常量/密钥”还是“理解交互顺序”。

## 重点检查

## 入口与调用关系

- WASM 从哪里加载、由谁调用、输入输出如何传递。
- JS 与 WASM 是否共同参与校验，还是其中一个只是包装层。

## 常量与提示

- 先找字符串、错误文案、隐藏路径、调试标志、字典表。
- 对纯前端题，不要忽略浏览器存储、请求参数、路由片段和资源命名。

## 状态机与交互

- 某些题目关键不在算法，而在触发顺序、事件链、路由状态、时间检查。
- 如果服务端也参与了一部分校验，要及时回到 Web 主线。

## 建议工具

- `strings`
- `http-framework-test`
- `ghidra`

## 输出要求

- 关键资源清单
- 校验主链
- 关键常量/函数/模块
- 下一步深挖方向

