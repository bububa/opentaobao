package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取促销规则列表 APIRequest
alibaba.alsc.crm.promotion.list

获取品牌的促销规则列表
*/
type AlibabaAlscCrmPromotionListRequest struct {
    model.Params

    // 获取促销规则请求参数
    promotionFacadeOpenReq   *PromotionFacadeOpenReq 

}

func NewAlibabaAlscCrmPromotionListRequest() *AlibabaAlscCrmPromotionListRequest{
    return &AlibabaAlscCrmPromotionListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmPromotionListRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.promotion.list"
}

func (r AlibabaAlscCrmPromotionListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmPromotionListRequest) SetPromotionFacadeOpenReq(promotionFacadeOpenReq *PromotionFacadeOpenReq) error {
    r.promotionFacadeOpenReq = promotionFacadeOpenReq
    r.Set("promotion_facade_open_req", promotionFacadeOpenReq)
    return nil
}

func (r AlibabaAlscCrmPromotionListRequest) GetPromotionFacadeOpenReq() *PromotionFacadeOpenReq {
    return r.promotionFacadeOpenReq
}

