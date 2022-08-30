package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectSortnoAPIRequest 新房排序值同步 API请求
// alibaba.alihouse.newhome.project.sortno
//
// 新房排序值同步
type AlibabaAlihouseNewhomeProjectSortnoAPIRequest struct {
	model.Params
	// 排序值请求对象
	_updateSortNoDto *UpdateSortNoDto
}

// NewAlibabaAlihouseNewhomeProjectSortnoRequest 初始化AlibabaAlihouseNewhomeProjectSortnoAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectSortnoRequest() *AlibabaAlihouseNewhomeProjectSortnoAPIRequest {
	return &AlibabaAlihouseNewhomeProjectSortnoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectSortnoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.sortno"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectSortnoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUpdateSortNoDto is UpdateSortNoDto Setter
// 排序值请求对象
func (r *AlibabaAlihouseNewhomeProjectSortnoAPIRequest) SetUpdateSortNoDto(_updateSortNoDto *UpdateSortNoDto) error {
	r._updateSortNoDto = _updateSortNoDto
	r.Set("update_sort_no_dto", _updateSortNoDto)
	return nil
}

// GetUpdateSortNoDto UpdateSortNoDto Getter
func (r AlibabaAlihouseNewhomeProjectSortnoAPIRequest) GetUpdateSortNoDto() *UpdateSortNoDto {
	return r._updateSortNoDto
}
