package zvalidator

// ValidatorHandler 自定义验证函数
type ValidatorHandler func(value any, extra ...any) bool

// Rule 验证规则
type Rule struct {
	Type       string             // 验证规则类型
	Message    string             // 自定义错误信息
	Value      any                // 用于传递验证所需要的参数,例如 min的最小值
	Validators []ValidatorHandler // 验证函数列表
}

type Rules map[string][]Rule

const (
	REQUIRED = "required"
	MIN      = "min"
	MAX      = "max"
	EMAIL    = "email"
	NUMERIC  = "numeric"
)
