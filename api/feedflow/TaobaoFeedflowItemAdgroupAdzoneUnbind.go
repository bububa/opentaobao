package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
信息流单元内解绑资源位 
taobao.feedflow.item.adgroup.adzone.unbind

信息流单元内解绑资源位
*/
func TaobaoFeedflowItemAdgroupAdzoneUnbind(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest, session string) (*feedflow.TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
