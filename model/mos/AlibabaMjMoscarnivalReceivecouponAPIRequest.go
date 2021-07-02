package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMoscarnivalReceivecouponAPIRequest 根据手机号码领券 API请求
// alibaba.mj.moscarnival.receivecoupon
//
// 根据手机号码领券
type AlibabaMjMoscarnivalReceivecouponAPIRequest struct {
	model.Params
	// 手机号码
	_mobile string
	// 活动id
	_activityId int64
}

// NewAlibabaMjMoscarnivalReceivecouponRequest 初始化AlibabaMjMoscarnivalReceivecouponAPIRequest对象
func NewAlibabaMjMoscarnivalReceivecouponRequest() *AlibabaMjMoscarnivalReceivecouponAPIRequest {
	return &AlibabaMjMoscarnivalReceivecouponAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMoscarnivalReceivecouponAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.moscarnival.receivecoupon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMoscarnivalReceivecouponAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Mobile Setter
// 手机号码
func (r *AlibabaMjMoscarnivalReceivecouponAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// Get Mobile Getter
func (r AlibabaMjMoscarnivalReceivecouponAPIRequest) GetMobile() string {
	return r._mobile
}

// Set is ActivityId Setter
// 活动id
func (r *AlibabaMjMoscarnivalReceivecouponAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r AlibabaMjMoscarnivalReceivecouponAPIRequest) GetActivityId() int64 {
	return r._activityId
}
