package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
获取卖家最低折扣 
taobao.ump.promotion.global.discount.get

提供卖家最低折扣查询功能
*/
func TaobaoUmpPromotionGlobalDiscountGet(clt *core.SDKClient, req *promotion.TaobaoUmpPromotionGlobalDiscountGetRequest, session string) (*promotion.TaobaoUmpPromotionGlobalDiscountGetResponse, error) {
    var resp promotion.TaobaoUmpPromotionGlobalDiscountGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
