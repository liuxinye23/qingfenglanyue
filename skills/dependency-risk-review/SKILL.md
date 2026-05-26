---
name: dependency-risk-review
description: 审计依赖版本、锁文件、供应链暴露面与构建流程风险的技能包。
metadata:
  version: "1.0.0"
  tags:
    - dependency
    - sbom
    - supply-chain
    - lockfile
    - package
  triggers:
    - dependency
    - package risk
    - lockfile
    - vulnerable library
    - supply chain
    - sbom
  recommended_tools:
    - trivy
    - exec
    - checkov
  role_hints:
    - secure-code-review
    - cloud-security-audit
    - 容器安全
  stages:
    - recon
    - verify
    - report
  autoload_priority: 14
---

# 依赖与供应链风险审计

## 何时使用

当目标包含应用代码库、镜像、制品、锁文件、构建脚本、包管理器配置或内部依赖仓库时使用本技能。

## 快速流程

1. 确认生态与资产：`package-lock`、`poetry.lock`、`requirements`、`go.sum`、镜像清单、制品仓库。
2. 区分直接依赖、传递依赖、开发依赖、构建依赖和运行时依赖。
3. 审核版本固定策略、来源可信度、过时组件、弃用组件、镜像基线。
4. 检查供应链控制：私仓、镜像源、构建签名、最小镜像、变更审计。

## 重点检查

## 版本与锁定

- 是否只在描述文件声明范围版本，缺少可重现锁定。
- 关键安全组件是否长时间停留在已弃用或不再维护的版本。

## 来源与信任

- 是否混用公共源、镜像源、私有包和临时脚本安装。
- 构建链是否依赖未经审计的下载步骤、在线脚本或不稳定镜像。

## 运行时暴露面

- 运行镜像中是否带入构建工具、测试依赖、调试包和多余 shell。
- 不同环境是否复用同一超大镜像基线。

## 建议工具

## trivy

- 用于依赖、镜像与文件系统层面的已知风险盘点。
- 输出要结合运行环境和组件可达性做二次判断，不要原样搬运漏洞列表。

## exec

- 用只读命令核对包管理器文件、锁文件、镜像基线、构建脚本与版本固定策略。

## checkov

- 当依赖风险与 IaC、镜像构建、云资源定义耦合时，用来补充配置面问题。

## 证据要求

- 记录组件名、版本、所在文件、用途、是否直接依赖、是否运行时可达。
- 对高危问题写明“为什么在当前系统里真正重要”。
- 区分可立即修复、需兼容验证、需架构替换三类建议。

## 修复建议方向

- 固定锁文件并建立更新窗口。
- 缩减运行时镜像与非必要依赖。
- 对公共源、私有源与构建下载流程建立信任边界和审计。

