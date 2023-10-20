package icbushowcase

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseStatusAPIRequest 橱窗状态 API请求
// alibaba.scbp.showcase.status
//
// 查询橱窗状态，如总数、可用数量
type AlibabaScbpShowcaseStatusAPIRequest struct {
	model.Params
}

// NewAlibabaScbpShowcaseStatusRequest 初始化AlibabaScbpShowcaseStatusAPIRequest对象
func NewAlibabaScbpShowcaseStatusRequest() *AlibabaScbpShowcaseStatusAPIRequest {
	return &AlibabaScbpShowcaseStatusAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpShowcaseStatusAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpShowcaseStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaScbpShowcaseStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpShowcaseStatusRequest()
	},
}

// GetAlibabaScbpShowcaseStatusRequest 从 sync.Pool 获取 AlibabaScbpShowcaseStatusAPIRequest
func GetAlibabaScbpShowcaseStatusAPIRequest() *AlibabaScbpShowcaseStatusAPIRequest {
	return poolAlibabaScbpShowcaseStatusAPIRequest.Get().(*AlibabaScbpShowcaseStatusAPIRequest)
}

// ReleaseAlibabaScbpShowcaseStatusAPIRequest 将 AlibabaScbpShowcaseStatusAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpShowcaseStatusAPIRequest(v *AlibabaScbpShowcaseStatusAPIRequest) {
	v.Reset()
	poolAlibabaScbpShowcaseStatusAPIRequest.Put(v)
}
