# rwplus-shorturl

短URL跳转长URL

## 获取shorten

curl -X POST "http://localhost:9000/j/shorten" -d "url=http://www.baidu.com"

**NOTE：** 如果已经生成过shorten，则从数据中匹配长url，返回旧的shorten。如果从未生成过shorten，则新生成一个shorten。返回新的shorten

## 获取长URL

curl -X POST "http://localhost:9000/j/shorten" -d "shorten=abcde"
