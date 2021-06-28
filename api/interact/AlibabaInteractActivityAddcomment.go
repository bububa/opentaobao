package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
微淘评论接口 
alibaba.interact.activity.addcomment

发表评论，并返回楼层
*/
func AlibabaInteractActivityAddcomment(clt *core.SDKClient, req *interact.AlibabaInteractActivityAddcommentRequest, session string) (*interact.AlibabaInteractActivityAddcommentAPIResponse, error) {
    var resp interact.AlibabaInteractActivityAddcommentAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
