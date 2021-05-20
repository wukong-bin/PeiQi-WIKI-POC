# 默安 幻阵蜜罐未授权访问 RCE

## 漏洞描述

默安 幻阵蜜罐存在部署页面未授权访问 ，可执行任意命令

## 漏洞影响

> [!NOTE]
>
> 默安 幻阵蜜罐

## 漏洞复现

产品页面

![](http://wikioss.peiqi.tech/vuln/mo-1.png)

安装页面如下

默安 幻阵蜜罐![](http://wikioss.peiqi.tech/vuln/mo-2.png)

刷新并抓包

![](http://wikioss.peiqi.tech/vuln/mo-3.png)

Drop掉 **/huanzhen/have_installed?**

![](http://wikioss.peiqi.tech/vuln/mo-4.png)

进入页面

![](http://wikioss.peiqi.tech/vuln/mo-5.png)

点击调试抓包

![](http://wikioss.peiqi.tech/vuln/mo-6.png)

执行其他命令

![](http://wikioss.peiqi.tech/vuln/mo-7.png)

点击一键诊断泄露 IP数据

![](http://wikioss.peiqi.tech/vuln/mo-8.png)