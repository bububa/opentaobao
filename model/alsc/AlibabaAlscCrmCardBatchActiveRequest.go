package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量激活卡 APIRequest
alibaba.alsc.crm.card.batch.active

批量激活卡
*/
type AlibabaAlscCrmCardBatchActiveRequest struct {
    model.Params

    // 请求对象
    paramBatchActiveCardOpenReq   *BatchActiveCardOpenReq 

}

func NewAlibabaAlscCrmCardBatchActiveRequest() *AlibabaAlscCrmCardBatchActiveRequest{
    return &AlibabaAlscCrmCardBatchActiveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCardBatchActiveRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.batch.active"
}

func (r AlibabaAlscCrmCardBatchActiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCardBatchActiveRequest) SetParamBatchActiveCardOpenReq(paramBatchActiveCardOpenReq *BatchActiveCardOpenReq) error {
    r.paramBatchActiveCardOpenReq = paramBatchActiveCardOpenReq
    r.Set("param_batch_active_card_open_req", paramBatchActiveCardOpenReq)
    return nil
}

func (r AlibabaAlscCrmCardBatchActiveRequest) GetParamBatchActiveCardOpenReq() *BatchActiveCardOpenReq {
    return r.paramBatchActiveCardOpenReq
}

