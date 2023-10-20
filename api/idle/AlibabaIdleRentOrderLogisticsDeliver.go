package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentorderlogisticsdeliver 创建揽收物流
// alibaba.idle.rent.order.logistics.deliver
//
// 创建揽收物流
// 商家去物流公司创建物流订单
func Alibabaidlerentorderlogisticsdeliver(clt *core.SDKClient, req *idle.AlibabaidlerentorderlogisticsdeliverAPIRequest, session string) (*idle.AlibabaidlerentorderlogisticsdeliverAPIResponse, error) {
	var resp idle.AlibabaidlerentorderlogisticsdeliverAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
