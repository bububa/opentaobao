package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
委托开庭服务查询 APIRequest
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

func NewAlibabaLegalSuitCourtEntrustGetRequest() *AlibabaLegalSuitCourtEntrustGetRequest{
    return &AlibabaLegalSuitCourtEntrustGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalSuitCourtEntrustGetRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.court.entrust.get"
}

func (r AlibabaLegalSuitCourtEntrustGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalSuitCourtEntrustGetRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

func (r AlibabaLegalSuitCourtEntrustGetRequest) GetEntrustId() int64 {
    return r.entrustId
}

func (r *AlibabaLegalSuitCourtEntrustGetRequest) SetSuitId(suitId int64) error {
    r.suitId = suitId
    r.Set("suit_id", suitId)
    return nil
}

func (r AlibabaLegalSuitCourtEntrustGetRequest) GetSuitId() int64 {
    return r.suitId
}

