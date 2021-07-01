package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabTbUserSkillOauthAPIRequest
用户技能 Oauth 授权（淘宝 openId） API请求
alibaba.ailab.tb.user.skill.oauth

定制机厂商，在用户配网完成后，厂商调用此接口，写入特定技能的 Oauth 信息 */
type AlibabaAilabTbUserSkillOauthAPIRequest struct {
	model.Params
	// taobao open id
	_taobaoId string
	// access token 过期时间，单位：秒
	_expireIn int64
	// access token
	_oauthAccessToken string
	// refresh token
	_refreshToken string
}

// New
