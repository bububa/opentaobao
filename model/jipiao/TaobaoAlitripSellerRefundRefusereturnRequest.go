package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】拒绝退票 APIRequest
taobao.alitrip.seller.refund.refusereturn

拒绝退票
*/
type TaobaoAlitripSellerRefundRefusereturnRequest struct {
    model.Params

    // 申请单ID
    applyId   int64 

    // 拒绝理由
    reason   string 

}

func NewTaobaoAlitripSellerRefundRefusereturnRequest() *TaobaoAlitripSellerRefundRefusereturnRequest{
    return &TaobaoAlitripSellerRefundRefusereturnRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripSellerRefundRefusereturnRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refund.refusereturn"
}

func (r TaobaoAlitripSellerRefundRefusereturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripSellerRefundRefusereturnRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripSellerRefundRefusereturnRequest) GetApplyId() int64 {
    return r.applyId
}

func (r *TaobaoAlitripSellerRefundRefusereturnRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r TaobaoAlitripSellerRefundRefusereturnRequest) GetReason() string {
    return r.reason
}

