package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘信息 APIRequest
alibaba.alihouse.newhome.project.submit

提交楼盘信息
*/
type AlibabaAlihouseNewhomeProjectSubmitRequest struct {
    model.Params

    // 楼盘对象
    ebbasProjectDto   *EbbasProjectDto 

}

func NewAlibabaAlihouseNewhomeProjectSubmitRequest() *AlibabaAlihouseNewhomeProjectSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.submit"
}

func (r AlibabaAlihouseNewhomeProjectSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectSubmitRequest) SetEbbasProjectDto(ebbasProjectDto *EbbasProjectDto) error {
    r.ebbasProjectDto = ebbasProjectDto
    r.Set("ebbas_project_dto", ebbasProjectDto)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectSubmitRequest) GetEbbasProjectDto() *EbbasProjectDto {
    return r.ebbasProjectDto
}

