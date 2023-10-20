package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharitycharitytimeuserquerythirduserhasauth 查询是否绑定3小时账号
// alibaba.charity.charitytime.user.querythirduserhasauth
//
// 查询是否绑定3小时账号
func Alibabacharitycharitytimeuserquerythirduserhasauth(clt *core.SDKClient, req *charity.AlibabacharitycharitytimeuserquerythirduserhasauthAPIRequest, session string) (*charity.AlibabacharitycharitytimeuserquerythirduserhasauthAPIResponse, error) {
	var resp charity.AlibabacharitycharitytimeuserquerythirduserhasauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
