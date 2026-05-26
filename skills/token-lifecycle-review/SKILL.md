---
name: token-lifecycle-review
description: 审计 JWT、刷新令牌、API Token 与 OAuth/OIDC 凭据生命周期的技能包。
metadata:
  version: "1.0.0"
  tags:
    - jwt
    - token
    - oauth
    - oidc
    - refresh-token
  triggers:
    - jwt
    - bearer token
    - refresh token
    - oauth
    - oidc
    - api token
    - token rotation
  recommended_tools:
    - jwt-analyzer
    - http-framework-test
    - nuclei
  role_hints:
    - API安全测试
    - 渗透测试
  stages:
    - recon
    - verify
    - report
  autoload_priority: 17
---

# Token 生命周期审计

## 何时使用

当系统使用 JWT、Opaque Token、刷新令牌、长期 API Token、OAuth/OIDC 登录或多端共享凭据时，优先调用本技能。

## 快速流程

1. 明确所有 Token 类型：访问令牌、刷新令牌、一次性票据、设备令牌、Webhook 密钥。
2. 记录每类 Token 的签发入口、有效期、续期方式、撤销条件和绑定对象。
3. 对登录、刷新、登出、密码变更、权限变更、租户切换后做一致性复测。
4. 检查声明、受众、作用域、设备绑定与轮换策略。

## 重点检查

## 签发与绑定

- Token 是否绑定正确的用户、租户、设备、客户端类型。
- 是否把高权限信息、内部环境信息、调试字段直接写入声明。
- 多客户端共用同一 Token 体系时，受众和来源是否可区分。

## 生命周期

- 访问令牌是否过长；刷新令牌是否缺乏轮换与撤销。
- 刷新后旧令牌是否保留过长并发窗口。
- 密码修改、角色变更、账号冻结、MFA 变更后，旧 Token 是否立即失效。

## 作用域与最小权限

- Scope / role / permissions 是否过宽。
- 不同业务域是否复用同一超大权限令牌。
- 机器到机器 Token 与用户态 Token 是否清晰分离。

## 传输与存储

- 浏览器端是否把长期 Token 暴露在不必要的位置，例如可被脚本访问的存储。
- 日志、错误响应、调试面板、前端配置中是否泄露真实 Token 值。

## 建议工具

## jwt-analyzer

- 解码头部和载荷，核对 `iss`、`aud`、`sub`、`exp`、`iat`、`nbf`、scope/role 字段。
- 比较普通用户、高权限用户、跨租户用户的声明差异。

## http-framework-test

- 复测登录、刷新、登出、权限变化后的关键接口。
- 对“旧 Token + 新状态”的组合进行只读验证，观察是否还被接受。

## nuclei

- 作为补充检查默认 OIDC 路径、调试端点、公开元数据、常见暴露接口。

## 证据要求

- 保存各类 Token 的声明摘要，不要在报告里泄露完整密文。
- 对刷新、撤销、登出前后同一接口做并排对比。
- 明确写出“旧 Token 是否仍可访问什么资源”。

## 修复建议方向

- 收紧有效期并实施刷新轮换。
- 在密码、MFA、角色与租户切换后统一吊销旧凭据。
- 按客户端类型与业务域拆分 Token 受众和最小权限。

