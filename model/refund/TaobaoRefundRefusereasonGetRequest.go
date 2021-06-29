package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取拒绝原因列表 API请求
taobao.refund.refusereason.get

获取商家拒绝原因列表
*/
type TaobaoRefundRefusereasonGetRequest struct {
    model.Params
    // 退款编号
    refundId   int64
    // 返回参数
    fields   string
    // 售中或售后
    refundPhase   string
}

// 初始化TaobaoRefundRefusereasonGetRequest对象
func NewTaobaoRefundRefusereasonGetRequest() *TaobaoRefundRefusereasonGetRequest{
    return &TaobaoRefundRefusereasonGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRefundRefusereasonGetRequest) GetApiMethodName() string {
    return "taobao.refund.refusereason.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRefundRefusereasonGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款编号
func (r *TaobaoRefundRefusereasonGetRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRefundRefusereasonGetRequest) GetRefundId() int64 {
    return r.refundId
}
// Fields Setter
// 返回参数
func (r *TaobaoRefundRefusereasonGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoRefundRefusereasonGetRequest) GetFields() string {
    return r.fields
}
// RefundPhase Setter
// 售中或售后
func (r *TaobaoRefundRefusereasonGetRequest) SetRefundPhase(refundPhase string) error {
    r.refundPhase = refundPhase
    r.Set("refund_phase", refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRefundRefusereasonGetRequest) GetRefundPhase() string {
    return r.refundPhase
}
