package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusTvmrefundorderSetAPIRequest
线下自助机逆向退款接口 API请求
taobao.bus.tvmrefundorder.set

汽车票线下自助机 逆向退票接口；用于已出票完成后，再发起退款（注意这是售后退款，如出票异常但是告诉我们出票成功，后续给客户退款，需要调用这个接口，一般开放给财务。出票过程中的失败，请直接调用出票接口并且传递false标志，我们会自动退款。） */
type TaobaoBusTvmrefundorderSetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
	// 分账退款明细
	_refundAccountInDetails []RefundAccountInDetail
	// 退款金额（单位分） 票金额
	_refundAmount int64
	// 退款原因
	_refundReason string
	// 批次号必须唯一，同一批次号只能退款一次 （多账号分润的该值 填写refundAccountInDetails中批次号的任意一个即可
	_refundBatchNo string
	// 保险退款详情
	_insuranceRefundDetails []InsuranceRefundDetail
}

// New
