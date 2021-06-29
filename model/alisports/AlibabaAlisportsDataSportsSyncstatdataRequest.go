package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户当天累积数据同步接口 API请求
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

// 初始化AlibabaAlisportsDataSportsSyncstatdataRequest对象
func NewAlibabaAlisportsDataSportsSyncstatdataRequest() *AlibabaAlisportsDataSportsSyncstatdataRequest{
    return &AlibabaAlisportsDataSportsSyncstatdataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetApiMethodName() string {
    return "alibaba.alisports.data.sports.syncstatdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAlispTime() string {
    return r.alispTime
}
// AlispSign Setter
// 签名
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAlispSign() string {
    return r.alispSign
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAliuid() string {
    return r.aliuid
}
// Steps Setter
// 运动步数
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetSteps(steps string) error {
    r.steps = steps
    r.Set("steps", steps)
    return nil
}

// Steps Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetSteps() string {
    return r.steps
}
// Calorie Setter
// 消耗卡路里 单位：卡
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetCalorie(calorie string) error {
    r.calorie = calorie
    r.Set("calorie", calorie)
    return nil
}

// Calorie Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetCalorie() string {
    return r.calorie
}
// Distance Setter
// 运动距离 单位：米
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDistance(distance string) error {
    r.distance = distance
    r.Set("distance", distance)
    return nil
}

// Distance Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDistance() string {
    return r.distance
}
// Time Setter
// 日期 格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetTime(time string) error {
    r.time = time
    r.Set("time", time)
    return nil
}

// Time Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetTime() string {
    return r.time
}
// DeviceType Setter
// 设备类型：1手环
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDeviceType() string {
    return r.deviceType
}
// DeviceName Setter
// 设备名
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDeviceName(deviceName string) error {
    r.deviceName = deviceName
    r.Set("device_name", deviceName)
    return nil
}

// DeviceName Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDeviceName() string {
    return r.deviceName
}
// DeviceModel Setter
// 设备型号
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

// DeviceModel Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDeviceModel() string {
    return r.deviceModel
}
