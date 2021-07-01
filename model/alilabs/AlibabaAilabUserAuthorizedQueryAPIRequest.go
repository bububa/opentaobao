package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabUserAuthorizedQueryAPIRequest
查询授权状态接口 API请求
alibaba.ailab.user.authorized.query

查询三方用户授权状态 */
type AlibabaAilabUserAuthorizedQueryAPIRequest struct {
	model.Params
	// 开放平台申请的schema
	_schemaKey string
	// 三方用户的唯一ID
	_merchantUserId string
}

// New
