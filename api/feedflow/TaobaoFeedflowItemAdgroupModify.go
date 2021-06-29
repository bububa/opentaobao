package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
信息流单元修改 
taobao.feedflow.item.adgroup.modify

信息流单元修改
*/
func TaobaoFeedflowItemAdgroupModify(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupModifyRequest, session string) (*feedflow.TaobaoFeedflowItemAdgroupModifyAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemAdgroupModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
