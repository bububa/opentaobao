package ju

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunitySubmittingTextAPIRequest 聚划算社群动态文案下发接口 API请求
// alibaba.jhs.community.submitting.text
//
// 聚划算社群动态文案下发接口
type AlibabaJhsCommunitySubmittingTextAPIRequest struct {
	model.Params
}

// NewAlibabaJhsCommunitySubmittingTextRequest 初始化AlibabaJhsCommunitySubmittingTextAPIRequest对象
func NewAlibabaJhsCommunitySubmittingTextRequest() *AlibabaJhsCommunitySubmittingTextAPIRequest {
	return &AlibabaJhsCommunitySubmittingTextAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJhsCommunitySubmittingTextAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJhsCommunitySubmittingTextAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.submitting.text"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJhsCommunitySubmittingTextAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJhsCommunitySubmittingTextAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaJhsCommunitySubmittingTextAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJhsCommunitySubmittingTextRequest()
	},
}

// GetAlibabaJhsCommunitySubmittingTextRequest 从 sync.Pool 获取 AlibabaJhsCommunitySubmittingTextAPIRequest
func GetAlibabaJhsCommunitySubmittingTextAPIRequest() *AlibabaJhsCommunitySubmittingTextAPIRequest {
	return poolAlibabaJhsCommunitySubmittingTextAPIRequest.Get().(*AlibabaJhsCommunitySubmittingTextAPIRequest)
}

// ReleaseAlibabaJhsCommunitySubmittingTextAPIRequest 将 AlibabaJhsCommunitySubmittingTextAPIRequest 放入 sync.Pool
func ReleaseAlibabaJhsCommunitySubmittingTextAPIRequest(v *AlibabaJhsCommunitySubmittingTextAPIRequest) {
	v.Reset()
	poolAlibabaJhsCommunitySubmittingTextAPIRequest.Put(v)
}
