package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
移除司机黑名单 APIRequest
alibaba.happytrip.taxi.driver.blacklist.remove

移除司机黑名单
*/
type AlibabaHappytripTaxiDriverBlacklistRemoveRequest struct {
    model.Params

    // 供应商单号
    orderId   string 

    // 用户唯一标识
    uid   string 

    // 司机唯一标识
    driverId   string 

}

func NewAlibabaHappytripTaxiDriverBlacklistRemoveRequest() *AlibabaHappytripTaxiDriverBlacklistRemoveRequest{
    return &AlibabaHappytripTaxiDriverBlacklistRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiDriverBlacklistRemoveRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.driver.blacklist.remove"
}

func (r AlibabaHappytripTaxiDriverBlacklistRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiDriverBlacklistRemoveRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiDriverBlacklistRemoveRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaHappytripTaxiDriverBlacklistRemoveRequest) SetUid(uid string) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

func (r AlibabaHappytripTaxiDriverBlacklistRemoveRequest) GetUid() string {
    return r.uid
}

func (r *AlibabaHappytripTaxiDriverBlacklistRemoveRequest) SetDriverId(driverId string) error {
    r.driverId = driverId
    r.Set("driver_id", driverId)
    return nil
}

func (r AlibabaHappytripTaxiDriverBlacklistRemoveRequest) GetDriverId() string {
    return r.driverId
}

