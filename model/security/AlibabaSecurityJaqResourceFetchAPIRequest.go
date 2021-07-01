package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqResourceFetchAPIRequest
获取资源文件 API请求
alibaba.security.jaq.resource.fetch

在前向化验证流程中提供资源文件服务 */
type AlibabaSecurityJaqResourceFetchAPIRequest struct {
	model.Params
	// 设备类型可能值有：android ios wp
	_deviceType string
	// 分辨率
	_dpi string
	// 语言类型 zh_CN en_US
	_lang string
}

// New
