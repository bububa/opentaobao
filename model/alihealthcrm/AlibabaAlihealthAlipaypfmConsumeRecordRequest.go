package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
记录用户每日消耗卡路里总量 API请求
alibaba.alihealth.alipaypfm.consume.record

记录用户每日消耗卡路里总量
*/
type AlibabaAlihealthAlipaypfmConsumeRecordRequest struct {
    model.Params
    // 用户健康ID
    _userId   int64
    // 用户消耗卡路里总量
    _energy   int64
    // 记录日期, 格式: yyyy-MM-dd
    _date   string
}

// 初始化AlibabaAlihealthAlipaypfmConsumeRecordRequest对象
func NewAlibabaAlihealthAlipaypfmConsumeRecordRequest() *AlibabaAlihealthAlipaypfmConsumeRecordRequest{
    return &AlibabaAlihealthAlipaypfmConsumeRecordRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlipaypfmConsumeRecordRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.consume.record"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlipaypfmConsumeRecordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户健康ID
func (r *AlibabaAlihealthAlipaypfmConsumeRecordRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthAlipaypfmConsumeRecordRequest) GetUserId() int64 {
    return r._userId
}
// Energy Setter
// 用户消耗卡路里总量
func (r *AlibabaAlihealthAlipaypfmConsumeRecordRequest) SetEnergy(_energy int64) error {
    r._energy = _energy
    r.Set("energy", _energy)
    return nil
}

// Energy Getter
func (r AlibabaAlihealthAlipaypfmConsumeRecordRequest) GetEnergy() int64 {
    return r._energy
}
// Date Setter
// 记录日期, 格式: yyyy-MM-dd
func (r *AlibabaAlihealthAlipaypfmConsumeRecordRequest) SetDate(_date string) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r AlibabaAlihealthAlipaypfmConsumeRecordRequest) GetDate() string {
    return r._date
}
