# TranslateAndSaveMysqlByXOR-M

百度翻译接口查询英文单词对应的中文翻译并放入数据库中保存，以便下一次相同查询的时候直接从数据库中查找

其中放置根目录下的配置文件格式如下

```ini
[api]
appid = <百度开放平台的appid>
key = <百度开放平台的key>
from = <源语言> # auto即为自动判断 
to = <目标语言>
[mysql]
user = <数据库用户名>
passwd = <数据库密码>
database = <数据库名>
[log]
level = Debug
```

## 带参数使用方法
第一个参数即为要翻译的外文
+ 不带参数使用方法
会自动翻译同级目录下`doc.txt`的文本