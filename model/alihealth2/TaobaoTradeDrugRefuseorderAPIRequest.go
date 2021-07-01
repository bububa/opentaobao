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
type TaobaoTradeDrugRefuseorderAPIRequest struct {
    model.Params
    // 订单ID
    _orderId   int64
    // 拒单原因
    _refuseReason   string
}

// 初始化TaobaoTradeDrugRefuseorderAPIRequest对象
func NewTaobaoTradeDrugRefuseorderRequest() *TaobaoTradeDrugRefuseorderAPIRequest{
    return &TaobaoTradeDrugRefuseorderAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetApiMethodName() string {
    return "taobao.trade.drug.refuseorder"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *TaobaoTradeDrugRefuseorderAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// RefuseReason Setter
// 拒单原因
func (r *TaobaoTradeDrugRefuseorderAPIRequest) SetRefuseReason(_refuseReason string) error {
    r._refuseReason = _refuseReason
    r.Set("refuse_reason", _refuseReason)
    return nil
}

// RefuseReason Getter
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetRefuseReason() string {
    return r._refuseReason
}
