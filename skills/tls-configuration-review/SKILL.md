---
name: tls-configuration-review
description: 审计 HTTPS、证书链、重定向、HSTS 与 TLS 配置一致性的技能包。
metadata:
  version: "1.0.0"
  tags:
    - tls
    - https
    - certificate
    - hsts
    - transport
  triggers:
    - tls
    - https
    - certificate
    - redirect
    - hsts
    - mixed content
  recommended_tools:
    - http-framework-test
    - nuclei
  role_hints:
    - Web应用扫描
    - API安全测试
    - 云安全审计
  stages:
    - recon
    - verify
    - report
  autoload_priority: 14
---

# TLS 与 HTTPS 配置审计

## 何时使用

当系统暴露 HTTPS、反向代理、多域名、多子域、下载站点、管理后台或移动端 API 时使用本技能。

## 快速流程

1. 梳理所有域名、子域、证书边界和跳转链。
2. 检查 HTTP→HTTPS 跳转、证书链、主机名匹配、HSTS、缓存代理行为。
3. 对首页、登录页、API、下载链接和静态资源分别核对是否一致。
4. 记录仅在某个子域、某个入口或某个 CDN 节点出现的偏差。

## 重点检查

## 入口与重定向

- 是否所有敏感入口都会稳定落到 HTTPS。
- 多跳重定向、跨域跳转、端口切换后是否丢失安全头或会话状态。

## 证书与域名

- 证书链是否完整，SAN 是否覆盖真实访问域。
- 过期、错配、临时证书、测试域证书误上生产要单独标注。

## HSTS 与浏览器策略

- 是否只在首页设置 HSTS，而登录页、后台或子域缺失。
- 是否在混合部署下错误地把未准备好的子域一起纳入严格策略。

## 内容与资源一致性

- 页面、下载、图片、脚本、API 是否仍引用明文资源。
- 代理/CDN/边缘节点是否对不同路径应用不同 TLS 与缓存策略。

## 建议工具

## http-framework-test

- 重点抓跳转链、首包时间、TLS 握手指标、关键响应头和最终资源落点。
- 对同一资源分别用 `http://` 与 `https://` 验证实际行为。

## nuclei

- 用于补充常见 TLS/证书/HSTS 暴露和基础配置错误检测。

## 证据要求

- 保存问题域名、目标路径、最终落点、证书摘要和关键响应头。
- 若问题仅出现在某类资源或某个节点，写清分布范围。

## 修复建议方向

- 统一敏感入口的 HTTPS 与跳转策略。
- 修复证书覆盖与链路不一致问题。
- 对 HSTS、混合内容和 CDN 边界做分层治理。

