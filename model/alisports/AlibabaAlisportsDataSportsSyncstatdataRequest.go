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
    _alispAppKey   string
    // 时间戳精确到秒
    _alispTime   string
    // 签名
    _alispSign   string
    // 阿里体育用户id
    _aliuid   string
    // 运动步数
    _steps   string
    // 消耗卡路里 单位：卡
    _calorie   string
    // 运动距离 单位：米
    _distance   string
    // 日期 格式：y-m-d h:i:s
    _time   string
    // 设备类型：1手环
    _deviceType   string
    // 设备名
    _deviceName   string
    // 设备型号
    _deviceModel   string
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
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAlispTime() string {
    return r._alispTime
}
// AlispSign Setter
// 签名
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAlispSign() string {
    return r._alispSign
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetAliuid(_aliuid string) error {
    r._aliuid = _aliuid
    r.Set("aliuid", _aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetAliuid() string {
    return r._aliuid
}
// Steps Setter
// 运动步数
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetSteps(_steps string) error {
    r._steps = _steps
    r.Set("steps", _steps)
    return nil
}

// Steps Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetSteps() string {
    return r._steps
}
// Calorie Setter
// 消耗卡路里 单位：卡
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetCalorie(_calorie string) error {
    r._calorie = _calorie
    r.Set("calorie", _calorie)
    return nil
}

// Calorie Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetCalorie() string {
    return r._calorie
}
// Distance Setter
// 运动距离 单位：米
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDistance(_distance string) error {
    r._distance = _distance
    r.Set("distance", _distance)
    return nil
}

// Distance Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDistance() string {
    return r._distance
}
// Time Setter
// 日期 格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetTime(_time string) error {
    r._time = _time
    r.Set("time", _time)
    return nil
}

// Time Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetTime() string {
    return r._time
}
// DeviceType Setter
// 设备类型：1手环
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDeviceType(_deviceType string) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDeviceType() string {
    return r._deviceType
}
// DeviceName Setter
// 设备名
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDeviceName(_deviceName string) error {
    r._deviceName = _deviceName
    r.Set("device_name", _deviceName)
    return nil
}

// DeviceName Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDeviceName() string {
    return r._deviceName
}
// DeviceModel Setter
// 设备型号
func (r *AlibabaAlisportsDataSportsSyncstatdataRequest) SetDeviceModel(_deviceModel string) error {
    r._deviceModel = _deviceModel
    r.Set("device_model", _deviceModel)
    return nil
}

// DeviceModel Getter
func (r AlibabaAlisportsDataSportsSyncstatdataRequest) GetDeviceModel() string {
    return r._deviceModel
}
