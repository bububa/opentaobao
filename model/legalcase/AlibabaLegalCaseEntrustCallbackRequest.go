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
    _entrustId   int64
    // 案件id
    _caseId   int64
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
func (r *AlibabaLegalCaseEntrustCallbackRequest) SetEntrustId(_entrustId int64) error {
    r._entrustId = _entrustId
    r.Set("entrust_id", _entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseEntrustCallbackRequest) GetEntrustId() int64 {
    return r._entrustId
}
// CaseId Setter
// 案件id
func (r *AlibabaLegalCaseEntrustCallbackRequest) SetCaseId(_caseId int64) error {
    r._caseId = _caseId
    r.Set("case_id", _caseId)
    return nil
}

// CaseId Getter
func (r AlibabaLegalCaseEntrustCallbackRequest) GetCaseId() int64 {
    return r._caseId
}
