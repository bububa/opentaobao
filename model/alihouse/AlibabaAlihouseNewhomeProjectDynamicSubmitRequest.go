package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘快讯 API请求
alibaba.alihouse.newhome.project.dynamic.submit

提交楼盘快讯
*/
type AlibabaAlihouseNewhomeProjectDynamicSubmitRequest struct {
    model.Params
    // 楼盘动态列表
    _projectDynamics   []ProjectDynamicDTO
}

// 初始化AlibabaAlihouseNewhomeProjectDynamicSubmitRequest对象
func NewAlibabaAlihouseNewhomeProjectDynamicSubmitRequest() *AlibabaAlihouseNewhomeProjectDynamicSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectDynamicSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectDynamicSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.dynamic.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectDynamicSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProjectDynamics Setter
// 楼盘动态列表
func (r *AlibabaAlihouseNewhomeProjectDynamicSubmitRequest) SetProjectDynamics(_projectDynamics []ProjectDynamicDTO) error {
    r._projectDynamics = _projectDynamics
    r.Set("project_dynamics", _projectDynamics)
    return nil
}

// ProjectDynamics Getter
func (r AlibabaAlihouseNewhomeProjectDynamicSubmitRequest) GetProjectDynamics() []ProjectDynamicDTO {
    return r._projectDynamics
}
