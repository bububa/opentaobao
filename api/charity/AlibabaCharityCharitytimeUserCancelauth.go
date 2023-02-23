package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeUserCancelauth 取消用户授权
// alibaba.charity.charitytime.user.cancelauth
//
// 取消用户授权
func AlibabaCharityCharitytimeUserCancelauth(clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeUserCancelauthAPIRequest, session string) (*charity.AlibabaCharityCharitytimeUserCancelauthAPIResponse, error) {
	var resp charity.AlibabaCharityCharitytimeUserCancelauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
