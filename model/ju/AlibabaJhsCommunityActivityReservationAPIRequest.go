package ju

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityActivityReservationAPIRequest 社群活动预约 API请求
// alibaba.jhs.community.activity.reservation
//
// 社群活动预约
type AlibabaJhsCommunityActivityReservationAPIRequest struct {
	model.Params
	// 会话token
	_token string
	// 活动id
	_activityId int64
}

// NewAlibabaJhsCommunityActivityReservationRequest 初始化AlibabaJhsCommunityActivityReservationAPIRequest对象
func NewAlibabaJhsCommunityActivityReservationRequest() *AlibabaJhsCommunityActivityReservationAPIRequest {
	return &AlibabaJhsCommunityActivityReservationAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJhsCommunityActivityReservationAPIRequest) Reset() {
	r._token = ""
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJhsCommunityActivityReservationAPIRequest) GetApiMethodName() string {
	return "alibaba.jhs.community.activity.reservation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJhsCommunityActivityReservationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJhsCommunityActivityReservationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 会话token
func (r *AlibabaJhsCommunityActivityReservationAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaJhsCommunityActivityReservationAPIRequest) GetToken() string {
	return r._token
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *AlibabaJhsCommunityActivityReservationAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabaJhsCommunityActivityReservationAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolAlibabaJhsCommunityActivityReservationAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJhsCommunityActivityReservationRequest()
	},
}

// GetAlibabaJhsCommunityActivityReservationRequest 从 sync.Pool 获取 AlibabaJhsCommunityActivityReservationAPIRequest
func GetAlibabaJhsCommunityActivityReservationAPIRequest() *AlibabaJhsCommunityActivityReservationAPIRequest {
	return poolAlibabaJhsCommunityActivityReservationAPIRequest.Get().(*AlibabaJhsCommunityActivityReservationAPIRequest)
}

// ReleaseAlibabaJhsCommunityActivityReservationAPIRequest 将 AlibabaJhsCommunityActivityReservationAPIRequest 放入 sync.Pool
func ReleaseAlibabaJhsCommunityActivityReservationAPIRequest(v *AlibabaJhsCommunityActivityReservationAPIRequest) {
	v.Reset()
	poolAlibabaJhsCommunityActivityReservationAPIRequest.Put(v)
}
