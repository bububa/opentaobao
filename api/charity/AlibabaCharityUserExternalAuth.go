package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityUserExternalAuth 外部用户授权
// alibaba.charity.user.external.auth
//
// 外部用户授权
func AlibabaCharityUserExternalAuth(clt *core.SDKClient, req *charity.AlibabaCharityUserExternalAuthAPIRequest, session string) (*charity.AlibabaCharityUserExternalAuthAPIResponse, error) {
	var resp charity.AlibabaCharityUserExternalAuthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
