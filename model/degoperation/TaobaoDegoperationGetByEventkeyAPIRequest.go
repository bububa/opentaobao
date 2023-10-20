package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodegoperationgetbyeventkeyAPIRequest 通用用户抽奖次数限制 API请求
// taobao.degoperation.get.by.eventkey
//
// 通用用户抽奖次数限制
type TaobaodegoperationgetbyeventkeyAPIRequest struct {
	model.Params
	// info
	_degAccessToken string
	// 活动后台配置appkey
	_degAppKey string
	// 活动后台配置eventkey
	_eventKey string
}

// NewTaobaodegoperationgetbyeventkeyRequest 初始化TaobaodegoperationgetbyeventkeyAPIRequest对象
func NewTaobaodegoperationgetbyeventkeyRequest() *TaobaodegoperationgetbyeventkeyAPIRequest {
	return &TaobaodegoperationgetbyeventkeyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodegoperationgetbyeventkeyAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.get.by.eventkey"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodegoperationgetbyeventkeyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodegoperationgetbyeventkeyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDegAccessToken is DegAccessToken Setter
// info
func (r *TaobaodegoperationgetbyeventkeyAPIRequest) SetDegAccessToken(_degAccessToken string) error {
	r._degAccessToken = _degAccessToken
	r.Set("deg_access_token", _degAccessToken)
	return nil
}

// GetDegAccessToken DegAccessToken Getter
func (r TaobaodegoperationgetbyeventkeyAPIRequest) GetDegAccessToken() string {
	return r._degAccessToken
}

// SetDegAppKey is DegAppKey Setter
// 活动后台配置appkey
func (r *TaobaodegoperationgetbyeventkeyAPIRequest) SetDegAppKey(_degAppKey string) error {
	r._degAppKey = _degAppKey
	r.Set("deg_app_key", _degAppKey)
	return nil
}

// GetDegAppKey DegAppKey Getter
func (r TaobaodegoperationgetbyeventkeyAPIRequest) GetDegAppKey() string {
	return r._degAppKey
}

// SetEventKey is EventKey Setter
// 活动后台配置eventkey
func (r *TaobaodegoperationgetbyeventkeyAPIRequest) SetEventKey(_eventKey string) error {
	r._eventKey = _eventKey
	r.Set("event_key", _eventKey)
	return nil
}

// GetEventKey EventKey Getter
func (r TaobaodegoperationgetbyeventkeyAPIRequest) GetEventKey() string {
	return r._eventKey
}
