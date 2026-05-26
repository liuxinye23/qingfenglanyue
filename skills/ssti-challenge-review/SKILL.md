---
name: ssti-challenge-review
description: 面向 CTF 与靶题中的模板注入场景，提供模板识别、上下文判断与验证路径的技能包。
metadata:
  version: "1.0.0"
  tags:
    - ctf
    - ssti
    - template
    - web
  triggers:
    - ssti
    - template injection
    - jinja2
    - twig
    - freemarker
    - server-side template
    - ctf web challenge
  recommended_tools:
    - http-framework-test
    - ffuf
    - nuclei
  role_hints:
    - CTF
    - Web应用扫描
    - 渗透测试
  stages:
    - triage
    - verify
    - solve
  autoload_priority: 18
---

# SSTI 题分流与验证

## 何时使用

当 Web 题包含模板渲染、邮件预览、主题/样式、自定义页面、错误页、用户资料渲染、服务端表达式或框架特征字符串时使用。

## 快速流程

1. 找到回显点、错误点和上下文类型：HTML 正文、属性、文本片段、模板语句块、邮件模板。
2. 判断是模板注入、普通反射，还是仅前端模板。
3. 确认模板引擎线索：错误栈、语法符号、框架特征、转义行为。
4. 用最小、只读、低副作用的方式验证执行与对象访问边界。

## 重点检查

## 引擎识别

- 观察报错信息、过滤行为、分隔符风格与上下文变化。
- 先判断题目是“识别引擎”还是“到达特定对象/变量”。

## 回显与执行边界

- 是否有回显、是否可控、是否仅在某个模板字段生效。
- 不要一上来假定存在高危执行；先确认对象访问、条件判断或简单表达式是否成立。

## 题目目标

- 有些题目只要求读取某个变量、配置或 flag 文件，不一定要做到通用执行。
- 如果题目更像业务逻辑题，及时回到 Web 主线，不要被“像 SSTI”的表象带偏。

## 建议工具

## http-framework-test

- 对模板参数做只读重放，保存错误信息、回显差异和上下文位置。

## ffuf

- 适合补充枚举模板相关入口、预览端点和隐藏页面。

## 输出要求

- 模板位置
- 可能的引擎线索
- 验证路径
- 最小复现步骤
- 取 flag 或关键证据的思路

