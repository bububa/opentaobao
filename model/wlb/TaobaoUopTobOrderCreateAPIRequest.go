package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ToB仓储发货 API请求
taobao.uop.tob.order.create

ToB仓储发货
*/
type TaobaoUopTobOrderCreateAPIRequest struct {
    model.Params
    // ERP出库对象
    _deliveryOrder   *DeliveryOrder
}

// 初始化TaobaoUopTobOrderCreateAPIRequest对象
func NewTaobaoUopTobOrderCreateRequest() *TaobaoUopTobOrderCreateAPIRequest{
    return &TaobaoUopTobOrderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUopTobOrderCreateAPIRequest) GetApiMethodName() string {
    return "taobao.uop.tob.order.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUopTobOrderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeliveryOrder Setter
// ERP出库对象
func (r *TaobaoUopTobOrderCreateAPIRequest) SetDeliveryOrder(_deliveryOrder *DeliveryOrder) error {
    r._deliveryOrder = _deliveryOrder
    r.Set("delivery_order", _deliveryOrder)
    return nil
}

// DeliveryOrder Getter
func (r TaobaoUopTobOrderCreateAPIRequest) GetDeliveryOrder() *DeliveryOrder {
    return r._deliveryOrder
}
