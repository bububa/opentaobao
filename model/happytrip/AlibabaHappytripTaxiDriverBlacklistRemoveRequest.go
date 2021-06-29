package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
移除司机黑名单 API请求
alibaba.happytrip.taxi.driver.blacklist.remove

移除司机黑名单
*/
type AlibabaHappytripTaxiDriverBlacklistRemoveRequest struct {
    model.Params
    // 供应商单号
    _orderId   string
    // 用户唯一标识
    _uid   string
    // 司机唯一标识
    _driverId   string
}

// 初始化AlibabaHappytripTaxiDriverBlacklistRemoveRequest对象
func NewAlibabaHappytripTaxiDriverBlacklistRemoveRequest() *AlibabaHappytripTaxiDriverBlacklistRemoveRequest{
    return &AlibabaHappytripTaxiDriverBlacklistRemoveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiDriverBlacklistRemoveRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.driver.blacklist.remove"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiDriverBlacklistRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 供应商单号
func (r *AlibabaHappytripTaxiDriverBlacklistRemoveRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiDriverBlacklistRemoveRequest) GetOrderId() string {
    return r._orderId
}
// Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistRemoveRequest) SetUid(_uid string) error {
    r._uid = _uid
    r.Set("uid", _uid)
    return nil
}

// Uid Getter
func (r AlibabaHappytripTaxiDriverBlacklistRemoveRequest) GetUid() string {
    return r._uid
}
// DriverId Setter
// 司机唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistRemoveRequest) SetDriverId(_driverId string) error {
    r._driverId = _driverId
    r.Set("driver_id", _driverId)
    return nil
}

// DriverId Getter
func (r AlibabaHappytripTaxiDriverBlacklistRemoveRequest) GetDriverId() string {
    return r._driverId
}
