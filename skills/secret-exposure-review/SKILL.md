---
name: secret-exposure-review
description: 审计前端、配置、备份路径与公开资源中的密钥、令牌与敏感配置暴露问题。
metadata:
  version: "1.0.0"
  tags:
    - secrets
    - api-key
    - config
    - frontend
    - exposure
  triggers:
    - secret
    - api key
    - token leak
    - exposed config
    - js bundle
    - backup file
  recommended_tools:
    - http-framework-test
    - nuclei
    - gau
    - waybackurls
  role_hints:
    - Web应用扫描
    - API安全测试
    - secure-code-review
  stages:
    - recon
    - verify
    - report
  autoload_priority: 16
---

# 密钥与敏感配置暴露审计

## 何时使用

当目标包含前端构建产物、公开配置文件、移动端接口、备份资源、历史路径或第三方集成配置时，使用本技能。

## 快速流程

1. 枚举静态资源、构建产物、配置文件、历史 URL 和备份命名模式。
2. 审阅是否存在真实 API Key、访问令牌、Webhook 密钥、数据库连接串、调试开关。
3. 区分“公开标识符”和“可直接滥用的秘密”，避免误报。
4. 记录暴露位置、可见范围、是否仍有效、可到达的后端资源。

## 重点检查

## 前端与静态资源

- `main.*.js`、配置 JSON、Source Map、调试页面、运行时环境变量回显。
- 第三方 SDK 初始化参数中，哪些只是公开站点标识，哪些是真正私密凭据。

## 备份与历史路径

- 常见备份命名、旧版本入口、临时导出文件、测试目录、历史快照。
- 历史 URL 中暴露的旧配置、旧接口或已下线资源是否仍可访问。

## 错误响应与日志

- 调试堆栈、代理错误页、集成失败提示、诊断面板是否回显敏感字段。
- 是否在响应头、下载文件名、注释或示例代码中泄露内部标识。

## 建议工具

## http-framework-test

- 拉取首页、关键 JS、配置路径、错误路径并保存响应头与正文摘要。
- 对静态资源使用只读方式检查，不做修改型验证。

## nuclei

- 用于补充常见敏感文件、备份扩展名、默认配置暴露模板。

## gau / waybackurls

- 从历史路径中找回旧资源、旧配置、旧后台入口，再回源验证是否仍在线。

## 证据要求

- 记录暴露位置、响应状态、片段摘要和失效时间判断依据。
- 若为真实秘密，只展示脱敏前后缀，不直接写完整值。
- 写清楚该秘密的潜在影响范围，而不是只说“存在 key”。

## 修复建议方向

- 把真实秘密移出前端与公开资源。
- 对备份、历史产物、Source Map 和调试页做访问控制或下线清理。
- 对第三方集成区分公开标识符与私密凭据，避免混用。

