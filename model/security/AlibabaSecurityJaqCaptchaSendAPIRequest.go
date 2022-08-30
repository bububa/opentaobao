package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqCaptchaSendAPIRequest 聚安全安全验证发起接口 API请求
// alibaba.security.jaq.captcha.send
//
// 聚安全安全验证发起
type AlibabaSecurityJaqCaptchaSendAPIRequest struct {
	model.Params
	// 扩展字段，格式为JSON字符串，用于传递“滚小球”等验证方式所需的额外入参，例如屏幕尺寸等，请参考示例
	_extendData string
	// 风险识别接口返回的信息口令
	_infoToken string
	// 协议版本号
	_protocolVersion string
	// 软token签名
	_rsign string
	// 软token索引
	_rtkenIndex string
	// UMID token
	_utoken string
	// 验证码发送渠道类型 1-短信 3-邮件
	_captchaType int64
}

// NewAlibabaSecurityJaqCaptchaSendRequest 初始化AlibabaSecurityJaqCaptchaSendAPIRequest对象
func NewAlibabaSecurityJaqCaptchaSendRequest() *AlibabaSecurityJaqCaptchaSendAPIRequest {
	return &AlibabaSecurityJaqCaptchaSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqCaptchaSendAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.captcha.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqCaptchaSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExtendData is ExtendData Setter
// 扩展字段，格式为JSON字符串，用于传递“滚小球”等验证方式所需的额外入参，例如屏幕尺寸等，请参考示例
func (r *AlibabaSecurityJaqCaptchaSendAPIRequest) SetExtendData(_extendData string) error {
	r._extendData = _extendData
	r.Set("extend_data", _extendData)
	return nil
}

// GetExtendData ExtendData Getter
func (r AlibabaSecurityJaqCaptchaSendAPIRequest) GetExtendData() string {
	return r._extendData
}

// SetInfoToken is InfoToken Setter
// 风险识别接口返回的信息口令
func (r *AlibabaSecurityJaqCaptchaSendAPIRequest) SetInfoToken(_infoToken string) error {
	r._infoToken = _infoToken
	r.Set("info_token", _infoToken)
	return nil
}

// GetInfoToken InfoToken Getter
func (r AlibabaSecurityJaqCaptchaSendAPIRequest) GetInfoToken() string {
	return r._infoToken
}

// SetProtocolVersion is ProtocolVersion Setter
// 协议版本号
func (r *AlibabaSecurityJaqCaptchaSendAPIRequest) SetProtocolVersion(_protocolVersion string) error {
	r._protocolVersion = _protocolVersion
	r.Set("protocol_version", _protocolVersion)
	return nil
}

// GetProtocolVersion ProtocolVersion Getter
func (r AlibabaSecurityJaqCaptchaSendAPIRequest) GetProtocolVersion() string {
	return r._protocolVersion
}

// SetRsign is Rsign Setter
// 软token签名
func (r *AlibabaSecurityJaqCaptchaSendAPIRequest) SetRsign(_rsign string) error {
	r._rsign = _rsign
	r.Set("rsign", _rsign)
	return nil
}

// GetRsign Rsign Getter
func (r AlibabaSecurityJaqCaptchaSendAPIRequest) GetRsign() string {
	return r._rsign
}

// SetRtkenIndex is RtkenIndex Setter
// 软token索引
func (r *AlibabaSecurityJaqCaptchaSendAPIRequest) SetRtkenIndex(_rtkenIndex string) error {
	r._rtkenIndex = _rtkenIndex
	r.Set("rtken_index", _rtkenIndex)
	return nil
}

// GetRtkenIndex RtkenIndex Getter
func (r AlibabaSecurityJaqCaptchaSendAPIRequest) GetRtkenIndex() string {
	return r._rtkenIndex
}

// SetUtoken is Utoken Setter
// UMID token
func (r *AlibabaSecurityJaqCaptchaSendAPIRequest) SetUtoken(_utoken string) error {
	r._utoken = _utoken
	r.Set("utoken", _utoken)
	return nil
}

// GetUtoken Utoken Getter
func (r AlibabaSecurityJaqCaptchaSendAPIRequest) GetUtoken() string {
	return r._utoken
}

// SetCaptchaType is CaptchaType Setter
// 验证码发送渠道类型 1-短信 3-邮件
func (r *AlibabaSecurityJaqCaptchaSendAPIRequest) SetCaptchaType(_captchaType int64) error {
	r._captchaType = _captchaType
	r.Set("captcha_type", _captchaType)
	return nil
}

// GetCaptchaType CaptchaType Getter
func (r AlibabaSecurityJaqCaptchaSendAPIRequest) GetCaptchaType() int64 {
	return r._captchaType
}
