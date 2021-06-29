package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康020拒单 APIRequest
taobao.trade.drug.refuseorder

阿里健康020拒单
*/
type TaobaoTradeDrugRefuseorderRequest struct {
    model.Params

    // 订单ID
    orderId   int64 

    // 拒单原因
    refuseReason   string 

}

func NewTaobaoTradeDrugRefuseorderRequest() *TaobaoTradeDrugRefuseorderRequest{
    return &TaobaoTradeDrugRefuseorderRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeDrugRefuseorderRequest) GetApiMethodName() string {
    return "taobao.trade.drug.refuseorder"
}

func (r TaobaoTradeDrugRefuseorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeDrugRefuseorderRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoTradeDrugRefuseorderRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoTradeDrugRefuseorderRequest) SetRefuseReason(refuseReason string) error {
    r.refuseReason = refuseReason
    r.Set("refuse_reason", refuseReason)
    return nil
}

func (r TaobaoTradeDrugRefuseorderRequest) GetRefuseReason() string {
    return r.refuseReason
}

