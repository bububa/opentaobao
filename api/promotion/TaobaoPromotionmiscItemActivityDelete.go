package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
删除无条件单品优惠活动 
taobao.promotionmisc.item.activity.delete

删除无条件单品优惠活动
*/
func TaobaoPromotionmiscItemActivityDelete(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscItemActivityDeleteRequest, session string) (*promotion.TaobaoPromotionmiscItemActivityDeleteResponse, error) {
    var resp promotion.TaobaoPromotionmiscItemActivityDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
