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
type TaobaoUmpPromotionGlobalDiscountGetRequest struct {
    model.Params
}

// 初始化TaobaoUmpPromotionGlobalDiscountGetRequest对象
func NewTaobaoUmpPromotionGlobalDiscountGetRequest() *TaobaoUmpPromotionGlobalDiscountGetRequest{
    return &TaobaoUmpPromotionGlobalDiscountGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpPromotionGlobalDiscountGetRequest) GetApiMethodName() string {
    return "taobao.ump.promotion.global.discount.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpPromotionGlobalDiscountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
