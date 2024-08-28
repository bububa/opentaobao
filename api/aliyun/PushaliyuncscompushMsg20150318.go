package aliyun

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// PushAliyuncsComPushMsg20150318 消息推送
// push.aliyuncs.com.pushMsg.2015-03-18
//
// 消息推送  ,支持指定用户/账号/广播等模式
func PushAliyuncsComPushMsg20150318(ctx context.Context, clt *core.SDKClient, req *aliyun.PushAliyuncsComPushMsg20150318APIRequest, resp *aliyun.PushAliyuncsComPushMsg20150318APIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
