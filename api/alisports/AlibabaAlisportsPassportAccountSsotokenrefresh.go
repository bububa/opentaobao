package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportaccountssotokenrefresh sso_token刷新
// alibaba.alisports.passport.account.ssotokenrefresh
//
// sso_token刷新
func Alibabaalisportspassportaccountssotokenrefresh(clt *core.SDKClient, req *alisports.AlibabaalisportspassportaccountssotokenrefreshAPIRequest, session string) (*alisports.AlibabaalisportspassportaccountssotokenrefreshAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportaccountssotokenrefreshAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
