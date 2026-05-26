---
name: ctf-skills
description: 面向 CTF 题目的分类、分流、初始取证与解题工作流技能包，覆盖 Web、Crypto、Pwn、Reverse、Forensics 与 Misc 场景。
metadata:
  version: "1.0.0"
  tags:
    - ctf
    - challenge
    - web
    - crypto
    - pwn
    - reverse
    - forensics
    - misc
  triggers:
    - ctf
    - challenge
    - flag
    - pwn
    - reverse
    - web challenge
    - crypto challenge
    - forensics
    - misc challenge
  recommended_tools:
    - http-framework-test
    - sqlmap
    - ffuf
    - nuclei
    - strings
    - binwalk
    - exiftool
    - foremost
    - volatility3
    - ghidra
    - radare2
    - gdb
    - pwntools
    - hashcat
    - john
    - zsteg
  role_hints:
    - CTF
    - 渗透测试
    - 二进制分析
    - 数字取证
  stages:
    - triage
    - verify
    - solve
    - report
  autoload_priority: 20
---

# CTF 通用解题技能

## 何时使用

当用户明确提到 CTF、flag、题目附件、靶题、题解、远程 challenge、Pwn / Reverse / Crypto / Web / Forensics / Misc 分类，或给出一个不确定类别的题目样本时，优先使用本技能。

本技能的职责不是直接替代专项方法，而是先把题目分流到正确方向，并给出稳定的首轮工作流。

## 快速分流

1. 先识别输入形态：URL、端口、压缩包、二进制、脚本、图片、音频、流量包、内存/磁盘镜像、源码、加密文本。
2. 再识别主要题型：
   - Web：HTTP 入口、表单、JWT、模板、上传、接口、前端逻辑
   - Crypto：密文、签名、哈希、随机数、数论、编码与协议细节
   - Pwn：ELF、崩溃、远程交互、栈堆格式化字符串、沙箱逃逸
   - Reverse：二进制、APK、WASM、壳/混淆、虚拟机、校验逻辑
   - Forensics：图片、PDF、PCAP、内存、日志、注册表、磁盘、元数据
   - Misc：编码谜题、pyjail、bashjail、二维码、逻辑题、混合题
3. 如果题型不清晰，先做最轻量的三步：
   - 文件类型与字符串
   - 元数据与结构
   - 只读网络或入口探测

## 首轮工作流

## 通用 triage

- 明确目标：拿 flag、拿 shell、恢复明文、绕过校验、找到隐藏数据。
- 保存原始样本与题目说明，不要一开始就覆盖文件。
- 对附件记录文件名、大小、类型、压缩层级、可执行性和外部依赖。

## Web 题

- 先枚举入口、参数、身份态、调试路径和静态资源。
- 关注模板注入、鉴权缺陷、文件读写、接口参数、前端暗藏逻辑。
- 常用起手工具：`http-framework-test`、`ffuf`、`sqlmap`、`nuclei`。

## Crypto 题

- 先分清是编码、古典、对称、非对称、签名、哈希还是随机数问题。
- 把题目里的常量、密文格式、块长、模数、指数、nonce、IV、已知明文整理清楚。
- 避免在题型未判明前盲跑暴力。

## Pwn 题

- 先确认架构、保护、输入面、崩溃点、远程交互方式。
- 关注 `checksec` 类信息、函数调用链、栈堆对象、格式串与 syscall 面。
- 常用起手工具：`strings`、`gdb`、`pwntools`、`ropper` / `ROPgadget`、`one-gadget`。

## Reverse 题

- 先找入口函数、核心校验、密钥来源、常量表、状态机与反调试。
- 确认是“理解算法”还是“提取硬编码秘密”，不要过早做大规模动态分析。
- 常用起手工具：`strings`、`ghidra`、`radare2`、`objdump`。

## Forensics 题

- 先做文件指纹、元数据、嵌套结构、时间线、隐藏数据面检查。
- 对镜像和流量保持只读分析，必要时复制工作副本。
- 常用起手工具：`binwalk`、`exiftool`、`foremost`、`volatility3`、`zsteg`。

## Misc 题

- 把输入拆成“编码/规则/执行环境/交互限制”四部分。
- 对 pyjail、bashjail、逻辑题、二维码、音频图像混合题，优先找最小突破口。

## 何时切换方向

- Web：探测结果与业务逻辑严重不符，且 HTTP 面没有更多入口时，回头检查前端脚本或题目附件。
- Crypto：样本更像编码/混淆而非真正密码学时，不要继续按高阶密码方向硬攻。
- Pwn：如果核心难点其实是弄清程序做什么，应切到 reverse。
- Reverse：如果已经拿到清晰漏洞原语，再转到 pwn 完成利用。
- Forensics：如果附件真实本体是脚本/二进制，应及时转到 reverse 或 misc。

## 证据与输出

- 记录每一步的观察、假设、验证结果和下一步分流理由。
- 题解应包含：
  - 输入与题型判断
  - 首轮 triage 结果
  - 关键突破点
  - 复现步骤
  - flag 获取过程

## 约束

- 优先使用轻量、可复现、可解释的路径。
- 没有证据时，不要把题目强行归到复杂题型。
- 当某条路线连续两轮没有产生新信息，就回到分流点重新判断。

