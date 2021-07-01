package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqOcrImageSyncDetectAPIRequest
聚安全图文识别同步检测接口 API请求
alibaba.security.jaq.ocr.image.sync.detect

图像字符识别同步检测接口 */
type AlibabaSecurityJaqOcrImageSyncDetectAPIRequest struct {
	model.Params
	// 待检测图像链接
	_imageUrl string
}

// New
