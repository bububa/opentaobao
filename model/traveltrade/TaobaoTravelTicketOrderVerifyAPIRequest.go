package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTravelTicketOrderVerifyAPIRequest
飞猪门票核销通知 API请求
taobao.travel.ticket.order.verify

系统商通过TOP接口调用通知飞猪门票核销情况 */
type TaobaoTravelTicketOrderVerifyAPIRequest struct {
	model.Params
	// 下单订单ID
	_orderId int64
	// 外部订单ID
	_outOrderId string
	// 使用凭证信息
	_voucherInfos []VoucherInfoDto
	// 供应商核销回调类型：0表示使用本次核销数量（常规），1表示使用总核销数量（已使用+本次）
	_writeOffType int64
}

// New
