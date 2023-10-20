package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeUserCancelauth 取消用户授权
// alibaba.charity.charitytime.user.cancelauth
//
// 取消用户授权
func AlibabaCharityCharitytimeUserCancelauth(clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeUserCancelauthAPIRequest, resp *charity.AlibabaCharityCharitytimeUserCancelauthAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
