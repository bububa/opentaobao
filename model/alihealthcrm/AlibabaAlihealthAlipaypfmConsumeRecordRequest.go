package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
记录用户每日消耗卡路里总量 APIRequest
alibaba.alihealth.alipaypfm.consume.record

记录用户每日消耗卡路里总量
*/
type AlibabaAlihealthAlipaypfmConsumeRecordRequest struct {
    model.Params

    // 用户健康ID
    userId   int64 

    // 用户消耗卡路里总量
    energy   int64 

    // 记录日期, 格式: yyyy-MM-dd
    date   string 

}

func NewAlibabaAlihealthAlipaypfmConsumeRecordRequest() *AlibabaAlihealthAlipaypfmConsumeRecordRequest{
    return &AlibabaAlihealthAlipaypfmConsumeRecordRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthAlipaypfmConsumeRecordRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.consume.record"
}

func (r AlibabaAlihealthAlipaypfmConsumeRecordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthAlipaypfmConsumeRecordRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlihealthAlipaypfmConsumeRecordRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaAlihealthAlipaypfmConsumeRecordRequest) SetEnergy(energy int64) error {
    r.energy = energy
    r.Set("energy", energy)
    return nil
}

func (r AlibabaAlihealthAlipaypfmConsumeRecordRequest) GetEnergy() int64 {
    return r.energy
}

func (r *AlibabaAlihealthAlipaypfmConsumeRecordRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

func (r AlibabaAlihealthAlipaypfmConsumeRecordRequest) GetDate() string {
    return r.date
}

