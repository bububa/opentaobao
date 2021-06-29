package middleclaims

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际化中台服务域接收理赔账单 APIRequest
alibaba.middle.claimsbill.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔打款的处理后，将该打款结果返回至服务域
*/
type AlibabaMiddleClaimsbillReceiveRequest struct {
    model.Params

    // 理赔账单实体
    claimsBillDto   *ClaimsBillDto 

}

func NewAlibabaMiddleClaimsbillReceiveRequest() *AlibabaMiddleClaimsbillReceiveRequest{
    return &AlibabaMiddleClaimsbillReceiveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMiddleClaimsbillReceiveRequest) GetApiMethodName() string {
    return "alibaba.middle.claimsbill.receive"
}

func (r AlibabaMiddleClaimsbillReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMiddleClaimsbillReceiveRequest) SetClaimsBillDto(claimsBillDto *ClaimsBillDto) error {
    r.claimsBillDto = claimsBillDto
    r.Set("claims_bill_dto", claimsBillDto)
    return nil
}

func (r AlibabaMiddleClaimsbillReceiveRequest) GetClaimsBillDto() *ClaimsBillDto {
    return r.claimsBillDto
}

