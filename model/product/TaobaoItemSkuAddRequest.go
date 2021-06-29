package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加SKU API请求
taobao.item.sku.add

新增一个sku到num_iid指定的商品中 <br/>传入的iid所对应的商品必须属于当前会话的用户
*/
type TaobaoItemSkuAddRequest struct {
    model.Params
    // Sku所属商品数字id。必选
    numIid   int64
    // Sku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
    properties   string
    // Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数
    quantity   int64
    // Sku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分
    price   float64
    // Sku的商家外部id
    outerId   string
    // sku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功
    itemPrice   float64
    // Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
    lang   string
    // 忽略警告提示.
    ignorewarning   string
}

// 初始化TaobaoItemSkuAddRequest对象
func NewTaobaoItemSkuAddRequest() *TaobaoItemSkuAddRequest{
    return &TaobaoItemSkuAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemSkuAddRequest) GetApiMethodName() string {
    return "taobao.item.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemSkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIid Setter
// Sku所属商品数字id。必选
func (r *TaobaoItemSkuAddRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemSkuAddRequest) GetNumIid() int64 {
    return r.numIid
}
// Properties Setter
// Sku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
func (r *TaobaoItemSkuAddRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

// Properties Getter
func (r TaobaoItemSkuAddRequest) GetProperties() string {
    return r.properties
}
// Quantity Setter
// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数
func (r *TaobaoItemSkuAddRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

// Quantity Getter
func (r TaobaoItemSkuAddRequest) GetQuantity() int64 {
    return r.quantity
}
// Price Setter
// Sku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分
func (r *TaobaoItemSkuAddRequest) SetPrice(price float64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

// Price Getter
func (r TaobaoItemSkuAddRequest) GetPrice() float64 {
    return r.price
}
// OuterId Setter
// Sku的商家外部id
func (r *TaobaoItemSkuAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoItemSkuAddRequest) GetOuterId() string {
    return r.outerId
}
// ItemPrice Setter
// sku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功
func (r *TaobaoItemSkuAddRequest) SetItemPrice(itemPrice float64) error {
    r.itemPrice = itemPrice
    r.Set("item_price", itemPrice)
    return nil
}

// ItemPrice Getter
func (r TaobaoItemSkuAddRequest) GetItemPrice() float64 {
    return r.itemPrice
}
// Lang Setter
// Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
func (r *TaobaoItemSkuAddRequest) SetLang(lang string) error {
    r.lang = lang
    r.Set("lang", lang)
    return nil
}

// Lang Getter
func (r TaobaoItemSkuAddRequest) GetLang() string {
    return r.lang
}
// Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoItemSkuAddRequest) SetIgnorewarning(ignorewarning string) error {
    r.ignorewarning = ignorewarning
    r.Set("ignorewarning", ignorewarning)
    return nil
}

// Ignorewarning Getter
func (r TaobaoItemSkuAddRequest) GetIgnorewarning() string {
    return r.ignorewarning
}
