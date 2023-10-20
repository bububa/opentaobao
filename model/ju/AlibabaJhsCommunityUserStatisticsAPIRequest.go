package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajhscommunityuserstatisticsAPIRequest 聚划算社群用户行为上报 API请求
// alibaba.jhs.community.user.statistics
//
// 聚划算社群用户行为上报
type AlibabajhscommunityuserstatisticsAPIRequest struct {
	model.Params
	// 用户会话
	_token string
	// 扩展参数
	_extend string
}

// NewAlibabajhscommunityuserstatisticsRequest 初始化AlibabajhscommunityuserstatisticsAPIRequest对象
func NewAlibabajhscommunityuserstatisticsRequest() *AlibabajhscommunityuserstatisticsAPIRequest {
	return &AlibabajhscommunityuserstatisticsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajhscommunityuserstatisticsAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.user.statistics"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajhscommunityuserstatisticsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajhscommunityuserstatisticsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 用户会话
func (r *AlibabajhscommunityuserstatisticsAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabajhscommunityuserstatisticsAPIRequest) GetToken() string {
	return r._token
}

// SetExtend is Extend Setter
// 扩展参数
func (r *AlibabajhscommunityuserstatisticsAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r AlibabajhscommunityuserstatisticsAPIRequest) GetExtend() string {
	return r._extend
}
