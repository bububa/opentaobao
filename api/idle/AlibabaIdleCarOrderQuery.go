package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlecarorderquery 二手车寄卖查询订单接口
// alibaba.idle.car.order.query
//
// 二手车寄卖查询订单接口
func Alibabaidlecarorderquery(clt *core.SDKClient, req *idle.AlibabaidlecarorderqueryAPIRequest, session string) (*idle.AlibabaidlecarorderqueryAPIResponse, error) {
	var resp idle.AlibabaidlecarorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
