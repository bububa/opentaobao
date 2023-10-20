package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cTradeRefund 零售+请求退款
// alibaba.nlife.b2c.trade.refund
//
// 零售+平台请求退款接口，在零售+平台不会有资金流变动，只是订单状态的更新
func AlibabaNlifeB2cTradeRefund(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradeRefundAPIRequest, resp *nlife.AlibabaNlifeB2cTradeRefundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
