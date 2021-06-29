package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量开卡（售卡） API请求
alibaba.alsc.crm.card.batch.sell

批量开卡（售卡）
*/
type AlibabaAlscCrmCardBatchSellRequest struct {
    model.Params
    // 请求对象
    paramBatchOpenCardOpenReq   *BatchOpenCardOpenReq
}

// 初始化AlibabaAlscCrmCardBatchSellRequest对象
func NewAlibabaAlscCrmCardBatchSellRequest() *AlibabaAlscCrmCardBatchSellRequest{
    return &AlibabaAlscCrmCardBatchSellRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBatchSellRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.batch.sell"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBatchSellRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBatchOpenCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardBatchSellRequest) SetParamBatchOpenCardOpenReq(paramBatchOpenCardOpenReq *BatchOpenCardOpenReq) error {
    r.paramBatchOpenCardOpenReq = paramBatchOpenCardOpenReq
    r.Set("param_batch_open_card_open_req", paramBatchOpenCardOpenReq)
    return nil
}

// ParamBatchOpenCardOpenReq Getter
func (r AlibabaAlscCrmCardBatchSellRequest) GetParamBatchOpenCardOpenReq() *BatchOpenCardOpenReq {
    return r.paramBatchOpenCardOpenReq
}
