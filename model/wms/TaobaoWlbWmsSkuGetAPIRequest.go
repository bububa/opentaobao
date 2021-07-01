package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsSkuGetAPIRequest
商品信息查询 API请求
taobao.wlb.wms.sku.get

商品信息查询 */
type TaobaoWlbWmsSkuGetAPIRequest struct {
	model.Params
	// 菜鸟商品ID,与itemcode必须有一个值不为空
	_itemCode string
	// 商家商品编码,与itemid必须有一个值不为空
	_itemId string
	// 货主ID
	_ownerUserId string
}

// New
