package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasecourttimeupdateAPIRequest 开庭时间变更 API请求
// alibaba.legal.case.court.time.update
//
// 修改案件的开庭时间
type AlibabalegalcasecourttimeupdateAPIRequest struct {
	model.Params
	// 开庭时间
	_courtTime string
	// 案件id
	_caseId int64
	// 委托id
	_entrustId int64
}

// NewAlibabalegalcasecourttimeupdateRequest 初始化AlibabalegalcasecourttimeupdateAPIRequest对象
func NewAlibabalegalcasecourttimeupdateRequest() *AlibabalegalcasecourttimeupdateAPIRequest {
	return &AlibabalegalcasecourttimeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalcasecourttimeupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.court.time.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalcasecourttimeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalcasecourttimeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtTime is CourtTime Setter
// 开庭时间
func (r *AlibabalegalcasecourttimeupdateAPIRequest) SetCourtTime(_courtTime string) error {
	r._courtTime = _courtTime
	r.Set("court_time", _courtTime)
	return nil
}

// GetCourtTime CourtTime Getter
func (r AlibabalegalcasecourttimeupdateAPIRequest) GetCourtTime() string {
	return r._courtTime
}

// SetCaseId is CaseId Setter
// 案件id
func (r *AlibabalegalcasecourttimeupdateAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabalegalcasecourttimeupdateAPIRequest) GetCaseId() int64 {
	return r._caseId
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabalegalcasecourttimeupdateAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabalegalcasecourttimeupdateAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}
