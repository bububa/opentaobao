package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AccountaliyuncscomgetPubKey20130701 获取用户公钥
// account.aliyuncs.com.GetPubKey.2013-07-01
//
// 根据用户的appkey查询用户的pubkey
func AccountaliyuncscomgetPubKey20130701(clt *core.SDKClient, req *user.AccountaliyuncscomgetPubKey20130701APIRequest, session string) (*user.AccountaliyuncscomgetPubKey20130701APIResponse, error) {
	var resp user.AccountaliyuncscomgetPubKey20130701APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
