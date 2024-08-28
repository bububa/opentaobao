package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountSsotokenrefresh sso_token刷新
// alibaba.alisports.passport.account.ssotokenrefresh
//
// sso_token刷新
func AlibabaAlisportsPassportAccountSsotokenrefresh(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountSsotokenrefreshAPIRequest, resp *alisports.AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
