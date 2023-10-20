package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityUserExternalAuthCancel 取消外部授权
// alibaba.charity.user.external.auth.cancel
//
// 取消外部授权
func AlibabaCharityUserExternalAuthCancel(clt *core.SDKClient, req *charity.AlibabaCharityUserExternalAuthCancelAPIRequest, resp *charity.AlibabaCharityUserExternalAuthCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
