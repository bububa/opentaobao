package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户每日摄入卡路里总量回传接口 APIRequest
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

func NewAlibabaAlihealthAlipaypfmDietRecordRequest() *AlibabaAlihealthAlipaypfmDietRecordRequest{
    return &AlibabaAlihealthAlipaypfmDietRecordRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthAlipaypfmDietRecordRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.diet.record"
}

func (r AlibabaAlihealthAlipaypfmDietRecordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthAlipaypfmDietRecordRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlihealthAlipaypfmDietRecordRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaAlihealthAlipaypfmDietRecordRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

func (r AlibabaAlihealthAlipaypfmDietRecordRequest) GetDate() string {
    return r.date
}

func (r *AlibabaAlihealthAlipaypfmDietRecordRequest) SetEnergy(energy int64) error {
    r.energy = energy
    r.Set("energy", energy)
    return nil
}

func (r AlibabaAlihealthAlipaypfmDietRecordRequest) GetEnergy() int64 {
    return r.energy
}

