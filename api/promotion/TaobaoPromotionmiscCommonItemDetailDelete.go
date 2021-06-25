package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
删除通用单品优惠详情 
taobao.promotionmisc.common.item.detail.delete

删除通用单品优惠详情。
*/
func TaobaoPromotionmiscCommonItemDetailDelete(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemDetailDeleteRequest, session string) (*promotion.TaobaoPromotionmiscCommonItemDetailDeleteResponse, error) {
    var resp promotion.TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
