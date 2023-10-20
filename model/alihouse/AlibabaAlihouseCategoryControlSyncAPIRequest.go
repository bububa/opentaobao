package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
