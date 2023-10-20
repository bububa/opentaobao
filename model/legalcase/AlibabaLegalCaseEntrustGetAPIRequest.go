package legalcase

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseEntrustGetAPIRequest 委托 API请求
// alibaba.legal.case.entrust.get
//
// 获取委托案件的基本信息
type AlibabaLegalCaseEntrustGetAPIRequest struct {
	model.Params
	// 委托id
	_entrustId int64
}

// NewAlibabaLegalCaseEntrustGetRequest 初始化AlibabaLegalCaseEntrustGetAPIRequest对象
func NewAlibabaLegalCaseEntrustGetRequest() *AlibabaLegalCaseEntrustGetAPIRequest {
	return &AlibabaLegalCaseEntrustGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalCaseEntrustGetAPIRequest) Reset() {
	r._entrustId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseEntrustGetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.entrust.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseEntrustGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalCaseEntrustGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseEntrustGetAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabaLegalCaseEntrustGetAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

var poolAlibabaLegalCaseEntrustGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalCaseEntrustGetRequest()
	},
}

// GetAlibabaLegalCaseEntrustGetRequest 从 sync.Pool 获取 AlibabaLegalCaseEntrustGetAPIRequest
func GetAlibabaLegalCaseEntrustGetAPIRequest() *AlibabaLegalCaseEntrustGetAPIRequest {
	return poolAlibabaLegalCaseEntrustGetAPIRequest.Get().(*AlibabaLegalCaseEntrustGetAPIRequest)
}

// ReleaseAlibabaLegalCaseEntrustGetAPIRequest 将 AlibabaLegalCaseEntrustGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalCaseEntrustGetAPIRequest(v *AlibabaLegalCaseEntrustGetAPIRequest) {
	v.Reset()
	poolAlibabaLegalCaseEntrustGetAPIRequest.Put(v)
}
