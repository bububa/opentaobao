package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcaseentrustcallbackAPIRequest 委托回调接口 API请求
// alibaba.legal.case.entrust.callback
//
// 委托回调接口
type AlibabalegalcaseentrustcallbackAPIRequest struct {
	model.Params
	// 委托id
	_entrustId int64
	// 案件id
	_caseId int64
}

// NewAlibabalegalcaseentrustcallbackRequest 初始化AlibabalegalcaseentrustcallbackAPIRequest对象
func NewAlibabalegalcaseentrustcallbackRequest() *AlibabalegalcaseentrustcallbackAPIRequest {
	return &AlibabalegalcaseentrustcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalcaseentrustcallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.entrust.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalcaseentrustcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalcaseentrustcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabalegalcaseentrustcallbackAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabalegalcaseentrustcallbackAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetCaseId is CaseId Setter
// 案件id
func (r *AlibabalegalcaseentrustcallbackAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabalegalcaseentrustcallbackAPIRequest) GetCaseId() int64 {
	return r._caseId
}
