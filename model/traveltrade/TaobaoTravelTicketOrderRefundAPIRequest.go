package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTravelTicketOrderRefundAPIRequest
飞猪门票退票结果通知 API请求
taobao.travel.ticket.order.refund

门票系统商通过TOP接口通知飞猪门票是否退票成功，以及退票数量。 */
type TaobaoTravelTicketOrderRefundAPIRequest struct {
	model.Params
	// 下单时订单ID
	_orderId int64
	// 退票结果；1: 退票成功；2: 退票失败
	_refundStatus int64
	// 退票失败理由
	_refundFailureReason string
}

// New
