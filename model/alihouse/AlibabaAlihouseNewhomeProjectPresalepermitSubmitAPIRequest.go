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
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest struct {
    model.Params
    // 预售证对象
    _preSalePermitDto   *ProjectPreSalePermitDto
}

// 初始化AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectPresalepermitSubmitRequest() *AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest{
    return &AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.presalepermit.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PreSalePermitDto Setter
// 预售证对象
func (r *AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest) SetPreSalePermitDto(_preSalePermitDto *ProjectPreSalePermitDto) error {
    r._preSalePermitDto = _preSalePermitDto
    r.Set("pre_sale_permit_dto", _preSalePermitDto)
    return nil
}

// PreSalePermitDto Getter
func (r AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest) GetPreSalePermitDto() *ProjectPreSalePermitDto {
    return r._preSalePermitDto
}
