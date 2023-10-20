package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// Alibabanlifeb2ctradecancel 零售+平台取消订单
// alibaba.nlife.b2c.trade.cancel
//
// 零售+平台取消订单接口
func Alibabanlifeb2ctradecancel(clt *core.SDKClient, req *nlife.Alibabanlifeb2ctradecancelAPIRequest, session string) (*nlife.Alibabanlifeb2ctradecancelAPIResponse, error) {
	var resp nlife.Alibabanlifeb2ctradecancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
