package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// Pushaliyuncscompush20150518 云推送指令API
// push.aliyuncs.com.push.20150518
//
// 阿里云推送新增API，允许一条推送指令同时发布到多个终端上。
func Pushaliyuncscompush20150518(clt *core.SDKClient, req *aliyun.Pushaliyuncscompush20150518APIRequest, session string) (*aliyun.Pushaliyuncscompush20150518APIResponse, error) {
	var resp aliyun.Pushaliyuncscompush20150518APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
