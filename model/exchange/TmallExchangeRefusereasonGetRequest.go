package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取拒绝换货原因列表 APIRequest
tmall.exchange.refusereason.get

获取拒绝换货原因列表
*/
type TmallExchangeRefusereasonGetRequest struct {
    model.Params

    // 换货单号ID
    disputeId   int64 

    // 返回字段
    fields   []string 

    // 换货申请类型：0-任意类型；1-售中；2-售后
    disputeType   int64 

}

func NewTmallExchangeRefusereasonGetRequest() *TmallExchangeRefusereasonGetRequest{
    return &TmallExchangeRefusereasonGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallExchangeRefusereasonGetRequest) GetApiMethodName() string {
    return "tmall.exchange.refusereason.get"
}

func (r TmallExchangeRefusereasonGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallExchangeRefusereasonGetRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

func (r TmallExchangeRefusereasonGetRequest) GetDisputeId() int64 {
    return r.disputeId
}

func (r *TmallExchangeRefusereasonGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TmallExchangeRefusereasonGetRequest) GetFields() []string {
    return r.fields
}

func (r *TmallExchangeRefusereasonGetRequest) SetDisputeType(disputeType int64) error {
    r.disputeType = disputeType
    r.Set("dispute_type", disputeType)
    return nil
}

func (r TmallExchangeRefusereasonGetRequest) GetDisputeType() int64 {
    return r.disputeType
}

