package icburfq

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqMyequityAPIRequest 我的权益 API请求
// alibaba.icbu.rfq.myequity
//
// 查询供应商权益接口
type AlibabaIcbuRfqMyequityAPIRequest struct {
	model.Params
}

// NewAlibabaIcbuRfqMyequityRequest 初始化AlibabaIcbuRfqMyequityAPIRequest对象
func NewAlibabaIcbuRfqMyequityRequest() *AlibabaIcbuRfqMyequityAPIRequest {
	return &AlibabaIcbuRfqMyequityAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuRfqMyequityAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqMyequityAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfq.myequity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqMyequityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuRfqMyequityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaIcbuRfqMyequityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuRfqMyequityRequest()
	},
}

// GetAlibabaIcbuRfqMyequityRequest 从 sync.Pool 获取 AlibabaIcbuRfqMyequityAPIRequest
func GetAlibabaIcbuRfqMyequityAPIRequest() *AlibabaIcbuRfqMyequityAPIRequest {
	return poolAlibabaIcbuRfqMyequityAPIRequest.Get().(*AlibabaIcbuRfqMyequityAPIRequest)
}

// ReleaseAlibabaIcbuRfqMyequityAPIRequest 将 AlibabaIcbuRfqMyequityAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuRfqMyequityAPIRequest(v *AlibabaIcbuRfqMyequityAPIRequest) {
	v.Reset()
	poolAlibabaIcbuRfqMyequityAPIRequest.Put(v)
}
