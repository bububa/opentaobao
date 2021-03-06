package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountSsotokenvalidate sso_token验证
// alibaba.alisports.passport.account.ssotokenvalidate
//
// sso_token验证
func AlibabaAlisportsPassportAccountSsotokenvalidate(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest, session string) (*alisports.AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse, error) {
	var resp alisports.AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
