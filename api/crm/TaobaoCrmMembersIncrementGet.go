package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMembersIncrementGet 增量获取卖家会员
// taobao.crm.members.increment.get
//
// 增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100
func TaobaoCrmMembersIncrementGet(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmMembersIncrementGetAPIRequest, resp *crm.TaobaoCrmMembersIncrementGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
