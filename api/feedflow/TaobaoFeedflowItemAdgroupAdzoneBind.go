package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
信息流单元内绑定资源位 
taobao.feedflow.item.adgroup.adzone.bind

信息流单元内绑定资源位
*/
func TaobaoFeedflowItemAdgroupAdzoneBind(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest, session string) (*feedflow.TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
