package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
修改关联的活动权益 
taobao.promotion.benefit.activity.update

修改卖家活动中关联的对应的权益。
*/
func TaobaoPromotionBenefitActivityUpdate(clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityUpdateAPIRequest, session string) (*promotion.TaobaoPromotionBenefitActivityUpdateAPIResponse, error) {
    var resp promotion.TaobaoPromotionBenefitActivityUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
