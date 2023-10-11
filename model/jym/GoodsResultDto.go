package jym

// GoodsResultDto 结构体
type GoodsResultDto struct {
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 状态码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否下架成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
