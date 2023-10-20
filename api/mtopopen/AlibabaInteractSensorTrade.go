package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Alibabainteractsensortrade 交易组件
// alibaba.interact.sensor.trade
//
// 交易流程
func Alibabainteractsensortrade(clt *core.SDKClient, req *mtopopen.AlibabainteractsensortradeAPIRequest, session string) (*mtopopen.AlibabainteractsensortradeAPIResponse, error) {
	var resp mtopopen.AlibabainteractsensortradeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
