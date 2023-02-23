package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemCategoryUpdateAPIRequest 修改类目 API请求
// alibaba.wdk.item.category.update
//
// 修改类目
type AlibabaWdkItemCategoryUpdateAPIRequest struct {
	model.Params
	// 入参
	_bean *UpdateCategoryRequestBean
}

// NewAlibabaWdkItemCategoryUpdateRequest 初始化AlibabaWdkItemCategoryUpdateAPIRequest对象
func NewAlibabaWdkItemCategoryUpdateRequest() *AlibabaWdkItemCategoryUpdateAPIRequest {
	return &AlibabaWdkItemCategoryUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemCategoryUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.category.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemCategoryUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemCategoryUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBean is Bean Setter
// 入参
func (r *AlibabaWdkItemCategoryUpdateAPIRequest) SetBean(_bean *UpdateCategoryRequestBean) error {
	r._bean = _bean
	r.Set("bean", _bean)
	return nil
}

// GetBean Bean Getter
func (r AlibabaWdkItemCategoryUpdateAPIRequest) GetBean() *UpdateCategoryRequestBean {
	return r._bean
}
