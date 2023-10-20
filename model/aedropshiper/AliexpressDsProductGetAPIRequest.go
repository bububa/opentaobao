package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressdsproductgetAPIRequest 商品信息查询 API请求
// aliexpress.ds.product.get
//
// 商品信息查询
type AliexpressdsproductgetAPIRequest struct {
	model.Params
	// Country
	_shipToCountry string
	// Target currency
	_targetCurrency string
	// Target language
	_targetLanguage string
	// Item ID
	_productId int64
}

// NewAliexpressdsproductgetRequest 初始化AliexpressdsproductgetAPIRequest对象
func NewAliexpressdsproductgetRequest() *AliexpressdsproductgetAPIRequest {
	return &AliexpressdsproductgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressdsproductgetAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressdsproductgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressdsproductgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShipToCountry is ShipToCountry Setter
// Country
func (r *AliexpressdsproductgetAPIRequest) SetShipToCountry(_shipToCountry string) error {
	r._shipToCountry = _shipToCountry
	r.Set("ship_to_country", _shipToCountry)
	return nil
}

// GetShipToCountry ShipToCountry Getter
func (r AliexpressdsproductgetAPIRequest) GetShipToCountry() string {
	return r._shipToCountry
}

// SetTargetCurrency is TargetCurrency Setter
// Target currency
func (r *AliexpressdsproductgetAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressdsproductgetAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetTargetLanguage is TargetLanguage Setter
// Target language
func (r *AliexpressdsproductgetAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressdsproductgetAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetProductId is ProductId Setter
// Item ID
func (r *AliexpressdsproductgetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AliexpressdsproductgetAPIRequest) GetProductId() int64 {
	return r._productId
}
