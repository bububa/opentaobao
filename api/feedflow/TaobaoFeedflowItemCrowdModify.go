package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
覆盖单元下同类型定向人群 
taobao.feedflow.item.crowd.modify

覆盖单元下同类型定向人群
*/
func TaobaoFeedflowItemCrowdModify(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdModifyAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCrowdModifyAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemCrowdModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
