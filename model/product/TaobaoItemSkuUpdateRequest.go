package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新SKU信息 API请求
taobao.item.sku.update

*更新一个sku的数据 <br/>*需要更新的sku通过属性properties进行匹配查找 <br/>*商品的数量和价格必须大于等于0 <br/>*sku记录会更新到指定的num_iid对应的商品中 <br/>*num_iid对应的商品必须属于当前的会话用户
*/
type TaobaoItemSkuUpdateRequest struct {
    model.Params
    // Sku所属商品数字id，可通过 taobao.item.get 获取
    numIid   int64
    // Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充。<br/>如果包含自定义属性，则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号，
    properties   string
    // Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
    quantity   int64
    // Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
    price   float64
    // Sku的商家外部id
    outerId   string
    // sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功
    itemPrice   float64
    // Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
    lang   string
    // 忽略警告提示.
    ignorewarning   string
}

// 初始化TaobaoItemSkuUpdateRequest对象
func NewTaobaoItemSkuUpdateRequest() *TaobaoItemSkuUpdateRequest{
    return &TaobaoItemSkuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemSkuUpdateRequest) GetApiMethodName() string {
    return "taobao.item.sku.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemSkuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIid Setter
// Sku所属商品数字id，可通过 taobao.item.get 获取
func (r *TaobaoItemSkuUpdateRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemSkuUpdateRequest) GetNumIid() int64 {
    return r.numIid
}
// Properties Setter
// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充。<br/>如果包含自定义属性，则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号，
func (r *TaobaoItemSkuUpdateRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

// Properties Getter
func (r TaobaoItemSkuUpdateRequest) GetProperties() string {
    return r.properties
}
// Quantity Setter
// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
func (r *TaobaoItemSkuUpdateRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

// Quantity Getter
func (r TaobaoItemSkuUpdateRequest) GetQuantity() int64 {
    return r.quantity
}
// Price Setter
// Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
func (r *TaobaoItemSkuUpdateRequest) SetPrice(price float64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

// Price Getter
func (r TaobaoItemSkuUpdateRequest) GetPrice() float64 {
    return r.price
}
// OuterId Setter
// Sku的商家外部id
func (r *TaobaoItemSkuUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoItemSkuUpdateRequest) GetOuterId() string {
    return r.outerId
}
// ItemPrice Setter
// sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功
func (r *TaobaoItemSkuUpdateRequest) SetItemPrice(itemPrice float64) error {
    r.itemPrice = itemPrice
    r.Set("item_price", itemPrice)
    return nil
}

// ItemPrice Getter
func (r TaobaoItemSkuUpdateRequest) GetItemPrice() float64 {
    return r.itemPrice
}
// Lang Setter
// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
func (r *TaobaoItemSkuUpdateRequest) SetLang(lang string) error {
    r.lang = lang
    r.Set("lang", lang)
    return nil
}

// Lang Getter
func (r TaobaoItemSkuUpdateRequest) GetLang() string {
    return r.lang
}
// Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoItemSkuUpdateRequest) SetIgnorewarning(ignorewarning string) error {
    r.ignorewarning = ignorewarning
    r.Set("ignorewarning", ignorewarning)
    return nil
}

// Ignorewarning Getter
func (r TaobaoItemSkuUpdateRequest) GetIgnorewarning() string {
    return r.ignorewarning
}
