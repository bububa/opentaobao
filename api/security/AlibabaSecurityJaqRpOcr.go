package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqRpOcr 聚安全实人认证证件OCR识别功能接口
// alibaba.security.jaq.rp.ocr
//
// 聚安全实人认证证件OCR识别功能接口
func AlibabaSecurityJaqRpOcr(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpOcrAPIRequest, resp *security.AlibabaSecurityJaqRpOcrAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
