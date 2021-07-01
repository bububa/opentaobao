package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
更新关联活动有效时间 
taobao.promotion.benefit.activity.time.update

更新关联权益的活动有效时间
*/
func TaobaoPromotionBenefitActivityTimeUpdate(clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityTimeUpdateAPIRequest, session string) (*promotion.TaobaoPromotionBenefitActivityTimeUpdateAPIResponse, error) {
    var resp promotion.TaobaoPromotionBenefitActivityTimeUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
