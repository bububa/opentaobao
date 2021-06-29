package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔退款详情 API请求
taobao.refund.get

获取单笔退款详情
*/
type TaobaoRefundGetRequest struct {
    model.Params
    // 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
    _fields   []string
    // 退款单号
    _refundId   int64
}

// 初始化TaobaoRefundGetRequest对象
func NewTaobaoRefundGetRequest() *TaobaoRefundGetRequest{
    return &TaobaoRefundGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRefundGetRequest) GetApiMethodName() string {
    return "taobao.refund.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
func (r *TaobaoRefundGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoRefundGetRequest) GetFields() []string {
    return r._fields
}
// RefundId Setter
// 退款单号
func (r *TaobaoRefundGetRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRefundGetRequest) GetRefundId() int64 {
    return r._refundId
}
