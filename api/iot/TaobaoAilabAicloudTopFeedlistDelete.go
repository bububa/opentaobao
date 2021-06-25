package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
删除单条对话流信息 
taobao.ailab.aicloud.top.feedlist.delete

删除指定的某一条对话流信息
*/
func TaobaoAilabAicloudTopFeedlistDelete(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopFeedlistDeleteRequest, session string) (*iot.TaobaoAilabAicloudTopFeedlistDeleteResponse, error) {
    var resp iot.TaobaoAilabAicloudTopFeedlistDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
