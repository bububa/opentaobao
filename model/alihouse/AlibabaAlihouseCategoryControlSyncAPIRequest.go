package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseCategoryControlSyncAPIRequest 类目权限上翻 API请求
// alibaba.alihouse.category.control.sync
//
// 类目权限上翻
type AlibabaAlihouseCategoryControlSyncAPIRequest struct {
	model.Params
	// 入参
	_categoryControlDto *CategoryControlDto
}

// NewAlibabaAlihouseCategoryControlSyncRequest 初始化AlibabaAlihouseCategoryControlSyncAPIRequest对象
func NewAlibabaAlihouseCategoryControlSyncRequest() *AlibabaAlihouseCategoryControlSyncAPIRequest {
	return &AlibabaAlihouseCategoryControlSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseCategoryControlSyncAPIRequest) Reset() {
	r._categoryControlDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseCategoryControlSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.category.control.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseCategoryControlSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseCategoryControlSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryControlDto is CategoryControlDto Setter
// 入参
func (r *AlibabaAlihouseCategoryControlSyncAPIRequest) SetCategoryControlDto(_categoryControlDto *CategoryControlDto) error {
	r._categoryControlDto = _categoryControlDto
	r.Set("category_control_dto", _categoryControlDto)
	return nil
}

// GetCategoryControlDto CategoryControlDto Getter
func (r AlibabaAlihouseCategoryControlSyncAPIRequest) GetCategoryControlDto() *CategoryControlDto {
	return r._categoryControlDto
}

var poolAlibabaAlihouseCategoryControlSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseCategoryControlSyncRequest()
	},
}

// GetAlibabaAlihouseCategoryControlSyncRequest 从 sync.Pool 获取 AlibabaAlihouseCategoryControlSyncAPIRequest
func GetAlibabaAlihouseCategoryControlSyncAPIRequest() *AlibabaAlihouseCategoryControlSyncAPIRequest {
	return poolAlibabaAlihouseCategoryControlSyncAPIRequest.Get().(*AlibabaAlihouseCategoryControlSyncAPIRequest)
}

// ReleaseAlibabaAlihouseCategoryControlSyncAPIRequest 将 AlibabaAlihouseCategoryControlSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseCategoryControlSyncAPIRequest(v *AlibabaAlihouseCategoryControlSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseCategoryControlSyncAPIRequest.Put(v)
}
