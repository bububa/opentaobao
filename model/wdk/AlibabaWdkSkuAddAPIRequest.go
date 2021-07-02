package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuAddAPIRequest 新增商品 API请求
// alibaba.wdk.sku.add
//
// 创建RT门店商品或DC商品
type AlibabaWdkSkuAddAPIRequest struct {
	model.Params
	// 商品列表
	_paramList []SkuDo
}

// NewAlibabaWdkSkuAddRequest 初始化AlibabaWdkSkuAddAPIRequest对象
func NewAlibabaWdkSkuAddRequest() *AlibabaWdkSkuAddAPIRequest {
	return &AlibabaWdkSkuAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamList is ParamList Setter
// 商品列表
func (r *AlibabaWdkSkuAddAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaWdkSkuAddAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}
