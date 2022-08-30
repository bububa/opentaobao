package scbp

// RelevantProductDto 结构体
type RelevantProductDto struct {
	// 推荐理由
	ReasonList []string `json:"reason_list,omitempty" xml:"reason_list>string,omitempty"`
	// 品图片链接
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 星级
	QsStar int64 `json:"qs_star,omitempty" xml:"qs_star,omitempty"`
	// 品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 操作优推品得主键id
	AdGroupId int64 `json:"ad_group_id,omitempty" xml:"ad_group_id,omitempty"`
	// 是否优推
	IsPreferential bool `json:"is_preferential,omitempty" xml:"is_preferential,omitempty"`
}
