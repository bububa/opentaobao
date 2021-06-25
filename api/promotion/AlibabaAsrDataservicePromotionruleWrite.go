package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
业务优惠规则写入 
alibaba.asr.dataservice.promotionrule.write

星巴克优惠规则写入
*/
func AlibabaAsrDataservicePromotionruleWrite(clt *core.SDKClient, req *promotion.AlibabaAsrDataservicePromotionruleWriteRequest, session string) (*promotion.AlibabaAsrDataservicePromotionruleWriteResponse, error) {
    var resp promotion.AlibabaAsrDataservicePromotionruleWriteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
