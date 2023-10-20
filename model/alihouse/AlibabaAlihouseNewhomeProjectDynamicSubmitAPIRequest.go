package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest 提交楼盘快讯 API请求
// alibaba.alihouse.newhome.project.dynamic.submit
//
// 提交楼盘快讯
type AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest struct {
	model.Params
	// 楼盘动态列表
	_projectDynamics []ProjectDynamicDto
}

// NewAlibabaalihousenewhomeprojectdynamicsubmitRequest 初始化AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest对象
func NewAlibabaalihousenewhomeprojectdynamicsubmitRequest() *AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest {
	return &AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.dynamic.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectDynamics is ProjectDynamics Setter
// 楼盘动态列表
func (r *AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest) SetProjectDynamics(_projectDynamics []ProjectDynamicDto) error {
	r._projectDynamics = _projectDynamics
	r.Set("project_dynamics", _projectDynamics)
	return nil
}

// GetProjectDynamics ProjectDynamics Getter
func (r AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest) GetProjectDynamics() []ProjectDynamicDto {
	return r._projectDynamics
}
