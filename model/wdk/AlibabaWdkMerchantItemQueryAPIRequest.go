package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantitemqueryAPIRequest 商家商品查询 API请求
// alibaba.wdk.merchant.item.query
//
// 商家商品查询
type AlibabawdkmerchantitemqueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 商家编码
	_merchantCode string
}

// NewAlibabawdkmerchantitemqueryRequest 初始化AlibabawdkmerchantitemqueryAPIRequest对象
func NewAlibabawdkmerchantitemqueryRequest() *AlibabawdkmerchantitemqueryAPIRequest {
	return &AlibabawdkmerchantitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmerchantitemqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmerchantitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmerchantitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabawdkmerchantitemqueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkmerchantitemqueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetMerchantCode is MerchantCode Setter
// 商家编码
func (r *AlibabawdkmerchantitemqueryAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabawdkmerchantitemqueryAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}
