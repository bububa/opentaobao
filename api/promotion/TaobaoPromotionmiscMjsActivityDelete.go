package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
删除满就送活动 
taobao.promotionmisc.mjs.activity.delete

删除满就送活动
*/
func TaobaoPromotionmiscMjsActivityDelete(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscMjsActivityDeleteRequest, session string) (*promotion.TaobaoPromotionmiscMjsActivityDeleteResponse, error) {
    var resp promotion.TaobaoPromotionmiscMjsActivityDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
