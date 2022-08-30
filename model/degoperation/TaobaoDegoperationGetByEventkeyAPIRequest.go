package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationGetByEventkeyAPIRequest 通用用户抽奖次数限制 API请求
// taobao.degoperation.get.by.eventkey
//
// 通用用户抽奖次数限制
type TaobaoDegoperationGetByEventkeyAPIRequest struct {
	model.Params
	// info
	_degAccessToken string
	// 活动后台配置appkey
	_degAppKey string
	// 活动后台配置eventkey
	_eventKey string
}

// NewTaobaoDegoperationGetByEventkeyRequest 初始化TaobaoDegoperationGetByEventkeyAPIRequest对象
func NewTaobaoDegoperationGetByEventkeyRequest() *TaobaoDegoperationGetByEventkeyAPIRequest {
	return &TaobaoDegoperationGetByEventkeyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationGetByEventkeyAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.get.by.eventkey"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationGetByEventkeyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDegAccessToken is DegAccessToken Setter
// info
func (r *TaobaoDegoperationGetByEventkeyAPIRequest) SetDegAccessToken(_degAccessToken string) error {
	r._degAccessToken = _degAccessToken
	r.Set("deg_access_token", _degAccessToken)
	return nil
}

// GetDegAccessToken DegAccessToken Getter
func (r TaobaoDegoperationGetByEventkeyAPIRequest) GetDegAccessToken() string {
	return r._degAccessToken
}

// SetDegAppKey is DegAppKey Setter
// 活动后台配置appkey
func (r *TaobaoDegoperationGetByEventkeyAPIRequest) SetDegAppKey(_degAppKey string) error {
	r._degAppKey = _degAppKey
	r.Set("deg_app_key", _degAppKey)
	return nil
}

// GetDegAppKey DegAppKey Getter
func (r TaobaoDegoperationGetByEventkeyAPIRequest) GetDegAppKey() string {
	return r._degAppKey
}

// SetEventKey is EventKey Setter
// 活动后台配置eventkey
func (r *TaobaoDegoperationGetByEventkeyAPIRequest) SetEventKey(_eventKey string) error {
	r._eventKey = _eventKey
	r.Set("event_key", _eventKey)
	return nil
}

// GetEventKey EventKey Getter
func (r TaobaoDegoperationGetByEventkeyAPIRequest) GetEventKey() string {
	return r._eventKey
}
