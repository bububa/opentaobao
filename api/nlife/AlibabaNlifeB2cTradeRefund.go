package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

/* AlibabaNlifeB2cTradeRefund
零售+请求退款
alibaba.nlife.b2c.trade.refund

零售+平台请求退款接口，在零售+平台不会有资金流变动，只是订单状态的更新 */
func AlibabaNlifeB2cTradeRefund(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradeRefundAPIRequest, session string) (*nlife.AlibabaNlifeB2cTradeRefundAPIResponse, error) {
	var resp nlife.AlibabaNlifeB2cTradeRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
