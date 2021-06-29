package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单可购买用户标记 APIRequest
taobao.opentrade.special.users.mark

对于专属下单的交易场景，根据openid标记用户可购买商品
*/
type TaobaoOpentradeSpecialUsersMarkRequest struct {
    model.Params

    // 是否目标用户，传入true后，用户可购买商品
    hit   bool 

    // 本次待标记的用户列表，多个以逗号(,)分割，最大20个
    openUserIds   []string 

    // 商品ID
    itemId   int64 

    // 商品SKU ID，不存在传0
    skuId   int64 

    // 用户状态，可任意传入，后续查询返回
    status   string 

    // 单次购买最大限购数量
    limitNum   int64 

}

func NewTaobaoOpentradeSpecialUsersMarkRequest() *TaobaoOpentradeSpecialUsersMarkRequest{
    return &TaobaoOpentradeSpecialUsersMarkRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeSpecialUsersMarkRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.users.mark"
}

func (r TaobaoOpentradeSpecialUsersMarkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeSpecialUsersMarkRequest) SetHit(hit bool) error {
    r.hit = hit
    r.Set("hit", hit)
    return nil
}

func (r TaobaoOpentradeSpecialUsersMarkRequest) GetHit() bool {
    return r.hit
}

func (r *TaobaoOpentradeSpecialUsersMarkRequest) SetOpenUserIds(openUserIds []string) error {
    r.openUserIds = openUserIds
    r.Set("open_user_ids", openUserIds)
    return nil
}

func (r TaobaoOpentradeSpecialUsersMarkRequest) GetOpenUserIds() []string {
    return r.openUserIds
}

func (r *TaobaoOpentradeSpecialUsersMarkRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOpentradeSpecialUsersMarkRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoOpentradeSpecialUsersMarkRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoOpentradeSpecialUsersMarkRequest) GetSkuId() int64 {
    return r.skuId
}

func (r *TaobaoOpentradeSpecialUsersMarkRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoOpentradeSpecialUsersMarkRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoOpentradeSpecialUsersMarkRequest) SetLimitNum(limitNum int64) error {
    r.limitNum = limitNum
    r.Set("limit_num", limitNum)
    return nil
}

func (r TaobaoOpentradeSpecialUsersMarkRequest) GetLimitNum() int64 {
    return r.limitNum
}

