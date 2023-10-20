package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// Alibabanlifeb2ctradepay 零售+平台支付订单
// alibaba.nlife.b2c.trade.pay
//
// 零售+平台支付接口，外部商户调用此接口告知零售+支付结果，保持订单状态同步
func Alibabanlifeb2ctradepay(clt *core.SDKClient, req *nlife.Alibabanlifeb2ctradepayAPIRequest, session string) (*nlife.Alibabanlifeb2ctradepayAPIResponse, error) {
	var resp nlife.Alibabanlifeb2ctradepayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
