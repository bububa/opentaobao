package youkuott

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/youkuott"
)

/* 
运营商一般订单同步 
youku.ott.kitty.commonorder.sync

运营商一般订单同步
*/
func YoukuOttKittyCommonorderSync(clt *core.SDKClient, req *youkuott.YoukuOttKittyCommonorderSyncRequest, session string) (*youkuott.YoukuOttKittyCommonorderSyncAPIResponse, error) {
    var resp youkuott.YoukuOttKittyCommonorderSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
