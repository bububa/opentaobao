package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改预约 API请求
alibaba.alihealth.booking.reserve.modify

消费医疗统一预约平台，取消预约
*/
type AlibabaAlihealthBookingReserveModifyRequest struct {
    model.Params
    // modify
    modify   *IsvReserveRequest
}

// 初始化AlibabaAlihealthBookingReserveModifyRequest对象
func NewAlibabaAlihealthBookingReserveModifyRequest() *AlibabaAlihealthBookingReserveModifyRequest{
    return &AlibabaAlihealthBookingReserveModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveModifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Modify Setter
// modify
func (r *AlibabaAlihealthBookingReserveModifyRequest) SetModify(modify *IsvReserveRequest) error {
    r.modify = modify
    r.Set("modify", modify)
    return nil
}

// Modify Getter
func (r AlibabaAlihealthBookingReserveModifyRequest) GetModify() *IsvReserveRequest {
    return r.modify
}
