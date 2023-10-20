package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeUserQueryusercharityaccount 查询用户公益账户
// alibaba.charity.charitytime.user.queryusercharityaccount
//
// 查询用户公益账户
func AlibabaCharityCharitytimeUserQueryusercharityaccount(clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeUserQueryusercharityaccountAPIRequest, session string) (*charity.AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse, error) {
	var resp charity.AlibabaCharityCharitytimeUserQueryusercharityaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
