# TranslateAndSaveMysqlByXOR-M

百度翻译接口查询英文单词对应的中文翻译并放入数据库中保存，以便下一次相同查询的时候直接从数据库中查找

其中放置根目录下的配置文件格式如下

```ini
[api]
appid = <百度开放平台的appid>
key = <百度开放平台的key>
from = <源语言>
to = <目标语言>
[mysql]
user = <数据库用户名>
passwd = <数据库密码>
database = <数据库名>
```