package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbItemGetAPIRequest
根据商品ID获取商品信息 API请求
taobao.wlb.item.get

根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。 */
type TaobaoWlbItemGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// New
