package idleisv

// UserPermitCmd 结构体
type UserPermitCmd struct {
	// 当前用户的所属业务类型编码，优品&amp;开放平台业务 默认使用 IDLE_TOP
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
}
