package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单查询接口 API请求
alibaba.health.nr.logistics.waybill.get

商家登录后根据订单号查询物流单号及电子面单信息
*/
type AlibabaHealthNrLogisticsWaybillGetRequest struct {
    model.Params
    // 订单id
    orderId   int64
}

// 初始化AlibabaHealthNrLogisticsWaybillGetRequest对象
func NewAlibabaHealthNrLogisticsWaybillGetRequest() *AlibabaHealthNrLogisticsWaybillGetRequest{
    return &AlibabaHealthNrLogisticsWaybillGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrLogisticsWaybillGetRequest) GetApiMethodName() string {
    return "alibaba.health.nr.logistics.waybill.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrLogisticsWaybillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaHealthNrLogisticsWaybillGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthNrLogisticsWaybillGetRequest) GetOrderId() int64 {
    return r.orderId
}
