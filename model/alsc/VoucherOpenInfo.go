package alsc

// VoucherOpenInfo 结构体
type VoucherOpenInfo struct {
	// 面额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 失效时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 满足金额阀值，单位：分
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 券开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 券名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 券实例id
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 券模板id
	VoucherTemplateId string `json:"voucher_template_id,omitempty" xml:"voucher_template_id,omitempty"`
	// 优惠券类型，CASH：代金券，DISCOUNT：折扣券，GIFT：礼品券
	VoucherType string `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	// 折扣率
	DiscountRate string `json:"discount_rate,omitempty" xml:"discount_rate,omitempty"`
	// 商品id列表
	ItemIdList []string `json:"item_id_list,omitempty" xml:"item_id_list>string,omitempty"`
	// 门店id列表
	ShopIdList []string `json:"shop_id_list,omitempty" xml:"shop_id_list>string,omitempty"`
	// 每人限领，-1代表不限
	UserLimit int64 `json:"user_limit,omitempty" xml:"user_limit,omitempty"`
	// 适用商品范围 值：ALL，PART_AVAILABLE，PART_UNAVAILABLE * 说明：全部商品可用，部分商品可用，部分商品不可用
	ItemCoverage string `json:"item_coverage,omitempty" xml:"item_coverage,omitempty"`
	// 适用商品范围 值：ALL，PART_AVAILABLE * 说明：全部门店可用，部分门店可用
	ShopCoverage string `json:"shop_coverage,omitempty" xml:"shop_coverage,omitempty"`
	// {\"type\":0,\"settings\":[{\"dayStartTime\":\"00:00\",\"dayEndTime\":\"23:59\",\"week\":[]}]} type:0不限制，1限制 dayStartTime:开始时间 dayEndTime:结束时间 week:1,2,3,4,5,6,7
	AvailableTime string `json:"available_time,omitempty" xml:"available_time,omitempty"`
	// 商品信息
	ItemInfoList []VoucherItemInfo `json:"item_info_list,omitempty" xml:"item_info_list>voucher_item_info,omitempty"`
	// 门店信息
	ShopInfoList []VoucherShopInfo `json:"shop_info_list,omitempty" xml:"shop_info_list>voucher_shop_info,omitempty"`
	// 扩展字段,包含了礼品券的信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// DINE_IN堂食，TAKE_OUT外卖
	UseCondition string `json:"use_condition,omitempty" xml:"use_condition,omitempty"`
}
