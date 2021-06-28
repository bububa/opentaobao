package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
儿童音频列表 
taobao.ailab.aicloud.top.freelisten.childrenalbum

儿童音频列表
*/
func TaobaoAilabAicloudTopFreelistenChildrenalbum(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopFreelistenChildrenalbumRequest, session string) (*iot.TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
