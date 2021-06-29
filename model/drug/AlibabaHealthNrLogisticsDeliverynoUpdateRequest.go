package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传订单同城快递单号 API请求
alibaba.health.nr.logistics.deliveryno.update

上传订单同城快递单号
*/
type AlibabaHealthNrLogisticsDeliverynoUpdateRequest struct {
    model.Params
    // 订单ID
    _orderId   int64
    // 快递公司代码
    _cpCode   string
    // 快递单号
    _courierNo   string
    // 是否强制上传，1代表强制，其他值代表需要进行cp_code合法性校验
    _force   int64
}

// 初始化AlibabaHealthNrLogisticsDeliverynoUpdateRequest对象
func NewAlibabaHealthNrLogisticsDeliverynoUpdateRequest() *AlibabaHealthNrLogisticsDeliverynoUpdateRequest{
    return &AlibabaHealthNrLogisticsDeliverynoUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrLogisticsDeliverynoUpdateRequest) GetApiMethodName() string {
    return "alibaba.health.nr.logistics.deliveryno.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrLogisticsDeliverynoUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateRequest) GetOrderId() int64 {
    return r._orderId
}
// CpCode Setter
// 快递公司代码
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateRequest) SetCpCode(_cpCode string) error {
    r._cpCode = _cpCode
    r.Set("cp_code", _cpCode)
    return nil
}

// CpCode Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateRequest) GetCpCode() string {
    return r._cpCode
}
// CourierNo Setter
// 快递单号
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateRequest) SetCourierNo(_courierNo string) error {
    r._courierNo = _courierNo
    r.Set("courier_no", _courierNo)
    return nil
}

// CourierNo Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateRequest) GetCourierNo() string {
    return r._courierNo
}
// Force Setter
// 是否强制上传，1代表强制，其他值代表需要进行cp_code合法性校验
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateRequest) SetForce(_force int64) error {
    r._force = _force
    r.Set("force", _force)
    return nil
}

// Force Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateRequest) GetForce() int64 {
    return r._force
}
