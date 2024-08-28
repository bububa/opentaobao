package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoKoubeiTribeOpenVerifyCodeApply 口碑综合体手机号获取验证码
// taobao.koubei.tribe.open.verify.code.apply
//
// 口碑综合体通过手机号获取验证码对外开放接口
func TaobaoKoubeiTribeOpenVerifyCodeApply(ctx context.Context, clt *core.SDKClient, req *user.TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest, resp *user.TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
