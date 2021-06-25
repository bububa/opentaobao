package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
ISV报名官方活动(中心化流量) 
alibaba.interact.activity.apply

支持商家将使用isv创建的活动所对应的权益信息同步到手淘，供过滤是否在中心化流量入口透出
*/
func AlibabaInteractActivityApply(clt *core.SDKClient, req *interact.AlibabaInteractActivityApplyRequest, session string) (*interact.AlibabaInteractActivityApplyResponse, error) {
    var resp interact.AlibabaInteractActivityApplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
