package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqUrlScanAPIRequest
恶意网址检测接口 API请求
alibaba.security.jaq.url.scan

url扫描接口 */
type AlibabaSecurityJaqUrlScanAPIRequest struct {
	model.Params
	// 扫描参数
	_paramUrlScanParamList *UrlScanParamList
}

// New
