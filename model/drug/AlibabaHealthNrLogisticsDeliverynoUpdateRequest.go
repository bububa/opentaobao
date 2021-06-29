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
    orderId   int64
    // 快递公司代码
    cpCode   string
    // 快递单号
    courierNo   string
    // 是否强制上传，1代表强制，其他值代表需要进行cp_code合法性校验
    force   int64
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
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateRequest) GetOrderId() int64 {
    return r.orderId
}
// CpCode Setter
// 快递公司代码
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

// CpCode Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateRequest) GetCpCode() string {
    return r.cpCode
}
// CourierNo Setter
// 快递单号
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateRequest) SetCourierNo(courierNo string) error {
    r.courierNo = courierNo
    r.Set("courier_no", courierNo)
    return nil
}

// CourierNo Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateRequest) GetCourierNo() string {
    return r.courierNo
}
// Force Setter
// 是否强制上传，1代表强制，其他值代表需要进行cp_code合法性校验
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateRequest) SetForce(force int64) error {
    r.force = force
    r.Set("force", force)
    return nil
}

// Force Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateRequest) GetForce() int64 {
    return r.force
}
