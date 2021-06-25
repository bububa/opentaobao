package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AG物流签收状态写接口 APIRequest
taobao.nextone.logistics.sign.update

商家上传退货的签收状态给AG
*/
type TaobaoNextoneLogisticsSignUpdateRequest struct {
    model.Params

    // 退款编号
    refundId   int64 

    // 货物签收状态
    signStatus   int64 

}

func NewTaobaoNextoneLogisticsSignUpdateRequest() *TaobaoNextoneLogisticsSignUpdateRequest{
    return &TaobaoNextoneLogisticsSignUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoNextoneLogisticsSignUpdateRequest) GetApiMethodName() string {
    return "taobao.nextone.logistics.sign.update"
}

func (r TaobaoNextoneLogisticsSignUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoNextoneLogisticsSignUpdateRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoNextoneLogisticsSignUpdateRequest) GetRefundId() int64 {
    return r.refundId
}

func (r *TaobaoNextoneLogisticsSignUpdateRequest) SetSignStatus(signStatus int64) error {
    r.signStatus = signStatus
    r.Set("sign_status", signStatus)
    return nil
}

func (r TaobaoNextoneLogisticsSignUpdateRequest) GetSignStatus() int64 {
    return r.signStatus
}

