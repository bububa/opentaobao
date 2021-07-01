package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqPornImageSyncDetectAPIRequest
聚安全智能鉴黄同步检测接口 API请求
alibaba.security.jaq.porn.image.sync.detect

同步黄图图像检测接口 */
type AlibabaSecurityJaqPornImageSyncDetectAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// New
