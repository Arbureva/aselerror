[English](README-en.md)

# aselerror

`aselerror` 是一个专门处理“异常API接口响应”的错误处理包。

在传统的响应处理中，我们通常需要判断错误信息是否敏感，然后才能决定是否返回给前端。频繁使用`if`和`switch`语句可能不是最佳选择。
`aselerror`可以通过比较错误等级来简化这一过程，从而显著减少接口响应时间。


## 注意事项
- 由于`aselerror`设计较为简单，建议与`errors`包结合使用以获取堆栈信息。
- `aselerror`主要用于处理异常API接口响应。在正常情况下，建议避免使用它，因为错误信息应保持简洁明了。

## 安装

```
go get github.com/asel-go/aselerror
```

## 使用

请查看 `error_test.go` 文件 => [error_test.go](error_test.go)