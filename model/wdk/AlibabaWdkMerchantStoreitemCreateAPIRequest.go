package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantStoreitemCreateAPIRequest 新建门店商品 API请求
// alibaba.wdk.merchant.storeitem.create
//
// 新建门店商品
type AlibabaWdkMerchantStoreitemCreateAPIRequest struct {
	model.Params
	// 门店id
	_storeId string
	// 商品编码
	_skuCode string
	// 商家编码
	_merchantCode string
	// 新增门店商品参数，具体字段详见文档
	_params string
}

// NewAlibabaWdkMerchantStoreitemCreateRequest 初始化AlibabaWdkMerchantStoreitemCreateAPIRequest对象
func NewAlibabaWdkMerchantStoreitemCreateRequest() *AlibabaWdkMerchantStoreitemCreateAPIRequest {
	return &AlibabaWdkMerchantStoreitemCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantStoreitemCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.storeitem.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantStoreitemCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 门店id
func (r *AlibabaWdkMerchantStoreitemCreateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaWdkMerchantStoreitemCreateAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantStoreitemCreateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// Get SkuCode Getter
func (r AlibabaWdkMerchantStoreitemCreateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// Set is MerchantCode Setter
// 商家编码
func (r *AlibabaWdkMerchantStoreitemCreateAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// Get MerchantCode Getter
func (r AlibabaWdkMerchantStoreitemCreateAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// Set is Params Setter
// 新增门店商品参数，具体字段详见文档
func (r *AlibabaWdkMerchantStoreitemCreateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// Get Params Getter
func (r AlibabaWdkMerchantStoreitemCreateAPIRequest) GetParams() string {
	return r._params
}
