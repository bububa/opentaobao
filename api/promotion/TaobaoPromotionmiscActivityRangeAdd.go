package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
增加活动参与的商品 
taobao.promotionmisc.activity.range.add

增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
*/
func TaobaoPromotionmiscActivityRangeAdd(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscActivityRangeAddRequest, session string) (*promotion.TaobaoPromotionmiscActivityRangeAddResponse, error) {
    var resp promotion.TaobaoPromotionmiscActivityRangeAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
