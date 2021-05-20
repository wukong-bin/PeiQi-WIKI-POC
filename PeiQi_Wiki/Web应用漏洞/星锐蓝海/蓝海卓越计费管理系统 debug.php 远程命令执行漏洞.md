# 蓝海卓越计费管理系统 debug.php 远程命令执行漏洞

## 漏洞描述

蓝海卓越计费管理系统 debug.php 存在命令调试页面，导致攻击者可以远程命令执行

## 漏洞影响

> [!NOTE]
>
> 蓝海卓越计费管理系统

## FOFA

> [!NOTE]
>
> title=="蓝海卓越计费管理系统"

## 漏洞复现

登录页面如下

![](http://wikioss.peiqi.tech/vuln/lh-1.png)

漏洞代码

![](http://wikioss.peiqi.tech/vuln/lh-5.png)

访问 debug.php页面 远程调试命令执行

![](http://wikioss.peiqi.tech/vuln/lh-4.png)

