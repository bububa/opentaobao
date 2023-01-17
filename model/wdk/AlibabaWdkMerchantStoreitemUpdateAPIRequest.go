package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantStoreitemUpdateAPIRequest 修改门店商品 API请求
// alibaba.wdk.merchant.storeitem.update
//
// 修改门店商品
type AlibabaWdkMerchantStoreitemUpdateAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 商家编码
	_merchantCode string
	// 门店编码
	_storeId string
	// 修改参数的json
	_params string
}

// NewAlibabaWdkMerchantStoreitemUpdateRequest 初始化AlibabaWdkMerchantStoreitemUpdateAPIRequest对象
func NewAlibabaWdkMerchantStoreitemUpdateRequest() *AlibabaWdkMerchantStoreitemUpdateAPIRequest {
	return &AlibabaWdkMerchantStoreitemUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.storeitem.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetMerchantCode is MerchantCode Setter
// 商家编码
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetStoreId is StoreId Setter
// 门店编码
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetParams is Params Setter
// 修改参数的json
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetParams() string {
	return r._params
}
