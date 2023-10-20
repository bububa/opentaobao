package tbitem

import (
	"sync"
)

// LocalityLife 结构体
type LocalityLife struct {
	// 表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄
	ChooseLogis string `json:"choose_logis,omitempty" xml:"choose_logis,omitempty"`
	// 电子凭证业务属性
	Eticket string `json:"eticket,omitempty" xml:"eticket,omitempty"`
	// 电子交易凭证有效期，有三种格式：如果有效期为起止日期类型，此值为2012-08-06,2012-08-16 如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15
	Expirydate string `json:"expirydate,omitempty" xml:"expirydate,omitempty"`
	// 格式为 码商id:nick
	Merchant string `json:"merchant,omitempty" xml:"merchant,omitempty"`
	// 网点ID,在参数empty_fields里设置locality_life.network_id可删除网点ID
	NetworkId string `json:"network_id,omitempty" xml:"network_id,omitempty"`
	// 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担)
	Refundmafee string `json:"refundmafee,omitempty" xml:"refundmafee,omitempty"`
	// 核销打款:1代表核销打款,0代表非核销打款;在参数empty_fields里设置locality_life.verification可删除核销打款
	Verification string `json:"verification,omitempty" xml:"verification,omitempty"`
	// 电子凭证售中自动退款比例
	OnsaleAutoRefundRatio int64 `json:"onsale_auto_refund_ratio,omitempty" xml:"onsale_auto_refund_ratio,omitempty"`
	// 退款比例，百分比%前的数字，1-100的正整数值；在参数empty_fields里设置locality_life.refund_ratio可删除退款比例
	RefundRatio int64 `json:"refund_ratio,omitempty" xml:"refund_ratio,omitempty"`
}

var poolLocalityLife = sync.Pool{
	New: func() any {
		return new(LocalityLife)
	},
}

// GetLocalityLife() 从对象池中获取LocalityLife
func GetLocalityLife() *LocalityLife {
	return poolLocalityLife.Get().(*LocalityLife)
}

// ReleaseLocalityLife 释放LocalityLife
func ReleaseLocalityLife(v *LocalityLife) {
	v.ChooseLogis = ""
	v.Eticket = ""
	v.Expirydate = ""
	v.Merchant = ""
	v.NetworkId = ""
	v.Refundmafee = ""
	v.Verification = ""
	v.OnsaleAutoRefundRatio = 0
	v.RefundRatio = 0
	poolLocalityLife.Put(v)
}
