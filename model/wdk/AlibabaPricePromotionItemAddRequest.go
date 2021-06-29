package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增档期商品 API请求
alibaba.price.promotion.item.add

批量新增档期活动商品
*/
type AlibabaPricePromotionItemAddRequest struct {
    model.Params
    // 入参
    _promotionContent   *PromotionContent
}

// 初始化AlibabaPricePromotionItemAddRequest对象
func NewAlibabaPricePromotionItemAddRequest() *AlibabaPricePromotionItemAddRequest{
    return &AlibabaPricePromotionItemAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionItemAddRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.item.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionItemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PromotionContent Setter
// 入参
func (r *AlibabaPricePromotionItemAddRequest) SetPromotionContent(_promotionContent *PromotionContent) error {
    r._promotionContent = _promotionContent
    r.Set("promotion_content", _promotionContent)
    return nil
}

// PromotionContent Getter
func (r AlibabaPricePromotionItemAddRequest) GetPromotionContent() *PromotionContent {
    return r._promotionContent
}
