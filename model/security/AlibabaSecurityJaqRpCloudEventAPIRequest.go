package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpCloudEventAPIRequest
事件上报 API请求
alibaba.security.jaq.rp.cloud.event

事件上报接口 */
type AlibabaSecurityJaqRpCloudEventAPIRequest struct {
	model.Params
	// 认证token
	_verifyToken string
	// 事件编码
	_eventCode string
	// 事件信息
	_eventData string
}

// New
