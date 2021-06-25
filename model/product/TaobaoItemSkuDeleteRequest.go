package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除SKU APIRequest
taobao.item.sku.delete

删除一个sku的数据<br/>需要删除的sku通过属性properties进行匹配查找
*/
type TaobaoItemSkuDeleteRequest struct {
    model.Params

    // Sku所属商品数字id，可通过 taobao.item.get 获取。必选
    numIid   int64 

    // Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充
    properties   string 

    // sku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功
    itemPrice   float64 

    // sku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败
    itemNum   int64 

    // Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN
    lang   string 

    // 忽略警告提示.
    ignorewarning   string 

}

func NewTaobaoItemSkuDeleteRequest() *TaobaoItemSkuDeleteRequest{
    return &TaobaoItemSkuDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemSkuDeleteRequest) GetApiMethodName() string {
    return "taobao.item.sku.delete"
}

func (r TaobaoItemSkuDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemSkuDeleteRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemSkuDeleteRequest) GetNumIid() int64 {
    return r.numIid
}

func (r *TaobaoItemSkuDeleteRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

func (r TaobaoItemSkuDeleteRequest) GetProperties() string {
    return r.properties
}

func (r *TaobaoItemSkuDeleteRequest) SetItemPrice(itemPrice float64) error {
    r.itemPrice = itemPrice
    r.Set("item_price", itemPrice)
    return nil
}

func (r TaobaoItemSkuDeleteRequest) GetItemPrice() float64 {
    return r.itemPrice
}

func (r *TaobaoItemSkuDeleteRequest) SetItemNum(itemNum int64) error {
    r.itemNum = itemNum
    r.Set("item_num", itemNum)
    return nil
}

func (r TaobaoItemSkuDeleteRequest) GetItemNum() int64 {
    return r.itemNum
}

func (r *TaobaoItemSkuDeleteRequest) SetLang(lang string) error {
    r.lang = lang
    r.Set("lang", lang)
    return nil
}

func (r TaobaoItemSkuDeleteRequest) GetLang() string {
    return r.lang
}

func (r *TaobaoItemSkuDeleteRequest) SetIgnorewarning(ignorewarning string) error {
    r.ignorewarning = ignorewarning
    r.Set("ignorewarning", ignorewarning)
    return nil
}

func (r TaobaoItemSkuDeleteRequest) GetIgnorewarning() string {
    return r.ignorewarning
}

