package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量开卡（售卡） APIRequest
alibaba.alsc.crm.card.batch.sell

批量开卡（售卡）
*/
type AlibabaAlscCrmCardBatchSellRequest struct {
    model.Params

    // 请求对象
    paramBatchOpenCardOpenReq   *BatchOpenCardOpenReq 

}

func NewAlibabaAlscCrmCardBatchSellRequest() *AlibabaAlscCrmCardBatchSellRequest{
    return &AlibabaAlscCrmCardBatchSellRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCardBatchSellRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.batch.sell"
}

func (r AlibabaAlscCrmCardBatchSellRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCardBatchSellRequest) SetParamBatchOpenCardOpenReq(paramBatchOpenCardOpenReq *BatchOpenCardOpenReq) error {
    r.paramBatchOpenCardOpenReq = paramBatchOpenCardOpenReq
    r.Set("param_batch_open_card_open_req", paramBatchOpenCardOpenReq)
    return nil
}

func (r AlibabaAlscCrmCardBatchSellRequest) GetParamBatchOpenCardOpenReq() *BatchOpenCardOpenReq {
    return r.paramBatchOpenCardOpenReq
}

