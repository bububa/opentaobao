package alitripbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbpchannelcrowqueryAPIRequest 人群匹配 API请求
// alitrip.bp.channel.crow.query
//
// 判断用户是否在圈选的人群中
type AlitripbpchannelcrowqueryAPIRequest struct {
	model.Params
	// 查询入参
	_queryParam *ExamineOuterUserRequest
}

// NewAlitripbpchannelcrowqueryRequest 初始化AlitripbpchannelcrowqueryAPIRequest对象
func NewAlitripbpchannelcrowqueryRequest() *AlitripbpchannelcrowqueryAPIRequest {
	return &AlitripbpchannelcrowqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbpchannelcrowqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.bp.channel.crow.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbpchannelcrowqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbpchannelcrowqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryParam is QueryParam Setter
// 查询入参
func (r *AlitripbpchannelcrowqueryAPIRequest) SetQueryParam(_queryParam *ExamineOuterUserRequest) error {
	r._queryParam = _queryParam
	r.Set("query_param", _queryParam)
	return nil
}

// GetQueryParam QueryParam Getter
func (r AlitripbpchannelcrowqueryAPIRequest) GetQueryParam() *ExamineOuterUserRequest {
	return r._queryParam
}
