package fenxiao

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallSupplychainChannelProductPriceGetAPIRequest) Reset() {
	r._productId = 0
	r._priceType = 0
	r._skuId = 0
	r._channelCode = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TmallSupplychainChannelProductPriceGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetPriceType is PriceType Setter
// 区域价、指导价
func (r *TmallSupplychainChannelProductPriceGetAPIRequest) SetPriceType(_priceType int64) error {
	r._priceType = _priceType
	r.Set("price_type", _priceType)
	return nil
}

// GetPriceType PriceType Getter
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetPriceType() int64 {
	return r._priceType
}

// SetSkuId is SkuId Setter
// SKU ID
func (r *TmallSupplychainChannelProductPriceGetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetChannelCode is ChannelCode Setter
// 渠道ID（台湾市场/供销平台/大农业…..）
func (r *TmallSupplychainChannelProductPriceGetAPIRequest) SetChannelCode(_channelCode int64) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TmallSupplychainChannelProductPriceGetAPIRequest) GetChannelCode() int64 {
	return r._channelCode
}

var poolTmallSupplychainChannelProductPriceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallSupplychainChannelProductPriceGetRequest()
	},
}

// GetTmallSupplychainChannelProductPriceGetRequest 从 sync.Pool 获取 TmallSupplychainChannelProductPriceGetAPIRequest
func GetTmallSupplychainChannelProductPriceGetAPIRequest() *TmallSupplychainChannelProductPriceGetAPIRequest {
	return poolTmallSupplychainChannelProductPriceGetAPIRequest.Get().(*TmallSupplychainChannelProductPriceGetAPIRequest)
}

// ReleaseTmallSupplychainChannelProductPriceGetAPIRequest 将 TmallSupplychainChannelProductPriceGetAPIRequest 放入 sync.Pool
func ReleaseTmallSupplychainChannelProductPriceGetAPIRequest(v *TmallSupplychainChannelProductPriceGetAPIRequest) {
	v.Reset()
	poolTmallSupplychainChannelProductPriceGetAPIRequest.Put(v)
}
