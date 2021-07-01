package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
获取促销规则列表 
alibaba.alsc.crm.promotion.list

获取品牌的促销规则列表
*/
func AlibabaAlscCrmPromotionList(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPromotionListAPIRequest, session string) (*alsc.AlibabaAlscCrmPromotionListAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmPromotionListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
