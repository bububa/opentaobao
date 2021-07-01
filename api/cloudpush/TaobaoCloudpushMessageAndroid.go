package cloudpush

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cloudpush"
)

/* 
百川云推送发送消息给android 
taobao.cloudpush.message.android

百川用户使用云推送发送消息给android
*/
func TaobaoCloudpushMessageAndroid(clt *core.SDKClient, req *cloudpush.TaobaoCloudpushMessageAndroidAPIRequest, session string) (*cloudpush.TaobaoCloudpushMessageAndroidAPIResponse, error) {
    var resp cloudpush.TaobaoCloudpushMessageAndroidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
