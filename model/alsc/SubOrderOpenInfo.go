package alsc

// SubOrderOpenInfo 结构体
type SubOrderOpenInfo struct {
	// 子单业务上下文
	BizContext string `json:"biz_context,omitempty" xml:"biz_context,omitempty"`
	// 子单业务状态
	BizStatus string `json:"biz_status,omitempty" xml:"biz_status,omitempty"`
	// 子单状态中文描述
	BizStatusDesc string `json:"biz_status_desc,omitempty" xml:"biz_status_desc,omitempty"`
	// consumeDesc
	ConsumeDesc string `json:"consume_desc,omitempty" xml:"consume_desc,omitempty"`
	// 子单扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 图片地址
	ItemPicUrl string `json:"item_pic_url,omitempty" xml:"item_pic_url,omitempty"`
	// 子单记录的时间提示
	ValidDesc string `json:"valid_desc,omitempty" xml:"valid_desc,omitempty"`
	// 有效期结束时间
	ValidEnd string `json:"valid_end,omitempty" xml:"valid_end,omitempty"`
	// 有效期开始时间
	ValidStart string `json:"valid_start,omitempty" xml:"valid_start,omitempty"`
	// 子单资金明细
	FundOpenInfo *FundOpenInfo `json:"fund_open_info,omitempty" xml:"fund_open_info,omitempty"`
	// 单价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
