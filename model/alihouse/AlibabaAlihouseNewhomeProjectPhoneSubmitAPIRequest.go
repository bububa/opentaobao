package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest 提交楼盘电话 API请求
// alibaba.alihouse.newhome.project.phone.submit
//
// 提交楼盘电话
type AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest struct {
	model.Params
	// 楼盘电话
	_projectPhoneDto *ProjectPhoneDto
}

// NewAlibabaAlihouseNewhomeProjectPhoneSubmitRequest 初始化AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectPhoneSubmitRequest() *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.phone.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProjectPhoneDto is ProjectPhoneDto Setter
// 楼盘电话
func (r *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) SetProjectPhoneDto(_projectPhoneDto *ProjectPhoneDto) error {
	r._projectPhoneDto = _projectPhoneDto
	r.Set("project_phone_dto", _projectPhoneDto)
	return nil
}

// GetProjectPhoneDto ProjectPhoneDto Getter
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) GetProjectPhoneDto() *ProjectPhoneDto {
	return r._projectPhoneDto
}
