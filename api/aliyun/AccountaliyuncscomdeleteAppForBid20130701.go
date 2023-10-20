package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountAliyuncsComDeleteAppForBid20130701 运营商删除用户的appkey
// account.aliyuncs.com.DeleteAppForBid.2013-07-01
//
// 删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
func AccountAliyuncsComDeleteAppForBid20130701(clt *core.SDKClient, req *aliyun.AccountAliyuncsComDeleteAppForBid20130701APIRequest, resp *aliyun.AccountAliyuncsComDeleteAppForBid20130701APIResponse, session string) error {
	return clt.Post(req, resp, session)
}
