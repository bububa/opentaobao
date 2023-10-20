package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthuicuserinfohealthidgetAPIRequest 获取健康id API请求
// alibaba.alihealth.uic.userinfo.healthid.get
//
// 根据支付宝用户ID获取用户健康ID
type AlibabaalihealthuicuserinfohealthidgetAPIRequest struct {
	model.Params
	// 支付宝用户ID
	_alipayUserId string
}

// NewAlibabaalihealthuicuserinfohealthidgetRequest 初始化AlibabaalihealthuicuserinfohealthidgetAPIRequest对象
func NewAlibabaalihealthuicuserinfohealthidgetRequest() *AlibabaalihealthuicuserinfohealthidgetAPIRequest {
	return &AlibabaalihealthuicuserinfohealthidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthuicuserinfohealthidgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.uic.userinfo.healthid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthuicuserinfohealthidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthuicuserinfohealthidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayUserId is AlipayUserId Setter
// 支付宝用户ID
func (r *AlibabaalihealthuicuserinfohealthidgetAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r AlibabaalihealthuicuserinfohealthidgetAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}
