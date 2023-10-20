package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdAccountBalanceGetAPIRequest 查询账户余额 API请求
// alibaba.scbp.ad.account.balance.get
//
// 查询推广账户余额
type AlibabaScbpAdAccountBalanceGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAdAccountBalanceGetRequest 初始化AlibabaScbpAdAccountBalanceGetAPIRequest对象
func NewAlibabaScbpAdAccountBalanceGetRequest() *AlibabaScbpAdAccountBalanceGetAPIRequest {
	return &AlibabaScbpAdAccountBalanceGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdAccountBalanceGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdAccountBalanceGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.account.balance.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdAccountBalanceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdAccountBalanceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaScbpAdAccountBalanceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdAccountBalanceGetRequest()
	},
}

// GetAlibabaScbpAdAccountBalanceGetRequest 从 sync.Pool 获取 AlibabaScbpAdAccountBalanceGetAPIRequest
func GetAlibabaScbpAdAccountBalanceGetAPIRequest() *AlibabaScbpAdAccountBalanceGetAPIRequest {
	return poolAlibabaScbpAdAccountBalanceGetAPIRequest.Get().(*AlibabaScbpAdAccountBalanceGetAPIRequest)
}

// ReleaseAlibabaScbpAdAccountBalanceGetAPIRequest 将 AlibabaScbpAdAccountBalanceGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdAccountBalanceGetAPIRequest(v *AlibabaScbpAdAccountBalanceGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdAccountBalanceGetAPIRequest.Put(v)
}
