package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest) Reset() {
	r._projectDynamics = r._projectDynamics[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.dynamic.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectDynamicSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectDynamicSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest
func GetAlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest() *AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest 将 AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest(v *AlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectDynamicSubmitAPIRequest.Put(v)
}
