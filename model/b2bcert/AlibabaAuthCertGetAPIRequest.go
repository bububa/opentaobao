package b2bcert

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAuthCertGetAPIRequest
获取证书数据 API请求
alibaba.auth.cert.get

获取证书数据 */
type AlibabaAuthCertGetAPIRequest struct {
	model.Params
	// 认证商
	_provider string
	// 证书数据
	_receiveInfo string
}

// New
