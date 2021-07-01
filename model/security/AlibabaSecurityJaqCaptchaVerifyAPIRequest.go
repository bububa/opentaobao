package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqCaptchaVerifyAPIRequest
聚安全安全验证检查接口 API请求
alibaba.security.jaq.captcha.verify

聚安全安全验证检查 */
type AlibabaSecurityJaqCaptchaVerifyAPIRequest struct {
	model.Params
	// 验证码发送渠道类型 1-短信 2-语音 3-邮件
	_captchaType int64
	// 扩展字段，格式为JSON字符串，用于传递“短信验证”等验证方式所需的额外入参，例如用户输入的验证码等，格式及JSON字段key定义请参考示例
	_extendData string
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
	// 风险识别接口返回的信息口令
	_infoToken string
}

// New
