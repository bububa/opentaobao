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
type TaobaoSimbaAdgroupsItemExistRequest struct {
    model.Params
    // 主人昵称
    nick   string
    // 推广计划Id
    campaignId   int64
    // 商品Id
    itemId   int64
    // 产品类型 101001005 代表普通推广，101001014代表销量明星
    productId   int64
}

// 初始化TaobaoSimbaAdgroupsItemExistRequest对象
func NewTaobaoSimbaAdgroupsItemExistRequest() *TaobaoSimbaAdgroupsItemExistRequest{
    return &TaobaoSimbaAdgroupsItemExistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupsItemExistRequest) GetApiMethodName() string {
    return "taobao.simba.adgroups.item.exist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupsItemExistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupsItemExistRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupsItemExistRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaAdgroupsItemExistRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaAdgroupsItemExistRequest) GetCampaignId() int64 {
    return r.campaignId
}
// ItemId Setter
// 商品Id
func (r *TaobaoSimbaAdgroupsItemExistRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoSimbaAdgroupsItemExistRequest) GetItemId() int64 {
    return r.itemId
}
// ProductId Setter
// 产品类型 101001005 代表普通推广，101001014代表销量明星
func (r *TaobaoSimbaAdgroupsItemExistRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoSimbaAdgroupsItemExistRequest) GetProductId() int64 {
    return r.productId
}
