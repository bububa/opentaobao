package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// PushaliyuncscompushNotification20150318 推送通知
// push.aliyuncs.com.pushNotification.2015-03-18
//
// pushNotification
func PushaliyuncscompushNotification20150318(clt *core.SDKClient, req *aliyun.PushaliyuncscompushNotification20150318APIRequest, session string) (*aliyun.PushaliyuncscompushNotification20150318APIResponse, error) {
	var resp aliyun.PushaliyuncscompushNotification20150318APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
