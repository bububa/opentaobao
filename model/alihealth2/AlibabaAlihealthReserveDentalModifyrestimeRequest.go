package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改预约时间 API请求
alibaba.alihealth.reserve.dental.modifyrestime

修改预约时间
*/
type AlibabaAlihealthReserveDentalModifyrestimeRequest struct {
    model.Params
    // 预约单ID
    _reserveId   int64
    // 预约时间
    _reserveTime   string
}

// 初始化AlibabaAlihealthReserveDentalModifyrestimeRequest对象
func NewAlibabaAlihealthReserveDentalModifyrestimeRequest() *AlibabaAlihealthReserveDentalModifyrestimeRequest{
    return &AlibabaAlihealthReserveDentalModifyrestimeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalModifyrestimeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reserve.dental.modifyrestime"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalModifyrestimeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReserveId Setter
// 预约单ID
func (r *AlibabaAlihealthReserveDentalModifyrestimeRequest) SetReserveId(_reserveId int64) error {
    r._reserveId = _reserveId
    r.Set("reserve_id", _reserveId)
    return nil
}

// ReserveId Getter
func (r AlibabaAlihealthReserveDentalModifyrestimeRequest) GetReserveId() int64 {
    return r._reserveId
}
// ReserveTime Setter
// 预约时间
func (r *AlibabaAlihealthReserveDentalModifyrestimeRequest) SetReserveTime(_reserveTime string) error {
    r._reserveTime = _reserveTime
    r.Set("reserve_time", _reserveTime)
    return nil
}

// ReserveTime Getter
func (r AlibabaAlihealthReserveDentalModifyrestimeRequest) GetReserveTime() string {
    return r._reserveTime
}
