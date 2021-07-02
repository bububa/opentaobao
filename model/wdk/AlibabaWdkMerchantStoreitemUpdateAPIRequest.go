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
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// Get SkuCode Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// Set is MerchantCode Setter
// 商家编码
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// Get MerchantCode Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// Set is StoreId Setter
// 门店编码
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is Params Setter
// 修改参数的json
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// Get Params Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetParams() string {
	return r._params
}
