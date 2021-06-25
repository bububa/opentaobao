package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家最低折扣 APIRequest
taobao.ump.promotion.global.discount.get

提供卖家最低折扣查询功能
*/
type TaobaoUmpPromotionGlobalDiscountGetRequest struct {
    model.Params

}

func NewTaobaoUmpPromotionGlobalDiscountGetRequest() *TaobaoUmpPromotionGlobalDiscountGetRequest{
    return &TaobaoUmpPromotionGlobalDiscountGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpPromotionGlobalDiscountGetRequest) GetApiMethodName() string {
    return "taobao.ump.promotion.global.discount.get"
}

func (r TaobaoUmpPromotionGlobalDiscountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


