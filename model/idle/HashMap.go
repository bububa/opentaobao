package idle

import (
	"sync"
)

// HashMap 结构体
type HashMap struct {
	// 上面区间订单
	ShipMailNo string `json:"ship_mail_no,omitempty" xml:"ship_mail_no,omitempty"`
	// 估价价格 单位分
	QaAmount string `json:"qa_amount,omitempty" xml:"qa_amount,omitempty"`
	// 参考价格起始价 单位分
	ReferenceStartPrice string `json:"reference_start_price,omitempty" xml:"reference_start_price,omitempty"`
	// 参考价格终止价 单位分
	ReferenceEndPrice string `json:"reference_end_price,omitempty" xml:"reference_end_price,omitempty"`
	// 验货报告id
	Report string `json:"report,omitempty" xml:"report,omitempty"`
	// 服务商最终成交价
	ConfirmFee string `json:"confirm_fee,omitempty" xml:"confirm_fee,omitempty"`
	// 服务商给买家发货订单号
	Ac2buyerMailNo string `json:"ac2buyer_mail_no,omitempty" xml:"ac2buyer_mail_no,omitempty"`
	// 服务商关单原因
	BuyerCloseReason string `json:"buyer_close_reason,omitempty" xml:"buyer_close_reason,omitempty"`
	// 服务商退回单号
	RefundMailNo string `json:"refund_mail_no,omitempty" xml:"refund_mail_no,omitempty"`
	// 质检工程师姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 质检工程师电话
	ContactMobile string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
}

var poolHashMap = sync.Pool{
	New: func() any {
		return new(HashMap)
	},
}

// GetHashMap() 从对象池中获取HashMap
func GetHashMap() *HashMap {
	return poolHashMap.Get().(*HashMap)
}

// ReleaseHashMap 释放HashMap
func ReleaseHashMap(v *HashMap) {
	v.ShipMailNo = ""
	v.QaAmount = ""
	v.ReferenceStartPrice = ""
	v.ReferenceEndPrice = ""
	v.Report = ""
	v.ConfirmFee = ""
	v.Ac2buyerMailNo = ""
	v.BuyerCloseReason = ""
	v.RefundMailNo = ""
	v.ContactName = ""
	v.ContactMobile = ""
	poolHashMap.Put(v)
}
