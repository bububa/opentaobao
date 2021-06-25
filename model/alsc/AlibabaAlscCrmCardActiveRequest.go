package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标准激活卡 APIRequest
alibaba.alsc.crm.card.active

激活卡
*/
type AlibabaAlscCrmCardActiveRequest struct {
    model.Params

    // 请求参数
    paramActiveCardOpenReq   *ActiveCardOpenReq 

}

func NewAlibabaAlscCrmCardActiveRequest() *AlibabaAlscCrmCardActiveRequest{
    return &AlibabaAlscCrmCardActiveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCardActiveRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.active"
}

func (r AlibabaAlscCrmCardActiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCardActiveRequest) SetParamActiveCardOpenReq(paramActiveCardOpenReq *ActiveCardOpenReq) error {
    r.paramActiveCardOpenReq = paramActiveCardOpenReq
    r.Set("param_active_card_open_req", paramActiveCardOpenReq)
    return nil
}

func (r AlibabaAlscCrmCardActiveRequest) GetParamActiveCardOpenReq() *ActiveCardOpenReq {
    return r.paramActiveCardOpenReq
}

