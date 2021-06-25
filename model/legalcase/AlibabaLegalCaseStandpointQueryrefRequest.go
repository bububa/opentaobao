package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推送口径信息 APIRequest
alibaba.legal.case.standpoint.queryref

查询推送口径信息
*/
type AlibabaLegalCaseStandpointQueryrefRequest struct {
    model.Params

    // 案件id
    suitId   int64 

    // 委托id
    entrustId   int64 

    // 查询的口径id列表
    standpointIds   []Number 

}

func NewAlibabaLegalCaseStandpointQueryrefRequest() *AlibabaLegalCaseStandpointQueryrefRequest{
    return &AlibabaLegalCaseStandpointQueryrefRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalCaseStandpointQueryrefRequest) GetApiMethodName() string {
    return "alibaba.legal.case.standpoint.queryref"
}

func (r AlibabaLegalCaseStandpointQueryrefRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalCaseStandpointQueryrefRequest) SetSuitId(suitId int64) error {
    r.suitId = suitId
    r.Set("suit_id", suitId)
    return nil
}

func (r AlibabaLegalCaseStandpointQueryrefRequest) GetSuitId() int64 {
    return r.suitId
}

func (r *AlibabaLegalCaseStandpointQueryrefRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

func (r AlibabaLegalCaseStandpointQueryrefRequest) GetEntrustId() int64 {
    return r.entrustId
}

func (r *AlibabaLegalCaseStandpointQueryrefRequest) SetStandpointIds(standpointIds []Number) error {
    r.standpointIds = standpointIds
    r.Set("standpoint_ids", standpointIds)
    return nil
}

func (r AlibabaLegalCaseStandpointQueryrefRequest) GetStandpointIds() []Number {
    return r.standpointIds
}

