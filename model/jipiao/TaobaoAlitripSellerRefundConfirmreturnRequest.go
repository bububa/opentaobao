package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】确认退票 APIRequest
taobao.alitrip.seller.refund.confirmreturn

确认退票
*/
type TaobaoAlitripSellerRefundConfirmreturnRequest struct {
    model.Params

    // 退票申请单
    applyId   int64 

}

func NewTaobaoAlitripSellerRefundConfirmreturnRequest() *TaobaoAlitripSellerRefundConfirmreturnRequest{
    return &TaobaoAlitripSellerRefundConfirmreturnRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripSellerRefundConfirmreturnRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refund.confirmreturn"
}

func (r TaobaoAlitripSellerRefundConfirmreturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripSellerRefundConfirmreturnRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripSellerRefundConfirmreturnRequest) GetApplyId() int64 {
    return r.applyId
}

