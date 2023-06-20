## http2ok

### 00 功能

此工具开发目的是攻防场景下的资产梳理。

主要功能如下：

- [x] 请求协议自动判断
- [x] 多线程并发
- [x] 可指定请求路径
- [x] 获取请求状态码、标题等
- [x] 域名备案查询【 [ICP/IP地址/域名信息备案管理系统](https://beian.miit.gov.cn/#/Integrated/recordQuery) 备案查询接口】，支持单独查询和url域名查询
- [x] 实现代理功能，支持http和socks5
- [x] 目标文件导入，支持ip:port和domain/url两种格式
- [x] 结果导出



### 01 参数

```
Usage of http2ok.exe:
  -checkproxy
        check the proxy is valid
  -code string
        set success code,set '*' will match all code, -code 200,302 (default "*")
  -i string
        icp check target,eg: -i www.example.com
  -icp
        set url icp check flag
  -icpsize string
        icppagesize,eg:-icpsize 10 (default "40")
  -if string
        icpfile,eg: -if domain.txt
  -o string
        Outputfile (default "result.txt")
  -path string
        http request path (default "/")
  -proxy string
        set poc proxy,eg: -proxy http://127.0.0.1:8080
  -socks5 string
        set socks5 proxy, will be used in tcp connection, timeout setting will not work
  -t int
        Thread nums (default 600)
  -tf string
        target file,eg: -tf ip.txt
  -time int
        Set timeout (default 3)
  -u string
        url
```

### 02 使用示例

#### 1 单个URL检测

使用`-u`指定单个url进行检测，单URL指定path时替换路径而非追加到url后面，以下示例中若指定path="/ccc.php"，则最终路径为`http://example.com/ccc.php`

```
http2ok.exe -u http://example.com/aaa/bbb.jsp
```

#### 2 批量目标检测

使用`-tf`参数进行批量目标检测，包括状态码，title，ICP备案等

```
http2ok.exe -tf target.txt -code 200,302,403
```

使用`-code`指定成功状态码，默认为`*`即所有状态码都为成功

```
http2ok.exe -tf target.txt -code 200,302,403
```

使用`-path`指定检测路径，默认为根目录`/`

```
http2ok.exe -tf target.txt -code 200,302,403 -path "/admin/"
```

#### 3 仅ICP备案查询

使用`-icp`指定ICP查询关键字，支持域名、IP、企业名称，使用`-icpsize`指定页面条数，默认40条

```
http2ok.exe -icp keyword -icpsize 100
```

使用`-if`参数进行批量ICP备案查询

```
http2ok.exe -if target.txt
```

#### 4 文件保存

target检测结果保存至`result.txt`，可使用`-o`指定导出目录，成功url保存至`okurl.txt`方便其他工具使用,备案查询结果保存至`ICP.txt`。

### 03 参考文档

1、参考`shadow1ng`大佬的[fscan](https://github.com/shadow1ng/fscan) 项目，中关于目标格式化、协议判定及代理功能

2、参考`fghwett`大佬的[icp](https://github.com/fghwett/icp)项目，在此基础上实现批量查询，支持域名、IP、企业名称查询

### todo

- [ ] 备案查询多页结果，目前仅支持当前页，默认40条

