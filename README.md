## CI/CD 概念介绍

在软件开发的过程中，会经历多次的迭代更新。每次更新，有的是为了修复历史的bug，有的是为了引入新的功能。伴随这这些新的改动，必然也会导致代码的变动，而代码的变动则会导致项目出现bug；或者新代码虽然没有问题，但是代码写的很烂，对未来维护项目带来困难。

以前并不是没有CI/CD流程，而是以前的流程是人工的，繁琐的。只要是人工的，每次都必然是不稳定的，都是手工艺品，正如同我们编写代码一样，总是充满着各种问题，并且耗时甚久。

而为了解决上面的问题，我们需要尽量引入**自动化**工具和流程，为我们减轻审核代码负担，减轻部署项目的时间，以便于更快的交付新功能，修复遗留的bug，并且每次都能稳定的成功部署项目，把软件交付给我们的用户。

#### CI 持续集成
集成指的是开发过程中，新代码逐渐通过各种检测合入项目主干的过程。
在这个过程中，一般会做这样一些事情：

1. 代码检查(review)
    1.1 风格检查:
    > 对代码规范进行检查，看看是不是有长函数，大小写不规范，公开函数不写注释这样的问题。

    1.2 安全扫描:
    > 检查代码中是否有安全漏洞，比如内存溢出，SQL注入，密钥泄露等风险。

    1.3 功能走查:
    
> 一般是请另一名开发人员检查代码逻辑是否有问题，代码实现方法是否合理。
    
2. 集成测试
    
    > 测试新代码的程序是否能正常工作，是否和预期一致。

#### CD 持续部署
部署就是把编译成功的程序成功交付给用户的过程。网站型软件就是部署到生产服务器，也有客户端软件就是上传到应用商店等。由于这个部分根据不同的项目，公司，环境五花八门，自然也难以统一展示，因此本项目注重于展示上面的CI部分。

## Github Action
在线文档：https://docs.github.com/cn/actions