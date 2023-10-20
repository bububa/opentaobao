package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Taobaoidlerecyclerefundapply 闲鱼回收交易退款申请V2
// taobao.idle.recycle.refund.apply
//
// 回收商买家申请退款
func Taobaoidlerecyclerefundapply(clt *core.SDKClient, req *idle.TaobaoidlerecyclerefundapplyAPIRequest, session string) (*idle.TaobaoidlerecyclerefundapplyAPIResponse, error) {
	var resp idle.TaobaoidlerecyclerefundapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
