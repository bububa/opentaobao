package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentorderpackage 确认揽收商品
// alibaba.idle.rent.order.package
//
// 确认揽收
func Alibabaidlerentorderpackage(clt *core.SDKClient, req *idle.AlibabaidlerentorderpackageAPIRequest, session string) (*idle.AlibabaidlerentorderpackageAPIResponse, error) {
	var resp idle.AlibabaidlerentorderpackageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
