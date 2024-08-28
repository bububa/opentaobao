package jym

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// TaobaoJymMemberRealnameVerifyWithoutuid 用户实名认证
// taobao.jym.member.realname.verify.withoutuid
//
// 开放用户实名认证接口使用
func TaobaoJymMemberRealnameVerifyWithoutuid(ctx context.Context, clt *core.SDKClient, req *jym.TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest, resp *jym.TaobaoJymMemberRealnameVerifyWithoutuidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
