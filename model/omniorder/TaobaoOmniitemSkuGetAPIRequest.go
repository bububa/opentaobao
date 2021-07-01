package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemSkuGetAPIRequest
获取全渠道门店商品sku API请求
taobao.omniitem.sku.get

通过skuId或者skuOutId查询全渠道门店商品sku信息 */
type TaobaoOmniitemSkuGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// skuId
	_skuId int64
	// sku商家编码
	_skuOuterId string
}

// New
