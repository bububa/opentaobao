package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户睡眠数据同步接口 API请求
alibaba.alisports.data.sports.syncsleepdata

阿里体育数据中心用户睡眠数据同步接口
*/
type AlibabaAlisportsDataSportsSyncsleepdataRequest struct {
    model.Params
    // 应用appkey
    _alispAppKey   string
    // 入睡时间，格式：y-m-d h:i:s
    _stime   string
    // 清醒时间，单位：小时
    _soberTime   string
    // 浅度睡眠时间，单位：小时
    _shallowTime   string
    // 深度睡眠时间，单位：小时
    _deepTime   string
    // 睡眠总时间，单位：小时
    _allTime   string
    // 阿里体育用户id
    _aliuid   string
    // 接口签名
    _alispSign   string
    // 时间戳精确到秒
    _alispTime   string
    // 醒来时间，格式：y-m-d h:i:s
    _etime   string
}

// 初始化AlibabaAlisportsDataSportsSyncsleepdataRequest对象
func NewAlibabaAlisportsDataSportsSyncsleepdataRequest() *AlibabaAlisportsDataSportsSyncsleepdataRequest{
    return &AlibabaAlisportsDataSportsSyncsleepdataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetApiMethodName() string {
    return "alibaba.alisports.data.sports.syncsleepdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// Stime Setter
// 入睡时间，格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetStime(_stime string) error {
    r._stime = _stime
    r.Set("stime", _stime)
    return nil
}

// Stime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetStime() string {
    return r._stime
}
// SoberTime Setter
// 清醒时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetSoberTime(_soberTime string) error {
    r._soberTime = _soberTime
    r.Set("sober_time", _soberTime)
    return nil
}

// SoberTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetSoberTime() string {
    return r._soberTime
}
// ShallowTime Setter
// 浅度睡眠时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetShallowTime(_shallowTime string) error {
    r._shallowTime = _shallowTime
    r.Set("shallow_time", _shallowTime)
    return nil
}

// ShallowTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetShallowTime() string {
    return r._shallowTime
}
// DeepTime Setter
// 深度睡眠时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetDeepTime(_deepTime string) error {
    r._deepTime = _deepTime
    r.Set("deep_time", _deepTime)
    return nil
}

// DeepTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetDeepTime() string {
    return r._deepTime
}
// AllTime Setter
// 睡眠总时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAllTime(_allTime string) error {
    r._allTime = _allTime
    r.Set("all_time", _allTime)
    return nil
}

// AllTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAllTime() string {
    return r._allTime
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAliuid(_aliuid string) error {
    r._aliuid = _aliuid
    r.Set("aliuid", _aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAliuid() string {
    return r._aliuid
}
// AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAlispSign() string {
    return r._alispSign
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAlispTime() string {
    return r._alispTime
}
// Etime Setter
// 醒来时间，格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetEtime(_etime string) error {
    r._etime = _etime
    r.Set("etime", _etime)
    return nil
}

// Etime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetEtime() string {
    return r._etime
}
