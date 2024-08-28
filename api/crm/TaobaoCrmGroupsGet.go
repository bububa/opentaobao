package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGroupsGet 查询卖家的分组
// taobao.crm.groups.get
//
// 查询卖家的分组，返回查询到的分组列表，分页返回分组
func TaobaoCrmGroupsGet(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmGroupsGetAPIRequest, resp *crm.TaobaoCrmGroupsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
