package cloudpush

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cloudpush"
)

/* 
百川云推送发送消息给ios 
taobao.cloudpush.message.ios

百川云推送发送消息给iOS设备.
*/
func TaobaoCloudpushMessageIos(clt *core.SDKClient, req *cloudpush.TaobaoCloudpushMessageIosRequest, session string) (*cloudpush.TaobaoCloudpushMessageIosAPIResponse, error) {
    var resp cloudpush.TaobaoCloudpushMessageIosAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
