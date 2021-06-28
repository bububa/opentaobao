package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询无条件单品优惠活动列表 
taobao.promotionmisc.item.activity.list.get

查询无条件单品优惠活动列表
*/
func TaobaoPromotionmiscItemActivityListGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscItemActivityListGetRequest, session string) (*promotion.TaobaoPromotionmiscItemActivityListGetAPIResponse, error) {
    var resp promotion.TaobaoPromotionmiscItemActivityListGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
