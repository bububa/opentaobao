package ieagency

import (
	"sync"
)

// RefundOrderSimpleVo 结构体
type RefundOrderSimpleVo struct {
	// 新老模型（V1:老模型，V2：新模型）
	ModelVersion string `json:"model_version,omitempty" xml:"model_version,omitempty"`
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 10: &#34;已提交待处理&#34;,WAI20:&#34;待人工处理&#34;
	RefundBizStatus int64 `json:"refund_biz_status,omitempty" xml:"refund_biz_status,omitempty"`
	// 退票单ID
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	// 申请单状态(WAIT(1,&#34;待处理&#34;), AGREED(2, &#34;已同意&#34;),REFUSE(3, &#34;已拒绝&#34;),PROCESS(6, &#34;已受理&#34;), SUCCESS(7, &#34;已退款&#34;))
	RefundOrderStatus int64 `json:"refund_order_status,omitempty" xml:"refund_order_status,omitempty"`
	// 申请单支付状态(   INIT(1, &#34;初始化&#34;),     REFUND_FAIL(2, &#34;退款失败&#34;),     REFUND_SUCCESS(3, &#34;退款成功&#34;))
	RefundPayStatus int64 `json:"refund_pay_status,omitempty" xml:"refund_pay_status,omitempty"`
}

var poolRefundOrderSimpleVo = sync.Pool{
	New: func() any {
		return new(RefundOrderSimpleVo)
	},
}

// GetRefundOrderSimpleVo() 从对象池中获取RefundOrderSimpleVo
func GetRefundOrderSimpleVo() *RefundOrderSimpleVo {
	return poolRefundOrderSimpleVo.Get().(*RefundOrderSimpleVo)
}

// ReleaseRefundOrderSimpleVo 释放RefundOrderSimpleVo
func ReleaseRefundOrderSimpleVo(v *RefundOrderSimpleVo) {
	v.ModelVersion = ""
	v.AgentId = 0
	v.OrderId = 0
	v.RefundBizStatus = 0
	v.RefundOrderId = 0
	v.RefundOrderStatus = 0
	v.RefundPayStatus = 0
	poolRefundOrderSimpleVo.Put(v)
}
