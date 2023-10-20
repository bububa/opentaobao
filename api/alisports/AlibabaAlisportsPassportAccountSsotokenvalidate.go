package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportaccountssotokenvalidate sso_token验证
// alibaba.alisports.passport.account.ssotokenvalidate
//
// sso_token验证
func Alibabaalisportspassportaccountssotokenvalidate(clt *core.SDKClient, req *alisports.AlibabaalisportspassportaccountssotokenvalidateAPIRequest, session string) (*alisports.AlibabaalisportspassportaccountssotokenvalidateAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportaccountssotokenvalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
