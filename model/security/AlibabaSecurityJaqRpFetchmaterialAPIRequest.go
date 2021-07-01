package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpFetchmaterialAPIRequest
聚安全实人认证获取结果接口 API请求
alibaba.security.jaq.rp.fetchmaterial

聚安全实人认证获取结果接口 */
type AlibabaSecurityJaqRpFetchmaterialAPIRequest struct {
	model.Params
	// 消息服务推送的key
	_securityKey string
}

// New
