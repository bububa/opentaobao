package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqCaptchaVerifyAPIRequest 聚安全安全验证检查接口 API请求
// alibaba.security.jaq.captcha.verify
//
// 聚安全安全验证检查
type AlibabaSecurityJaqCaptchaVerifyAPIRequest struct {
	model.Params
	// 扩展字段，格式为JSON字符串，用于传递“短信验证”等验证方式所需的额外入参，例如用户输入的验证码等，格式及JSON字段key定义请参考示例
	_extendData string
	// 风险识别接口返回的信息口令
	_infoToken string
	// 协议版本号
	_protocolVersion string
	// 软token签名
	_rsign string
	// 软token索引
	_rtkenIndex string
	// 验证码会话ID
	_sessionId string
	// UMID token
	_utoken string
	// 验证码发送渠道类型 1-短信 2-语音 3-邮件
	_captchaType int64
}

// NewAlibabaSecurityJaqCaptchaVerifyRequest 初始化AlibabaSecurityJaqCaptchaVerifyAPIRequest对象
func NewAlibabaSecurityJaqCaptchaVerifyRequest() *AlibabaSecurityJaqCaptchaVerifyAPIRequest {
	return &AlibabaSecurityJaqCaptchaVerifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqCaptchaVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.captcha.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqCaptchaVerifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExtendData is ExtendData Setter
// 扩展字段，格式为JSON字符串，用于传递“短信验证”等验证方式所需的额外入参，例如用户输入的验证码等，格式及JSON字段key定义请参考示例
func (r *AlibabaSecurityJaqCaptchaVerifyAPIRequest) SetExtendData(_extendData string) error {
	r._extendData = _extendData
	r.Set("extend_data", _extendData)
	return nil
}

// GetExtendData ExtendData Getter
func (r AlibabaSecurityJaqCaptchaVerifyAPIRequest) GetExtendData() string {
	return r._extendData
}

// SetInfoToken is InfoToken Setter
// 风险识别接口返回的信息口令
func (r *AlibabaSecurityJaqCaptchaVerifyAPIRequest) SetInfoToken(_infoToken string) error {
	r._infoToken = _infoToken
	r.Set("info_token", _infoToken)
	return nil
}

// GetInfoToken InfoToken Getter
func (r AlibabaSecurityJaqCaptchaVerifyAPIRequest) GetInfoToken() string {
	return r._infoToken
}

// SetProtocolVersion is ProtocolVersion Setter
// 协议版本号
func (r *AlibabaSecurityJaqCaptchaVerifyAPIRequest) SetProtocolVersion(_protocolVersion string) error {
	r._protocolVersion = _protocolVersion
	r.Set("protocol_version", _protocolVersion)
	return nil
}

// GetProtocolVersion ProtocolVersion Getter
func (r AlibabaSecurityJaqCaptchaVerifyAPIRequest) GetProtocolVersion() string {
	return r._protocolVersion
}

// SetRsign is Rsign Setter
// 软token签名
func (r *AlibabaSecurityJaqCaptchaVerifyAPIRequest) SetRsign(_rsign string) error {
	r._rsign = _rsign
	r.Set("rsign", _rsign)
	return nil
}

// GetRsign Rsign Getter
func (r AlibabaSecurityJaqCaptchaVerifyAPIRequest) GetRsign() string {
	return r._rsign
}

// SetRtkenIndex is RtkenIndex Setter
// 软token索引
func (r *AlibabaSecurityJaqCaptchaVerifyAPIRequest) SetRtkenIndex(_rtkenIndex string) error {
	r._rtkenIndex = _rtkenIndex
	r.Set("rtken_index", _rtkenIndex)
	return nil
}

// GetRtkenIndex RtkenIndex Getter
func (r AlibabaSecurityJaqCaptchaVerifyAPIRequest) GetRtkenIndex() string {
	return r._rtkenIndex
}

// SetSessionId is SessionId Setter
// 验证码会话ID
func (r *AlibabaSecurityJaqCaptchaVerifyAPIRequest) SetSessionId(_sessionId string) error {
	r._sessionId = _sessionId
	r.Set("session_id", _sessionId)
	return nil
}

// GetSessionId SessionId Getter
func (r AlibabaSecurityJaqCaptchaVerifyAPIRequest) GetSessionId() string {
	return r._sessionId
}

// SetUtoken is Utoken Setter
// UMID token
func (r *AlibabaSecurityJaqCaptchaVerifyAPIRequest) SetUtoken(_utoken string) error {
	r._utoken = _utoken
	r.Set("utoken", _utoken)
	return nil
}

// GetUtoken Utoken Getter
func (r AlibabaSecurityJaqCaptchaVerifyAPIRequest) GetUtoken() string {
	return r._utoken
}

// SetCaptchaType is CaptchaType Setter
// 验证码发送渠道类型 1-短信 2-语音 3-邮件
func (r *AlibabaSecurityJaqCaptchaVerifyAPIRequest) SetCaptchaType(_captchaType int64) error {
	r._captchaType = _captchaType
	r.Set("captcha_type", _captchaType)
	return nil
}

// GetCaptchaType CaptchaType Getter
func (r AlibabaSecurityJaqCaptchaVerifyAPIRequest) GetCaptchaType() int64 {
	return r._captchaType
}
