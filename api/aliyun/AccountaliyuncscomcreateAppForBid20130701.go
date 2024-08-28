package aliyun

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountAliyuncsComCreateAppForBid20130701 给当前渠道下的用户创建app
// account.aliyuncs.com.CreateAppForBid.2013-07-01
//
// 给自己渠道下的用户创建app
func AccountAliyuncsComCreateAppForBid20130701(ctx context.Context, clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateAppForBid20130701APIRequest, resp *aliyun.AccountAliyuncsComCreateAppForBid20130701APIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
