package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户每日摄入卡路里总量回传接口 API请求
alibaba.alihealth.alipaypfm.diet.record

用户每日摄入卡路里总量回传接口
*/
type AlibabaAlihealthAlipaypfmDietRecordRequest struct {
    model.Params
    // 用户健康ID
    userId   int64
    // 记录日期，format：yyyy-MM-dd
    date   string
    // 累积摄入卡路里
    energy   int64
}

// 初始化AlibabaAlihealthAlipaypfmDietRecordRequest对象
func NewAlibabaAlihealthAlipaypfmDietRecordRequest() *AlibabaAlihealthAlipaypfmDietRecordRequest{
    return &AlibabaAlihealthAlipaypfmDietRecordRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlipaypfmDietRecordRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.diet.record"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlipaypfmDietRecordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户健康ID
func (r *AlibabaAlihealthAlipaypfmDietRecordRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthAlipaypfmDietRecordRequest) GetUserId() int64 {
    return r.userId
}
// Date Setter
// 记录日期，format：yyyy-MM-dd
func (r *AlibabaAlihealthAlipaypfmDietRecordRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

// Date Getter
func (r AlibabaAlihealthAlipaypfmDietRecordRequest) GetDate() string {
    return r.date
}
// Energy Setter
// 累积摄入卡路里
func (r *AlibabaAlihealthAlipaypfmDietRecordRequest) SetEnergy(energy int64) error {
    r.energy = energy
    r.Set("energy", energy)
    return nil
}

// Energy Getter
func (r AlibabaAlihealthAlipaypfmDietRecordRequest) GetEnergy() int64 {
    return r.energy
}
