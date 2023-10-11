package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityUserStatisticsAPIRequest 聚划算社群用户行为上报 API请求
// alibaba.jhs.community.user.statistics
//
// 聚划算社群用户行为上报
type AlibabaJhsCommunityUserStatisticsAPIRequest struct {
	model.Params
	// 用户会话
	_token string
	// 扩展参数
	_extend string
}

// NewAlibabaJhsCommunityUserStatisticsRequest 初始化AlibabaJhsCommunityUserStatisticsAPIRequest对象
func NewAlibabaJhsCommunityUserStatisticsRequest() *AlibabaJhsCommunityUserStatisticsAPIRequest {
	return &AlibabaJhsCommunityUserStatisticsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJhsCommunityUserStatisticsAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.user.statistics"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJhsCommunityUserStatisticsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJhsCommunityUserStatisticsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 用户会话
func (r *AlibabaJhsCommunityUserStatisticsAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaJhsCommunityUserStatisticsAPIRequest) GetToken() string {
	return r._token
}

// SetExtend is Extend Setter
// 扩展参数
func (r *AlibabaJhsCommunityUserStatisticsAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r AlibabaJhsCommunityUserStatisticsAPIRequest) GetExtend() string {
	return r._extend
}
