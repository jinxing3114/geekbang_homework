## 学习方法：
1. 工作日在乘坐公共交通观看课程视频，对于有疑问或者不理解的反复观看视频，做到有印象、有概念，下班后在实操。
2. 先按照老师的教程方法，严格操作，争取能尽快看到效果。如有问题查资料看文档，无法解决就通过群和周末直播请教。
3. 然后在老师教程基础上调整参数，步骤，扩大或缩小范围，以此检测可能会出现的问题。
4. 重点还是尽量能在实际工作中使用，这样学习效果，收获经验都会非常好。
## 学习心得：
1. 搭建两套小集群环境，一个为测试开发环境，一个为稳定长期使用，两套环境轮流使用。
2. 学习期间使用测试开发环境，出现不可恢复错误不至于导致学习受阻，也可大胆做实验，可在灾难现场排查问题，更多提升处理问题能力。
3. 尝试后可控操作在稳定环境使用，有助于呈现所有功能模块聚合在一起效果，更多的模拟真实环境会出现的问题。
## 学习总结：
1. 课程内容很丰富，含有大量理论知识，需要花费较多时间理解。
2. 实操内容也很多，但是课程并没有全部演示，课余时间需要认真跟教程实操。
3. 如有问题及时和助教老师沟通，不管是课程内容，还是自己操作碰到的问题。
4. 每周领教直播最好都参加，日常的课程是录播，不能及时反馈交流，领教直播是每周总结后教学，更能及时的反馈我们碰到的问题。

## 个人分享(仅供参考)
1. ### go语言应用
    1. //go:embed，资源文件编译嵌入到应用程序中，降低部署分发难度，减少依赖资源异常 
    2. 将使用到的资源文件编译进程序中，比如、web应用的模板文件、静态资源文件(css/js/img)，默认配置文件、数据库sql文件等。 一个可执行文件，包含主程序和附加资源文件。
    3. atomic包，sync.Map，sync.Mutex等，goroutine安全。在竞态下更有利于解决高并发问题。
    4. 包名internal，限制外部引用访问等。对外部提供基础库或项目底层代码。防止未经允许直接调用导致程序错误。
    5. GODEBUG=gctrace=1 ./main 查看应用gc频率和次数，高频gc会影响应用性能。
    6. 编译指定平台，如不指定，可能会随打包环境不同，打包出可执行平台不同。关闭cgo，不依赖外部库。减少异常因素出现。 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
2. ### docker镜像打包小技巧
   1. 尽量使用最小的基础镜像，分发镜像速度更快。
   2. scratch基础镜像，空镜像，占用空间最小，但没有可用工具，不利于排查问题。
   3. alpine镜像，轻量(5MB左右)、安全、提供APK包管理，可根据需求在镜像打包期间安装其他软件。
   4. busybox镜像，工具箱集成很多linux命令和工具、轻量(2MB左右)，无包管理器。
   5. 拉取镜像验证sha256签名，对于相同tag镜像，可能会存在修改或覆盖，导致相同tag不同镜像。
   6. 多阶段构建，屏蔽相互之间依赖、环境、配置等差异带来的影响。
3. ### kubernetes环境
   1. docker desktop自带kubernetes，安装较方便，本地调试更容易，出现问题重置环境即可。
