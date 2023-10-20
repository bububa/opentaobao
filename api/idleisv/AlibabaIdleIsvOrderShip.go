package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvordership 闲鱼订单服务商物流发货
// alibaba.idle.isv.order.ship
//
// 闲鱼开放平台服务商订单发货接口
func Alibabaidleisvordership(clt *core.SDKClient, req *idleisv.AlibabaidleisvordershipAPIRequest, session string) (*idleisv.AlibabaidleisvordershipAPIResponse, error) {
	var resp idleisv.AlibabaidleisvordershipAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
