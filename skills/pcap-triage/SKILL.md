---
name: pcap-triage
description: 面向 CTF 流量题的首轮 triage 技能包，用于梳理 PCAP/网络日志中的协议、时间线与可疑对象。
metadata:
  version: "1.0.0"
  tags:
    - ctf
    - pcap
    - network
    - forensics
    - traffic
  triggers:
    - pcap
    - traffic
    - packet capture
    - wireshark
    - network forensics
    - ctf traffic
  recommended_tools:
    - strings
    - foremost
    - exiftool
  role_hints:
    - CTF
    - 数字取证
  stages:
    - triage
    - verify
    - solve
  autoload_priority: 18
---

# PCAP 流量题首轮分析

## 何时使用

当题目给出 `.pcap`、`.pcapng`、导出的网络日志或“某台主机流量”样本时使用。

## 快速流程

1. 先看整体时间线、会话数量、协议分布、异常端口和异常体量。
2. 再定位题目重点：明文凭据、文件传输、隧道、恶意流量、DNS、HTTP、TLS、内网横移痕迹。
3. 如果不能立刻判断题型，先做只读提取：字符串、对象、可见文件名、主机名、域名。

## 重点检查

## 通信画像

- 谁和谁通信、何时开始、何时异常、哪类协议最突出。
- 是否存在周期性 beacon、单向长连接、短时大量失败、异常域名。

## 明文与对象

- HTTP、FTP、SMTP、POP、IMAP、Telnet 等是否直接泄露关键信息。
- 是否有明显文件名、下载路径、上传对象、附件、图片、压缩包。

## DNS 与名称信息

- 可疑域名、拼写变体、随机子域、内外网混合查询。
- DNS 往往能给出题目真正聚焦的主机与时间点。

## 输出要求

- 时间线摘要
- 关键主机/域名/端口
- 最可疑的 3 个会话或对象
- 下一步深挖方向

