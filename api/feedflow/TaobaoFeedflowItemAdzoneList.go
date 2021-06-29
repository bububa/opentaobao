package feedflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/feedflow"
)

/* 
批量查询可用广告位列表 
taobao.feedflow.item.adzone.list

批量查询可用广告位列表
*/
func TaobaoFeedflowItemAdzoneList(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdzoneListRequest, session string) (*feedflow.TaobaoFeedflowItemAdzoneListAPIResponse, error) {
    var resp feedflow.TaobaoFeedflowItemAdzoneListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
