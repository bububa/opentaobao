package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
删除营销活动 
taobao.ump.activity.delete

删除营销活动。对应的活动详情等将会被全部删除。
*/
func TaobaoUmpActivityDelete(clt *core.SDKClient, req *promotion.TaobaoUmpActivityDeleteAPIRequest, session string) (*promotion.TaobaoUmpActivityDeleteAPIResponse, error) {
    var resp promotion.TaobaoUmpActivityDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
