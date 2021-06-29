package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加司机黑名单 API请求
alibaba.happytrip.taxi.driver.blacklist.add

实现用户1对1永久拉黑司机，如果不支持永久拉黑，则在自动解禁黑名单司机时需回调通知欢行
*/
type AlibabaHappytripTaxiDriverBlacklistAddRequest struct {
    model.Params
    // 供应商单号
    _orderId   string
    // 用户唯一标识
    _uid   string
    // 司机唯一标识
    _driverId   string
}

// 初始化AlibabaHappytripTaxiDriverBlacklistAddRequest对象
func NewAlibabaHappytripTaxiDriverBlacklistAddRequest() *AlibabaHappytripTaxiDriverBlacklistAddRequest{
    return &AlibabaHappytripTaxiDriverBlacklistAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.driver.blacklist.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 供应商单号
func (r *AlibabaHappytripTaxiDriverBlacklistAddRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetOrderId() string {
    return r._orderId
}
// Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistAddRequest) SetUid(_uid string) error {
    r._uid = _uid
    r.Set("uid", _uid)
    return nil
}

// Uid Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetUid() string {
    return r._uid
}
// DriverId Setter
// 司机唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistAddRequest) SetDriverId(_driverId string) error {
    r._driverId = _driverId
    r.Set("driver_id", _driverId)
    return nil
}

// DriverId Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetDriverId() string {
    return r._driverId
}
