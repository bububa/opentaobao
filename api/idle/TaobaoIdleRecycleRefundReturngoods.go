package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

/* TaobaoIdleRecycleRefundReturngoods
闲鱼回收退货V2
taobao.idle.recycle.refund.returngoods

回收商买家退货，填写退货运单号 */
func TaobaoIdleRecycleRefundReturngoods(clt *core.SDKClient, req *idle.TaobaoIdleRecycleRefundReturngoodsAPIRequest, session string) (*idle.TaobaoIdleRecycleRefundReturngoodsAPIResponse, error) {
	var resp idle.TaobaoIdleRecycleRefundReturngoodsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
