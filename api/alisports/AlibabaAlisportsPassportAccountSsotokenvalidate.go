package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountSsotokenvalidate sso_token验证
// alibaba.alisports.passport.account.ssotokenvalidate
//
// sso_token验证
func AlibabaAlisportsPassportAccountSsotokenvalidate(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest, resp *alisports.AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
