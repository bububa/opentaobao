package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
信息流删除创意 
taobao.feedflow.item.creative.delete

信息流删除创意
*/
func TaobaoFeedflowItemCreativeDelete(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCreativeDeleteAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCreativeDeleteAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemCreativeDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
