package wdklogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自提业务-车辆到达上报车牌号 APIRequest
alibaba.wdk.logistics.pus.pickup.cararrived

自提业务-汽车自提,车辆到达上报车牌号
*/
type AlibabaWdkLogisticsPusPickupCararrivedRequest struct {
    model.Params

    // 自提点
    stationCode   string 

    // 车牌号
    carNum   string 

}

func NewAlibabaWdkLogisticsPusPickupCararrivedRequest() *AlibabaWdkLogisticsPusPickupCararrivedRequest{
    return &AlibabaWdkLogisticsPusPickupCararrivedRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkLogisticsPusPickupCararrivedRequest) GetApiMethodName() string {
    return "alibaba.wdk.logistics.pus.pickup.cararrived"
}

func (r AlibabaWdkLogisticsPusPickupCararrivedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkLogisticsPusPickupCararrivedRequest) SetStationCode(stationCode string) error {
    r.stationCode = stationCode
    r.Set("station_code", stationCode)
    return nil
}

func (r AlibabaWdkLogisticsPusPickupCararrivedRequest) GetStationCode() string {
    return r.stationCode
}

func (r *AlibabaWdkLogisticsPusPickupCararrivedRequest) SetCarNum(carNum string) error {
    r.carNum = carNum
    r.Set("car_num", carNum)
    return nil
}

func (r AlibabaWdkLogisticsPusPickupCararrivedRequest) GetCarNum() string {
    return r.carNum
}

