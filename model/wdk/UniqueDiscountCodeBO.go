package wdk

// UniqueDiscountCodeBo 结构体
type UniqueDiscountCodeBo struct {
	// 唯一码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}
