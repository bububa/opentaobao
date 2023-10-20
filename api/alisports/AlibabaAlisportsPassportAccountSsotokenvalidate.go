package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountSsotokenvalidate sso_token验证
// alibaba.alisports.passport.account.ssotokenvalidate
//
// sso_token验证
func AlibabaAlisportsPassportAccountSsotokenvalidate(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest, resp *alisports.AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
