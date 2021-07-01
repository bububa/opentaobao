package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRegionPriceQueryAPIRequest
区域价格查询 API请求
taobao.region.price.query

区域价格查询 */
type TaobaoRegionPriceQueryAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 无sku可传0
	_skuId int64
	// 不传则返回所有设置的区域价格
	_regionalPriceDtos []RegionalPriceDto
}

// New
