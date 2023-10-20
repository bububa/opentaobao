package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbItemCombinationGet 根据商品id查询商品组合关系
// taobao.wlb.item.combination.get
//
// 根据商品id查询商品组合关系
func TaobaoWlbItemCombinationGet(clt *core.SDKClient, req *wlb.TaobaoWlbItemCombinationGetAPIRequest, session string) (*wlb.TaobaoWlbItemCombinationGetAPIResponse, error) {
	var resp wlb.TaobaoWlbItemCombinationGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
