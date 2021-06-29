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
    caseId   int64
    // 委托id
    entrustId   int64
    // 开庭时间
    courtTime   string
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
func (r *AlibabaLegalCaseCourtTimeUpdateRequest) SetCaseId(caseId int64) error {
    r.caseId = caseId
    r.Set("case_id", caseId)
    return nil
}

// CaseId Getter
func (r AlibabaLegalCaseCourtTimeUpdateRequest) GetCaseId() int64 {
    return r.caseId
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseCourtTimeUpdateRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseCourtTimeUpdateRequest) GetEntrustId() int64 {
    return r.entrustId
}
// CourtTime Setter
// 开庭时间
func (r *AlibabaLegalCaseCourtTimeUpdateRequest) SetCourtTime(courtTime string) error {
    r.courtTime = courtTime
    r.Set("court_time", courtTime)
    return nil
}

// CourtTime Getter
func (r AlibabaLegalCaseCourtTimeUpdateRequest) GetCourtTime() string {
    return r.courtTime
}
