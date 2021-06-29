package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交KA合作楼盘 APIRequest
alibaba.alihouse.newhome.project.cooperation.submit

提交KA合作楼盘
*/
type AlibabaAlihouseNewhomeProjectCooperationSubmitRequest struct {
    model.Params

    // ka合作对象
    projectCooperationDto   *ProjectCooperationDto 

}

func NewAlibabaAlihouseNewhomeProjectCooperationSubmitRequest() *AlibabaAlihouseNewhomeProjectCooperationSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectCooperationSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectCooperationSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.cooperation.submit"
}

func (r AlibabaAlihouseNewhomeProjectCooperationSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectCooperationSubmitRequest) SetProjectCooperationDto(projectCooperationDto *ProjectCooperationDto) error {
    r.projectCooperationDto = projectCooperationDto
    r.Set("project_cooperation_dto", projectCooperationDto)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectCooperationSubmitRequest) GetProjectCooperationDto() *ProjectCooperationDto {
    return r.projectCooperationDto
}

