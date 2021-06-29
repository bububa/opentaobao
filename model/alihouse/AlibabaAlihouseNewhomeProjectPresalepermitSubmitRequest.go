package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交预售证 API请求
alibaba.alihouse.newhome.project.presalepermit.submit

提交楼盘预售证信息
*/
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest struct {
    model.Params
    // 预售证对象
    preSalePermitDto   *ProjectPreSalePermitDto
}

// 初始化AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest对象
func NewAlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest() *AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.presalepermit.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PreSalePermitDto Setter
// 预售证对象
func (r *AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest) SetPreSalePermitDto(preSalePermitDto *ProjectPreSalePermitDto) error {
    r.preSalePermitDto = preSalePermitDto
    r.Set("pre_sale_permit_dto", preSalePermitDto)
    return nil
}

// PreSalePermitDto Getter
func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest) GetPreSalePermitDto() *ProjectPreSalePermitDto {
    return r.preSalePermitDto
}
