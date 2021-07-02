package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCombineskuUpdateAPIRequest 组合商品更新接口 API请求
// alibaba.wdk.sku.combinesku.update
//
// 组合商品修改接口
type AlibabaWdkSkuCombineskuUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramList []SkuDo
}

// NewAlibabaWdkSkuCombineskuUpdateRequest 初始化AlibabaWdkSkuCombineskuUpdateAPIRequest对象
func NewAlibabaWdkSkuCombineskuUpdateRequest() *AlibabaWdkSkuCombineskuUpdateAPIRequest {
	return &AlibabaWdkSkuCombineskuUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCombineskuUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.combinesku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCombineskuUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamList is ParamList Setter
// 请求参数
func (r *AlibabaWdkSkuCombineskuUpdateAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaWdkSkuCombineskuUpdateAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}
