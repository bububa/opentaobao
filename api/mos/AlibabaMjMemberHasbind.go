package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMemberHasbind 喵街会员是否绑定
// alibaba.mj.member.hasbind
//
// 喵街检测用户是否为数字化会员
func AlibabaMjMemberHasbind(clt *core.SDKClient, req *mos.AlibabaMjMemberHasbindAPIRequest, resp *mos.AlibabaMjMemberHasbindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
