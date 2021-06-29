package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
定向分日数据查询 
taobao.feedflow.item.crowd.rptdailylist

定向分日数据查询
*/
func TaobaoFeedflowItemCrowdRptdailylist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdRptdailylistRequest, session string) (*feedflow.TaobaoFeedflowItemCrowdRptdailylistAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemCrowdRptdailylistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
