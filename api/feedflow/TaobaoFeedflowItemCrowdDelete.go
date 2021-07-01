package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
删除单品人群 
taobao.feedflow.item.crowd.delete

删除单品人群
*/
func TaobaoFeedflowItemCrowdDelete(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdDeleteAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCrowdDeleteAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemCrowdDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
