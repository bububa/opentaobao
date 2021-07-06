package jipiao

// Itinerary 结构体
type Itinerary struct {
	// 收件人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 收件人手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 收件人备用手机号
	MobileBak string `json:"mobile_bak,omitempty" xml:"mobile_bak,omitempty"`
	// 收件人地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 行程单价格，单位：分
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 物流公司代码
	CompanyCode string `json:"company_code,omitempty" xml:"company_code,omitempty"`
	// 邮寄单号
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
	// 行程单号
	ItineraryNo string `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
	// 邮寄时间
	SendDate string `json:"send_date,omitempty" xml:"send_date,omitempty"`
	// 扩展字段
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 支付宝交易号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 行程单类型：6，快递
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 行程单订单的状态 0：未付款 1：已付款 2：转交易成功 3：已邮寄 4：已取消
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 淘宝主键id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
