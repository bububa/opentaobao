package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询通用单品优惠详情列表 
taobao.promotionmisc.common.item.detail.list.get

查询通用单品优惠详情列表。
*/
func TaobaoPromotionmiscCommonItemDetailListGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemDetailListGetAPIRequest, session string) (*promotion.TaobaoPromotionmiscCommonItemDetailListGetAPIResponse, error) {
    var resp promotion.TaobaoPromotionmiscCommonItemDetailListGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
