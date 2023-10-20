package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest 提交KA合作楼盘 API请求
// alibaba.alihouse.newhome.project.cooperation.submit
//
// 提交KA合作楼盘
type AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest struct {
	model.Params
	// ka合作对象
	_projectCooperationDto *ProjectCooperationDto
}

// NewAlibabaAlihouseNewhomeProjectCooperationSubmitRequest 初始化AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectCooperationSubmitRequest() *AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) Reset() {
	r._projectCooperationDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.cooperation.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectCooperationDto is ProjectCooperationDto Setter
// ka合作对象
func (r *AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) SetProjectCooperationDto(_projectCooperationDto *ProjectCooperationDto) error {
	r._projectCooperationDto = _projectCooperationDto
	r.Set("project_cooperation_dto", _projectCooperationDto)
	return nil
}

// GetProjectCooperationDto ProjectCooperationDto Getter
func (r AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) GetProjectCooperationDto() *ProjectCooperationDto {
	return r._projectCooperationDto
}

var poolAlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectCooperationSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectCooperationSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest
func GetAlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest() *AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest 将 AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest(v *AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest.Put(v)
}
