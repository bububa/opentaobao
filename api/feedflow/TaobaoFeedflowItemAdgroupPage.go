package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
查询单元列表 
taobao.feedflow.item.adgroup.page

通过计划id查询单元信息
*/
func TaobaoFeedflowItemAdgroupPage(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupPageAPIRequest, session string) (*feedflow.TaobaoFeedflowItemAdgroupPageAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemAdgroupPageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
