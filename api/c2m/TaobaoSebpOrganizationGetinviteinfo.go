package c2m

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/c2m"
)

// TaobaoSebpOrganizationGetinviteinfo 淘小铺机构上下级关系
// taobao.sebp.organization.getinviteinfo
//
// 机构人员获取机构上下级关系信息
func TaobaoSebpOrganizationGetinviteinfo(ctx context.Context, clt *core.SDKClient, req *c2m.TaobaoSebpOrganizationGetinviteinfoAPIRequest, resp *c2m.TaobaoSebpOrganizationGetinviteinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
