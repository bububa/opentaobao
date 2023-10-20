package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// Alibabanlifeb2ctraderefund 零售+请求退款
// alibaba.nlife.b2c.trade.refund
//
// 零售+平台请求退款接口，在零售+平台不会有资金流变动，只是订单状态的更新
func Alibabanlifeb2ctraderefund(clt *core.SDKClient, req *nlife.Alibabanlifeb2ctraderefundAPIRequest, session string) (*nlife.Alibabanlifeb2ctraderefundAPIResponse, error) {
	var resp nlife.Alibabanlifeb2ctraderefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
