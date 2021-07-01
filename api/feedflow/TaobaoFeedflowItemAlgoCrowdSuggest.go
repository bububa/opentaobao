package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
单品人群建议出价 
taobao.feedflow.item.algo.crowd.suggest

给超级推荐的广告主查看建议出价
*/
func TaobaoFeedflowItemAlgoCrowdSuggest(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest, session string) (*feedflow.TaobaoFeedflowItemAlgoCrowdSuggestAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemAlgoCrowdSuggestAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
