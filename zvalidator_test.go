package zvalidator

import (
	"regexp"
	"testing"
)

func TestValidateStruct(t *testing.T) {
	age := 25
	validAddress := &Address{
		City:   "Beijing",
		Street: "Long St",
	}

	inValidator := func(value any, rawData map[string]any, rule Rule) bool {
		if len(rule.In) == 0 {
			panic("in validator need In ")
		}

		for _, v := range rule.In {
			if v == value {
				return true
			}
		}
		return false
	}

	RegisterValidator("in", inValidator)

	testCases := []struct {
		name     string
		input    User
		rules    Rules
		wantErr  bool
		errCount int // 预期的错误数量
	}{
		{
			name: "Valid User",
			input: User{
				Name:            "John Doe",
				Age:             &age,
				Email:           "john.doe@example.com",
				Password:        "password123",
				ConfirmPassword: "password123",
				City:            "北京",
				Address: Address{
					City:   "Beijing",
					Street: "Long St",
				},
				PtrAddress: validAddress,
			},
			rules: Rules{
				"name": {
					RequiredRule("姓名不能为空"),
					MinRule(2, "姓名长度至少为 2"),
					MaxRule(16, "姓名长度不能超过 16"),
				},
				"age": {
					MinRule(18, "年龄不能小于 18"),
					MaxRule(120, "年龄不能大于 120"),
				},
				"email": {
					RequiredRule("Email 不能为空"),
					Rule{
						Type:    "email",
						Message: "请输入有效的电子邮件地址",
						CustomValidator: func(data any, rawData map[string]any) bool {
							valueStr, ok := data.(string)
							if !ok {
								return false
							}
							pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
							matched, err := regexp.MatchString(pattern, valueStr)
							if err != nil {
								panic(err) // 实际应用中应该返回错误
							}
							return matched
						},
					},
				},
				"password": {
					RequiredRule("密码不能为空"),
					RangeRule(8, 16, "密码长度应该在 8 到 16 之间"),
				},
				"confirm_password": {
					RequiredRule("确认密码不能为空"),
					Rule{
						Type:    "confirmPassword", // 使用注册的验证函数名
						Message: "两次输入的密码不一致",
						CustomValidator: func(data any, rawData map[string]any) bool {
							password, ok := rawData["password"]
							if !ok {
								return false
							}
							return password == data
						},
					},
				},
				"city": {
					Rule{
						Type:    "in",
						Message: "选择的城市无效",
						In:      []any{"上海", "深圳", "广州"},
					},
				},
				"address.city": {
					RequiredRule("城市不能为空"),
				},
				"address.street": {
					MinRule(3, "街道长度至少为 3"),
				},
				"ptr_address.city": {
					RequiredRule("城市不能为空"),
				},
				"ptr_address.street": {
					MinRule(3, "街道长度至少为 3"),
				},
			},
			wantErr:  false,
			errCount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateStruct(tc.input, tc.rules)
			// if err != nil {
			// 	if validationErros, ok := err.(ValidationErrors); ok {
			// 		t.Log(tc.name, "errors: ", validationErros.GetErrors())
			// 	} else {
			// 		t.Log("error: ", err)
			// 	}
			// }
			if (err != nil) == tc.wantErr {
				t.Errorf("ValidateStruct() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr {
				if validationErrors, ok := err.(ValidationErrors); ok {
					if len(validationErrors.GetErrors()) != tc.errCount {
						t.Errorf("ValidateStruct() expected %d errors, but got %d", tc.errCount, len(validationErrors.GetErrors()))
					}
				} else {
					t.Errorf("ValidateStruct() expected ValidationErrors, but got %T", err)
				}
			}
		})
	}
}
