package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemCategoryUpdateAPIRequest) Reset() {
	r._bean = nil
	r.Params.ToZero()
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

var poolAlibabaWdkItemCategoryUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemCategoryUpdateRequest()
	},
}

// GetAlibabaWdkItemCategoryUpdateRequest 从 sync.Pool 获取 AlibabaWdkItemCategoryUpdateAPIRequest
func GetAlibabaWdkItemCategoryUpdateAPIRequest() *AlibabaWdkItemCategoryUpdateAPIRequest {
	return poolAlibabaWdkItemCategoryUpdateAPIRequest.Get().(*AlibabaWdkItemCategoryUpdateAPIRequest)
}

// ReleaseAlibabaWdkItemCategoryUpdateAPIRequest 将 AlibabaWdkItemCategoryUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemCategoryUpdateAPIRequest(v *AlibabaWdkItemCategoryUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemCategoryUpdateAPIRequest.Put(v)
}
