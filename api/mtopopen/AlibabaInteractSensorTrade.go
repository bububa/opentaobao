package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

/* AlibabaInteractSensorTrade
交易组件
alibaba.interact.sensor.trade

交易流程 */
func AlibabaInteractSensorTrade(clt *core.SDKClient, req *mtopopen.AlibabaInteractSensorTradeAPIRequest, session string) (*mtopopen.AlibabaInteractSensorTradeAPIResponse, error) {
	var resp mtopopen.AlibabaInteractSensorTradeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
