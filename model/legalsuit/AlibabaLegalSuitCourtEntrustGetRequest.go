package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
委托开庭服务查询 API请求
alibaba.legal.suit.court.entrust.get

查询委托开庭信息
*/
type AlibabaLegalSuitCourtEntrustGetRequest struct {
    model.Params
    // 委托ID
    entrustId   int64
    // 案件ID
    suitId   int64
}

// 初始化AlibabaLegalSuitCourtEntrustGetRequest对象
func NewAlibabaLegalSuitCourtEntrustGetRequest() *AlibabaLegalSuitCourtEntrustGetRequest{
    return &AlibabaLegalSuitCourtEntrustGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtEntrustGetRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.court.entrust.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtEntrustGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntrustId Setter
// 委托ID
func (r *AlibabaLegalSuitCourtEntrustGetRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalSuitCourtEntrustGetRequest) GetEntrustId() int64 {
    return r.entrustId
}
// SuitId Setter
// 案件ID
func (r *AlibabaLegalSuitCourtEntrustGetRequest) SetSuitId(suitId int64) error {
    r.suitId = suitId
    r.Set("suit_id", suitId)
    return nil
}

// SuitId Getter
func (r AlibabaLegalSuitCourtEntrustGetRequest) GetSuitId() int64 {
    return r.suitId
}
