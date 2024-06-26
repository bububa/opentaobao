package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest 案场活动楼盘维护 API请求
// alibaba.alihouse.newhome.casefield.activity.project.submit
//
// 案场活动楼盘维护
type AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest struct {
	model.Params
	// 请求对象
	_caseFieldActivityProjectsDto *CaseFieldActivityProjectsDto
}

// NewAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitRequest 初始化AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitRequest() *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) Reset() {
	r._caseFieldActivityProjectsDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.casefield.activity.project.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCaseFieldActivityProjectsDto is CaseFieldActivityProjectsDto Setter
// 请求对象
func (r *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) SetCaseFieldActivityProjectsDto(_caseFieldActivityProjectsDto *CaseFieldActivityProjectsDto) error {
	r._caseFieldActivityProjectsDto = _caseFieldActivityProjectsDto
	r.Set("case_field_activity_projects_dto", _caseFieldActivityProjectsDto)
	return nil
}

// GetCaseFieldActivityProjectsDto CaseFieldActivityProjectsDto Getter
func (r AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) GetCaseFieldActivityProjectsDto() *CaseFieldActivityProjectsDto {
	return r._caseFieldActivityProjectsDto
}

var poolAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest
func GetAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest() *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest 将 AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest(v *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest.Put(v)
}
