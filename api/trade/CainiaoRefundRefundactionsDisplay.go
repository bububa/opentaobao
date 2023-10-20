package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// CainiaoRefundRefundactionsDisplay 退货退款操作的展示信息(展现给买家)
// cainiao.refund.refundactions.display
//
// 退货退款操作的展示信息(展现给买家)
func CainiaoRefundRefundactionsDisplay(clt *core.SDKClient, req *trade.CainiaoRefundRefundactionsDisplayAPIRequest, resp *trade.CainiaoRefundRefundactionsDisplayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
