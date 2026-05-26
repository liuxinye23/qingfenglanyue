---
name: elf-pwn-triage
description: 面向 CTF ELF Pwn 题的首轮 triage 技能包，帮助快速判断架构、保护、输入面与利用方向。
metadata:
  version: "1.0.0"
  tags:
    - ctf
    - pwn
    - elf
    - binary
    - exploitation
  triggers:
    - elf
    - pwn
    - stack overflow
    - heap
    - format string
    - ret2libc
    - ctf binary
  recommended_tools:
    - strings
    - gdb
    - pwntools
    - ropper
    - one-gadget
  role_hints:
    - CTF
    - 二进制分析
    - 渗透测试
  stages:
    - triage
    - verify
    - solve
  autoload_priority: 19
---

# ELF Pwn 题首轮 triage

## 何时使用

当题目提供 ELF、菜单程序、远程 nc 端口、明显崩溃点或交互式本地/远程程序时使用。

## 快速流程

1. 确认架构、位数、动态/静态链接、是否有符号与 libc 依赖。
2. 识别保护：栈保护、PIE、NX、RELRO、FORTIFY、沙箱或 seccomp 痕迹。
3. 梳理输入面：菜单、长度字段、格式化输出、文件读写、索引、指针操作。
4. 先判断更接近栈、堆、格式串还是逻辑型原语。

## 重点检查

## 可控面

- 哪个输入真正决定长度、索引、地址、格式串或生命周期。
- 程序是否在多个阶段复用同一缓冲区或对象。

## 信息泄露

- 是否存在地址、libc、栈、堆、指针、错误信息等泄露源。
- 很多题先要做的是“拿到稳定泄露”，不是立刻构造最终利用链。

## 远程约束

- 远程与本地的 libc、超时、交互方式、菜单节奏是否一致。
- 如果远程只允许少数轮交互，优先优化最短利用路径。

## 建议工具

- `strings`
- `gdb`
- `pwntools`
- `ropper`
- `one-gadget`

## 输出要求

- 架构与保护概况
- 主要输入原语
- 泄露面与潜在利用方向
- 本地/远程差异

