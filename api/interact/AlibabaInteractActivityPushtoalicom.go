package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
小铺isv推广流量活动到流量聚乐部 
alibaba.interact.activity.pushtoalicom

涉及到流量包的小铺isv，将活动推送到流量聚乐部
*/
func AlibabaInteractActivityPushtoalicom(clt *core.SDKClient, req *interact.AlibabaInteractActivityPushtoalicomRequest, session string) (*interact.AlibabaInteractActivityPushtoalicomResponse, error) {
    var resp interact.AlibabaInteractActivityPushtoalicomAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
