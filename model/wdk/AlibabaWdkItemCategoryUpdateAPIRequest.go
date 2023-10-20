package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemcategoryupdateAPIRequest 修改类目 API请求
// alibaba.wdk.item.category.update
//
// 修改类目
type AlibabawdkitemcategoryupdateAPIRequest struct {
	model.Params
	// 入参
	_bean *UpdateCategoryRequestBean
}

// NewAlibabawdkitemcategoryupdateRequest 初始化AlibabawdkitemcategoryupdateAPIRequest对象
func NewAlibabawdkitemcategoryupdateRequest() *AlibabawdkitemcategoryupdateAPIRequest {
	return &AlibabawdkitemcategoryupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemcategoryupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.category.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemcategoryupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemcategoryupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBean is Bean Setter
// 入参
func (r *AlibabawdkitemcategoryupdateAPIRequest) SetBean(_bean *UpdateCategoryRequestBean) error {
	r._bean = _bean
	r.Set("bean", _bean)
	return nil
}

// GetBean Bean Getter
func (r AlibabawdkitemcategoryupdateAPIRequest) GetBean() *UpdateCategoryRequestBean {
	return r._bean
}
