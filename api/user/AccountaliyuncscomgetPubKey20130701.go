package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AccountAliyuncsComGetPubKey20130701 获取用户公钥
// account.aliyuncs.com.GetPubKey.2013-07-01
//
// 根据用户的appkey查询用户的pubkey
func AccountAliyuncsComGetPubKey20130701(clt *core.SDKClient, req *user.AccountAliyuncsComGetPubKey20130701APIRequest, resp *user.AccountAliyuncsComGetPubKey20130701APIResponse, session string) error {
	return clt.Post(req, resp, session)
}
