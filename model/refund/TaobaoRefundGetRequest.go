package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔退款详情 APIRequest
taobao.refund.get

获取单笔退款详情
*/
type TaobaoRefundGetRequest struct {
    model.Params

    // 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
    fields   []String 

    // 退款单号
    refundId   int64 

}

func NewTaobaoRefundGetRequest() *TaobaoRefundGetRequest{
    return &TaobaoRefundGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRefundGetRequest) GetApiMethodName() string {
    return "taobao.refund.get"
}

func (r TaobaoRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRefundGetRequest) SetFields(fields []String) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoRefundGetRequest) GetFields() []String {
    return r.fields
}

func (r *TaobaoRefundGetRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoRefundGetRequest) GetRefundId() int64 {
    return r.refundId
}

