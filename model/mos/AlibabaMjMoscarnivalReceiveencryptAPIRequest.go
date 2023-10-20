package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjmoscarnivalreceiveencryptAPIRequest 根据加密手机号领券 API请求
// alibaba.mj.moscarnival.receiveencrypt
//
// 根据加密手机号领券
type AlibabamjmoscarnivalreceiveencryptAPIRequest struct {
	model.Params
	// 加密手机号码
	_mobile string
	// 活动id
	_activityId int64
}

// NewAlibabamjmoscarnivalreceiveencryptRequest 初始化AlibabamjmoscarnivalreceiveencryptAPIRequest对象
func NewAlibabamjmoscarnivalreceiveencryptRequest() *AlibabamjmoscarnivalreceiveencryptAPIRequest {
	return &AlibabamjmoscarnivalreceiveencryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjmoscarnivalreceiveencryptAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.moscarnival.receiveencrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjmoscarnivalreceiveencryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjmoscarnivalreceiveencryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMobile is Mobile Setter
// 加密手机号码
func (r *AlibabamjmoscarnivalreceiveencryptAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabamjmoscarnivalreceiveencryptAPIRequest) GetMobile() string {
	return r._mobile
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *AlibabamjmoscarnivalreceiveencryptAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabamjmoscarnivalreceiveencryptAPIRequest) GetActivityId() int64 {
	return r._activityId
}
