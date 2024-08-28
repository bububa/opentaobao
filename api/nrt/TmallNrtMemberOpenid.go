package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtMemberOpenid 根据会员手机查询openId
// tmall.nrt.member.openid
//
// 根据会员手机查询openId
func TmallNrtMemberOpenid(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtMemberOpenidAPIRequest, resp *nrt.TmallNrtMemberOpenidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
