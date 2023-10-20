package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtMemberOpenid 根据会员手机查询openId
// tmall.nrt.member.openid
//
// 根据会员手机查询openId
func TmallNrtMemberOpenid(clt *core.SDKClient, req *nrt.TmallNrtMemberOpenidAPIRequest, resp *nrt.TmallNrtMemberOpenidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
