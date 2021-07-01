package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家最低折扣 API请求
taobao.ump.promotion.global.discount.get

提供卖家最低折扣查询功能
*/
type TaobaoUmpPromotionGlobalDiscountGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoUmpPromotionGlobalDiscountGetAPIRequest对象
func NewTaobaoUmpPromotionGlobalDiscountGetRequest() *TaobaoUmpPromotionGlobalDiscountGetAPIRequest{
    return &TaobaoUmpPromotionGlobalDiscountGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpPromotionGlobalDiscountGetAPIRequest) GetApiMethodName() string {
    return "taobao.ump.promotion.global.discount.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpPromotionGlobalDiscountGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
