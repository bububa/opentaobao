package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractSensorTrade 交易组件
// alibaba.interact.sensor.trade
//
// 交易流程
func AlibabaInteractSensorTrade(clt *core.SDKClient, req *mtopopen.AlibabaInteractSensorTradeAPIRequest, resp *mtopopen.AlibabaInteractSensorTradeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
