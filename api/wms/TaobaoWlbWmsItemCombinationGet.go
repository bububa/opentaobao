package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsItemCombinationGet 查询组合商品的组合关系
// taobao.wlb.wms.item.combination.get
//
// 查询组合商品的组合关系
func TaobaoWlbWmsItemCombinationGet(clt *core.SDKClient, req *wms.TaobaoWlbWmsItemCombinationGetAPIRequest, resp *wms.TaobaoWlbWmsItemCombinationGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
