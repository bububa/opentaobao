package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeUserQuerythirduserhasauth 查询是否绑定3小时账号
// alibaba.charity.charitytime.user.querythirduserhasauth
//
// 查询是否绑定3小时账号
func AlibabaCharityCharitytimeUserQuerythirduserhasauth(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIRequest, resp *charity.AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
