package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantItemCreatedraftAPIRequest 新建商品草稿 API请求
// alibaba.wdk.merchant.item.createdraft
//
// 新建商品草稿erp接口
type AlibabaWdkMerchantItemCreatedraftAPIRequest struct {
	model.Params
	// 商品信息json
	_params string
}

// NewAlibabaWdkMerchantItemCreatedraftRequest 初始化AlibabaWdkMerchantItemCreatedraftAPIRequest对象
func NewAlibabaWdkMerchantItemCreatedraftRequest() *AlibabaWdkMerchantItemCreatedraftAPIRequest {
	return &AlibabaWdkMerchantItemCreatedraftAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantItemCreatedraftAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.item.createdraft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantItemCreatedraftAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParams is Params Setter
// 商品信息json
func (r *AlibabaWdkMerchantItemCreatedraftAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaWdkMerchantItemCreatedraftAPIRequest) GetParams() string {
	return r._params
}
