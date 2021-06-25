package cloudpush

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cloudpush"
)

/* 
百川用户使用云推送高级推送接口 
taobao.cloudpush.push

百川用户使用云推送高级推送接口
*/
func TaobaoCloudpushPush(clt *core.SDKClient, req *cloudpush.TaobaoCloudpushPushRequest, session string) (*cloudpush.TaobaoCloudpushPushResponse, error) {
    var resp cloudpush.TaobaoCloudpushPushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
