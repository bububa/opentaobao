package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Taobaoidlerecyclerefundreturngoods 闲鱼回收退货V2
// taobao.idle.recycle.refund.returngoods
//
// 回收商买家退货，填写退货运单号
func Taobaoidlerecyclerefundreturngoods(clt *core.SDKClient, req *idle.TaobaoidlerecyclerefundreturngoodsAPIRequest, session string) (*idle.TaobaoidlerecyclerefundreturngoodsAPIResponse, error) {
	var resp idle.TaobaoidlerecyclerefundreturngoodsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
