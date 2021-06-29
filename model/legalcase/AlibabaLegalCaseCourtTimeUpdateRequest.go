package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开庭时间变更 API请求
alibaba.legal.case.court.time.update

修改案件的开庭时间
*/
type AlibabaLegalCaseCourtTimeUpdateRequest struct {
    model.Params
    // 案件id
    _caseId   int64
    // 委托id
    _entrustId   int64
    // 开庭时间
    _courtTime   string
}

// 初始化AlibabaLegalCaseCourtTimeUpdateRequest对象
func NewAlibabaLegalCaseCourtTimeUpdateRequest() *AlibabaLegalCaseCourtTimeUpdateRequest{
    return &AlibabaLegalCaseCourtTimeUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseCourtTimeUpdateRequest) GetApiMethodName() string {
    return "alibaba.legal.case.court.time.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseCourtTimeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CaseId Setter
// 案件id
func (r *AlibabaLegalCaseCourtTimeUpdateRequest) SetCaseId(_caseId int64) error {
    r._caseId = _caseId
    r.Set("case_id", _caseId)
    return nil
}

// CaseId Getter
func (r AlibabaLegalCaseCourtTimeUpdateRequest) GetCaseId() int64 {
    return r._caseId
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseCourtTimeUpdateRequest) SetEntrustId(_entrustId int64) error {
    r._entrustId = _entrustId
    r.Set("entrust_id", _entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseCourtTimeUpdateRequest) GetEntrustId() int64 {
    return r._entrustId
}
// CourtTime Setter
// 开庭时间
func (r *AlibabaLegalCaseCourtTimeUpdateRequest) SetCourtTime(_courtTime string) error {
    r._courtTime = _courtTime
    r.Set("court_time", _courtTime)
    return nil
}

// CourtTime Getter
func (r AlibabaLegalCaseCourtTimeUpdateRequest) GetCourtTime() string {
    return r._courtTime
}
