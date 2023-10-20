package film

// FcodeMerchantSendCodeRq 结构体
type FcodeMerchantSendCodeRq struct {
	// 外部业务用户id
	OutUid string `json:"out_uid,omitempty" xml:"out_uid,omitempty"`
	// 发券码商品mixId
	MixId string `json:"mix_id,omitempty" xml:"mix_id,omitempty"`
	// 外部商户发码的外部业务号
	ExtOrderId string `json:"ext_order_id,omitempty" xml:"ext_order_id,omitempty"`
	// 外部下单时间
	OrderTime string `json:"order_time,omitempty" xml:"order_time,omitempty"`
	// 用户ID类型 TAOBAO 或者 TAOBAO_NAME
	UserIdType string `json:"user_id_type,omitempty" xml:"user_id_type,omitempty"`
	// 外部商户标示
	PartnerCode string `json:"partner_code,omitempty" xml:"partner_code,omitempty"`
	// 外面用户昵称
	OutUserName string `json:"out_user_name,omitempty" xml:"out_user_name,omitempty"`
	// 请求属性字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 淘宝用户ID列表，用|分割
	UserIdList string `json:"user_id_list,omitempty" xml:"user_id_list,omitempty"`
	// 每个用户发码的数量
	EachNum int64 `json:"each_num,omitempty" xml:"each_num,omitempty"`
	// 发码总数
	Number int64 `json:"number,omitempty" xml:"number,omitempty"`
}
