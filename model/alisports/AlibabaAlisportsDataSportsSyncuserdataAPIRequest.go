package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户个人信息同步接口 API请求
alibaba.alisports.data.sports.syncuserdata

阿里体育数据中心用户个人信息同步接口
*/
type AlibabaAlisportsDataSportsSyncuserdataAPIRequest struct {
    model.Params
    // 应用appkey
    _alispAppKey   string
    // 基础代谢率
    _metabolize   string
    // 蛋白质含量
    _protein   string
    // 骨量
    _bone   string
    // 水分率
    _water   string
    // 肌肉率
    _muscle   string
    // 体脂率
    _fat   string
    // 静息心率，单位：次/每分
    _heartRate   int64
    // 体重，单位kg
    _weight   string
    // 身高，单位m
    _height   string
    // 年龄
    _age   int64
    // 三方主键id，唯一标识数据
    _dataId   string
    // 阿里体育用户id
    _aliuid   string
    // 接口签名
    _alispSign   string
    // 时间戳精确到秒
    _alispTime   string
    // 日期 格式：y-m-d h:i:s
    _time   string
}

// 初始化AlibabaAlisportsDataSportsSyncuserdataAPIRequest对象
func NewAlibabaAlisportsDataSportsSyncuserdataRequest() *AlibabaAlisportsDataSportsSyncuserdataAPIRequest{
    return &AlibabaAlisportsDataSportsSyncuserdataAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetApiMethodName() string {
    return "alibaba.alisports.data.sports.syncuserdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// Metabolize Setter
// 基础代谢率
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetMetabolize(_metabolize string) error {
    r._metabolize = _metabolize
    r.Set("metabolize", _metabolize)
    return nil
}

// Metabolize Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetMetabolize() string {
    return r._metabolize
}
// Protein Setter
// 蛋白质含量
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetProtein(_protein string) error {
    r._protein = _protein
    r.Set("protein", _protein)
    return nil
}

// Protein Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetProtein() string {
    return r._protein
}
// Bone Setter
// 骨量
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetBone(_bone string) error {
    r._bone = _bone
    r.Set("bone", _bone)
    return nil
}

// Bone Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetBone() string {
    return r._bone
}
// Water Setter
// 水分率
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetWater(_water string) error {
    r._water = _water
    r.Set("water", _water)
    return nil
}

// Water Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetWater() string {
    return r._water
}
// Muscle Setter
// 肌肉率
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetMuscle(_muscle string) error {
    r._muscle = _muscle
    r.Set("muscle", _muscle)
    return nil
}

// Muscle Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetMuscle() string {
    return r._muscle
}
// Fat Setter
// 体脂率
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetFat(_fat string) error {
    r._fat = _fat
    r.Set("fat", _fat)
    return nil
}

// Fat Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetFat() string {
    return r._fat
}
// HeartRate Setter
// 静息心率，单位：次/每分
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetHeartRate(_heartRate int64) error {
    r._heartRate = _heartRate
    r.Set("heart_rate", _heartRate)
    return nil
}

// HeartRate Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetHeartRate() int64 {
    return r._heartRate
}
// Weight Setter
// 体重，单位kg
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetWeight(_weight string) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetWeight() string {
    return r._weight
}
// Height Setter
// 身高，单位m
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetHeight(_height string) error {
    r._height = _height
    r.Set("height", _height)
    return nil
}

// Height Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetHeight() string {
    return r._height
}
// Age Setter
// 年龄
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetAge(_age int64) error {
    r._age = _age
    r.Set("age", _age)
    return nil
}

// Age Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetAge() int64 {
    return r._age
}
// DataId Setter
// 三方主键id，唯一标识数据
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetDataId(_dataId string) error {
    r._dataId = _dataId
    r.Set("data_id", _dataId)
    return nil
}

// DataId Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetDataId() string {
    return r._dataId
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetAliuid(_aliuid string) error {
    r._aliuid = _aliuid
    r.Set("aliuid", _aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetAliuid() string {
    return r._aliuid
}
// AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetAlispSign() string {
    return r._alispSign
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetAlispTime() string {
    return r._alispTime
}
// Time Setter
// 日期 格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncuserdataAPIRequest) SetTime(_time string) error {
    r._time = _time
    r.Set("time", _time)
    return nil
}

// Time Getter
func (r AlibabaAlisportsDataSportsSyncuserdataAPIRequest) GetTime() string {
    return r._time
}
