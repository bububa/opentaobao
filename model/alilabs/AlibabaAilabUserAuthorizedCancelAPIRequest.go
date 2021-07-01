package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabUserAuthorizedCancelAPIRequest
取消账号授权 API请求
alibaba.ailab.user.authorized.cancel

三方用户取消授权给天猫精灵用户 */
type AlibabaAilabUserAuthorizedCancelAPIRequest struct {
	model.Params
	// 三方用户的唯一ID
	_merchantUserId string
	// 开放平台申请的schema
	_schemaKey string
}

// New
