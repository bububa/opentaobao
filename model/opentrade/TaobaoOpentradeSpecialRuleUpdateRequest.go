package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单更新限购规则 API请求
taobao.opentrade.special.rule.update

对于专属下单的交易场景更新限购规则
*/
type TaobaoOpentradeSpecialRuleUpdateRequest struct {
    model.Params
    // 最大限购数量
    limitNum   int64
    // 商品id列表
    itemIds   []int64
}

// 初始化TaobaoOpentradeSpecialRuleUpdateRequest对象
func NewTaobaoOpentradeSpecialRuleUpdateRequest() *TaobaoOpentradeSpecialRuleUpdateRequest{
    return &TaobaoOpentradeSpecialRuleUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialRuleUpdateRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.rule.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialRuleUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LimitNum Setter
// 最大限购数量
func (r *TaobaoOpentradeSpecialRuleUpdateRequest) SetLimitNum(limitNum int64) error {
    r.limitNum = limitNum
    r.Set("limit_num", limitNum)
    return nil
}

// LimitNum Getter
func (r TaobaoOpentradeSpecialRuleUpdateRequest) GetLimitNum() int64 {
    return r.limitNum
}
// ItemIds Setter
// 商品id列表
func (r *TaobaoOpentradeSpecialRuleUpdateRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoOpentradeSpecialRuleUpdateRequest) GetItemIds() []int64 {
    return r.itemIds
}
