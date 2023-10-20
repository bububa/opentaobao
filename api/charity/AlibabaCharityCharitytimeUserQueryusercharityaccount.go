package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharitycharitytimeuserqueryusercharityaccount 查询用户公益账户
// alibaba.charity.charitytime.user.queryusercharityaccount
//
// 查询用户公益账户
func Alibabacharitycharitytimeuserqueryusercharityaccount(clt *core.SDKClient, req *charity.AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest, session string) (*charity.AlibabacharitycharitytimeuserqueryusercharityaccountAPIResponse, error) {
	var resp charity.AlibabacharitycharitytimeuserqueryusercharityaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
