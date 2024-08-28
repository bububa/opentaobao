package mei

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// TmallCrmMemberPointChange 会员积分变更
// tmall.crm.member.point.change
//
// 品牌CRM项目中，会员积分变更接口。
func TmallCrmMemberPointChange(ctx context.Context, clt *core.SDKClient, req *mei.TmallCrmMemberPointChangeAPIRequest, resp *mei.TmallCrmMemberPointChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
