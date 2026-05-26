---
name: stego-triage
description: 面向 CTF 隐写题的首轮分流技能包，覆盖图片、音频、压缩包与容器格式中的隐藏数据线索。
metadata:
  version: "1.0.0"
  tags:
    - ctf
    - stego
    - image
    - audio
    - forensics
  triggers:
    - stego
    - hidden data
    - image challenge
    - audio challenge
    - zsteg
    - metadata puzzle
  recommended_tools:
    - exiftool
    - binwalk
    - foremost
    - zsteg
    - strings
  role_hints:
    - CTF
    - 数字取证
  stages:
    - triage
    - verify
    - solve
  autoload_priority: 17
---

# 隐写题首轮分流

## 何时使用

当题目给出图片、音频、视频、PDF、压缩包，或提示“隐藏信息”“看起来正常但有异常体积/元数据/尾部数据”时使用。

## 快速流程

1. 看文件类型、大小、后缀、实际 magic、元数据。
2. 检查是否存在嵌套文件、尾部附加、异常块、异常颜色或通道。
3. 判断题目更像“直接藏数据”还是“文件本身就是线索”。

## 重点检查

## 元数据

- 作者、时间、设备、注释、坐标、缩略图、软件名。
- 异常元数据本身可能就是 flag 入口或下一步提示。

## 容器结构

- 图片与压缩格式中是否嵌有第二层文件或多余数据。
- 有些题目不是复杂隐写，只是拼接、伪装或简单结构篡改。

## 文本与可见痕迹

- 文件里是否直接残留路径、口令提示、脚本片段、假 flag。
- 对音频/图像题先确认有没有更直接的文字或压缩层线索，再决定是否做更深分析。

## 建议工具

- `exiftool`
- `binwalk`
- `foremost`
- `zsteg`
- `strings`

## 输出要求

- 文件指纹
- 元数据异常
- 嵌套/附加对象
- 最可能的下一步

