package tmallnr

// TagRespDto 结构体
type TagRespDto struct {
	// 失败描述
	Descs []string `json:"descs,omitempty" xml:"descs>string,omitempty"`
	// 失败商品编码
	FailIds []int64 `json:"fail_ids,omitempty" xml:"fail_ids>int64,omitempty"`
	// 成功商品编码
	SuccessIds []int64 `json:"success_ids,omitempty" xml:"success_ids>int64,omitempty"`
}
