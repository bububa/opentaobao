package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGrademktMemberDetailDelete 会员等级营销-删除商品等级营销明细
// taobao.crm.grademkt.member.detail.delete
//
// 删除商品等级营销明细
func TaobaoCrmGrademktMemberDetailDelete(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmGrademktMemberDetailDeleteAPIRequest, resp *crm.TaobaoCrmGrademktMemberDetailDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
