package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康020拒单 API请求
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

// 初始化TaobaoTradeDrugRefuseorderRequest对象
func NewTaobaoTradeDrugRefuseorderRequest() *TaobaoTradeDrugRefuseorderRequest{
    return &TaobaoTradeDrugRefuseorderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugRefuseorderRequest) GetApiMethodName() string {
    return "taobao.trade.drug.refuseorder"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugRefuseorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *TaobaoTradeDrugRefuseorderRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoTradeDrugRefuseorderRequest) GetOrderId() int64 {
    return r.orderId
}
// RefuseReason Setter
// 拒单原因
func (r *TaobaoTradeDrugRefuseorderRequest) SetRefuseReason(refuseReason string) error {
    r.refuseReason = refuseReason
    r.Set("refuse_reason", refuseReason)
    return nil
}

// RefuseReason Getter
func (r TaobaoTradeDrugRefuseorderRequest) GetRefuseReason() string {
    return r.refuseReason
}
