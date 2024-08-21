# zvalidator

**zvalidator** 是一个轻量级的 Go 语言参数 验证库, 提供简洁易用的 API, 帮助你轻松的进行数据验证。

## 特性

* **支持多种验证规则:** 包括必填验证、最小值/最大值验证、范围验证、正则表达式验证等。
* **自定义验证器:** 支持自定义验证函数,可以满足各种特殊的验证需求。
* **支持嵌套结构体:** 可以验证嵌套结构体中的字段。
* **易于是用:** 提供辅助函数,简化规则定义。

## 安装 

`go get github.com/zR-Zr/zvalidator.git`

## 使用方法

### 定义验证规则
```go
rules := zvalidator.Rules{
    "name": {
        zvalidator.RequiredRule("姓名不能为空"),
        zvalidator.MinRule(2, "姓名长度至少为 2"),
        zvalidator.MaxRule(16, "姓名长度最多为 16"),
    },
    // ... 其他字段的验证规则
}
```

### 执行验证

```go
data := map[string]interface{}{
    // ... 待验证的数据 ...
}

isValid, errors := zvalidator.Validate(data, rules)

if isValid {
    // 验证通过
} else {
    // 验证失败，处理错误信息
    fmt.Println(errors)
}
```

### 自定义验证器

```go
ules := zvalidator.Rules{
    "email": {
        zvalidator.Rule{
            Type:    "email",
            Message: "请输入有效的电子邮件地址",
            CustomValidator: func(value any, rawData map[string]any) bool {
                // 自定义验证逻辑
                return true
            },
        },
    },
}
```

### 嵌套结构体
```go
rules := zvalidator.Rules{
    "address.city": {
        zvalidator.RequiredRule("城市不能为空"),
    },
}

data := map[string]interface{}{
    "address": map[string]interface{}{
        "city": "Beijing",
    },
}
```

## 许可证

MIT License
