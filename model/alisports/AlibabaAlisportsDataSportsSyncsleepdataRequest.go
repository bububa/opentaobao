package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户睡眠数据同步接口 APIRequest
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

func NewAlibabaAlisportsDataSportsSyncsleepdataRequest() *AlibabaAlisportsDataSportsSyncsleepdataRequest{
    return &AlibabaAlisportsDataSportsSyncsleepdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetApiMethodName() string {
    return "alibaba.alisports.data.sports.syncsleepdata"
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetStime(stime string) error {
    r.stime = stime
    r.Set("stime", stime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetStime() string {
    return r.stime
}

func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetSoberTime(soberTime string) error {
    r.soberTime = soberTime
    r.Set("sober_time", soberTime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetSoberTime() string {
    return r.soberTime
}

func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetShallowTime(shallowTime string) error {
    r.shallowTime = shallowTime
    r.Set("shallow_time", shallowTime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetShallowTime() string {
    return r.shallowTime
}

func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetDeepTime(deepTime string) error {
    r.deepTime = deepTime
    r.Set("deep_time", deepTime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetDeepTime() string {
    return r.deepTime
}

func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAllTime(allTime string) error {
    r.allTime = allTime
    r.Set("all_time", allTime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAllTime() string {
    return r.allTime
}

func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAliuid() string {
    return r.aliuid
}

func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAlispSign() string {
    return r.alispSign
}

func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsDataSportsSyncsleepdataRequest) SetEtime(etime string) error {
    r.etime = etime
    r.Set("etime", etime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncsleepdataRequest) GetEtime() string {
    return r.etime
}

