package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】退票申请单详情 APIRequest
taobao.alitrip.seller.refund.get

查询退票申请单详情
*/
type TaobaoAlitripSellerRefundGetRequest struct {
    model.Params

    // 申请单ID
    applyId   int64 

}

func NewTaobaoAlitripSellerRefundGetRequest() *TaobaoAlitripSellerRefundGetRequest{
    return &TaobaoAlitripSellerRefundGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripSellerRefundGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refund.get"
}

func (r TaobaoAlitripSellerRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripSellerRefundGetRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripSellerRefundGetRequest) GetApplyId() int64 {
    return r.applyId
}

