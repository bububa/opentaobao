package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousecategorycontrolsyncAPIRequest 类目权限上翻 API请求
// alibaba.alihouse.category.control.sync
//
// 类目权限上翻
type AlibabaalihousecategorycontrolsyncAPIRequest struct {
	model.Params
	// 入参
	_categoryControlDto *CategoryControlDto
}

// NewAlibabaalihousecategorycontrolsyncRequest 初始化AlibabaalihousecategorycontrolsyncAPIRequest对象
func NewAlibabaalihousecategorycontrolsyncRequest() *AlibabaalihousecategorycontrolsyncAPIRequest {
	return &AlibabaalihousecategorycontrolsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousecategorycontrolsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.category.control.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousecategorycontrolsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousecategorycontrolsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryControlDto is CategoryControlDto Setter
// 入参
func (r *AlibabaalihousecategorycontrolsyncAPIRequest) SetCategoryControlDto(_categoryControlDto *CategoryControlDto) error {
	r._categoryControlDto = _categoryControlDto
	r.Set("category_control_dto", _categoryControlDto)
	return nil
}

// GetCategoryControlDto CategoryControlDto Getter
func (r AlibabaalihousecategorycontrolsyncAPIRequest) GetCategoryControlDto() *CategoryControlDto {
	return r._categoryControlDto
}
