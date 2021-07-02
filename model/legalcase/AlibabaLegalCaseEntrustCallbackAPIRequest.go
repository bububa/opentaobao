package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseEntrustCallbackAPIRequest 委托回调接口 API请求
// alibaba.legal.case.entrust.callback
//
// 委托回调接口
type AlibabaLegalCaseEntrustCallbackAPIRequest struct {
	model.Params
	// 委托id
	_entrustId int64
	// 案件id
	_caseId int64
}

// NewAlibabaLegalCaseEntrustCallbackRequest 初始化AlibabaLegalCaseEntrustCallbackAPIRequest对象
func NewAlibabaLegalCaseEntrustCallbackRequest() *AlibabaLegalCaseEntrustCallbackAPIRequest {
	return &AlibabaLegalCaseEntrustCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseEntrustCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.entrust.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseEntrustCallbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseEntrustCallbackAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// Get EntrustId Getter
func (r AlibabaLegalCaseEntrustCallbackAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// Set is CaseId Setter
// 案件id
func (r *AlibabaLegalCaseEntrustCallbackAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// Get CaseId Getter
func (r AlibabaLegalCaseEntrustCallbackAPIRequest) GetCaseId() int64 {
	return r._caseId
}
