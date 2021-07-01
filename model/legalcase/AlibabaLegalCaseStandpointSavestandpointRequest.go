package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增反馈口径 API请求
alibaba.legal.case.standpoint.savestandpoint

新增反馈口径 ,从外部接受反馈的口径
*/
type AlibabaLegalCaseStandpointSavestandpointAPIRequest struct {
    model.Params
    // 答辩口径
    _defenseCaliber   string
    // 口径描述
    _standpointDesc   string
    // 案件id
    _suitId   int64
    // 委托id
    _entrustId   int64
    // 提交人
    _submitPeople   string
}

// 初始化AlibabaLegalCaseStandpointSavestandpointAPIRequest对象
func NewAlibabaLegalCaseStandpointSavestandpointRequest() *AlibabaLegalCaseStandpointSavestandpointAPIRequest{
    return &AlibabaLegalCaseStandpointSavestandpointAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetApiMethodName() string {
    return "alibaba.legal.case.standpoint.savestandpoint"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DefenseCaliber Setter
// 答辩口径
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) SetDefenseCaliber(_defenseCaliber string) error {
    r._defenseCaliber = _defenseCaliber
    r.Set("defense_caliber", _defenseCaliber)
    return nil
}

// DefenseCaliber Getter
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetDefenseCaliber() string {
    return r._defenseCaliber
}
// StandpointDesc Setter
// 口径描述
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) SetStandpointDesc(_standpointDesc string) error {
    r._standpointDesc = _standpointDesc
    r.Set("standpoint_desc", _standpointDesc)
    return nil
}

// StandpointDesc Getter
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetStandpointDesc() string {
    return r._standpointDesc
}
// SuitId Setter
// 案件id
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) SetSuitId(_suitId int64) error {
    r._suitId = _suitId
    r.Set("suit_id", _suitId)
    return nil
}

// SuitId Getter
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetSuitId() int64 {
    return r._suitId
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) SetEntrustId(_entrustId int64) error {
    r._entrustId = _entrustId
    r.Set("entrust_id", _entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetEntrustId() int64 {
    return r._entrustId
}
// SubmitPeople Setter
// 提交人
func (r *AlibabaLegalCaseStandpointSavestandpointAPIRequest) SetSubmitPeople(_submitPeople string) error {
    r._submitPeople = _submitPeople
    r.Set("submit_people", _submitPeople)
    return nil
}

// SubmitPeople Getter
func (r AlibabaLegalCaseStandpointSavestandpointAPIRequest) GetSubmitPeople() string {
    return r._submitPeople
}
