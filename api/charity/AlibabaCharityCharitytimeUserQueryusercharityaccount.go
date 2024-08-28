package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeUserQueryusercharityaccount 查询用户公益账户
// alibaba.charity.charitytime.user.queryusercharityaccount
//
// 查询用户公益账户
func AlibabaCharityCharitytimeUserQueryusercharityaccount(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeUserQueryusercharityaccountAPIRequest, resp *charity.AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
