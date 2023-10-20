package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleCarOrderQuery 二手车寄卖查询订单接口
// alibaba.idle.car.order.query
//
// 二手车寄卖查询订单接口
func AlibabaIdleCarOrderQuery(clt *core.SDKClient, req *idle.AlibabaIdleCarOrderQueryAPIRequest, session string) (*idle.AlibabaIdleCarOrderQueryAPIResponse, error) {
	var resp idle.AlibabaIdleCarOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
