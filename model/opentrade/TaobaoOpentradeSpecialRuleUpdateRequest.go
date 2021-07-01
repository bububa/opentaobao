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
type TaobaoOpentradeSpecialRuleUpdateAPIRequest struct {
    model.Params
    // 最大限购数量
    _limitNum   int64
    // 商品id列表
    _itemIds   []int64
}

// 初始化TaobaoOpentradeSpecialRuleUpdateAPIRequest对象
func NewTaobaoOpentradeSpecialRuleUpdateRequest() *TaobaoOpentradeSpecialRuleUpdateAPIRequest{
    return &TaobaoOpentradeSpecialRuleUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialRuleUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.rule.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialRuleUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LimitNum Setter
// 最大限购数量
func (r *TaobaoOpentradeSpecialRuleUpdateAPIRequest) SetLimitNum(_limitNum int64) error {
    r._limitNum = _limitNum
    r.Set("limit_num", _limitNum)
    return nil
}

// LimitNum Getter
func (r TaobaoOpentradeSpecialRuleUpdateAPIRequest) GetLimitNum() int64 {
    return r._limitNum
}
// ItemIds Setter
// 商品id列表
func (r *TaobaoOpentradeSpecialRuleUpdateAPIRequest) SetItemIds(_itemIds []int64) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoOpentradeSpecialRuleUpdateAPIRequest) GetItemIds() []int64 {
    return r._itemIds
}
