package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseCategoryControlQueryAPIRequest 类目权限查询 API请求
// alibaba.alihouse.category.control.query
//
// 类目权限查询
type AlibabaAlihouseCategoryControlQueryAPIRequest struct {
	model.Params
	// dto
	_categoryControlDto *CategoryControlDto
}

// NewAlibabaAlihouseCategoryControlQueryRequest 初始化AlibabaAlihouseCategoryControlQueryAPIRequest对象
func NewAlibabaAlihouseCategoryControlQueryRequest() *AlibabaAlihouseCategoryControlQueryAPIRequest {
	return &AlibabaAlihouseCategoryControlQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseCategoryControlQueryAPIRequest) Reset() {
	r._categoryControlDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseCategoryControlQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.category.control.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseCategoryControlQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseCategoryControlQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryControlDto is CategoryControlDto Setter
// dto
func (r *AlibabaAlihouseCategoryControlQueryAPIRequest) SetCategoryControlDto(_categoryControlDto *CategoryControlDto) error {
	r._categoryControlDto = _categoryControlDto
	r.Set("category_control_dto", _categoryControlDto)
	return nil
}

// GetCategoryControlDto CategoryControlDto Getter
func (r AlibabaAlihouseCategoryControlQueryAPIRequest) GetCategoryControlDto() *CategoryControlDto {
	return r._categoryControlDto
}

var poolAlibabaAlihouseCategoryControlQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseCategoryControlQueryRequest()
	},
}

// GetAlibabaAlihouseCategoryControlQueryRequest 从 sync.Pool 获取 AlibabaAlihouseCategoryControlQueryAPIRequest
func GetAlibabaAlihouseCategoryControlQueryAPIRequest() *AlibabaAlihouseCategoryControlQueryAPIRequest {
	return poolAlibabaAlihouseCategoryControlQueryAPIRequest.Get().(*AlibabaAlihouseCategoryControlQueryAPIRequest)
}

// ReleaseAlibabaAlihouseCategoryControlQueryAPIRequest 将 AlibabaAlihouseCategoryControlQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseCategoryControlQueryAPIRequest(v *AlibabaAlihouseCategoryControlQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseCategoryControlQueryAPIRequest.Put(v)
}
