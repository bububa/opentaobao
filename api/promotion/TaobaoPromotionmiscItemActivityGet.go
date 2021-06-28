package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询无条件单品优惠活动 
taobao.promotionmisc.item.activity.get

查询无条件单品优惠活动
*/
func TaobaoPromotionmiscItemActivityGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscItemActivityGetRequest, session string) (*promotion.TaobaoPromotionmiscItemActivityGetAPIResponse, error) {
    var resp promotion.TaobaoPromotionmiscItemActivityGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
