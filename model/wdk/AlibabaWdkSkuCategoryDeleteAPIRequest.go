package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCategoryDeleteAPIRequest 商家类目删除接口 API请求
// alibaba.wdk.sku.category.delete
//
// 商家类目删除接口
type AlibabaWdkSkuCategoryDeleteAPIRequest struct {
	model.Params
	// 类目删除请求模型
	_param *CategoryDo
}

// NewAlibabaWdkSkuCategoryDeleteRequest 初始化AlibabaWdkSkuCategoryDeleteAPIRequest对象
func NewAlibabaWdkSkuCategoryDeleteRequest() *AlibabaWdkSkuCategoryDeleteAPIRequest {
	return &AlibabaWdkSkuCategoryDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCategoryDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.category.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCategoryDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 类目删除请求模型
func (r *AlibabaWdkSkuCategoryDeleteAPIRequest) SetParam(_param *CategoryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuCategoryDeleteAPIRequest) GetParam() *CategoryDo {
	return r._param
}
