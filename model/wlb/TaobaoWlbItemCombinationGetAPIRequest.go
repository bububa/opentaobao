package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbItemCombinationGetAPIRequest
根据商品id查询商品组合关系 API请求
taobao.wlb.item.combination.get

根据商品id查询商品组合关系 */
type TaobaoWlbItemCombinationGetAPIRequest struct {
	model.Params
	// 要查询的组合商品id
	_itemId int64
}

// New
