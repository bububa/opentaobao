package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加司机黑名单 APIRequest
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

func NewAlibabaHappytripTaxiDriverBlacklistAddRequest() *AlibabaHappytripTaxiDriverBlacklistAddRequest{
    return &AlibabaHappytripTaxiDriverBlacklistAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.driver.blacklist.add"
}

func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiDriverBlacklistAddRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaHappytripTaxiDriverBlacklistAddRequest) SetUid(uid string) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetUid() string {
    return r.uid
}

func (r *AlibabaHappytripTaxiDriverBlacklistAddRequest) SetDriverId(driverId string) error {
    r.driverId = driverId
    r.Set("driver_id", driverId)
    return nil
}

func (r AlibabaHappytripTaxiDriverBlacklistAddRequest) GetDriverId() string {
    return r.driverId
}

