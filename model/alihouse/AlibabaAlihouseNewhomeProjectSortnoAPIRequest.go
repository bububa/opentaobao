package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectSortnoAPIRequest) Reset() {
	r._updateSortNoDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectSortnoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.sortno"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectSortnoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectSortnoAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseNewhomeProjectSortnoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectSortnoRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectSortnoRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectSortnoAPIRequest
func GetAlibabaAlihouseNewhomeProjectSortnoAPIRequest() *AlibabaAlihouseNewhomeProjectSortnoAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectSortnoAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectSortnoAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectSortnoAPIRequest 将 AlibabaAlihouseNewhomeProjectSortnoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectSortnoAPIRequest(v *AlibabaAlihouseNewhomeProjectSortnoAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectSortnoAPIRequest.Put(v)
}
