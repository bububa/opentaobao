package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerecycleorderperform 回收订单履约V2
// alibaba.idle.recycle.order.perform
//
// 闲鱼回收业务中,外部服务商作为买家 需要驱动交易节点变化
func Alibabaidlerecycleorderperform(clt *core.SDKClient, req *idle.AlibabaidlerecycleorderperformAPIRequest, session string) (*idle.AlibabaidlerecycleorderperformAPIResponse, error) {
	var resp idle.AlibabaidlerecycleorderperformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
