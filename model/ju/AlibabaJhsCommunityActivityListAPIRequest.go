package ju

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityActivityListAPIRequest 聚划算用增淘外社群服务活动列表 API请求
// alibaba.jhs.community.activity.list
//
// 聚划算用增淘外社群服务活动列表
type AlibabaJhsCommunityActivityListAPIRequest struct {
	model.Params
}

// NewAlibabaJhsCommunityActivityListRequest 初始化AlibabaJhsCommunityActivityListAPIRequest对象
func NewAlibabaJhsCommunityActivityListRequest() *AlibabaJhsCommunityActivityListAPIRequest {
	return &AlibabaJhsCommunityActivityListAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJhsCommunityActivityListAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJhsCommunityActivityListAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.activity.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJhsCommunityActivityListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJhsCommunityActivityListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaJhsCommunityActivityListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJhsCommunityActivityListRequest()
	},
}

// GetAlibabaJhsCommunityActivityListRequest 从 sync.Pool 获取 AlibabaJhsCommunityActivityListAPIRequest
func GetAlibabaJhsCommunityActivityListAPIRequest() *AlibabaJhsCommunityActivityListAPIRequest {
	return poolAlibabaJhsCommunityActivityListAPIRequest.Get().(*AlibabaJhsCommunityActivityListAPIRequest)
}

// ReleaseAlibabaJhsCommunityActivityListAPIRequest 将 AlibabaJhsCommunityActivityListAPIRequest 放入 sync.Pool
func ReleaseAlibabaJhsCommunityActivityListAPIRequest(v *AlibabaJhsCommunityActivityListAPIRequest) {
	v.Reset()
	poolAlibabaJhsCommunityActivityListAPIRequest.Put(v)
}
