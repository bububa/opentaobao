package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品库存信息 API请求
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

// 初始化TaobaoInventoryQueryRequest对象
func NewTaobaoInventoryQueryRequest() *TaobaoInventoryQueryRequest{
    return &TaobaoInventoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryQueryRequest) GetApiMethodName() string {
    return "taobao.inventory.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScItemIds Setter
// 后端商品ID 列表，控制到50个
func (r *TaobaoInventoryQueryRequest) SetScItemIds(scItemIds string) error {
    r.scItemIds = scItemIds
    r.Set("sc_item_ids", scItemIds)
    return nil
}

// ScItemIds Getter
func (r TaobaoInventoryQueryRequest) GetScItemIds() string {
    return r.scItemIds
}
// ScItemCodes Setter
// 后端商品的商家编码列表，控制到50个
func (r *TaobaoInventoryQueryRequest) SetScItemCodes(scItemCodes string) error {
    r.scItemCodes = scItemCodes
    r.Set("sc_item_codes", scItemCodes)
    return nil
}

// ScItemCodes Getter
func (r TaobaoInventoryQueryRequest) GetScItemCodes() string {
    return r.scItemCodes
}
// SellerNick Setter
// 卖家昵称
func (r *TaobaoInventoryQueryRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

// SellerNick Getter
func (r TaobaoInventoryQueryRequest) GetSellerNick() string {
    return r.sellerNick
}
// StoreCodes Setter
// 仓库列表:GLY001^GLY002
func (r *TaobaoInventoryQueryRequest) SetStoreCodes(storeCodes string) error {
    r.storeCodes = storeCodes
    r.Set("store_codes", storeCodes)
    return nil
}

// StoreCodes Getter
func (r TaobaoInventoryQueryRequest) GetStoreCodes() string {
    return r.storeCodes
}
