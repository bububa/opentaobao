package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuUpdateAPIRequest 更新商品 API请求
// alibaba.wdk.sku.update
//
// 开放商品更新接口
type AlibabaWdkSkuUpdateAPIRequest struct {
	model.Params
	// 参数列表
	_paramList []SkuDo
}

// NewAlibabaWdkSkuUpdateRequest 初始化AlibabaWdkSkuUpdateAPIRequest对象
func NewAlibabaWdkSkuUpdateRequest() *AlibabaWdkSkuUpdateAPIRequest {
	return &AlibabaWdkSkuUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamList Setter
// 参数列表
func (r *AlibabaWdkSkuUpdateAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// Get ParamList Getter
func (r AlibabaWdkSkuUpdateAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}
