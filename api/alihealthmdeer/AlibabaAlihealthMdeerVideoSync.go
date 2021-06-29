package alihealthmdeer

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthmdeer"
)

/* 
合作伙伴视频同步给医知鹿接口 
alibaba.alihealth.mdeer.video.sync

合伙做伴内容同步接口，用来视频同步
*/
func AlibabaAlihealthMdeerVideoSync(clt *core.SDKClient, req *alihealthmdeer.AlibabaAlihealthMdeerVideoSyncRequest, session string) (*alihealthmdeer.AlibabaAlihealthMdeerVideoSyncAPIResponse, error) {
    var resp alihealthmdeer.AlibabaAlihealthMdeerVideoSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
