package jipiao

// TripFlightPassenger 结构体
type TripFlightPassenger struct {
	// 乘机人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 乘机人证件类型：0，身份证；1，护照；3，军人证；4，回乡证；5，台胞证；6，港澳台胞；10，警官证；11，士兵证件
	CertType int64 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// 乘机人证件号码
	CertNo string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 乘机人类型：0，成人；1，儿童；
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 乘机人生日：yyyy-mm-dd
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 常旅客卡号
	TripCardNo string `json:"trip_card_no,omitempty" xml:"trip_card_no,omitempty"`
	// PNR信息
	Pnr string `json:"pnr,omitempty" xml:"pnr,omitempty"`
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 扩展字段
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 政策id，淘宝系统政策唯一键
	PolicyId int64 `json:"policy_id,omitempty" xml:"policy_id,omitempty"`
	// 机票政策类型：0，默认；1，自定义
	PolicyType int64 `json:"policy_type,omitempty" xml:"policy_type,omitempty"`
	// 备注信息，政策中的备注信息
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 销售价格，单位：分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 航班机建费，单位：分
	Fee int64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// 航班燃油税，单位：分
	Tax int64 `json:"tax,omitempty" xml:"tax,omitempty"`
	// 航班舱位代码
	CabinCode string `json:"cabin_code,omitempty" xml:"cabin_code,omitempty"`
	// 舱位等级：0，头等舱(F)；1，商务舱(C)；2，经济舱(Y)
	CabinClass int64 `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 退改签
	Tuigaiqian string `json:"tuigaiqian,omitempty" xml:"tuigaiqian,omitempty"`
	// ei项
	Ei string `json:"ei,omitempty" xml:"ei,omitempty"`
	// 当前乘机人的保险分润金额，单位：分
	InsurePromotionPrice int64 `json:"insure_promotion_price,omitempty" xml:"insure_promotion_price,omitempty"`
	// 强制保险金额，单位：分
	ForceInsurePrice int64 `json:"force_insure_price,omitempty" xml:"force_insure_price,omitempty"`
}
