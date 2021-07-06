package fundplatform

// CardFetchDetailDto 结构体
type CardFetchDetailDto struct {
	// 制卡模板号
	TemplateNo string `json:"template_no,omitempty" xml:"template_no,omitempty"`
	// 制卡数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 售价,单位为分。不填写则使用模板上配置的默认售价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}
