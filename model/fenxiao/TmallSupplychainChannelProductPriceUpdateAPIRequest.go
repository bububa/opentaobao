package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallSupplychainChannelProductPriceUpdateAPIRequest
渠道价格更新接口 API请求
tmall.supplychain.channel.product.price.update

更新渠道产品价格 */
type TmallSupplychainChannelProductPriceUpdateAPIRequest struct {
	model.Params
	// 币种，非必填，仅支持当商品记为外币价格时使用
	_currencyType string
	// 产品数字ID
	_productId int64
	// 1.指导价(默认) 2.区域价
	_priceType int64
	// 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
	_skuPrice string
	// 产品价格，必填
	_productPrice string
	// SKU ID
	_skuId int64
	// 渠道编码
	_channelCode int64
}

// New
