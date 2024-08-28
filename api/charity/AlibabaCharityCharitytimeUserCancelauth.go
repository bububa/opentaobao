package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeUserCancelauth 取消用户授权
// alibaba.charity.charitytime.user.cancelauth
//
// 取消用户授权
func AlibabaCharityCharitytimeUserCancelauth(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeUserCancelauthAPIRequest, resp *charity.AlibabaCharityCharitytimeUserCancelauthAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
