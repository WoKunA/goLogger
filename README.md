# Logger

这是一款使用 Go 语言编写的可以自定义 log 的程序，打印所有 struct 格式的日志（这将对 ES 有好处，当然你也可以自己定义输出方式）,解决了日志过大导致磁盘不足的问题（如果你不想写一些定期脚本）。
支持个性化定制，你可以按照日期来分割日志，便于你接下来的分析。接下来的各种功能后续补上（当然，读者大大也可以提，绝对立马就做）。

## 使用说明（目前支持的方法）
我们默认了一个 type Template func(v interface{}) string，你也可以将你自己写的函数挂载上去。
*  自定义 log 地址
*  按日期分割

## 超级快速开始
直接引用 logger.DefaultLog.Log() ，就可以打印日志了，日志在当前目录的 data/log/ 下，每隔 24 小时，换一个日志文件打印日志，
每份日志保留 7 天。
```
    person := &Person{
		Age: 10,
	}
    logger.DefaultLog.Log(person)
```

## 快速开始
*   修改 logger.yml 文件
    *   Piece   : 这是分片，表示你需要将原本一份日志分成几分
    *   Timer   : 这是时间，表示需要过多久将日志输出下一份日志文件中
    *   FileFlag: 这是文件 Flag，表示如何打开文件，以后会支持 os.O_CREATE|os.O_TRUNC|os.O_WRONLY 形式
    *   LoopLogFile: 表示是否将日志切分，如果为 false 将不会创建多个文件供日志输出
    *   FileName   : 输出完整的路径和文件，如果不输入，则默认为当前文件夹下的 data/log/ 里，没有也会创建文件夹
然后
```
    person := &Person{
		Age: 10,
	}
    logger.DefaultLog.Log(person)
```
 
*   自定义输出方式
    *   编写一个符合 fn func(v interface{}) string 的函数传入 SetTemplete
```
    person := &Person{
		Age: 10,
	}

    func LoadToLogger(v interface{}) string {
	    TransformJson, err := json.Marshal(v)
	    if err != nil {
	    	panic(err)
	    }
	    return string(TransformJson)
    }
    logger.DefaultLog.SetTemplete(LoadToLogger)
    logger.DefaultLog.Log(person)
```