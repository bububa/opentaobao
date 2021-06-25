package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标准开卡流程 APIRequest
alibaba.alsc.crm.card.open

标准开卡流程
*/
type AlibabaAlscCrmCardOpenRequest struct {
    model.Params

    // 开卡参数
    paramOpenCardStandardOpenReq   *OpenCardStandardOpenReq 

}

func NewAlibabaAlscCrmCardOpenRequest() *AlibabaAlscCrmCardOpenRequest{
    return &AlibabaAlscCrmCardOpenRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCardOpenRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.open"
}

func (r AlibabaAlscCrmCardOpenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCardOpenRequest) SetParamOpenCardStandardOpenReq(paramOpenCardStandardOpenReq *OpenCardStandardOpenReq) error {
    r.paramOpenCardStandardOpenReq = paramOpenCardStandardOpenReq
    r.Set("param_open_card_standard_open_req", paramOpenCardStandardOpenReq)
    return nil
}

func (r AlibabaAlscCrmCardOpenRequest) GetParamOpenCardStandardOpenReq() *OpenCardStandardOpenReq {
    return r.paramOpenCardStandardOpenReq
}

