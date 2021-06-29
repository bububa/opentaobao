package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘电话 APIRequest
alibaba.alihouse.newhome.project.phone.submit

提交楼盘电话
*/
type AlibabaAlihouseNewhomeProjectPhoneSubmitRequest struct {
    model.Params

    // 楼盘电话
    projectPhoneDto   *ProjectPhoneDto 

}

func NewAlibabaAlihouseNewhomeProjectPhoneSubmitRequest() *AlibabaAlihouseNewhomeProjectPhoneSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectPhoneSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectPhoneSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.phone.submit"
}

func (r AlibabaAlihouseNewhomeProjectPhoneSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectPhoneSubmitRequest) SetProjectPhoneDto(projectPhoneDto *ProjectPhoneDto) error {
    r.projectPhoneDto = projectPhoneDto
    r.Set("project_phone_dto", projectPhoneDto)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectPhoneSubmitRequest) GetProjectPhoneDto() *ProjectPhoneDto {
    return r.projectPhoneDto
}

