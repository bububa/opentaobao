package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// PushAliyuncsComPushNotification20150318 推送通知
// push.aliyuncs.com.pushNotification.2015-03-18
//
// pushNotification
func PushAliyuncsComPushNotification20150318(clt *core.SDKClient, req *aliyun.PushAliyuncsComPushNotification20150318APIRequest, resp *aliyun.PushAliyuncsComPushNotification20150318APIResponse, session string) error {
	return clt.Post(req, resp, session)
}
