package alsc

import (
	"sync"
)

// RefundInfo 结构体
type RefundInfo struct {
	// 退款商品
	RefundItemList []Item `json:"refund_item_list,omitempty" xml:"refund_item_list>item,omitempty"`
	// 退款流水号（支付宝单号、微信单号、三方单号）
	RefundPayNo string `json:"refund_pay_no,omitempty" xml:"refund_pay_no,omitempty"`
	// 退款状态：  INITIAL(&#34;INITIAL&#34;, &#34;退款初始状态&#34;),  PROCESSING(&#34;PROCESSING&#34;, &#34;处理中状态&#34;),  SUCCESS(&#34;SUCCESS&#34;, &#34;退款成功状态&#34;),  FAIL(&#34;FAIL&#34;, &#34;退款失败状态&#34;);
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 退款时间
	RefundTime string `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	// 退款类型 ：  整单-ALL  部分-PART
	RefundType string `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	// 退款单号，必填
	OutRefundNo string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
	// 退款金额
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}

var poolRefundInfo = sync.Pool{
	New: func() any {
		return new(RefundInfo)
	},
}

// GetRefundInfo() 从对象池中获取RefundInfo
func GetRefundInfo() *RefundInfo {
	return poolRefundInfo.Get().(*RefundInfo)
}

// ReleaseRefundInfo 释放RefundInfo
func ReleaseRefundInfo(v *RefundInfo) {
	v.RefundItemList = v.RefundItemList[:0]
	v.RefundPayNo = ""
	v.RefundStatus = ""
	v.RefundTime = ""
	v.RefundType = ""
	v.OutRefundNo = ""
	v.RefundAmount = 0
	poolRefundInfo.Put(v)
}
