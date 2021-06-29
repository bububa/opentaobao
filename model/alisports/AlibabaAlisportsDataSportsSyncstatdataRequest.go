package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户当天累积数据同步接口 APIRequest
alibaba.alisports.data.sports.syncstatdata

阿里体育数据中心用户当天累积数据同步接口
*/
type AlibabaAlisportsDataSportsSyncstatdataRequest struct {
    model.Params

    // 应用appkey
    alispAppKey   string 

    // 时间戳精确到秒
    alispTime   string 

    // 签名
    alispSign   string 

    // 阿里体育用户id
    aliuid   string 

    // 运动步数
    steps   string 

    // 消耗卡路里 单位：卡
    calorie   string 

    // 运动距离 单位：米
    distance   string 

    // 日期 格式：y-m-d h:i:s
    time   string 

    // 设备类型：1手环
    deviceType   string 

    // 设备名
    deviceName   string 

    // 设备型号
    deviceModel   string 

}

func NewAlibabaAlisportsDataSportsSyncstatdataRequest() *AlibabaAlisportsDataSportsSyncstatdataRequest{
    return &AlibabaAlisportsDataSportsSyncstatdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetApiMethodName() string {
    return "alibaba.alisports.data.sports.syncstatdata"
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAlispSign() string {
    return r.alispSign
}

func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAliuid() string {
    return r.aliuid
}

func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetSteps(steps string) error {
    r.steps = steps
    r.Set("steps", steps)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetSteps() string {
    return r.steps
}

func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetCalorie(calorie string) error {
    r.calorie = calorie
    r.Set("calorie", calorie)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetCalorie() string {
    return r.calorie
}

func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDistance(distance string) error {
    r.distance = distance
    r.Set("distance", distance)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDistance() string {
    return r.distance
}

func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetTime(time string) error {
    r.time = time
    r.Set("time", time)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetTime() string {
    return r.time
}

func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDeviceType() string {
    return r.deviceType
}

func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDeviceName(deviceName string) error {
    r.deviceName = deviceName
    r.Set("device_name", deviceName)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDeviceName() string {
    return r.deviceName
}

func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDeviceModel() string {
    return r.deviceModel
}

