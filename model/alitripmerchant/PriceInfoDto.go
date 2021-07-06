package alitripmerchant

// PriceInfoDto 结构体
type PriceInfoDto struct {
	// 早餐信息 -1 含早、0无早、1单早、2、双早
	Breakfasts []string `json:"breakfasts,omitempty" xml:"breakfasts>string,omitempty"`
	// 详情页会员权益标签
	HotelDinamicLabels []HotelDinamicLabelDto `json:"hotel_dinamic_labels,omitempty" xml:"hotel_dinamic_labels>hotel_dinamic_label_dto,omitempty"`
	// 税费 单位元
	TaxAndFeeToString string `json:"tax_and_fee_to_string,omitempty" xml:"tax_and_fee_to_string,omitempty"`
	// 含税价 均价 单位元
	DinamicOriginalPriceToString string `json:"dinamic_original_price_to_string,omitempty" xml:"dinamic_original_price_to_string,omitempty"`
	// 退订规则
	RefundRules string `json:"refund_rules,omitempty" xml:"refund_rules,omitempty"`
	// 可够取消险
	CancelInsuranceIcon string `json:"cancel_insurance_icon,omitempty" xml:"cancel_insurance_icon,omitempty"`
	// 净价均价 不含税 单位元
	RatePriceToString string `json:"rate_price_to_string,omitempty" xml:"rate_price_to_string,omitempty"`
	// 会员等级
	MemberLevel string `json:"member_level,omitempty" xml:"member_level,omitempty"`
	// 币种CNY 人民币
	PriceType string `json:"price_type,omitempty" xml:"price_type,omitempty"`
	// 早餐
	BreakfastDesc string `json:"breakfast_desc,omitempty" xml:"breakfast_desc,omitempty"`
	// 取消政策模型同小搜一致即可，入住日往前XX小时
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// 不含税总价 保留到元
	RoomTotalPriceToString string `json:"room_total_price_to_string,omitempty" xml:"room_total_price_to_string,omitempty"`
	// 取消险文案
	CancelInsuranceDesc string `json:"cancel_insurance_desc,omitempty" xml:"cancel_insurance_desc,omitempty"`
	// 商品定价规则标题
	RpName string `json:"rp_name,omitempty" xml:"rp_name,omitempty"`
	// 库存数量
	StockNumberDes string `json:"stock_number_des,omitempty" xml:"stock_number_des,omitempty"`
	// 首住标签
	FirstLive string `json:"first_live,omitempty" xml:"first_live,omitempty"`
	// 副会员商品等级
	MemberLevelV2 string `json:"member_level_v2,omitempty" xml:"member_level_v2,omitempty"`
	// 是否担保商品
	IsGuarantee int64 `json:"is_guarantee,omitempty" xml:"is_guarantee,omitempty"`
	// 库存剩余数量
	StockNumber int64 `json:"stock_number,omitempty" xml:"stock_number,omitempty"`
	// 商品id
	Gid int64 `json:"gid,omitempty" xml:"gid,omitempty"`
	// 取消政策枚举 01-不可取消 02-免费取消 03-阶梯取消 04-限时免费取消
	CancelType int64 `json:"cancel_type,omitempty" xml:"cancel_type,omitempty"`
	// 含税价 均价 单位分
	DinamicOriginalPrice int64 `json:"dinamic_original_price,omitempty" xml:"dinamic_original_price,omitempty"`
	// 是否已售空
	IsSellOut int64 `json:"is_sell_out,omitempty" xml:"is_sell_out,omitempty"`
	// 不含税总价 单位分
	RoomTotalPrice int64 `json:"room_total_price,omitempty" xml:"room_total_price,omitempty"`
	// Rate Plan ID商品定价规则
	RpId int64 `json:"rp_id,omitempty" xml:"rp_id,omitempty"`
	// 商品最小可售卖粒度
	RateId int64 `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	// 1:全额支付;5:前台面付;
	PaymentType int64 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// 净价均价 不含税
	RatePrice int64 `json:"rate_price,omitempty" xml:"rate_price,omitempty"`
	// 税费 单位分
	TaxAndFee int64 `json:"tax_and_fee,omitempty" xml:"tax_and_fee,omitempty"`
	// 卖家房型id
	OuterRoomId int64 `json:"outer_room_id,omitempty" xml:"outer_room_id,omitempty"`
	// 表示当前价格允许几人入住
	AllowPersonNumber int64 `json:"allow_person_number,omitempty" xml:"allow_person_number,omitempty"`
	// 是否售空
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否是会员房
	MemberRoom bool `json:"member_room,omitempty" xml:"member_room,omitempty"`
	// 是否是副会员房
	MemberRoomV2 bool `json:"member_room_v2,omitempty" xml:"member_room_v2,omitempty"`
}
