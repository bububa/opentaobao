package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
判断用户是否收藏某个店铺 
alibaba.interact.open.isattention

判断用户是否收藏某个店铺
*/
func AlibabaInteractOpenIsattention(clt *core.SDKClient, req *interact.AlibabaInteractOpenIsattentionRequest, session string) (*interact.AlibabaInteractOpenIsattentionAPIResponse, error) {
    var resp interact.AlibabaInteractOpenIsattentionAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
