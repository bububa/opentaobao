package cloudgame

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameOpenidQueryAPIRequest 咖哒用户ID查询 API请求
// alibaba.cloudgame.openid.query
//
// 云游戏业务需要提供查询用户信息的TOP能力
type AlibabaCloudgameOpenidQueryAPIRequest struct {
	model.Params
	// 入参model
	_requestParam *HavanaUserIdQueryRequest
}

// NewAlibabaCloudgameOpenidQueryRequest 初始化AlibabaCloudgameOpenidQueryAPIRequest对象
func NewAlibabaCloudgameOpenidQueryRequest() *AlibabaCloudgameOpenidQueryAPIRequest {
	return &AlibabaCloudgameOpenidQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCloudgameOpenidQueryAPIRequest) Reset() {
	r._requestParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameOpenidQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.openid.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameOpenidQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameOpenidQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestParam is RequestParam Setter
// 入参model
func (r *AlibabaCloudgameOpenidQueryAPIRequest) SetRequestParam(_requestParam *HavanaUserIdQueryRequest) error {
	r._requestParam = _requestParam
	r.Set("request_param", _requestParam)
	return nil
}

// GetRequestParam RequestParam Getter
func (r AlibabaCloudgameOpenidQueryAPIRequest) GetRequestParam() *HavanaUserIdQueryRequest {
	return r._requestParam
}

var poolAlibabaCloudgameOpenidQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCloudgameOpenidQueryRequest()
	},
}

// GetAlibabaCloudgameOpenidQueryRequest 从 sync.Pool 获取 AlibabaCloudgameOpenidQueryAPIRequest
func GetAlibabaCloudgameOpenidQueryAPIRequest() *AlibabaCloudgameOpenidQueryAPIRequest {
	return poolAlibabaCloudgameOpenidQueryAPIRequest.Get().(*AlibabaCloudgameOpenidQueryAPIRequest)
}

// ReleaseAlibabaCloudgameOpenidQueryAPIRequest 将 AlibabaCloudgameOpenidQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaCloudgameOpenidQueryAPIRequest(v *AlibabaCloudgameOpenidQueryAPIRequest) {
	v.Reset()
	poolAlibabaCloudgameOpenidQueryAPIRequest.Put(v)
}
