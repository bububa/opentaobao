package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbItemMapGetAPIRequest
根据物流宝商品ID查询商品映射关系 API请求
taobao.wlb.item.map.get

根据物流宝商品ID查询商品映射关系 */
type TaobaoWlbItemMapGetAPIRequest struct {
	model.Params
	// 要查询映射关系的物流宝商品id
	_itemId int64
}

// New
