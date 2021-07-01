package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqCaptchaSendAPIRequest
聚安全安全验证发起接口 API请求
alibaba.security.jaq.captcha.send

聚安全安全验证发起 */
type AlibabaSecurityJaqCaptchaSendAPIRequest struct {
	model.Params
	// 验证码发送渠道类型 1-短信 3-邮件
	_captchaType int64
	// 扩展字段，格式为JSON字符串，用于传递“滚小球”等验证方式所需的额外入参，例如屏幕尺寸等，请参考示例
	_extendData string
	// 协议版本号
	_protocolVersion string
	// 软token签名
	_rsign string
	// 软token索引
	_rtkenIndex string
	// UMID token
	_utoken string
	// 风险识别接口返回的信息口令
	_infoToken string
}

// New
