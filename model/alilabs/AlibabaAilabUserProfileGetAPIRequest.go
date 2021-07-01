package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabUserProfileGetAPIRequest
查询用户信息 API请求
alibaba.ailab.user.profile.get

提供天猫精灵用户头像、昵称的查询接口，供本田车载天猫精灵使用 */
type AlibabaAilabUserProfileGetAPIRequest struct {
	model.Params
	// open uid
	_openUid string
	// client id
	_clientId string
}

// New
