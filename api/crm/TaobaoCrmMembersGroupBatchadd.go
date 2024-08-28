package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMembersGroupBatchadd 给一批会员添加一个分组
// taobao.crm.members.group.batchadd
//
// 为一批会员添加分组，接口返回添加是否成功,如至少有一个会员的分组添加成功，接口就返回成功，否则返回失败，如果当前会员已经拥有当前分组，则直接跳过
func TaobaoCrmMembersGroupBatchadd(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmMembersGroupBatchaddAPIRequest, resp *crm.TaobaoCrmMembersGroupBatchaddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
