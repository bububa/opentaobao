package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品已生效营销活动更新规则 API请求
taobao.item.promotion.rule.get

获取商品已生效的更新规则信息，主要包含库存禁止修改，商品一口价禁止修改，库存减少锁定等规则生效信息
*/
type TaobaoItemPromotionRuleGetRequest struct {
    model.Params
    // 商品ID
    itemId   int64
}

// 初始化TaobaoItemPromotionRuleGetRequest对象
func NewTaobaoItemPromotionRuleGetRequest() *TaobaoItemPromotionRuleGetRequest{
    return &TaobaoItemPromotionRuleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemPromotionRuleGetRequest) GetApiMethodName() string {
    return "taobao.item.promotion.rule.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemPromotionRuleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TaobaoItemPromotionRuleGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoItemPromotionRuleGetRequest) GetItemId() int64 {
    return r.itemId
}
