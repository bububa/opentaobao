package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.casefield.activity.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
