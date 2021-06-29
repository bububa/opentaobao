package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取OpenMall退款单详情 APIRequest
taobao.openmall.refund.get

获取OpenMall退款单详情
*/
type TaobaoOpenmallRefundGetRequest struct {
    model.Params

    // 渠道商身份
    distributor   string 

    // 退款单ID
    refundId   int64 

}

func NewTaobaoOpenmallRefundGetRequest() *TaobaoOpenmallRefundGetRequest{
    return &TaobaoOpenmallRefundGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallRefundGetRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.get"
}

func (r TaobaoOpenmallRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallRefundGetRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallRefundGetRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallRefundGetRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoOpenmallRefundGetRequest) GetRefundId() int64 {
    return r.refundId
}

