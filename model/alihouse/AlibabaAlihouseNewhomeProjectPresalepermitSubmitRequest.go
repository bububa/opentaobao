package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交预售证 APIRequest
alibaba.alihouse.newhome.project.presalepermit.submit

提交楼盘预售证信息
*/
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest struct {
    model.Params

    // 预售证对象
    preSalePermitDto   *ProjectPreSalePermitDto 

}

func NewAlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest() *AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.presalepermit.submit"
}

func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest) SetPreSalePermitDto(preSalePermitDto *ProjectPreSalePermitDto) error {
    r.preSalePermitDto = preSalePermitDto
    r.Set("pre_sale_permit_dto", preSalePermitDto)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest) GetPreSalePermitDto() *ProjectPreSalePermitDto {
    return r.preSalePermitDto
}

