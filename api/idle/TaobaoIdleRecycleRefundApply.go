package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// TaobaoIdleRecycleRefundApply 闲鱼回收交易退款申请V2
// taobao.idle.recycle.refund.apply
//
// 回收商买家申请退款
func TaobaoIdleRecycleRefundApply(clt *core.SDKClient, req *idle.TaobaoIdleRecycleRefundApplyAPIRequest, resp *idle.TaobaoIdleRecycleRefundApplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
