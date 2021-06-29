package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品SKU的价格 API请求
taobao.item.sku.price.update

更新商品SKU的价格
*/
type TaobaoItemSkuPriceUpdateRequest struct {
    model.Params
    // Sku所属商品数字id，可通过 taobao.item.get 获取
    _numIid   int64
    // Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
    _properties   string
    // Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
    _quantity   int64
    // Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
    _price   float64
    // Sku的商家外部id
    _outerId   string
    // sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功
    _itemPrice   float64
    // Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
    _lang   string
    // 忽略警告提示.
    _ignorewarning   string
}

// 初始化TaobaoItemSkuPriceUpdateRequest对象
func NewTaobaoItemSkuPriceUpdateRequest() *TaobaoItemSkuPriceUpdateRequest{
    return &TaobaoItemSkuPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemSkuPriceUpdateRequest) GetApiMethodName() string {
    return "taobao.item.sku.price.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemSkuPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIid Setter
// Sku所属商品数字id，可通过 taobao.item.get 获取
func (r *TaobaoItemSkuPriceUpdateRequest) SetNumIid(_numIid int64) error {
    r._numIid = _numIid
    r.Set("num_iid", _numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemSkuPriceUpdateRequest) GetNumIid() int64 {
    return r._numIid
}
// Properties Setter
// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
func (r *TaobaoItemSkuPriceUpdateRequest) SetProperties(_properties string) error {
    r._properties = _properties
    r.Set("properties", _properties)
    return nil
}

// Properties Getter
func (r TaobaoItemSkuPriceUpdateRequest) GetProperties() string {
    return r._properties
}
// Quantity Setter
// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
func (r *TaobaoItemSkuPriceUpdateRequest) SetQuantity(_quantity int64) error {
    r._quantity = _quantity
    r.Set("quantity", _quantity)
    return nil
}

// Quantity Getter
func (r TaobaoItemSkuPriceUpdateRequest) GetQuantity() int64 {
    return r._quantity
}
// Price Setter
// Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
func (r *TaobaoItemSkuPriceUpdateRequest) SetPrice(_price float64) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r TaobaoItemSkuPriceUpdateRequest) GetPrice() float64 {
    return r._price
}
// OuterId Setter
// Sku的商家外部id
func (r *TaobaoItemSkuPriceUpdateRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoItemSkuPriceUpdateRequest) GetOuterId() string {
    return r._outerId
}
// ItemPrice Setter
// sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功
func (r *TaobaoItemSkuPriceUpdateRequest) SetItemPrice(_itemPrice float64) error {
    r._itemPrice = _itemPrice
    r.Set("item_price", _itemPrice)
    return nil
}

// ItemPrice Getter
func (r TaobaoItemSkuPriceUpdateRequest) GetItemPrice() float64 {
    return r._itemPrice
}
// Lang Setter
// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
func (r *TaobaoItemSkuPriceUpdateRequest) SetLang(_lang string) error {
    r._lang = _lang
    r.Set("lang", _lang)
    return nil
}

// Lang Getter
func (r TaobaoItemSkuPriceUpdateRequest) GetLang() string {
    return r._lang
}
// Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoItemSkuPriceUpdateRequest) SetIgnorewarning(_ignorewarning string) error {
    r._ignorewarning = _ignorewarning
    r.Set("ignorewarning", _ignorewarning)
    return nil
}

// Ignorewarning Getter
func (r TaobaoItemSkuPriceUpdateRequest) GetIgnorewarning() string {
    return r._ignorewarning
}
