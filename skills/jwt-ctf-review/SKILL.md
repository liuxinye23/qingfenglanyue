---
name: jwt-ctf-review
description: 面向 CTF 中 JWT、会话票据与签名逻辑题目的分流、声明审阅与验证工作流技能包。
metadata:
  version: "1.0.0"
  tags:
    - ctf
    - jwt
    - token
    - web
    - auth
  triggers:
    - jwt
    - jws
    - jwe
    - bearer token
    - auth cookie
    - session token
    - ctf web auth
  recommended_tools:
    - jwt-analyzer
    - http-framework-test
    - ffuf
    - nuclei
  role_hints:
    - CTF
    - API安全测试
    - Web应用扫描
  stages:
    - triage
    - verify
    - solve
  autoload_priority: 19
---

# JWT 与票据类 CTF 题

## 何时使用

当题目包含 JWT、会话票据、登录态 Cookie、SSO/OAuth 风格流程、签名字段或“只差认证就能过”的 Web 题时使用本技能。

## 首轮 triage

1. 明确令牌类型：JWT、Opaque Token、伪 JWT、加密票据、会话 Cookie。
2. 解码头部与载荷，记录算法、声明、时效、受众、角色字段。
3. 区分“逻辑缺陷题”和“实现细节题”。
4. 对登录前后、不同角色、不同入口的请求做并排对比。

## 重点检查

## 声明与业务逻辑

- 是否存在明显的角色、租户、用户标识、调试字段。
- 题目更像“改声明触发逻辑”还是“绕过签名校验”。
- 令牌是否与 URL 参数、Header、Cookie 三者之一耦合。

## 流程与入口

- 登录、刷新、登出、个人资料、管理员接口是否共用同一授权来源。
- 是否存在只在某一路由读取某个声明的题目设计。

## 表现形式

- 有些 CTF 不会给你真实 JWT，而是给一个“看起来像 JWT 的三段文本”，先确认格式和编码细节。
- 如果题目强调头部、kid、jku、x5u、alg 等字段，优先检查头部控制面。

## 建议工具

## jwt-analyzer

- 先做解码和结构核对，确认是不是标准 JWT。
- 对比匿名、普通用户、管理员三类令牌差异。

## http-framework-test

- 固定重放关键接口，观察不同令牌或 Cookie 下的响应差异。

## ffuf / nuclei

- 在认证边界不清晰时补充枚举路径和隐藏接口，但不要偏离题目主线。

## 输出要求

- 令牌结构与关键信息
- 题目真正依赖的授权点
- 复现步骤
- 取 flag 的最小路径

