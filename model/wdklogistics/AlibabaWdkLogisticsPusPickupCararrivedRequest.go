package wdklogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自提业务-车辆到达上报车牌号 API请求
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

// 初始化AlibabaWdkLogisticsPusPickupCararrivedRequest对象
func NewAlibabaWdkLogisticsPusPickupCararrivedRequest() *AlibabaWdkLogisticsPusPickupCararrivedRequest{
    return &AlibabaWdkLogisticsPusPickupCararrivedRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkLogisticsPusPickupCararrivedRequest) GetApiMethodName() string {
    return "alibaba.wdk.logistics.pus.pickup.cararrived"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkLogisticsPusPickupCararrivedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StationCode Setter
// 自提点
func (r *AlibabaWdkLogisticsPusPickupCararrivedRequest) SetStationCode(stationCode string) error {
    r.stationCode = stationCode
    r.Set("station_code", stationCode)
    return nil
}

// StationCode Getter
func (r AlibabaWdkLogisticsPusPickupCararrivedRequest) GetStationCode() string {
    return r.stationCode
}
// CarNum Setter
// 车牌号
func (r *AlibabaWdkLogisticsPusPickupCararrivedRequest) SetCarNum(carNum string) error {
    r.carNum = carNum
    r.Set("car_num", carNum)
    return nil
}

// CarNum Getter
func (r AlibabaWdkLogisticsPusPickupCararrivedRequest) GetCarNum() string {
    return r.carNum
}
