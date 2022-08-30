package alitripbp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBpChannelCrowQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.bp.channel.crow.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBpChannelCrowQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
