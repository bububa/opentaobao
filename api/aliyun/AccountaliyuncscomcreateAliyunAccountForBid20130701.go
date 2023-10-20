package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountAliyuncsComCreateAliyunAccountForBid20130701 为bid用户创建账号
// account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01
//
// 给指定的bid创建账号，同时账号打上owner bid的标记
func AccountAliyuncsComCreateAliyunAccountForBid20130701(clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest, resp *aliyun.AccountAliyuncsComCreateAliyunAccountForBid20130701APIResponse, session string) error {
	return clt.Post(req, resp, session)
}
