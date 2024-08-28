package aliyun

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountAliyuncsComCreateApp20130701 给指定用户创建appkey
// account.aliyuncs.com.CreateApp.2013-07-01
//
// 为某个用户创建appkey
func AccountAliyuncsComCreateApp20130701(ctx context.Context, clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateApp20130701APIRequest, resp *aliyun.AccountAliyuncsComCreateApp20130701APIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
