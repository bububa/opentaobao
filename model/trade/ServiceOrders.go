package trade

// ServiceOrders 结构体
type ServiceOrders struct {
	// 服务详情的URL地址
	ServiceDetailUrl string `json:"service_detail_url,omitempty" xml:"service_detail_url,omitempty"`
	// 服务价格，精确到小数点后两位：单位:元
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 商品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 服务子订单总费用
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 最近退款的id
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 服务图片地址
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// 支持家装类物流的类型
	TmserSpuCode string `json:"tmser_spu_code,omitempty" xml:"tmser_spu_code,omitempty"`
	// 虚拟服务子订单订单号
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	// 服务所属的交易订单号。如果服务为一年包换，则item_oid这笔订单享受改服务的保护
	ItemOid int64 `json:"item_oid,omitempty" xml:"item_oid,omitempty"`
	// 服务数字id
	ServiceId int64 `json:"service_id,omitempty" xml:"service_id,omitempty"`
	// 购买数量，取值范围为大于0的整数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}
