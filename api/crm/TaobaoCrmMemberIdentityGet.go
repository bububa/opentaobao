package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMemberIdentityGet 会员身份识别
// taobao.crm.member.identity.get
//
// 用来识别该用户是否是商家会员
func TaobaoCrmMemberIdentityGet(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmMemberIdentityGetAPIRequest, resp *crm.TaobaoCrmMemberIdentityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
