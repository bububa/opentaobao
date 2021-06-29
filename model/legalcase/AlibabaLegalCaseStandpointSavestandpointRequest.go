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
type AlibabaLegalCaseStandpointSavestandpointRequest struct {
    model.Params
    // 答辩口径
    defenseCaliber   string
    // 口径描述
    standpointDesc   string
    // 案件id
    suitId   int64
    // 委托id
    entrustId   int64
    // 提交人
    submitPeople   string
}

// 初始化AlibabaLegalCaseStandpointSavestandpointRequest对象
func NewAlibabaLegalCaseStandpointSavestandpointRequest() *AlibabaLegalCaseStandpointSavestandpointRequest{
    return &AlibabaLegalCaseStandpointSavestandpointRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseStandpointSavestandpointRequest) GetApiMethodName() string {
    return "alibaba.legal.case.standpoint.savestandpoint"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseStandpointSavestandpointRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DefenseCaliber Setter
// 答辩口径
func (r *AlibabaLegalCaseStandpointSavestandpointRequest) SetDefenseCaliber(defenseCaliber string) error {
    r.defenseCaliber = defenseCaliber
    r.Set("defense_caliber", defenseCaliber)
    return nil
}

// DefenseCaliber Getter
func (r AlibabaLegalCaseStandpointSavestandpointRequest) GetDefenseCaliber() string {
    return r.defenseCaliber
}
// StandpointDesc Setter
// 口径描述
func (r *AlibabaLegalCaseStandpointSavestandpointRequest) SetStandpointDesc(standpointDesc string) error {
    r.standpointDesc = standpointDesc
    r.Set("standpoint_desc", standpointDesc)
    return nil
}

// StandpointDesc Getter
func (r AlibabaLegalCaseStandpointSavestandpointRequest) GetStandpointDesc() string {
    return r.standpointDesc
}
// SuitId Setter
// 案件id
func (r *AlibabaLegalCaseStandpointSavestandpointRequest) SetSuitId(suitId int64) error {
    r.suitId = suitId
    r.Set("suit_id", suitId)
    return nil
}

// SuitId Getter
func (r AlibabaLegalCaseStandpointSavestandpointRequest) GetSuitId() int64 {
    return r.suitId
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseStandpointSavestandpointRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseStandpointSavestandpointRequest) GetEntrustId() int64 {
    return r.entrustId
}
// SubmitPeople Setter
// 提交人
func (r *AlibabaLegalCaseStandpointSavestandpointRequest) SetSubmitPeople(submitPeople string) error {
    r.submitPeople = submitPeople
    r.Set("submit_people", submitPeople)
    return nil
}

// SubmitPeople Getter
func (r AlibabaLegalCaseStandpointSavestandpointRequest) GetSubmitPeople() string {
    return r.submitPeople
}
