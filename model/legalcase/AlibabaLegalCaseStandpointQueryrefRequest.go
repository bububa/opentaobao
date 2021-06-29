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
    _suitId   int64
    // 委托id
    _entrustId   int64
    // 查询的口径id列表
    _standpointIds   []int64
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
func (r *AlibabaLegalCaseStandpointQueryrefRequest) SetSuitId(_suitId int64) error {
    r._suitId = _suitId
    r.Set("suit_id", _suitId)
    return nil
}

// SuitId Getter
func (r AlibabaLegalCaseStandpointQueryrefRequest) GetSuitId() int64 {
    return r._suitId
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseStandpointQueryrefRequest) SetEntrustId(_entrustId int64) error {
    r._entrustId = _entrustId
    r.Set("entrust_id", _entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseStandpointQueryrefRequest) GetEntrustId() int64 {
    return r._entrustId
}
// StandpointIds Setter
// 查询的口径id列表
func (r *AlibabaLegalCaseStandpointQueryrefRequest) SetStandpointIds(_standpointIds []int64) error {
    r._standpointIds = _standpointIds
    r.Set("standpoint_ids", _standpointIds)
    return nil
}

// StandpointIds Getter
func (r AlibabaLegalCaseStandpointQueryrefRequest) GetStandpointIds() []int64 {
    return r._standpointIds
}
