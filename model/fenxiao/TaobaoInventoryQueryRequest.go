package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品库存信息 APIRequest
taobao.inventory.query

建议使用新接口：tmall.inventory.query.forstore ，新ISV不推荐使用。
商家查询商品总体库存信息
*/
type TaobaoInventoryQueryRequest struct {
    model.Params

    // 后端商品ID 列表，控制到50个
    scItemIds   string 

    // 后端商品的商家编码列表，控制到50个
    scItemCodes   string 

    // 卖家昵称
    sellerNick   string 

    // 仓库列表:GLY001^GLY002
    storeCodes   string 

}

func NewTaobaoInventoryQueryRequest() *TaobaoInventoryQueryRequest{
    return &TaobaoInventoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryQueryRequest) GetApiMethodName() string {
    return "taobao.inventory.query"
}

func (r TaobaoInventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryQueryRequest) SetScItemIds(scItemIds string) error {
    r.scItemIds = scItemIds
    r.Set("sc_item_ids", scItemIds)
    return nil
}

func (r TaobaoInventoryQueryRequest) GetScItemIds() string {
    return r.scItemIds
}

func (r *TaobaoInventoryQueryRequest) SetScItemCodes(scItemCodes string) error {
    r.scItemCodes = scItemCodes
    r.Set("sc_item_codes", scItemCodes)
    return nil
}

func (r TaobaoInventoryQueryRequest) GetScItemCodes() string {
    return r.scItemCodes
}

func (r *TaobaoInventoryQueryRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

func (r TaobaoInventoryQueryRequest) GetSellerNick() string {
    return r.sellerNick
}

func (r *TaobaoInventoryQueryRequest) SetStoreCodes(storeCodes string) error {
    r.storeCodes = storeCodes
    r.Set("store_codes", storeCodes)
    return nil
}

func (r TaobaoInventoryQueryRequest) GetStoreCodes() string {
    return r.storeCodes
}

