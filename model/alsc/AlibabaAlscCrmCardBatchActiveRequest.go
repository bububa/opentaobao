package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量激活卡 API请求
alibaba.alsc.crm.card.batch.active

批量激活卡
*/
type AlibabaAlscCrmCardBatchActiveRequest struct {
    model.Params
    // 请求对象
    _paramBatchActiveCardOpenReq   *BatchActiveCardOpenReq
}

// 初始化AlibabaAlscCrmCardBatchActiveRequest对象
func NewAlibabaAlscCrmCardBatchActiveRequest() *AlibabaAlscCrmCardBatchActiveRequest{
    return &AlibabaAlscCrmCardBatchActiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBatchActiveRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.batch.active"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBatchActiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBatchActiveCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardBatchActiveRequest) SetParamBatchActiveCardOpenReq(_paramBatchActiveCardOpenReq *BatchActiveCardOpenReq) error {
    r._paramBatchActiveCardOpenReq = _paramBatchActiveCardOpenReq
    r.Set("param_batch_active_card_open_req", _paramBatchActiveCardOpenReq)
    return nil
}

// ParamBatchActiveCardOpenReq Getter
func (r AlibabaAlscCrmCardBatchActiveRequest) GetParamBatchActiveCardOpenReq() *BatchActiveCardOpenReq {
    return r._paramBatchActiveCardOpenReq
}
