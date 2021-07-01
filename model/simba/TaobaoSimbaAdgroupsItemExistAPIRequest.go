package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品是否推广 API请求
taobao.simba.adgroups.item.exist

判断在一个推广计划中是否已经推广了一个商品
*/
type TaobaoSimbaAdgroupsItemExistAPIRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广计划Id
    _campaignId   int64
    // 商品Id
    _itemId   int64
    // 产品类型 101001005 代表普通推广，101001014代表销量明星
    _productId   int64
}

// 初始化TaobaoSimbaAdgroupsItemExistAPIRequest对象
func NewTaobaoSimbaAdgroupsItemExistRequest() *TaobaoSimbaAdgroupsItemExistAPIRequest{
    return &TaobaoSimbaAdgroupsItemExistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupsItemExistAPIRequest) GetApiMethodName() string {
    return "taobao.simba.adgroups.item.exist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupsItemExistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupsItemExistAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupsItemExistAPIRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaAdgroupsItemExistAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaAdgroupsItemExistAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// ItemId Setter
// 商品Id
func (r *TaobaoSimbaAdgroupsItemExistAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoSimbaAdgroupsItemExistAPIRequest) GetItemId() int64 {
    return r._itemId
}
// ProductId Setter
// 产品类型 101001005 代表普通推广，101001014代表销量明星
func (r *TaobaoSimbaAdgroupsItemExistAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoSimbaAdgroupsItemExistAPIRequest) GetProductId() int64 {
    return r._productId
}
