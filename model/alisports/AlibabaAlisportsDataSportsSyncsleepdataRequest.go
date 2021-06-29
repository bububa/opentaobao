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
    alispAppKey   string
    // 入睡时间，格式：y-m-d h:i:s
    stime   string
    // 清醒时间，单位：小时
    soberTime   string
    // 浅度睡眠时间，单位：小时
    shallowTime   string
    // 深度睡眠时间，单位：小时
    deepTime   string
    // 睡眠总时间，单位：小时
    allTime   string
    // 阿里体育用户id
    aliuid   string
    // 接口签名
    alispSign   string
    // 时间戳精确到秒
    alispTime   string
    // 醒来时间，格式：y-m-d h:i:s
    etime   string
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
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// Stime Setter
// 入睡时间，格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetStime(stime string) error {
    r.stime = stime
    r.Set("stime", stime)
    return nil
}

// Stime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetStime() string {
    return r.stime
}
// SoberTime Setter
// 清醒时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetSoberTime(soberTime string) error {
    r.soberTime = soberTime
    r.Set("sober_time", soberTime)
    return nil
}

// SoberTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetSoberTime() string {
    return r.soberTime
}
// ShallowTime Setter
// 浅度睡眠时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetShallowTime(shallowTime string) error {
    r.shallowTime = shallowTime
    r.Set("shallow_time", shallowTime)
    return nil
}

// ShallowTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetShallowTime() string {
    return r.shallowTime
}
// DeepTime Setter
// 深度睡眠时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetDeepTime(deepTime string) error {
    r.deepTime = deepTime
    r.Set("deep_time", deepTime)
    return nil
}

// DeepTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetDeepTime() string {
    return r.deepTime
}
// AllTime Setter
// 睡眠总时间，单位：小时
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAllTime(allTime string) error {
    r.allTime = allTime
    r.Set("all_time", allTime)
    return nil
}

// AllTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAllTime() string {
    return r.allTime
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAliuid() string {
    return r.aliuid
}
// AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAlispSign() string {
    return r.alispSign
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAlispTime() string {
    return r.alispTime
}
// Etime Setter
// 醒来时间，格式：y-m-d h:i:s
func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetEtime(etime string) error {
    r.etime = etime
    r.Set("etime", etime)
    return nil
}

// Etime Getter
func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetEtime() string {
    return r.etime
}
