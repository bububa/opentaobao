package fundplatform

// AlibabaFundplatformAccountChargeNotifyStruct 结构体
type AlibabaFundplatformAccountChargeNotifyStruct struct {
	// 回传充值时传入的outBizId
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	OwnSign string `json:"own_sign,omitempty" xml:"own_sign,omitempty"`
	// 在资金平台中账户ID
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 资金平台分配的业务类型
	SubBizType int64 `json:"sub_biz_type,omitempty" xml:"sub_biz_type,omitempty"`
	// 充值金额，单位分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 用户在淘宝中userId
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
