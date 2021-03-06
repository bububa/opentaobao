package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRecycleOrderPerform 回收订单履约V2
// alibaba.idle.recycle.order.perform
//
// 闲鱼回收业务中,外部服务商作为买家 需要驱动交易节点变化
func AlibabaIdleRecycleOrderPerform(clt *core.SDKClient, req *idle.AlibabaIdleRecycleOrderPerformAPIRequest, session string) (*idle.AlibabaIdleRecycleOrderPerformAPIResponse, error) {
	var resp idle.AlibabaIdleRecycleOrderPerformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
