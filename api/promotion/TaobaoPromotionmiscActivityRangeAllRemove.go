package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
清空活动参与的商品 
taobao.promotionmisc.activity.range.all.remove

清空活动参与的商品
*/
func TaobaoPromotionmiscActivityRangeAllRemove(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscActivityRangeAllRemoveRequest, session string) (*promotion.TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse, error) {
    var resp promotion.TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
