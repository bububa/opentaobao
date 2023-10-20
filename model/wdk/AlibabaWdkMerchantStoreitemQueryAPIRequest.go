package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantstoreitemqueryAPIRequest 门店商品信心查询 API请求
// alibaba.wdk.merchant.storeitem.query
//
// 门店商品信心查询
type AlibabawdkmerchantstoreitemqueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 商家编码
	_merchantCode string
	// 门店编码
	_storeId string
}

// NewAlibabawdkmerchantstoreitemqueryRequest 初始化AlibabawdkmerchantstoreitemqueryAPIRequest对象
func NewAlibabawdkmerchantstoreitemqueryRequest() *AlibabawdkmerchantstoreitemqueryAPIRequest {
	return &AlibabawdkmerchantstoreitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmerchantstoreitemqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.storeitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmerchantstoreitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmerchantstoreitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabawdkmerchantstoreitemqueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkmerchantstoreitemqueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetMerchantCode is MerchantCode Setter
// 商家编码
func (r *AlibabawdkmerchantstoreitemqueryAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabawdkmerchantstoreitemqueryAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetStoreId is StoreId Setter
// 门店编码
func (r *AlibabawdkmerchantstoreitemqueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdkmerchantstoreitemqueryAPIRequest) GetStoreId() string {
	return r._storeId
}
