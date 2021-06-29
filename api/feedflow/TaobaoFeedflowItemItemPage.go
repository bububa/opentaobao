package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
信息流查看商品列表 
taobao.feedflow.item.item.page

信息流查看商品列表
*/
func TaobaoFeedflowItemItemPage(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemItemPageRequest, session string) (*feedflow.TaobaoFeedflowItemItemPageAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemItemPageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
