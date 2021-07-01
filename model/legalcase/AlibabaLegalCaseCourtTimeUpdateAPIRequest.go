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
type AlibabaLegalCaseCourtTimeUpdateAPIRequest struct {
    model.Params
    // 案件id
    _caseId   int64
    // 委托id
    _entrustId   int64
    // 开庭时间
    _courtTime   string
}

// 初始化AlibabaLegalCaseCourtTimeUpdateAPIRequest对象
func NewAlibabaLegalCaseCourtTimeUpdateRequest() *AlibabaLegalCaseCourtTimeUpdateAPIRequest{
    return &AlibabaLegalCaseCourtTimeUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.legal.case.court.time.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CaseId Setter
// 案件id
func (r *AlibabaLegalCaseCourtTimeUpdateAPIRequest) SetCaseId(_caseId int64) error {
    r._caseId = _caseId
    r.Set("case_id", _caseId)
    return nil
}

// CaseId Getter
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetCaseId() int64 {
    return r._caseId
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseCourtTimeUpdateAPIRequest) SetEntrustId(_entrustId int64) error {
    r._entrustId = _entrustId
    r.Set("entrust_id", _entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetEntrustId() int64 {
    return r._entrustId
}
// CourtTime Setter
// 开庭时间
func (r *AlibabaLegalCaseCourtTimeUpdateAPIRequest) SetCourtTime(_courtTime string) error {
    r._courtTime = _courtTime
    r.Set("court_time", _courtTime)
    return nil
}

// CourtTime Getter
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetCourtTime() string {
    return r._courtTime
}
