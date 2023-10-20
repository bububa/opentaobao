package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharitycharitytimeusercancelauth 取消用户授权
// alibaba.charity.charitytime.user.cancelauth
//
// 取消用户授权
func Alibabacharitycharitytimeusercancelauth(clt *core.SDKClient, req *charity.AlibabacharitycharitytimeusercancelauthAPIRequest, session string) (*charity.AlibabacharitycharitytimeusercancelauthAPIResponse, error) {
	var resp charity.AlibabacharitycharitytimeusercancelauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
