package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRegionPriceManageAPIRequest
编辑区域价格 API请求
taobao.region.price.manage

编辑区域价格 */
type TaobaoRegionPriceManageAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 无sku传0
	_skuId int64
	// 列表
	_regionalPriceDtos []RegionalPriceDto
	// true:全量, false:增量
	_isFull bool
}

// New
