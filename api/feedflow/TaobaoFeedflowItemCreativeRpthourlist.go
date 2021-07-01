package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
超级推荐【商品推广】创意分时报表查询 
taobao.feedflow.item.creative.rpthourlist

创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据
*/
func TaobaoFeedflowItemCreativeRpthourlist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCreativeRpthourlistAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCreativeRpthourlistAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemCreativeRpthourlistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
