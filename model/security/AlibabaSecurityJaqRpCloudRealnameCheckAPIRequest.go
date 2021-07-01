package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest
验证姓名和证件号 API请求
alibaba.security.jaq.rp.cloud.realname.check

验证姓名和证件号 */
type AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 要识别的信息
	_imageUrls string
	// 姓名
	_name string
	// 证件号
	_identityCode string
}

// New
