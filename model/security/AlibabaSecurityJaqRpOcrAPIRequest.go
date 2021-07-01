package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpOcrAPIRequest
聚安全实人认证证件OCR识别功能接口 API请求
alibaba.security.jaq.rp.ocr

聚安全实人认证证件OCR识别功能接口 */
type AlibabaSecurityJaqRpOcrAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 要识别的信息
	_imageUrls string
}

// New
