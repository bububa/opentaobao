package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
星巴克优惠规则查询 
alibaba.asr.dataservice.promotionrule.query

查询优惠规则，例如星巴克查询优惠规则
*/
func AlibabaAsrDataservicePromotionruleQuery(clt *core.SDKClient, req *promotion.AlibabaAsrDataservicePromotionruleQueryAPIRequest, session string) (*promotion.AlibabaAsrDataservicePromotionruleQueryAPIResponse, error) {
    var resp promotion.AlibabaAsrDataservicePromotionruleQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
