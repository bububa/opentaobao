package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerecycleordershow 闲鱼回收订单查询V1.1
// alibaba.idle.recycle.order.show
//
// 查询回收订单
func Alibabaidlerecycleordershow(clt *core.SDKClient, req *idle.AlibabaidlerecycleordershowAPIRequest, session string) (*idle.AlibabaidlerecycleordershowAPIResponse, error) {
	var resp idle.AlibabaidlerecycleordershowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
