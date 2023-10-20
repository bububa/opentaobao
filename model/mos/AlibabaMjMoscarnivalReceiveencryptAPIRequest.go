package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMoscarnivalReceiveencryptAPIRequest 根据加密手机号领券 API请求
// alibaba.mj.moscarnival.receiveencrypt
//
// 根据加密手机号领券
type AlibabaMjMoscarnivalReceiveencryptAPIRequest struct {
	model.Params
	// 加密手机号码
	_mobile string
	// 活动id
	_activityId int64
}

// NewAlibabaMjMoscarnivalReceiveencryptRequest 初始化AlibabaMjMoscarnivalReceiveencryptAPIRequest对象
func NewAlibabaMjMoscarnivalReceiveencryptRequest() *AlibabaMjMoscarnivalReceiveencryptAPIRequest {
	return &AlibabaMjMoscarnivalReceiveencryptAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjMoscarnivalReceiveencryptAPIRequest) Reset() {
	r._mobile = ""
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMoscarnivalReceiveencryptAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.moscarnival.receiveencrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMoscarnivalReceiveencryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjMoscarnivalReceiveencryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMobile is Mobile Setter
// 加密手机号码
func (r *AlibabaMjMoscarnivalReceiveencryptAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaMjMoscarnivalReceiveencryptAPIRequest) GetMobile() string {
	return r._mobile
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *AlibabaMjMoscarnivalReceiveencryptAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabaMjMoscarnivalReceiveencryptAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolAlibabaMjMoscarnivalReceiveencryptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjMoscarnivalReceiveencryptRequest()
	},
}

// GetAlibabaMjMoscarnivalReceiveencryptRequest 从 sync.Pool 获取 AlibabaMjMoscarnivalReceiveencryptAPIRequest
func GetAlibabaMjMoscarnivalReceiveencryptAPIRequest() *AlibabaMjMoscarnivalReceiveencryptAPIRequest {
	return poolAlibabaMjMoscarnivalReceiveencryptAPIRequest.Get().(*AlibabaMjMoscarnivalReceiveencryptAPIRequest)
}

// ReleaseAlibabaMjMoscarnivalReceiveencryptAPIRequest 将 AlibabaMjMoscarnivalReceiveencryptAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjMoscarnivalReceiveencryptAPIRequest(v *AlibabaMjMoscarnivalReceiveencryptAPIRequest) {
	v.Reset()
	poolAlibabaMjMoscarnivalReceiveencryptAPIRequest.Put(v)
}
