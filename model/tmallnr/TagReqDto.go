package tmallnr

// TagReqDto 结构体
type TagReqDto struct {
	// 商品编码列表
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 业务身份标识,dss定时送;FN蜂鸟,CN菜鸟
	BizIdentity string `json:"biz_identity,omitempty" xml:"biz_identity,omitempty"`
	// 信息追踪串,用于后续排查问题,以及与外部厂商对账等场景下使用
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 商品标操作类型，1:打标，2:去标
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 商家编码
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
