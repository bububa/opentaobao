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
    _refundId   int64
    // 返回参数
    _fields   string
    // 售中或售后
    _refundPhase   string
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
func (r *TaobaoRefundRefusereasonGetRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRefundRefusereasonGetRequest) GetRefundId() int64 {
    return r._refundId
}
// Fields Setter
// 返回参数
func (r *TaobaoRefundRefusereasonGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoRefundRefusereasonGetRequest) GetFields() string {
    return r._fields
}
// RefundPhase Setter
// 售中或售后
func (r *TaobaoRefundRefusereasonGetRequest) SetRefundPhase(_refundPhase string) error {
    r._refundPhase = _refundPhase
    r.Set("refund_phase", _refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRefundRefusereasonGetRequest) GetRefundPhase() string {
    return r._refundPhase
}
