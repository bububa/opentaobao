package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantItemUpdateAPIRequest 修改商家商品 API请求
// alibaba.wdk.merchant.item.update
//
// 修改商家商品
type AlibabaWdkMerchantItemUpdateAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 门店编码
	_merchantCode string
	// 修改字段的json
	_params string
}

// NewAlibabaWdkMerchantItemUpdateRequest 初始化AlibabaWdkMerchantItemUpdateAPIRequest对象
func NewAlibabaWdkMerchantItemUpdateRequest() *AlibabaWdkMerchantItemUpdateAPIRequest {
	return &AlibabaWdkMerchantItemUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantItemUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.item.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantItemUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantItemUpdateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkMerchantItemUpdateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetMerchantCode is MerchantCode Setter
// 门店编码
func (r *AlibabaWdkMerchantItemUpdateAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaWdkMerchantItemUpdateAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetParams is Params Setter
// 修改字段的json
func (r *AlibabaWdkMerchantItemUpdateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaWdkMerchantItemUpdateAPIRequest) GetParams() string {
	return r._params
}
