package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTagListAPIRequest 查询所有分组 API请求
// alibaba.scbp.tag.list
//
// 查询所有分组
type AlibabaScbpTagListAPIRequest struct {
	model.Params
}

// NewAlibabaScbpTagListRequest 初始化AlibabaScbpTagListAPIRequest对象
func NewAlibabaScbpTagListRequest() *AlibabaScbpTagListAPIRequest {
	return &AlibabaScbpTagListAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTagListAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagListAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.tag.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTagListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaScbpTagListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTagListRequest()
	},
}

// GetAlibabaScbpTagListRequest 从 sync.Pool 获取 AlibabaScbpTagListAPIRequest
func GetAlibabaScbpTagListAPIRequest() *AlibabaScbpTagListAPIRequest {
	return poolAlibabaScbpTagListAPIRequest.Get().(*AlibabaScbpTagListAPIRequest)
}

// ReleaseAlibabaScbpTagListAPIRequest 将 AlibabaScbpTagListAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTagListAPIRequest(v *AlibabaScbpTagListAPIRequest) {
	v.Reset()
	poolAlibabaScbpTagListAPIRequest.Put(v)
}
