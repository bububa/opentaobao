package aliyun

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyun"
)

/* 
消息推送 
push.aliyuncs.com.pushMsg.2015-03-18

消息推送  ,支持指定用户/账号/广播等模式
*/
func PushAliyuncsComPushMsg2015_03_18(clt *core.SDKClient, req *aliyun.PushAliyuncsComPushMsg2015_03_18Request, session string) (*aliyun.PushAliyuncsComPushMsg2015_03_18APIResponse, error) {
    var resp aliyun.PushAliyuncsComPushMsg2015_03_18APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
