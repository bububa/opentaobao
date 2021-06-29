package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推送口径信息 API请求
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
    standpointIds   []int64
}

// 初始化AlibabaLegalCaseStandpointQueryrefRequest对象
func NewAlibabaLegalCaseStandpointQueryrefRequest() *AlibabaLegalCaseStandpointQueryrefRequest{
    return &AlibabaLegalCaseStandpointQueryrefRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseStandpointQueryrefRequest) GetApiMethodName() string {
    return "alibaba.legal.case.standpoint.queryref"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseStandpointQueryrefRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SuitId Setter
// 案件id
func (r *AlibabaLegalCaseStandpointQueryrefRequest) SetSuitId(suitId int64) error {
    r.suitId = suitId
    r.Set("suit_id", suitId)
    return nil
}

// SuitId Getter
func (r AlibabaLegalCaseStandpointQueryrefRequest) GetSuitId() int64 {
    return r.suitId
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseStandpointQueryrefRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseStandpointQueryrefRequest) GetEntrustId() int64 {
    return r.entrustId
}
// StandpointIds Setter
// 查询的口径id列表
func (r *AlibabaLegalCaseStandpointQueryrefRequest) SetStandpointIds(standpointIds []int64) error {
    r.standpointIds = standpointIds
    r.Set("standpoint_ids", standpointIds)
    return nil
}

// StandpointIds Getter
func (r AlibabaLegalCaseStandpointQueryrefRequest) GetStandpointIds() []int64 {
    return r.standpointIds
}
