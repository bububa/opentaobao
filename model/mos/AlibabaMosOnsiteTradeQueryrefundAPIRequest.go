package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOnsiteTradeQueryrefundAPIRequest
退款查询 API请求
alibaba.mos.onsite.trade.queryrefund

商户可使用该接口查询退款请求是否执行成功。 */
type AlibabaMosOnsiteTradeQueryrefundAPIRequest struct {
	model.Params
	// 退款外部流水号
	_outRequestNo string
	// 订单号。可能为外部订单号，也可能为喵街订单号
	_orderNo string
}

// New
