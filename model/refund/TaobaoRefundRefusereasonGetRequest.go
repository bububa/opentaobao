package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取拒绝原因列表 APIRequest
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

func NewTaobaoRefundRefusereasonGetRequest() *TaobaoRefundRefusereasonGetRequest{
    return &TaobaoRefundRefusereasonGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRefundRefusereasonGetRequest) GetApiMethodName() string {
    return "taobao.refund.refusereason.get"
}

func (r TaobaoRefundRefusereasonGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRefundRefusereasonGetRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoRefundRefusereasonGetRequest) GetRefundId() int64 {
    return r.refundId
}

func (r *TaobaoRefundRefusereasonGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoRefundRefusereasonGetRequest) GetFields() string {
    return r.fields
}

func (r *TaobaoRefundRefusereasonGetRequest) SetRefundPhase(refundPhase string) error {
    r.refundPhase = refundPhase
    r.Set("refund_phase", refundPhase)
    return nil
}

func (r TaobaoRefundRefusereasonGetRequest) GetRefundPhase() string {
    return r.refundPhase
}

