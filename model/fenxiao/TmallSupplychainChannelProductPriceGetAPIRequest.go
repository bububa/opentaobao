package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallSupplychainChannelProductPriceGetAPIRequest
渠道价格查询接口 API请求
tmall.supplychain.channel.product.price.get

渠道价格查询接口 */
type TmallSupplychainChannelProductPriceGetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 区域价、指导价
	_priceType int64
	// SKU ID
	_skuId int64
	// 渠道ID（台湾市场/供销平台/大农业…..）
	_channelCode int64
}

// New
