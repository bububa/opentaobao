package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

/* TaobaoWlbItemGet
根据商品ID获取商品信息
taobao.wlb.item.get

根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。 */
func TaobaoWlbItemGet(clt *core.SDKClient, req *wlb.TaobaoWlbItemGetAPIRequest, session string) (*wlb.TaobaoWlbItemGetAPIResponse, error) {
	var resp wlb.TaobaoWlbItemGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
