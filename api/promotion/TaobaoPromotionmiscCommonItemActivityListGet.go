package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询通用单品优惠活动列表 
taobao.promotionmisc.common.item.activity.list.get

查询通用单品优惠活动列表。
*/
func TaobaoPromotionmiscCommonItemActivityListGet(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemActivityListGetRequest, session string) (*promotion.TaobaoPromotionmiscCommonItemActivityListGetResponse, error) {
    var resp promotion.TaobaoPromotionmiscCommonItemActivityListGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
