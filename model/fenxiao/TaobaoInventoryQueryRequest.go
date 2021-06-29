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
    _scItemIds   string
    // 后端商品的商家编码列表，控制到50个
    _scItemCodes   string
    // 卖家昵称
    _sellerNick   string
    // 仓库列表:GLY001^GLY002
    _storeCodes   string
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
func (r *TaobaoInventoryQueryRequest) SetScItemIds(_scItemIds string) error {
    r._scItemIds = _scItemIds
    r.Set("sc_item_ids", _scItemIds)
    return nil
}

// ScItemIds Getter
func (r TaobaoInventoryQueryRequest) GetScItemIds() string {
    return r._scItemIds
}
// ScItemCodes Setter
// 后端商品的商家编码列表，控制到50个
func (r *TaobaoInventoryQueryRequest) SetScItemCodes(_scItemCodes string) error {
    r._scItemCodes = _scItemCodes
    r.Set("sc_item_codes", _scItemCodes)
    return nil
}

// ScItemCodes Getter
func (r TaobaoInventoryQueryRequest) GetScItemCodes() string {
    return r._scItemCodes
}
// SellerNick Setter
// 卖家昵称
func (r *TaobaoInventoryQueryRequest) SetSellerNick(_sellerNick string) error {
    r._sellerNick = _sellerNick
    r.Set("seller_nick", _sellerNick)
    return nil
}

// SellerNick Getter
func (r TaobaoInventoryQueryRequest) GetSellerNick() string {
    return r._sellerNick
}
// StoreCodes Setter
// 仓库列表:GLY001^GLY002
func (r *TaobaoInventoryQueryRequest) SetStoreCodes(_storeCodes string) error {
    r._storeCodes = _storeCodes
    r.Set("store_codes", _storeCodes)
    return nil
}

// StoreCodes Getter
func (r TaobaoInventoryQueryRequest) GetStoreCodes() string {
    return r._storeCodes
}
