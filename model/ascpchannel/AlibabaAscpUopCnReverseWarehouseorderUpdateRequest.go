package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链中台逆向入库单修改服务 API请求
alibaba.ascp.uop.cn.reverse.warehouseorder.update

供应链中台逆向入库单修改服务
*/
type AlibabaAscpUopCnReverseWarehouseorderUpdateRequest struct {
    model.Params
    // 逆向入库单号
    _orderCode   string
    // 是否已经退款
    _refunded   bool
    // 退款原因
    _refundReason   string
}

// 初始化AlibabaAscpUopCnReverseWarehouseorderUpdateRequest对象
func NewAlibabaAscpUopCnReverseWarehouseorderUpdateRequest() *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest{
    return &AlibabaAscpUopCnReverseWarehouseorderUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.cn.reverse.warehouseorder.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 逆向入库单号
func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetOrderCode() string {
    return r._orderCode
}
// Refunded Setter
// 是否已经退款
func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) SetRefunded(_refunded bool) error {
    r._refunded = _refunded
    r.Set("refunded", _refunded)
    return nil
}

// Refunded Getter
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetRefunded() bool {
    return r._refunded
}
// RefundReason Setter
// 退款原因
func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) SetRefundReason(_refundReason string) error {
    r._refundReason = _refundReason
    r.Set("refund_reason", _refundReason)
    return nil
}

// RefundReason Getter
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetRefundReason() string {
    return r._refundReason
}
