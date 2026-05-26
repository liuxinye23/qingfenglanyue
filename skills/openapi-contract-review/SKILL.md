---
name: openapi-contract-review
description: 审计 OpenAPI/Swagger 契约的认证边界、字段约束、错误模型与批量变更风险的技能包。
metadata:
  version: "1.0.0"
  tags:
    - openapi
    - swagger
    - api
    - schema
    - contract
  triggers:
    - openapi
    - swagger
    - api schema
    - contract review
    - request body
    - response schema
  recommended_tools:
    - api-schema-analyzer
    - http-framework-test
    - arjun
  role_hints:
    - API安全测试
    - Web应用扫描
  stages:
    - recon
    - verify
    - report
  autoload_priority: 16
---

# OpenAPI 契约审计

## 何时使用

当目标提供 OpenAPI/Swagger 文档、导出的接口契约、前后端共享 schema，或大量 JSON API 需要结构化审计时，优先调用本技能。

## 快速流程

1. 先确认契约来源：线上公开文档、代码仓库、导出文件、网关配置。
2. 枚举所有路径、方法、认证方式、批量接口、导出接口、上传接口和管理员接口。
3. 对照 schema 审核字段约束、只读/只写边界、默认值、枚举、分页、排序、过滤与错误模型。
4. 将“契约允许的输入”与“服务端真实行为”做抽样核对。

## 重点检查

## 认证与授权边界

- 不同路径是否声明了不同安全需求，还是全部继承了过宽的默认配置。
- 管理员接口、批量操作、导出接口是否在契约里被弱化或漏标。

## 字段与对象模型

- 是否缺少关键字段约束，例如长度、格式、枚举、只读/只写、最小最大值。
- 是否把内部字段、角色字段、租户字段、审批字段直接暴露为可写。

## 批量与危险操作

- 批量更新、批量删除、导入导出、异步任务、回调配置等高影响接口是否有单独限制。
- 契约是否模糊地复用了通用请求体，导致危险字段“顺便可写”。

## 错误与兼容性

- 契约中的错误响应是否足够区分认证失败、权限不足、参数错误、资源不存在。
- 版本兼容策略、弃用接口和隐藏接口是否有可见约束。

## 建议工具

## api-schema-analyzer

- 用于快速审阅 schema 结构、字段、鉴权声明和批量写接口。

## http-framework-test

- 对高风险接口做只读抽样核对，确认线上行为与契约是否一致。

## arjun

- 在契约不完整或文档过旧时，补充挖掘隐藏参数与未文档化字段。

## 证据要求

- 记录路径、方法、认证方式、可疑字段和契约片段摘要。
- 区分“文档缺陷”和“真实安全缺陷”；若两者并存，要分别说明。
- 对高风险字段写明其可能导致的数据或权限影响。

## 修复建议方向

- 把安全边界直接表达在契约中，而不是只在实现里“默认处理”。
- 为高影响接口单独设计请求体与权限模型。
- 建立契约与线上行为的回归校验。

