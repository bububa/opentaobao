package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
增加收藏 
taobao.ailab.aicloud.top.like.add

将制定内容加入收藏
*/
func TaobaoAilabAicloudTopLikeAdd(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeAddRequest, session string) (*iot.TaobaoAilabAicloudTopLikeAddAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopLikeAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
