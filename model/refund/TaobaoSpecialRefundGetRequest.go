package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
特殊部分退纠纷单查询 APIRequest
taobao.special.refund.get

获取单笔特殊部分退的纠纷单查询
*/
type TaobaoSpecialRefundGetRequest struct {
    model.Params

    // 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
    fields   []String 

    // 退款单号
    refundId   int64 

}

func NewTaobaoSpecialRefundGetRequest() *TaobaoSpecialRefundGetRequest{
    return &TaobaoSpecialRefundGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSpecialRefundGetRequest) GetApiMethodName() string {
    return "taobao.special.refund.get"
}

func (r TaobaoSpecialRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSpecialRefundGetRequest) SetFields(fields []String) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoSpecialRefundGetRequest) GetFields() []String {
    return r.fields
}

func (r *TaobaoSpecialRefundGetRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoSpecialRefundGetRequest) GetRefundId() int64 {
    return r.refundId
}

