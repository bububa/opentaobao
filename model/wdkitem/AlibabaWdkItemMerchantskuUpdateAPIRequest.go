package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantskuUpdateAPIRequest 商家商品修改 API请求
// alibaba.wdk.item.merchantsku.update
//
// 商家商品修改
type AlibabaWdkItemMerchantskuUpdateAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 参数json
	_params string
	// 机构编码
	_orgCode string
}

// NewAlibabaWdkItemMerchantskuUpdateRequest 初始化AlibabaWdkItemMerchantskuUpdateAPIRequest对象
func NewAlibabaWdkItemMerchantskuUpdateRequest() *AlibabaWdkItemMerchantskuUpdateAPIRequest {
	return &AlibabaWdkItemMerchantskuUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantskuUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchantsku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantskuUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantskuUpdateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// Get SkuCode Getter
func (r AlibabaWdkItemMerchantskuUpdateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// Set is Params Setter
// 参数json
func (r *AlibabaWdkItemMerchantskuUpdateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// Get Params Getter
func (r AlibabaWdkItemMerchantskuUpdateAPIRequest) GetParams() string {
	return r._params
}

// Set is OrgCode Setter
// 机构编码
func (r *AlibabaWdkItemMerchantskuUpdateAPIRequest) SetOrgCode(_orgCode string) error {
	r._orgCode = _orgCode
	r.Set("org_code", _orgCode)
	return nil
}

// Get OrgCode Getter
func (r AlibabaWdkItemMerchantskuUpdateAPIRequest) GetOrgCode() string {
	return r._orgCode
}
