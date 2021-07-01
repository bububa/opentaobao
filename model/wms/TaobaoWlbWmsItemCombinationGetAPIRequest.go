package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsItemCombinationGetAPIRequest
查询组合商品的组合关系 API请求
taobao.wlb.wms.item.combination.get

查询组合商品的组合关系 */
type TaobaoWlbWmsItemCombinationGetAPIRequest struct {
	model.Params
	// 货品Id
	_itemid int64
}

// New
