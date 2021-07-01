package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AG退货入仓状态写接口 API请求
taobao.nextone.logistics.warehouse.update

商家上传退货入仓状态给ag
*/
type TaobaoNextoneLogisticsWarehouseUpdateAPIRequest struct {
    model.Params
    // 退款编号
    _refundId   int64
    // 退货入仓状态 1.已入仓
    _warehouseStatus   int64
}

// 初始化TaobaoNextoneLogisticsWarehouseUpdateAPIRequest对象
func NewTaobaoNextoneLogisticsWarehouseUpdateRequest() *TaobaoNextoneLogisticsWarehouseUpdateAPIRequest{
    return &TaobaoNextoneLogisticsWarehouseUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoNextoneLogisticsWarehouseUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.nextone.logistics.warehouse.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoNextoneLogisticsWarehouseUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款编号
func (r *TaobaoNextoneLogisticsWarehouseUpdateAPIRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoNextoneLogisticsWarehouseUpdateAPIRequest) GetRefundId() int64 {
    return r._refundId
}
// WarehouseStatus Setter
// 退货入仓状态 1.已入仓
func (r *TaobaoNextoneLogisticsWarehouseUpdateAPIRequest) SetWarehouseStatus(_warehouseStatus int64) error {
    r._warehouseStatus = _warehouseStatus
    r.Set("warehouse_status", _warehouseStatus)
    return nil
}

// WarehouseStatus Getter
func (r TaobaoNextoneLogisticsWarehouseUpdateAPIRequest) GetWarehouseStatus() int64 {
    return r._warehouseStatus
}
