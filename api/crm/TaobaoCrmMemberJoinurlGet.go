package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMemberJoinurlGet 会员入会地址获取
// taobao.crm.member.joinurl.get
//
// 会员入会地址获取
func TaobaoCrmMemberJoinurlGet(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmMemberJoinurlGetAPIRequest, resp *crm.TaobaoCrmMemberJoinurlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
