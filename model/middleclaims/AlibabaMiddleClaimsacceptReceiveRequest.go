package middleclaims

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际化中台服务域接收保险公司理赔受理结果 APIRequest
alibaba.middle.claimsaccept.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔受理结果的处理后，将该结果返回至服务域
*/
type AlibabaMiddleClaimsacceptReceiveRequest struct {
    model.Params

    // 理赔受理结果实体类
    claimsAcceptDto   *ClaimsAcceptDto 

}

func NewAlibabaMiddleClaimsacceptReceiveRequest() *AlibabaMiddleClaimsacceptReceiveRequest{
    return &AlibabaMiddleClaimsacceptReceiveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMiddleClaimsacceptReceiveRequest) GetApiMethodName() string {
    return "alibaba.middle.claimsaccept.receive"
}

func (r AlibabaMiddleClaimsacceptReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMiddleClaimsacceptReceiveRequest) SetClaimsAcceptDto(claimsAcceptDto *ClaimsAcceptDto) error {
    r.claimsAcceptDto = claimsAcceptDto
    r.Set("claims_accept_dto", claimsAcceptDto)
    return nil
}

func (r AlibabaMiddleClaimsacceptReceiveRequest) GetClaimsAcceptDto() *ClaimsAcceptDto {
    return r.claimsAcceptDto
}

