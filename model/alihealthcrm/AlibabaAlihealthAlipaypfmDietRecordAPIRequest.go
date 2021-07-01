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
type AlibabaAlihealthAlipaypfmDietRecordAPIRequest struct {
    model.Params
    // 用户健康ID
    _userId   int64
    // 记录日期，format：yyyy-MM-dd
    _date   string
    // 累积摄入卡路里
    _energy   int64
}

// 初始化AlibabaAlihealthAlipaypfmDietRecordAPIRequest对象
func NewAlibabaAlihealthAlipaypfmDietRecordRequest() *AlibabaAlihealthAlipaypfmDietRecordAPIRequest{
    return &AlibabaAlihealthAlipaypfmDietRecordAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlipaypfmDietRecordAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.diet.record"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlipaypfmDietRecordAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户健康ID
func (r *AlibabaAlihealthAlipaypfmDietRecordAPIRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthAlipaypfmDietRecordAPIRequest) GetUserId() int64 {
    return r._userId
}
// Date Setter
// 记录日期，format：yyyy-MM-dd
func (r *AlibabaAlihealthAlipaypfmDietRecordAPIRequest) SetDate(_date string) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r AlibabaAlihealthAlipaypfmDietRecordAPIRequest) GetDate() string {
    return r._date
}
// Energy Setter
// 累积摄入卡路里
func (r *AlibabaAlihealthAlipaypfmDietRecordAPIRequest) SetEnergy(_energy int64) error {
    r._energy = _energy
    r.Set("energy", _energy)
    return nil
}

// Energy Getter
func (r AlibabaAlihealthAlipaypfmDietRecordAPIRequest) GetEnergy() int64 {
    return r._energy
}
