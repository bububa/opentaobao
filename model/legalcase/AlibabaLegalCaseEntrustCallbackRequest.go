package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
委托回调接口 API请求
alibaba.legal.case.entrust.callback

委托回调接口
*/
type AlibabaLegalCaseEntrustCallbackRequest struct {
    model.Params
    // 委托id
    entrustId   int64
    // 案件id
    caseId   int64
}

// 初始化AlibabaLegalCaseEntrustCallbackRequest对象
func NewAlibabaLegalCaseEntrustCallbackRequest() *AlibabaLegalCaseEntrustCallbackRequest{
    return &AlibabaLegalCaseEntrustCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseEntrustCallbackRequest) GetApiMethodName() string {
    return "alibaba.legal.case.entrust.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseEntrustCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseEntrustCallbackRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseEntrustCallbackRequest) GetEntrustId() int64 {
    return r.entrustId
}
// CaseId Setter
// 案件id
func (r *AlibabaLegalCaseEntrustCallbackRequest) SetCaseId(caseId int64) error {
    r.caseId = caseId
    r.Set("case_id", caseId)
    return nil
}

// CaseId Getter
func (r AlibabaLegalCaseEntrustCallbackRequest) GetCaseId() int64 {
    return r.caseId
}
