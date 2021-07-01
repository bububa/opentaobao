package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询营销活动 
taobao.ump.activity.get

查询营销活动
*/
func TaobaoUmpActivityGet(clt *core.SDKClient, req *promotion.TaobaoUmpActivityGetAPIRequest, session string) (*promotion.TaobaoUmpActivityGetAPIResponse, error) {
    var resp promotion.TaobaoUmpActivityGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
