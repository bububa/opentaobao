package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabTbUserSkillOauth 用户技能 Oauth 授权（淘宝 openId）
// alibaba.ailab.tb.user.skill.oauth
//
// 定制机厂商，在用户配网完成后，厂商调用此接口，写入特定技能的 Oauth 信息
func AlibabaAilabTbUserSkillOauth(clt *core.SDKClient, req *alilabs.AlibabaAilabTbUserSkillOauthAPIRequest, resp *alilabs.AlibabaAilabTbUserSkillOauthAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
