package main

import (
	"fmt"
	"regexp"

	"github.com/zR-Zr/zvalidator.git"
)

// "age":             199,
// 		"email":           "johndoe@example.com",
// 		"password":        "passwor",
// 		"confirmPassword": "password123",
// 		"city":            "北京",
// 		"address": map[string]interface{}{
// 			"city":   " ", // 空字符串，会触发 required 验证
// 			"street": "Short St",
// 		},

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type User struct {
	Name            string   `json:"name"`
	Age             int      `json:"age"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	ConfirmPassword string   `json:"comfirm_password"`
	City            string   `json:"city"`
	Address         *Address `json:"address"`
}

func main() {

}

func structExample() {

}

func mapExample() {
	rules := zvalidator.Rules{
		"name": {
			zvalidator.RequiredRule("姓名不能为空"),
			zvalidator.MinRule(2, "姓名长度至少为 2"),
			zvalidator.MaxRule(16, "姓名长度最多为 16"),
		},
		"password": {
			zvalidator.RequiredRule("密码不能为空"),
			zvalidator.RangeRule(8, 16, "密码长度应该在 8 到 16 之间"),
		},
		"age": {
			zvalidator.MinRule(18, "年龄不能小于 18"),
			zvalidator.RangeRule(18, 120, "年龄应该在 18 到 120 之间"),
		},
		"email": {
			zvalidator.RequiredRule("Email 不能为空"),
			zvalidator.Rule{
				Type:    "email",
				Message: "请输入有效的电子邮件地址",
				CustomValidator: func(value any, rawData map[string]any) bool {
					valueStr := value.(string)

					// 使用正则表达式检查是否是 电子邮件格式
					pattern := `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
					matched, err := regexp.Match(pattern, []byte(valueStr))
					if err != nil {
						panic(err)
					}

					return matched
				},
			},
		},
		"address.city": {
			zvalidator.RequiredRule("城市不能为空"),
		},
	}

	data := map[string]interface{}{
		// "name":            "John Doe",
		"age":             199,
		"email":           "johndoe@example.com",
		"password":        "passwor",
		"confirmPassword": "password123",
		"city":            "北京",
		"address": map[string]interface{}{
			"city":   " ", // 空字符串，会触发 required 验证
			"street": "Short St",
		},
	}

	isValid, errors := zvalidator.Validate(data, rules)
	if isValid {
		println("验证通过！")
	} else {
		fmt.Println(errors)
	}
}

// func main() {
// 	error
// 	rules := zvalidator.Rules{
// 		"Name": {
// 			{Type: "required", Message: "姓名不能为空"},
// 			{Type: "min", Message: "姓名长度至少为 200", Value: 200},
// 		},
// 		"Age": {
// 			{Type: "numeric", Message: "年龄必须是数字"},
// 			{Type: "min", Message: "年龄不能小于 18", Value: 18},
// 			{Type: "max", Message: "年龄不能大于 120", Value: 120},
// 		},
// 		"Email": {
// 			{Type: "required", Message: "Email 不能为空"},
// 			{Type: "email", Message: "请输入有效的电子邮件地址"},
// 		},
// 		"Password": {
// 			{Type: "required", Message: "密码不能为空"},
// 			{Type: "min", Message: "密码长度至少为 6", Value: 6},
// 		},
// 		"ConfirmPassword": {
// 			{Type: "required", Message: "确认密码不能为空"},
// 			{
// 				Type:    "confirmPassword", // 使用注册的验证函数名
// 				Message: "两次输入的密码不一致",
// 			},
// 		},
// 		"City": {
// 			{
// 				Type:    "in",
// 				Message: "选择的城市无效",
// 				Value:   []interface{}{"北京", "上海", "广州", "深圳"}, // 注意这里的类型
// 			},
// 		},
// 		"Address.City": { // 使用点分隔符表示嵌套字段
// 			{Type: "required", Message: "城市不能为空"},
// 		},
// 		"Address.Street": {
// 			{Type: "min", Message: "街道长度至少为 3", Value: 3},
// 		},
// 	}

// 	data := map[string]interface{}{
// 		"Name":            "John Doe",
// 		"Age":             25,
// 		"Email":           "john.doe@example.com",
// 		"Password":        "password123",
// 		"ConfirmPassword": "password123",
// 		"City":            "北京",
// 		"Address": map[string]interface{}{
// 			"City":   " ", // 空字符串，会触发 required 验证
// 			"Street": "Short St",
// 		},
// 	}
// 	isValid, errors := zvalidator.Validate(data, rules)
// 	if isValid {
// 		fmt.Println("验证通过！")
// 	} else {
// 		fmt.Println("验证失败：")
// 		for _, err := range errors {
// 			fmt.Printf("字段：%s，错误码：%s，错误信息：%s\n", err.Field, err.Code, err.Message)
// 		}
// 	}
// }
