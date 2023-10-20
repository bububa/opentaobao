package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajhscommunityactivityreservationAPIRequest 社群活动预约 API请求
// alibaba.jhs.community.activity.reservation
//
// 社群活动预约
type AlibabajhscommunityactivityreservationAPIRequest struct {
	model.Params
	// 会话token
	_token string
	// 活动id
	_activityId int64
}

// NewAlibabajhscommunityactivityreservationRequest 初始化AlibabajhscommunityactivityreservationAPIRequest对象
func NewAlibabajhscommunityactivityreservationRequest() *AlibabajhscommunityactivityreservationAPIRequest {
	return &AlibabajhscommunityactivityreservationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajhscommunityactivityreservationAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.activity.reservation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajhscommunityactivityreservationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajhscommunityactivityreservationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 会话token
func (r *AlibabajhscommunityactivityreservationAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabajhscommunityactivityreservationAPIRequest) GetToken() string {
	return r._token
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *AlibabajhscommunityactivityreservationAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabajhscommunityactivityreservationAPIRequest) GetActivityId() int64 {
	return r._activityId
}
