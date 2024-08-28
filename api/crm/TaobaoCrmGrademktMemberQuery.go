package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGrademktMemberQuery 会员等级营销-会员关系查询
// taobao.crm.grademkt.member.query
//
// 商家通过该接口查询线上店铺会员。
func TaobaoCrmGrademktMemberQuery(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmGrademktMemberQueryAPIRequest, resp *crm.TaobaoCrmGrademktMemberQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
