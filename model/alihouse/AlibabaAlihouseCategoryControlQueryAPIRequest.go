package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousecategorycontrolqueryAPIRequest 类目权限查询 API请求
// alibaba.alihouse.category.control.query
//
// 类目权限查询
type AlibabaalihousecategorycontrolqueryAPIRequest struct {
	model.Params
	// dto
	_categoryControlDto *CategoryControlDto
}

// NewAlibabaalihousecategorycontrolqueryRequest 初始化AlibabaalihousecategorycontrolqueryAPIRequest对象
func NewAlibabaalihousecategorycontrolqueryRequest() *AlibabaalihousecategorycontrolqueryAPIRequest {
	return &AlibabaalihousecategorycontrolqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousecategorycontrolqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.category.control.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousecategorycontrolqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousecategorycontrolqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryControlDto is CategoryControlDto Setter
// dto
func (r *AlibabaalihousecategorycontrolqueryAPIRequest) SetCategoryControlDto(_categoryControlDto *CategoryControlDto) error {
	r._categoryControlDto = _categoryControlDto
	r.Set("category_control_dto", _categoryControlDto)
	return nil
}

// GetCategoryControlDto CategoryControlDto Getter
func (r AlibabaalihousecategorycontrolqueryAPIRequest) GetCategoryControlDto() *CategoryControlDto {
	return r._categoryControlDto
}
