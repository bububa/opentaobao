package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityActivityDetailsAPIRequest 社群活动详情 API请求
// alibaba.jhs.community.activity.details
//
// 社群活动详情
type AlibabaJhsCommunityActivityDetailsAPIRequest struct {
	model.Params
	// 用户会话登录token
	_token string
	// 社群活动id
	_activityId int64
}

// NewAlibabaJhsCommunityActivityDetailsRequest 初始化AlibabaJhsCommunityActivityDetailsAPIRequest对象
func NewAlibabaJhsCommunityActivityDetailsRequest() *AlibabaJhsCommunityActivityDetailsAPIRequest {
	return &AlibabaJhsCommunityActivityDetailsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJhsCommunityActivityDetailsAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.activity.details"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJhsCommunityActivityDetailsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJhsCommunityActivityDetailsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 用户会话登录token
func (r *AlibabaJhsCommunityActivityDetailsAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaJhsCommunityActivityDetailsAPIRequest) GetToken() string {
	return r._token
}

// SetActivityId is ActivityId Setter
// 社群活动id
func (r *AlibabaJhsCommunityActivityDetailsAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabaJhsCommunityActivityDetailsAPIRequest) GetActivityId() int64 {
	return r._activityId
}
