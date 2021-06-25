package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增档期商品 APIRequest
alibaba.price.promotion.item.add

批量新增档期活动商品
*/
type AlibabaPricePromotionItemAddRequest struct {
    model.Params

    // 入参
    promotionContent   *PromotionContent 

}

func NewAlibabaPricePromotionItemAddRequest() *AlibabaPricePromotionItemAddRequest{
    return &AlibabaPricePromotionItemAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPricePromotionItemAddRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.item.add"
}

func (r AlibabaPricePromotionItemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPricePromotionItemAddRequest) SetPromotionContent(promotionContent *PromotionContent) error {
    r.promotionContent = promotionContent
    r.Set("promotion_content", promotionContent)
    return nil
}

func (r AlibabaPricePromotionItemAddRequest) GetPromotionContent() *PromotionContent {
    return r.promotionContent
}

