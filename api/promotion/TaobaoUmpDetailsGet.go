package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询活动详情列表 
taobao.ump.details.get

分页查询优惠详情列表
*/
func TaobaoUmpDetailsGet(clt *core.SDKClient, req *promotion.TaobaoUmpDetailsGetAPIRequest, session string) (*promotion.TaobaoUmpDetailsGetAPIResponse, error) {
    var resp promotion.TaobaoUmpDetailsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
