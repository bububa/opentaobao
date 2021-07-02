package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductPriceGetAPIRequest 渠道价格查询接口 API请求
// tmall.supplychain.channel.product.price.get
//
// 渠道价格查询接口
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

// NewTmallSupplychainChannelProductPriceGetRequest 初始化TmallSupplychainChannelProductPriceGetAPIRequest对象
func NewTmallSupplychainChannelProductPriceGetRequest() *TmallSupplychainChannelProductPriceGetAPIRequest {
	return &TmallSupplychainChannelProductPriceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProductId Setter
// 产品ID
func (r *TmallSupplychainChannelProductPriceGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetProductId() int64 {
	return r._productId
}

// Set is PriceType Setter
// 区域价、指导价
func (r *TmallSupplychainChannelProductPriceGetAPIRequest) SetPriceType(_priceType int64) error {
	r._priceType = _priceType
	r.Set("price_type", _priceType)
	return nil
}

// Get PriceType Getter
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetPriceType() int64 {
	return r._priceType
}

// Set is SkuId Setter
// SKU ID
func (r *TmallSupplychainChannelProductPriceGetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// Set is ChannelCode Setter
// 渠道ID（台湾市场/供销平台/大农业…..）
func (r *TmallSupplychainChannelProductPriceGetAPIRequest) SetChannelCode(_channelCode int64) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// Get ChannelCode Getter
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetChannelCode() int64 {
	return r._channelCode
}
