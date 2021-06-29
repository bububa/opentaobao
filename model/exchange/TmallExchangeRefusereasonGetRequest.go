package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取拒绝换货原因列表 API请求
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

// 初始化TmallExchangeRefusereasonGetRequest对象
func NewTmallExchangeRefusereasonGetRequest() *TmallExchangeRefusereasonGetRequest{
    return &TmallExchangeRefusereasonGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeRefusereasonGetRequest) GetApiMethodName() string {
    return "tmall.exchange.refusereason.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeRefusereasonGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeRefusereasonGetRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeRefusereasonGetRequest) GetDisputeId() int64 {
    return r.disputeId
}
// Fields Setter
// 返回字段
func (r *TmallExchangeRefusereasonGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TmallExchangeRefusereasonGetRequest) GetFields() []string {
    return r.fields
}
// DisputeType Setter
// 换货申请类型：0-任意类型；1-售中；2-售后
func (r *TmallExchangeRefusereasonGetRequest) SetDisputeType(disputeType int64) error {
    r.disputeType = disputeType
    r.Set("dispute_type", disputeType)
    return nil
}

// DisputeType Getter
func (r TmallExchangeRefusereasonGetRequest) GetDisputeType() int64 {
    return r.disputeType
}
