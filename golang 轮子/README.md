## Read.go
Read.go中是一些反爬的函数，我只写了Readua,即随机生成User-Agent，需要的话，referer和XFF头也可以自己写相应的函数，然后去修改get和post函数。

## report.go
report.go中封装的是get请求和post请求的函数。
- get请求直接就是`url+value`请求即可
~~~
url := "https://www.baidu.com"
username := "tunan"
age := "18"
url = url + "?username=" + username + "&age=" + age 
response := lib.Get(url)

~~~
- post请求为了实现动态的传入参数，所以接收两个参数，第一个是url，第二个是post的数据，具体使用如下
~~~

  data := make(url.Values)
  url := "https://www.baidu.com"
  num := "18"
  data["username"] = []string{"tunan"}
  data["age"] = []string{num}
  
  response := lib.Post(url,data)
  
  //因为我习惯把这些文件放到lib包下，所以我的引用就是lib.xxx()
~~~

## json解析库
还有一个很好用的json解析库
~~~
https://github.com/tidwall/gjson
它上面有文档，可以看看，转为自己需要的数据。
~~~



