package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleautotradeisvorderstateprocess 闲鱼订单状态推进
// alibaba.idle.autotrade.isv.order.state.process
//
// 闲鱼订单状态推进
func Alibabaidleautotradeisvorderstateprocess(clt *core.SDKClient, req *idle.AlibabaidleautotradeisvorderstateprocessAPIRequest, session string) (*idle.AlibabaidleautotradeisvorderstateprocessAPIResponse, error) {
	var resp idle.AlibabaidleautotradeisvorderstateprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
