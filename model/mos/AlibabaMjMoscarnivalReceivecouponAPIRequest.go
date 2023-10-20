package mos

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjMoscarnivalReceivecouponAPIRequest) Reset() {
	r._mobile = ""
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMoscarnivalReceivecouponAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.moscarnival.receivecoupon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMoscarnivalReceivecouponAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjMoscarnivalReceivecouponAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMobile is Mobile Setter
// 手机号码
func (r *AlibabaMjMoscarnivalReceivecouponAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaMjMoscarnivalReceivecouponAPIRequest) GetMobile() string {
	return r._mobile
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *AlibabaMjMoscarnivalReceivecouponAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabaMjMoscarnivalReceivecouponAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolAlibabaMjMoscarnivalReceivecouponAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjMoscarnivalReceivecouponRequest()
	},
}

// GetAlibabaMjMoscarnivalReceivecouponRequest 从 sync.Pool 获取 AlibabaMjMoscarnivalReceivecouponAPIRequest
func GetAlibabaMjMoscarnivalReceivecouponAPIRequest() *AlibabaMjMoscarnivalReceivecouponAPIRequest {
	return poolAlibabaMjMoscarnivalReceivecouponAPIRequest.Get().(*AlibabaMjMoscarnivalReceivecouponAPIRequest)
}

// ReleaseAlibabaMjMoscarnivalReceivecouponAPIRequest 将 AlibabaMjMoscarnivalReceivecouponAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjMoscarnivalReceivecouponAPIRequest(v *AlibabaMjMoscarnivalReceivecouponAPIRequest) {
	v.Reset()
	poolAlibabaMjMoscarnivalReceivecouponAPIRequest.Put(v)
}
