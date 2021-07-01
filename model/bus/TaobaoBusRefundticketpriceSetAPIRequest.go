package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusRefundticketpriceSetAPIRequest
汽车票退款申请接口 API请求
taobao.bus.refundticketprice.set

汽车票代理商利用该接口申请退票 */
type TaobaoBusRefundticketpriceSetAPIRequest struct {
	model.Params
	// 退票申请入参
	_offlineRefundTicketRq *OfflineRefundTicketPriceRq
}

// New
