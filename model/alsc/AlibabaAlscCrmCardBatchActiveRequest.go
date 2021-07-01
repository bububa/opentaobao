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
type AlibabaAlscCrmCardBatchActiveAPIRequest struct {
    model.Params
    // 请求对象
    _paramBatchActiveCardOpenReq   *BatchActiveCardOpenReq
}

// 初始化AlibabaAlscCrmCardBatchActiveAPIRequest对象
func NewAlibabaAlscCrmCardBatchActiveRequest() *AlibabaAlscCrmCardBatchActiveAPIRequest{
    return &AlibabaAlscCrmCardBatchActiveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBatchActiveAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.batch.active"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBatchActiveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBatchActiveCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardBatchActiveAPIRequest) SetParamBatchActiveCardOpenReq(_paramBatchActiveCardOpenReq *BatchActiveCardOpenReq) error {
    r._paramBatchActiveCardOpenReq = _paramBatchActiveCardOpenReq
    r.Set("param_batch_active_card_open_req", _paramBatchActiveCardOpenReq)
    return nil
}

// ParamBatchActiveCardOpenReq Getter
func (r AlibabaAlscCrmCardBatchActiveAPIRequest) GetParamBatchActiveCardOpenReq() *BatchActiveCardOpenReq {
    return r._paramBatchActiveCardOpenReq
}
