package mei

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// TmallMeiCrmMemberGetbypaycode 支付码获取会员信息
// tmall.mei.crm.member.getbypaycode
//
// 通过支付码获取会员信息
func TmallMeiCrmMemberGetbypaycode(ctx context.Context, clt *core.SDKClient, req *mei.TmallMeiCrmMemberGetbypaycodeAPIRequest, resp *mei.TmallMeiCrmMemberGetbypaycodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
