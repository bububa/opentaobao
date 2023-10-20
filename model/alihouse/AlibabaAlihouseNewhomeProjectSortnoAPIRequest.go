package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectsortnoAPIRequest 新房排序值同步 API请求
// alibaba.alihouse.newhome.project.sortno
//
// 新房排序值同步
type AlibabaalihousenewhomeprojectsortnoAPIRequest struct {
	model.Params
	// 排序值请求对象
	_updateSortNoDto *UpdateSortNoDto
}

// NewAlibabaalihousenewhomeprojectsortnoRequest 初始化AlibabaalihousenewhomeprojectsortnoAPIRequest对象
func NewAlibabaalihousenewhomeprojectsortnoRequest() *AlibabaalihousenewhomeprojectsortnoAPIRequest {
	return &AlibabaalihousenewhomeprojectsortnoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectsortnoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.sortno"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectsortnoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectsortnoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateSortNoDto is UpdateSortNoDto Setter
// 排序值请求对象
func (r *AlibabaalihousenewhomeprojectsortnoAPIRequest) SetUpdateSortNoDto(_updateSortNoDto *UpdateSortNoDto) error {
	r._updateSortNoDto = _updateSortNoDto
	r.Set("update_sort_no_dto", _updateSortNoDto)
	return nil
}

// GetUpdateSortNoDto UpdateSortNoDto Getter
func (r AlibabaalihousenewhomeprojectsortnoAPIRequest) GetUpdateSortNoDto() *UpdateSortNoDto {
	return r._updateSortNoDto
}
