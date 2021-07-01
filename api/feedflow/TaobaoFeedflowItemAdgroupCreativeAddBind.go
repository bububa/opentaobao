package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
信息流新增并且绑定创意 
taobao.feedflow.item.adgroup.creative.add.bind

信息流新增并且绑定创意
*/
func TaobaoFeedflowItemAdgroupCreativeAddBind(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest, session string) (*feedflow.TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
