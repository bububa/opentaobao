package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest 案场活动维护 API请求
// alibaba.alihouse.newhome.casefield.activity.submit
//
// 案场活动维护
type AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest struct {
	model.Params
	// 请求对象
	_caseFieldActivityDto *CaseFieldActivityDto
}

// NewAlibabaAlihouseNewhomeCasefieldActivitySubmitRequest 初始化AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeCasefieldActivitySubmitRequest() *AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest {
	return &AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest) Reset() {
	r._caseFieldActivityDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.casefield.activity.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCaseFieldActivityDto is CaseFieldActivityDto Setter
// 请求对象
func (r *AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest) SetCaseFieldActivityDto(_caseFieldActivityDto *CaseFieldActivityDto) error {
	r._caseFieldActivityDto = _caseFieldActivityDto
	r.Set("case_field_activity_dto", _caseFieldActivityDto)
	return nil
}

// GetCaseFieldActivityDto CaseFieldActivityDto Getter
func (r AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest) GetCaseFieldActivityDto() *CaseFieldActivityDto {
	return r._caseFieldActivityDto
}

var poolAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeCasefieldActivitySubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeCasefieldActivitySubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest
func GetAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest() *AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest 将 AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest(v *AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest.Put(v)
}
