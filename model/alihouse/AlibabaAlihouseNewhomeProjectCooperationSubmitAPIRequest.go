package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.cooperation.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProjectCooperationDto Setter
// ka合作对象
func (r *AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) SetProjectCooperationDto(_projectCooperationDto *ProjectCooperationDto) error {
	r._projectCooperationDto = _projectCooperationDto
	r.Set("project_cooperation_dto", _projectCooperationDto)
	return nil
}

// Get ProjectCooperationDto Getter
func (r AlibabaAlihouseNewhomeProjectCooperationSubmitAPIRequest) GetProjectCooperationDto() *ProjectCooperationDto {
	return r._projectCooperationDto
}
