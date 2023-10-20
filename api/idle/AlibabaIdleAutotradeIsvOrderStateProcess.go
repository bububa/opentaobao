package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleAutotradeIsvOrderStateProcess 闲鱼订单状态推进
// alibaba.idle.autotrade.isv.order.state.process
//
// 闲鱼订单状态推进
func AlibabaIdleAutotradeIsvOrderStateProcess(clt *core.SDKClient, req *idle.AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest, session string) (*idle.AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse, error) {
	var resp idle.AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
