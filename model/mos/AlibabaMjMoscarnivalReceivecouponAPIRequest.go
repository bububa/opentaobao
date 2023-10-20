package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjmoscarnivalreceivecouponAPIRequest 根据手机号码领券 API请求
// alibaba.mj.moscarnival.receivecoupon
//
// 根据手机号码领券
type AlibabamjmoscarnivalreceivecouponAPIRequest struct {
	model.Params
	// 手机号码
	_mobile string
	// 活动id
	_activityId int64
}

// NewAlibabamjmoscarnivalreceivecouponRequest 初始化AlibabamjmoscarnivalreceivecouponAPIRequest对象
func NewAlibabamjmoscarnivalreceivecouponRequest() *AlibabamjmoscarnivalreceivecouponAPIRequest {
	return &AlibabamjmoscarnivalreceivecouponAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjmoscarnivalreceivecouponAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.moscarnival.receivecoupon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjmoscarnivalreceivecouponAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjmoscarnivalreceivecouponAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMobile is Mobile Setter
// 手机号码
func (r *AlibabamjmoscarnivalreceivecouponAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabamjmoscarnivalreceivecouponAPIRequest) GetMobile() string {
	return r._mobile
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *AlibabamjmoscarnivalreceivecouponAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabamjmoscarnivalreceivecouponAPIRequest) GetActivityId() int64 {
	return r._activityId
}
