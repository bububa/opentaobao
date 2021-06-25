package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
取消收藏 
taobao.ailab.aicloud.top.like.delete

取消收藏
*/
func TaobaoAilabAicloudTopLikeDelete(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeDeleteRequest, session string) (*iot.TaobaoAilabAicloudTopLikeDeleteResponse, error) {
    var resp iot.TaobaoAilabAicloudTopLikeDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
