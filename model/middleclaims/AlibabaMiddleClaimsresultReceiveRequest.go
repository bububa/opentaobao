package middleclaims

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际化中台服务域接收理赔结果 APIRequest
alibaba.middle.claimsresult.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔结果的处理后，将该结果返回至服务域
*/
type AlibabaMiddleClaimsresultReceiveRequest struct {
    model.Params

    // 理赔结果实体
    claimsResultDTO   *ClaimsResultDto 

}

func NewAlibabaMiddleClaimsresultReceiveRequest() *AlibabaMiddleClaimsresultReceiveRequest{
    return &AlibabaMiddleClaimsresultReceiveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMiddleClaimsresultReceiveRequest) GetApiMethodName() string {
    return "alibaba.middle.claimsresult.receive"
}

func (r AlibabaMiddleClaimsresultReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMiddleClaimsresultReceiveRequest) SetClaimsResultDTO(claimsResultDTO *ClaimsResultDto) error {
    r.claimsResultDTO = claimsResultDTO
    r.Set("claims_result_d_t_o", claimsResultDTO)
    return nil
}

func (r AlibabaMiddleClaimsresultReceiveRequest) GetClaimsResultDTO() *ClaimsResultDto {
    return r.claimsResultDTO
}

