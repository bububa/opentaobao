package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentorderquery 查询订单
// alibaba.idle.rent.order.query
//
// 查询订单信息
func Alibabaidlerentorderquery(clt *core.SDKClient, req *idle.AlibabaidlerentorderqueryAPIRequest, session string) (*idle.AlibabaidlerentorderqueryAPIResponse, error) {
	var resp idle.AlibabaidlerentorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
