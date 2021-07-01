package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
视频播放 
alibaba.interact.ui.video

Weex页面播放视频
*/
func AlibabaInteractUiVideo(clt *core.SDKClient, req *user.AlibabaInteractUiVideoAPIRequest, session string) (*user.AlibabaInteractUiVideoAPIResponse, error) {
    var resp user.AlibabaInteractUiVideoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
