package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpEffectAccountDateGetAPIRequest 获取最近报表生成时间 API请求
// alibaba.scbp.effect.account.date.get
//
// 获取最近报表生成时间,格式为yyyy-MM-dd
type AlibabaScbpEffectAccountDateGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpEffectAccountDateGetRequest 初始化AlibabaScbpEffectAccountDateGetAPIRequest对象
func NewAlibabaScbpEffectAccountDateGetRequest() *AlibabaScbpEffectAccountDateGetAPIRequest {
	return &AlibabaScbpEffectAccountDateGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpEffectAccountDateGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectAccountDateGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.effect.account.date.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectAccountDateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpEffectAccountDateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaScbpEffectAccountDateGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpEffectAccountDateGetRequest()
	},
}

// GetAlibabaScbpEffectAccountDateGetRequest 从 sync.Pool 获取 AlibabaScbpEffectAccountDateGetAPIRequest
func GetAlibabaScbpEffectAccountDateGetAPIRequest() *AlibabaScbpEffectAccountDateGetAPIRequest {
	return poolAlibabaScbpEffectAccountDateGetAPIRequest.Get().(*AlibabaScbpEffectAccountDateGetAPIRequest)
}

// ReleaseAlibabaScbpEffectAccountDateGetAPIRequest 将 AlibabaScbpEffectAccountDateGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpEffectAccountDateGetAPIRequest(v *AlibabaScbpEffectAccountDateGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpEffectAccountDateGetAPIRequest.Put(v)
}
