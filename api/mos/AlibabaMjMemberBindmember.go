package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMemberBindmember 绑定会员
// alibaba.mj.member.bindmember
//
// 用于绑定喵街数字化会员
func AlibabaMjMemberBindmember(clt *core.SDKClient, req *mos.AlibabaMjMemberBindmemberAPIRequest, resp *mos.AlibabaMjMemberBindmemberAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
