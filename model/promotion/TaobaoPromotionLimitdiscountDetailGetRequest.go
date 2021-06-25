package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
限时打折详情查询 APIRequest
taobao.promotion.limitdiscount.detail.get

限时打折详情查询。查询出指定限时打折的对应商品记录信息。
*/
type TaobaoPromotionLimitdiscountDetailGetRequest struct {
    model.Params

    // 限时打折ID。这个针对查询唯一限时打折情况。
    limitDiscountId   int64 

}

func NewTaobaoPromotionLimitdiscountDetailGetRequest() *TaobaoPromotionLimitdiscountDetailGetRequest{
    return &TaobaoPromotionLimitdiscountDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionLimitdiscountDetailGetRequest) GetApiMethodName() string {
    return "taobao.promotion.limitdiscount.detail.get"
}

func (r TaobaoPromotionLimitdiscountDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionLimitdiscountDetailGetRequest) SetLimitDiscountId(limitDiscountId int64) error {
    r.limitDiscountId = limitDiscountId
    r.Set("limit_discount_id", limitDiscountId)
    return nil
}

func (r TaobaoPromotionLimitdiscountDetailGetRequest) GetLimitDiscountId() int64 {
    return r.limitDiscountId
}

