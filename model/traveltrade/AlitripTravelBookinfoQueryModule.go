package traveltrade

// AlitripTravelBookinfoQueryModule 结构体
type AlitripTravelBookinfoQueryModule struct {
	// 出行人信息
	TravellerInfos []TravellerInfos `json:"traveller_infos,omitempty" xml:"traveller_infos>traveller_infos,omitempty"`
	// 时段结束时间
	EndPeriod string `json:"end_period,omitempty" xml:"end_period,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 付款时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 时段开始时间
	StartPeriod string `json:"start_period,omitempty" xml:"start_period,omitempty"`
	// 修改时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 预约单项备注
	BookCellRemark string `json:"book_cell_remark,omitempty" xml:"book_cell_remark,omitempty"`
	// 买家申请时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 预约出发日期
	BookedTravelTime string `json:"booked_travel_time,omitempty" xml:"booked_travel_time,omitempty"`
	// 下单时间
	OrderCreateTime string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// 预约记录对应的保险保单数据记录ID
	TravelInsAppyId string `json:"travel_ins_appy_id,omitempty" xml:"travel_ins_appy_id,omitempty"`
	// 商品标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 卖家审批时间
	ProcessTime string `json:"process_time,omitempty" xml:"process_time,omitempty"`
	// 预约单项名称（子套餐名称）
	BookCellName string `json:"book_cell_name,omitempty" xml:"book_cell_name,omitempty"`
	// 预约返程时间
	BookedTravelReturnTime string `json:"booked_travel_return_time,omitempty" xml:"booked_travel_return_time,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 下单时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 扩展属性（json结构，其中outerIdSKU=商家编码）
	Features string `json:"features,omitempty" xml:"features,omitempty"`
	// TC子订单号
	SubTcOrderId int64 `json:"sub_tc_order_id,omitempty" xml:"sub_tc_order_id,omitempty"`
	// 预约次数
	BookNum int64 `json:"book_num,omitempty" xml:"book_num,omitempty"`
	// 预约订单ID
	BookOrderId int64 `json:"book_order_id,omitempty" xml:"book_order_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 业务类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 预约ID(主键)
	BookInfoId int64 `json:"book_info_id,omitempty" xml:"book_info_id,omitempty"`
	// TC主订单号
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// 预约单项ID（子套餐ID）
	BookCellId int64 `json:"book_cell_id,omitempty" xml:"book_cell_id,omitempty"`
	// 申请状态
	BookStatus int64 `json:"book_status,omitempty" xml:"book_status,omitempty"`
	// 是否仅更新bookId字段
	OnlyUpateBookId bool `json:"only_upate_book_id,omitempty" xml:"only_upate_book_id,omitempty"`
}
