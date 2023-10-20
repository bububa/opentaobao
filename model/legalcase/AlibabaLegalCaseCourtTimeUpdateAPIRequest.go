package legalcase

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseCourtTimeUpdateAPIRequest 开庭时间变更 API请求
// alibaba.legal.case.court.time.update
//
// 修改案件的开庭时间
type AlibabaLegalCaseCourtTimeUpdateAPIRequest struct {
	model.Params
	// 开庭时间
	_courtTime string
	// 案件id
	_caseId int64
	// 委托id
	_entrustId int64
}

// NewAlibabaLegalCaseCourtTimeUpdateRequest 初始化AlibabaLegalCaseCourtTimeUpdateAPIRequest对象
func NewAlibabaLegalCaseCourtTimeUpdateRequest() *AlibabaLegalCaseCourtTimeUpdateAPIRequest {
	return &AlibabaLegalCaseCourtTimeUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalCaseCourtTimeUpdateAPIRequest) Reset() {
	r._courtTime = ""
	r._caseId = 0
	r._entrustId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.court.time.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtTime is CourtTime Setter
// 开庭时间
func (r *AlibabaLegalCaseCourtTimeUpdateAPIRequest) SetCourtTime(_courtTime string) error {
	r._courtTime = _courtTime
	r.Set("court_time", _courtTime)
	return nil
}

// GetCourtTime CourtTime Getter
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetCourtTime() string {
	return r._courtTime
}

// SetCaseId is CaseId Setter
// 案件id
func (r *AlibabaLegalCaseCourtTimeUpdateAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetCaseId() int64 {
	return r._caseId
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseCourtTimeUpdateAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabaLegalCaseCourtTimeUpdateAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

var poolAlibabaLegalCaseCourtTimeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalCaseCourtTimeUpdateRequest()
	},
}

// GetAlibabaLegalCaseCourtTimeUpdateRequest 从 sync.Pool 获取 AlibabaLegalCaseCourtTimeUpdateAPIRequest
func GetAlibabaLegalCaseCourtTimeUpdateAPIRequest() *AlibabaLegalCaseCourtTimeUpdateAPIRequest {
	return poolAlibabaLegalCaseCourtTimeUpdateAPIRequest.Get().(*AlibabaLegalCaseCourtTimeUpdateAPIRequest)
}

// ReleaseAlibabaLegalCaseCourtTimeUpdateAPIRequest 将 AlibabaLegalCaseCourtTimeUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalCaseCourtTimeUpdateAPIRequest(v *AlibabaLegalCaseCourtTimeUpdateAPIRequest) {
	v.Reset()
	poolAlibabaLegalCaseCourtTimeUpdateAPIRequest.Put(v)
}
