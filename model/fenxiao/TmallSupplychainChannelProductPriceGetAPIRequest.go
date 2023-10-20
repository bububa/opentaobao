package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductpricegetAPIRequest 渠道价格查询接口 API请求
// tmall.supplychain.channel.product.price.get
//
// 渠道价格查询接口
type TmallsupplychainchannelproductpricegetAPIRequest struct {
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

// NewTmallsupplychainchannelproductpricegetRequest 初始化TmallsupplychainchannelproductpricegetAPIRequest对象
func NewTmallsupplychainchannelproductpricegetRequest() *TmallsupplychainchannelproductpricegetAPIRequest {
	return &TmallsupplychainchannelproductpricegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallsupplychainchannelproductpricegetAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallsupplychainchannelproductpricegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallsupplychainchannelproductpricegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TmallsupplychainchannelproductpricegetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallsupplychainchannelproductpricegetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetPriceType is PriceType Setter
// 区域价、指导价
func (r *TmallsupplychainchannelproductpricegetAPIRequest) SetPriceType(_priceType int64) error {
	r._priceType = _priceType
	r.Set("price_type", _priceType)
	return nil
}

// GetPriceType PriceType Getter
func (r TmallsupplychainchannelproductpricegetAPIRequest) GetPriceType() int64 {
	return r._priceType
}

// SetSkuId is SkuId Setter
// SKU ID
func (r *TmallsupplychainchannelproductpricegetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TmallsupplychainchannelproductpricegetAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetChannelCode is ChannelCode Setter
// 渠道ID（台湾市场/供销平台/大农业…..）
func (r *TmallsupplychainchannelproductpricegetAPIRequest) SetChannelCode(_channelCode int64) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TmallsupplychainchannelproductpricegetAPIRequest) GetChannelCode() int64 {
	return r._channelCode
}
