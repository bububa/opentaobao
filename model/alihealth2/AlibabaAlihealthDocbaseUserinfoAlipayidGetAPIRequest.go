package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest 根据健康ID获取支付宝ID API请求
// alibaba.alihealth.docbase.userinfo.alipayid.get
//
// 根据健康ID获取支付宝ID
type AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest struct {
	model.Params
	// 阿里健康ID
	_alihealthUserId string
	// 渠道alipay taobao uc gaode
	_appChannel string
}

// NewAlibabaalihealthdocbaseuserinfoalipayidgetRequest 初始化AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest对象
func NewAlibabaalihealthdocbaseuserinfoalipayidgetRequest() *AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest {
	return &AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.docbase.userinfo.alipayid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlihealthUserId is AlihealthUserId Setter
// 阿里健康ID
func (r *AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest) SetAlihealthUserId(_alihealthUserId string) error {
	r._alihealthUserId = _alihealthUserId
	r.Set("alihealth_user_id", _alihealthUserId)
	return nil
}

// GetAlihealthUserId AlihealthUserId Getter
func (r AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest) GetAlihealthUserId() string {
	return r._alihealthUserId
}

// SetAppChannel is AppChannel Setter
// 渠道alipay taobao uc gaode
func (r *AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest) SetAppChannel(_appChannel string) error {
	r._appChannel = _appChannel
	r.Set("app_channel", _appChannel)
	return nil
}

// GetAppChannel AppChannel Getter
func (r AlibabaalihealthdocbaseuserinfoalipayidgetAPIRequest) GetAppChannel() string {
	return r._appChannel
}
