package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关闭OpenMall退款单 APIRequest
taobao.openmall.refund.close

关闭OpenMall退款单
*/
type TaobaoOpenmallRefundCloseRequest struct {
    model.Params

    // 渠道
    distributor   string 

    // 退款ID
    refundId   int64 

}

func NewTaobaoOpenmallRefundCloseRequest() *TaobaoOpenmallRefundCloseRequest{
    return &TaobaoOpenmallRefundCloseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallRefundCloseRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.close"
}

func (r TaobaoOpenmallRefundCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallRefundCloseRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallRefundCloseRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallRefundCloseRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoOpenmallRefundCloseRequest) GetRefundId() int64 {
    return r.refundId
}

