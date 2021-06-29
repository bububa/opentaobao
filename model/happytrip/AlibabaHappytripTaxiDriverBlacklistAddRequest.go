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
    orderId   string
    // 用户唯一标识
    uid   string
    // 司机唯一标识
    driverId   string
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
func (r *AlibabaHappytripTaxiDriverBlacklistAddRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetOrderId() string {
    return r.orderId
}
// Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistAddRequest) SetUid(uid string) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

// Uid Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetUid() string {
    return r.uid
}
// DriverId Setter
// 司机唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistAddRequest) SetDriverId(driverId string) error {
    r.driverId = driverId
    r.Set("driver_id", driverId)
    return nil
}

// DriverId Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetDriverId() string {
    return r.driverId
}
