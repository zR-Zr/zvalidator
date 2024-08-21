package optional

// Optional 是一个包装类型,用于区分字段是否被显示的设置
type Optional[T any] struct {
	value T
	set   bool
}

// New 返回一个设置了值的 Optional
func New[T any](value T) Optional[T] {
	return Optional[T]{value: value, set: true}
}

// Empty 返回一个未设置值的 Optional
func Empty[T any]() Optional[T] {
	var zero T
	return Optional[T]{value: zero, set: false}
}

// IsSet 判断字段是否被设置
func (o Optional[T]) IsSet() bool {
	return o.set
}

// Value 返回字段的值
func (o Optional[T]) Value() T {
	return o.value
}
