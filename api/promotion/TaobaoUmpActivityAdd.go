package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
新增优惠活动 
taobao.ump.activity.add

新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家）
*/
func TaobaoUmpActivityAdd(clt *core.SDKClient, req *promotion.TaobaoUmpActivityAddAPIRequest, session string) (*promotion.TaobaoUmpActivityAddAPIResponse, error) {
    var resp promotion.TaobaoUmpActivityAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
