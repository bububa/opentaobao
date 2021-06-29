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
    orderCode   string
    // 是否已经退款
    refunded   bool
    // 退款原因
    refundReason   string
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
func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetOrderCode() string {
    return r.orderCode
}
// Refunded Setter
// 是否已经退款
func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) SetRefunded(refunded bool) error {
    r.refunded = refunded
    r.Set("refunded", refunded)
    return nil
}

// Refunded Getter
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetRefunded() bool {
    return r.refunded
}
// RefundReason Setter
// 退款原因
func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) SetRefundReason(refundReason string) error {
    r.refundReason = refundReason
    r.Set("refund_reason", refundReason)
    return nil
}

// RefundReason Getter
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetRefundReason() string {
    return r.refundReason
}
