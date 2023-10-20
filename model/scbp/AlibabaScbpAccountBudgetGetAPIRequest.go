package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountBudgetGetAPIRequest 查询日消耗预算 API请求
// alibaba.scbp.account.budget.get
//
// 查询日消耗预算
type AlibabaScbpAccountBudgetGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAccountBudgetGetRequest 初始化AlibabaScbpAccountBudgetGetAPIRequest对象
func NewAlibabaScbpAccountBudgetGetRequest() *AlibabaScbpAccountBudgetGetAPIRequest {
	return &AlibabaScbpAccountBudgetGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAccountBudgetGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountBudgetGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.budget.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountBudgetGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAccountBudgetGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaScbpAccountBudgetGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAccountBudgetGetRequest()
	},
}

// GetAlibabaScbpAccountBudgetGetRequest 从 sync.Pool 获取 AlibabaScbpAccountBudgetGetAPIRequest
func GetAlibabaScbpAccountBudgetGetAPIRequest() *AlibabaScbpAccountBudgetGetAPIRequest {
	return poolAlibabaScbpAccountBudgetGetAPIRequest.Get().(*AlibabaScbpAccountBudgetGetAPIRequest)
}

// ReleaseAlibabaScbpAccountBudgetGetAPIRequest 将 AlibabaScbpAccountBudgetGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAccountBudgetGetAPIRequest(v *AlibabaScbpAccountBudgetGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpAccountBudgetGetAPIRequest.Put(v)
}
