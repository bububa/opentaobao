package aliyun

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyun"
)

/* 
推送通知 
push.aliyuncs.com.pushNotification.2015-03-18

pushNotification
*/
func PushAliyuncsComPushNotification2015_03_18(clt *core.SDKClient, req *aliyun.PushAliyuncsComPushNotification2015_03_18Request, session string) (*aliyun.PushAliyuncsComPushNotification2015_03_18APIResponse, error) {
    var resp aliyun.PushAliyuncsComPushNotification2015_03_18APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
