package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单可购买用户标记 API请求
taobao.opentrade.special.users.mark

对于专属下单的交易场景，根据openid标记用户可购买商品
*/
type TaobaoOpentradeSpecialUsersMarkAPIRequest struct {
    model.Params
    // 是否目标用户，传入true后，用户可购买商品
    _hit   bool
    // 本次待标记的用户列表，多个以逗号(,)分割，最大20个
    _openUserIds   []string
    // 商品ID
    _itemId   int64
    // 商品SKU ID，不存在传0
    _skuId   int64
    // 用户状态，可任意传入，后续查询返回
    _status   string
    // 单次购买最大限购数量
    _limitNum   int64
}

// 初始化TaobaoOpentradeSpecialUsersMarkAPIRequest对象
func NewTaobaoOpentradeSpecialUsersMarkRequest() *TaobaoOpentradeSpecialUsersMarkAPIRequest{
    return &TaobaoOpentradeSpecialUsersMarkAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialUsersMarkAPIRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.users.mark"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialUsersMarkAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hit Setter
// 是否目标用户，传入true后，用户可购买商品
func (r *TaobaoOpentradeSpecialUsersMarkAPIRequest) SetHit(_hit bool) error {
    r._hit = _hit
    r.Set("hit", _hit)
    return nil
}

// Hit Getter
func (r TaobaoOpentradeSpecialUsersMarkAPIRequest) GetHit() bool {
    return r._hit
}
// OpenUserIds Setter
// 本次待标记的用户列表，多个以逗号(,)分割，最大20个
func (r *TaobaoOpentradeSpecialUsersMarkAPIRequest) SetOpenUserIds(_openUserIds []string) error {
    r._openUserIds = _openUserIds
    r.Set("open_user_ids", _openUserIds)
    return nil
}

// OpenUserIds Getter
func (r TaobaoOpentradeSpecialUsersMarkAPIRequest) GetOpenUserIds() []string {
    return r._openUserIds
}
// ItemId Setter
// 商品ID
func (r *TaobaoOpentradeSpecialUsersMarkAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeSpecialUsersMarkAPIRequest) GetItemId() int64 {
    return r._itemId
}
// SkuId Setter
// 商品SKU ID，不存在传0
func (r *TaobaoOpentradeSpecialUsersMarkAPIRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoOpentradeSpecialUsersMarkAPIRequest) GetSkuId() int64 {
    return r._skuId
}
// Status Setter
// 用户状态，可任意传入，后续查询返回
func (r *TaobaoOpentradeSpecialUsersMarkAPIRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoOpentradeSpecialUsersMarkAPIRequest) GetStatus() string {
    return r._status
}
// LimitNum Setter
// 单次购买最大限购数量
func (r *TaobaoOpentradeSpecialUsersMarkAPIRequest) SetLimitNum(_limitNum int64) error {
    r._limitNum = _limitNum
    r.Set("limit_num", _limitNum)
    return nil
}

// LimitNum Getter
func (r TaobaoOpentradeSpecialUsersMarkAPIRequest) GetLimitNum() int64 {
    return r._limitNum
}
