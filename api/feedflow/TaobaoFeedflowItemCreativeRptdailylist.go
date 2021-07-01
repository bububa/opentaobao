package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
创意分日数据查询 
taobao.feedflow.item.creative.rptdailylist

创意分日数据查询
*/
func TaobaoFeedflowItemCreativeRptdailylist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCreativeRptdailylistAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCreativeRptdailylistAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemCreativeRptdailylistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
