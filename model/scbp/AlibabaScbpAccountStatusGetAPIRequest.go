package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountStatusGetAPIRequest 查询账户级别关键词推广状态 API请求
// alibaba.scbp.account.status.get
//
// 查询账户级别关键词推广状态
type AlibabaScbpAccountStatusGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAccountStatusGetRequest 初始化AlibabaScbpAccountStatusGetAPIRequest对象
func NewAlibabaScbpAccountStatusGetRequest() *AlibabaScbpAccountStatusGetAPIRequest {
	return &AlibabaScbpAccountStatusGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAccountStatusGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountStatusGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountStatusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAccountStatusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaScbpAccountStatusGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAccountStatusGetRequest()
	},
}

// GetAlibabaScbpAccountStatusGetRequest 从 sync.Pool 获取 AlibabaScbpAccountStatusGetAPIRequest
func GetAlibabaScbpAccountStatusGetAPIRequest() *AlibabaScbpAccountStatusGetAPIRequest {
	return poolAlibabaScbpAccountStatusGetAPIRequest.Get().(*AlibabaScbpAccountStatusGetAPIRequest)
}

// ReleaseAlibabaScbpAccountStatusGetAPIRequest 将 AlibabaScbpAccountStatusGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAccountStatusGetAPIRequest(v *AlibabaScbpAccountStatusGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpAccountStatusGetAPIRequest.Put(v)
}
