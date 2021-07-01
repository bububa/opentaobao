package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpCloudOcrCheckAPIRequest
ocr同时实名校验 API请求
alibaba.security.jaq.rp.cloud.ocr.check

聚安全实人认证证件OCR识别功能接口 */
type AlibabaSecurityJaqRpCloudOcrCheckAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 要识别的信息
	_imageUrls string
}

// New
