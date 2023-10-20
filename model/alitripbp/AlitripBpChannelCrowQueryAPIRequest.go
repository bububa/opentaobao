package alitripbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBpChannelCrowQueryAPIRequest 人群匹配 API请求
// alitrip.bp.channel.crow.query
//
// 判断用户是否在圈选的人群中
type AlitripBpChannelCrowQueryAPIRequest struct {
	model.Params
	// 查询入参
	_queryParam *ExamineOuterUserRequest
}

// NewAlitripBpChannelCrowQueryRequest 初始化AlitripBpChannelCrowQueryAPIRequest对象
func NewAlitripBpChannelCrowQueryRequest() *AlitripBpChannelCrowQueryAPIRequest {
	return &AlitripBpChannelCrowQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBpChannelCrowQueryAPIRequest) Reset() {
	r._queryParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBpChannelCrowQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.bp.channel.crow.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBpChannelCrowQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBpChannelCrowQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryParam is QueryParam Setter
// 查询入参
func (r *AlitripBpChannelCrowQueryAPIRequest) SetQueryParam(_queryParam *ExamineOuterUserRequest) error {
	r._queryParam = _queryParam
	r.Set("query_param", _queryParam)
	return nil
}

// GetQueryParam QueryParam Getter
func (r AlitripBpChannelCrowQueryAPIRequest) GetQueryParam() *ExamineOuterUserRequest {
	return r._queryParam
}

var poolAlitripBpChannelCrowQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBpChannelCrowQueryRequest()
	},
}

// GetAlitripBpChannelCrowQueryRequest 从 sync.Pool 获取 AlitripBpChannelCrowQueryAPIRequest
func GetAlitripBpChannelCrowQueryAPIRequest() *AlitripBpChannelCrowQueryAPIRequest {
	return poolAlitripBpChannelCrowQueryAPIRequest.Get().(*AlitripBpChannelCrowQueryAPIRequest)
}

// ReleaseAlitripBpChannelCrowQueryAPIRequest 将 AlitripBpChannelCrowQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripBpChannelCrowQueryAPIRequest(v *AlitripBpChannelCrowQueryAPIRequest) {
	v.Reset()
	poolAlitripBpChannelCrowQueryAPIRequest.Put(v)
}
