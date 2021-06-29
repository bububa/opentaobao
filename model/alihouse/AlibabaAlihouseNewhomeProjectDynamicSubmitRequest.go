package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘快讯 APIRequest
alibaba.alihouse.newhome.project.dynamic.submit

提交楼盘快讯
*/
type AlibabaAlihouseNewhomeProjectDynamicSubmitRequest struct {
    model.Params

    // 楼盘动态列表
    projectDynamics   []ProjectDynamicDto 

}

func NewAlibabaAlihouseNewhomeProjectDynamicSubmitRequest() *AlibabaAlihouseNewhomeProjectDynamicSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectDynamicSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectDynamicSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.dynamic.submit"
}

func (r AlibabaAlihouseNewhomeProjectDynamicSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectDynamicSubmitRequest) SetProjectDynamics(projectDynamics []ProjectDynamicDto) error {
    r.projectDynamics = projectDynamics
    r.Set("project_dynamics", projectDynamics)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectDynamicSubmitRequest) GetProjectDynamics() []ProjectDynamicDto {
    return r.projectDynamics
}

