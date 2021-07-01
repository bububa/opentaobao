package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
获取有权限的定向列表 
taobao.feedflow.item.target.validlist

获取有权限的定向列表
*/
func TaobaoFeedflowItemTargetValidlist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemTargetValidlistAPIRequest, session string) (*feedflow.TaobaoFeedflowItemTargetValidlistAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemTargetValidlistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
