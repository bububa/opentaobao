package legalcase

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalCaseEntrustCallbackAPIRequest) Reset() {
	r._entrustId = 0
	r._caseId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseEntrustCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.entrust.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseEntrustCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalCaseEntrustCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseEntrustCallbackAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabaLegalCaseEntrustCallbackAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetCaseId is CaseId Setter
// 案件id
func (r *AlibabaLegalCaseEntrustCallbackAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabaLegalCaseEntrustCallbackAPIRequest) GetCaseId() int64 {
	return r._caseId
}

var poolAlibabaLegalCaseEntrustCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalCaseEntrustCallbackRequest()
	},
}

// GetAlibabaLegalCaseEntrustCallbackRequest 从 sync.Pool 获取 AlibabaLegalCaseEntrustCallbackAPIRequest
func GetAlibabaLegalCaseEntrustCallbackAPIRequest() *AlibabaLegalCaseEntrustCallbackAPIRequest {
	return poolAlibabaLegalCaseEntrustCallbackAPIRequest.Get().(*AlibabaLegalCaseEntrustCallbackAPIRequest)
}

// ReleaseAlibabaLegalCaseEntrustCallbackAPIRequest 将 AlibabaLegalCaseEntrustCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalCaseEntrustCallbackAPIRequest(v *AlibabaLegalCaseEntrustCallbackAPIRequest) {
	v.Reset()
	poolAlibabaLegalCaseEntrustCallbackAPIRequest.Put(v)
}
