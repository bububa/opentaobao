package alihouse

// AstoreRespDto 结构体
type AstoreRespDto struct {
	// 链接
	ShopUrl string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
	// 装修入口
	DecorationUrl string `json:"decoration_url,omitempty" xml:"decoration_url,omitempty"`
	// 模版名称
	DecorationName string `json:"decoration_name,omitempty" xml:"decoration_name,omitempty"`
	// 外部目标ID
	OuterTargetId string `json:"outer_target_id,omitempty" xml:"outer_target_id,omitempty"`
	// id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部目标ID
	OuterTargetType int64 `json:"outer_target_type,omitempty" xml:"outer_target_type,omitempty"`
}
