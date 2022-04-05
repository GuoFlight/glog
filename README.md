# 简介

作者：京城郭少

# 关于项目

* 依赖logrus和file-rotatelogs

# 特点

* 每个level都会独立输出到不同的文件。
* 支持日志切割。
* 支持输出文件名、函数名和行号。
* 非DEBUG模式不输出到终端。

# 注意事项

* rotatelogs这个库对Win10老版本、win10以前的版本、go1.11以下下版本支持得很不好，可能出现无法创建软连接的问题。

# Question

问：为什么要自己实现一个输出行号的Hook？
> 答：因为logrus作者也不建议输出行号，这会降低程序的性能。这也是logrus迟迟不增加这项功能的原因。而我的Hook在INFO这个level上是不打日志的，所以性能的损耗可以忽略。

