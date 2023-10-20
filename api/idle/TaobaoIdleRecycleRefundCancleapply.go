package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Taobaoidlerecyclerefundcancleapply 闲鱼回收取消退款申请V2
// taobao.idle.recycle.refund.cancleapply
//
// 回收商的回收订单取消退款申请
func Taobaoidlerecyclerefundcancleapply(clt *core.SDKClient, req *idle.TaobaoidlerecyclerefundcancleapplyAPIRequest, session string) (*idle.TaobaoidlerecyclerefundcancleapplyAPIResponse, error) {
	var resp idle.TaobaoidlerecyclerefundcancleapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
