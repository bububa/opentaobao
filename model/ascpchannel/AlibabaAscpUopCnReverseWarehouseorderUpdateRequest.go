package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链中台逆向入库单修改服务 APIRequest
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

func NewAlibabaAscpUopCnReverseWarehouseorderUpdateRequest() *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest{
    return &AlibabaAscpUopCnReverseWarehouseorderUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.cn.reverse.warehouseorder.update"
}

func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) SetRefunded(refunded bool) error {
    r.refunded = refunded
    r.Set("refunded", refunded)
    return nil
}

func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetRefunded() bool {
    return r.refunded
}

func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) SetRefundReason(refundReason string) error {
    r.refundReason = refundReason
    r.Set("refund_reason", refundReason)
    return nil
}

func (r AlibabaAscpUopCnReverseWarehouseorderUpdateRequest) GetRefundReason() string {
    return r.refundReason
}

