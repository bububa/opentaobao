package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
特殊部分退纠纷单查询 API请求
taobao.special.refund.get

获取单笔特殊部分退的纠纷单查询
*/
type TaobaoSpecialRefundGetRequest struct {
    model.Params
    // 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
    _fields   []string
    // 退款单号
    _refundId   int64
}

// 初始化TaobaoSpecialRefundGetRequest对象
func NewTaobaoSpecialRefundGetRequest() *TaobaoSpecialRefundGetRequest{
    return &TaobaoSpecialRefundGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSpecialRefundGetRequest) GetApiMethodName() string {
    return "taobao.special.refund.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSpecialRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
func (r *TaobaoSpecialRefundGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoSpecialRefundGetRequest) GetFields() []string {
    return r._fields
}
// RefundId Setter
// 退款单号
func (r *TaobaoSpecialRefundGetRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoSpecialRefundGetRequest) GetRefundId() int64 {
    return r._refundId
}
