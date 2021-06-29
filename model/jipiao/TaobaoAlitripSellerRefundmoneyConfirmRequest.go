package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】确认退款 APIRequest
taobao.alitrip.seller.refundmoney.confirm

代理人确认退票申请单的退款
*/
type TaobaoAlitripSellerRefundmoneyConfirmRequest struct {
    model.Params

    // 申请单id
    applyId   int64 

}

func NewTaobaoAlitripSellerRefundmoneyConfirmRequest() *TaobaoAlitripSellerRefundmoneyConfirmRequest{
    return &TaobaoAlitripSellerRefundmoneyConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripSellerRefundmoneyConfirmRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refundmoney.confirm"
}

func (r TaobaoAlitripSellerRefundmoneyConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripSellerRefundmoneyConfirmRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripSellerRefundmoneyConfirmRequest) GetApplyId() int64 {
    return r.applyId
}

