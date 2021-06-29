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
    _orderId   int64
    // 拒单原因
    _refuseReason   string
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
func (r *TaobaoTradeDrugRefuseorderRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoTradeDrugRefuseorderRequest) GetOrderId() int64 {
    return r._orderId
}
// RefuseReason Setter
// 拒单原因
func (r *TaobaoTradeDrugRefuseorderRequest) SetRefuseReason(_refuseReason string) error {
    r._refuseReason = _refuseReason
    r.Set("refuse_reason", _refuseReason)
    return nil
}

// RefuseReason Getter
func (r TaobaoTradeDrugRefuseorderRequest) GetRefuseReason() string {
    return r._refuseReason
}
