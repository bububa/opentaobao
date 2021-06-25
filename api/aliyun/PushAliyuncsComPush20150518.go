package aliyun

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyun"
)

/* 
云推送指令API 
push.aliyuncs.com.push.20150518

阿里云推送新增API，允许一条推送指令同时发布到多个终端上。
*/
func PushAliyuncsComPush20150518(clt *core.SDKClient, req *aliyun.PushAliyuncsComPush20150518Request, session string) (*aliyun.PushAliyuncsComPush20150518Response, error) {
    var resp aliyun.PushAliyuncsComPush20150518APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
