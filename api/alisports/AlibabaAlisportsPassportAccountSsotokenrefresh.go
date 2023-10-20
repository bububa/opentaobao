package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountSsotokenrefresh sso_token刷新
// alibaba.alisports.passport.account.ssotokenrefresh
//
// sso_token刷新
func AlibabaAlisportsPassportAccountSsotokenrefresh(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountSsotokenrefreshAPIRequest, resp *alisports.AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
