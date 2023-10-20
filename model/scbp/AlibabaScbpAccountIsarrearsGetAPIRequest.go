package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountIsarrearsGetAPIRequest 查询关键词推广账户是否欠款 API请求
// alibaba.scbp.account.isarrears.get
//
// 查询关键词推广账户是否欠款
type AlibabaScbpAccountIsarrearsGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAccountIsarrearsGetRequest 初始化AlibabaScbpAccountIsarrearsGetAPIRequest对象
func NewAlibabaScbpAccountIsarrearsGetRequest() *AlibabaScbpAccountIsarrearsGetAPIRequest {
	return &AlibabaScbpAccountIsarrearsGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAccountIsarrearsGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountIsarrearsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.isarrears.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountIsarrearsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAccountIsarrearsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaScbpAccountIsarrearsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAccountIsarrearsGetRequest()
	},
}

// GetAlibabaScbpAccountIsarrearsGetRequest 从 sync.Pool 获取 AlibabaScbpAccountIsarrearsGetAPIRequest
func GetAlibabaScbpAccountIsarrearsGetAPIRequest() *AlibabaScbpAccountIsarrearsGetAPIRequest {
	return poolAlibabaScbpAccountIsarrearsGetAPIRequest.Get().(*AlibabaScbpAccountIsarrearsGetAPIRequest)
}

// ReleaseAlibabaScbpAccountIsarrearsGetAPIRequest 将 AlibabaScbpAccountIsarrearsGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAccountIsarrearsGetAPIRequest(v *AlibabaScbpAccountIsarrearsGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpAccountIsarrearsGetAPIRequest.Put(v)
}
