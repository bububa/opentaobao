package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
活动关联的权益详情获取 
taobao.promotion.benefit.activity.detail.get

活动关联的权益详情获取
*/
func TaobaoPromotionBenefitActivityDetailGet(clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityDetailGetRequest, session string) (*promotion.TaobaoPromotionBenefitActivityDetailGetResponse, error) {
    var resp promotion.TaobaoPromotionBenefitActivityDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
