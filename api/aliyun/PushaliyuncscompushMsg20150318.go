package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// PushaliyuncscompushMsg20150318 消息推送
// push.aliyuncs.com.pushMsg.2015-03-18
//
// 消息推送  ,支持指定用户/账号/广播等模式
func PushaliyuncscompushMsg20150318(clt *core.SDKClient, req *aliyun.PushaliyuncscompushMsg20150318APIRequest, session string) (*aliyun.PushaliyuncscompushMsg20150318APIResponse, error) {
	var resp aliyun.PushaliyuncscompushMsg20150318APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
