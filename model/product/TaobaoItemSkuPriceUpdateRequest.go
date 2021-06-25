package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品SKU的价格 APIRequest
taobao.item.sku.price.update

更新商品SKU的价格
*/
type TaobaoItemSkuPriceUpdateRequest struct {
    model.Params

    // Sku所属商品数字id，可通过 taobao.item.get 获取
    numIid   int64 

    // Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
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

func NewTaobaoItemSkuPriceUpdateRequest() *TaobaoItemSkuPriceUpdateRequest{
    return &TaobaoItemSkuPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemSkuPriceUpdateRequest) GetApiMethodName() string {
    return "taobao.item.sku.price.update"
}

func (r TaobaoItemSkuPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemSkuPriceUpdateRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemSkuPriceUpdateRequest) GetNumIid() int64 {
    return r.numIid
}

func (r *TaobaoItemSkuPriceUpdateRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

func (r TaobaoItemSkuPriceUpdateRequest) GetProperties() string {
    return r.properties
}

func (r *TaobaoItemSkuPriceUpdateRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

func (r TaobaoItemSkuPriceUpdateRequest) GetQuantity() int64 {
    return r.quantity
}

func (r *TaobaoItemSkuPriceUpdateRequest) SetPrice(price float64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

func (r TaobaoItemSkuPriceUpdateRequest) GetPrice() float64 {
    return r.price
}

func (r *TaobaoItemSkuPriceUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoItemSkuPriceUpdateRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoItemSkuPriceUpdateRequest) SetItemPrice(itemPrice float64) error {
    r.itemPrice = itemPrice
    r.Set("item_price", itemPrice)
    return nil
}

func (r TaobaoItemSkuPriceUpdateRequest) GetItemPrice() float64 {
    return r.itemPrice
}

func (r *TaobaoItemSkuPriceUpdateRequest) SetLang(lang string) error {
    r.lang = lang
    r.Set("lang", lang)
    return nil
}

func (r TaobaoItemSkuPriceUpdateRequest) GetLang() string {
    return r.lang
}

func (r *TaobaoItemSkuPriceUpdateRequest) SetIgnorewarning(ignorewarning string) error {
    r.ignorewarning = ignorewarning
    r.Set("ignorewarning", ignorewarning)
    return nil
}

func (r TaobaoItemSkuPriceUpdateRequest) GetIgnorewarning() string {
    return r.ignorewarning
}

