package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabUserProfileGet 查询用户信息
// alibaba.ailab.user.profile.get
//
// 提供天猫精灵用户头像、昵称的查询接口，供本田车载天猫精灵使用
func AlibabaAilabUserProfileGet(clt *core.SDKClient, req *alilabs.AlibabaAilabUserProfileGetAPIRequest, resp *alilabs.AlibabaAilabUserProfileGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
