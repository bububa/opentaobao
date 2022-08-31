package alihouse

// TokenCreateDto 结构体
type TokenCreateDto struct {
	// 授权码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否测试(1-是测试 0-不是测试)
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
