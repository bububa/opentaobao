package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Alibabainteractsensortradebuy 手淘下单能力开放
// alibaba.interact.sensor.trade.buy
//
// 交易流程鉴权
func Alibabainteractsensortradebuy(clt *core.SDKClient, req *mtopopen.AlibabainteractsensortradebuyAPIRequest, session string) (*mtopopen.AlibabainteractsensortradebuyAPIResponse, error) {
	var resp mtopopen.AlibabainteractsensortradebuyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
