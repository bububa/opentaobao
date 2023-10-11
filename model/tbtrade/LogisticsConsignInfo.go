package tbtrade

// LogisticsConsignInfo 结构体
type LogisticsConsignInfo struct {
	// 成分品发货时效
	CombineItem []CombineConsignInfo `json:"combine_item,omitempty" xml:"combine_item>combine_consign_info,omitempty"`
	// 时效关联的订单好
	RelatedId string `json:"related_id,omitempty" xml:"related_id,omitempty"`
	// 商家的预计发货时间
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 商家的预计发货时间
	RenderConsignTime string `json:"render_consign_time,omitempty" xml:"render_consign_time,omitempty"`
}
