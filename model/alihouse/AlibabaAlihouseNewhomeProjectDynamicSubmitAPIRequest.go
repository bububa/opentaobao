package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest 提交楼盘快讯 API请求
// alibaba.alihouse.newhome.project.dynamic.submit
//
// 提交楼盘快讯
type AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest struct {
	model.Params
	// 楼盘动态列表
	_projectDynamics []ProjectDynamicDto
}

// NewAlibabaAlihouseNewhomeProjectDynamicSubmitRequest 初始化AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectDynamicSubmitRequest() *AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.dynamic.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProjectDynamics is ProjectDynamics Setter
// 楼盘动态列表
func (r *AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest) SetProjectDynamics(_projectDynamics []ProjectDynamicDto) error {
	r._projectDynamics = _projectDynamics
	r.Set("project_dynamics", _projectDynamics)
	return nil
}

// GetProjectDynamics ProjectDynamics Getter
func (r AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest) GetProjectDynamics() []ProjectDynamicDto {
	return r._projectDynamics
}
