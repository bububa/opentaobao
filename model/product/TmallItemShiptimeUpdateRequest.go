package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品发货时间 API请求
tmall.item.shiptime.update

增加更新删除商品/SKU发货时间(支持同一商品下的SKU同时批量更新)
1.
    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 0 ---更新SKU
    },

    按照指定SKU更新指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空

2.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 0 --更新SKU
    },
    按照指定SKU删除指定SKU的发货时间

3.

    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 1 ---更新商品
    },

    更新商品级发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间

4.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 1 --更新商品
    },
    删除商品级的发货时间
*/
type TmallItemShiptimeUpdateRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
    // 被更新发货时间（商品级）；格式和具体设置的发货时间格式相关。绝对发货时间填写yyyy-MM-dd;相对发货时间填写数字。发货时间必须在当前时间后三天。如果设置的绝对时间小于当前时间的三天后，会清除该商品的发货时间设置。如果是相对时间小于3，则会提示出错。如果shiptimeType为0，要清除商品上的发货时间，该字段可以填任意字符,也可以不填。
    _shipTime   string
    // 被更新SKU的发货时间，后台会根据三个子属性去查找匹配的sku，如果找到就默认对sku进行更新，当无匹配sku且更新类型针对sku，会报错。
    _skuShipTimes   []UpdateSkuShipTime
    // 批量更新商品/SKU发货时间的备选项
    _option   *UpdateItemShipTimeOption
}

// 初始化TmallItemShiptimeUpdateRequest对象
func NewTmallItemShiptimeUpdateRequest() *TmallItemShiptimeUpdateRequest{
    return &TmallItemShiptimeUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemShiptimeUpdateRequest) GetApiMethodName() string {
    return "tmall.item.shiptime.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemShiptimeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TmallItemShiptimeUpdateRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallItemShiptimeUpdateRequest) GetItemId() int64 {
    return r._itemId
}
// ShipTime Setter
// 被更新发货时间（商品级）；格式和具体设置的发货时间格式相关。绝对发货时间填写yyyy-MM-dd;相对发货时间填写数字。发货时间必须在当前时间后三天。如果设置的绝对时间小于当前时间的三天后，会清除该商品的发货时间设置。如果是相对时间小于3，则会提示出错。如果shiptimeType为0，要清除商品上的发货时间，该字段可以填任意字符,也可以不填。
func (r *TmallItemShiptimeUpdateRequest) SetShipTime(_shipTime string) error {
    r._shipTime = _shipTime
    r.Set("ship_time", _shipTime)
    return nil
}

// ShipTime Getter
func (r TmallItemShiptimeUpdateRequest) GetShipTime() string {
    return r._shipTime
}
// SkuShipTimes Setter
// 被更新SKU的发货时间，后台会根据三个子属性去查找匹配的sku，如果找到就默认对sku进行更新，当无匹配sku且更新类型针对sku，会报错。
func (r *TmallItemShiptimeUpdateRequest) SetSkuShipTimes(_skuShipTimes []UpdateSkuShipTime) error {
    r._skuShipTimes = _skuShipTimes
    r.Set("sku_ship_times", _skuShipTimes)
    return nil
}

// SkuShipTimes Getter
func (r TmallItemShiptimeUpdateRequest) GetSkuShipTimes() []UpdateSkuShipTime {
    return r._skuShipTimes
}
// Option Setter
// 批量更新商品/SKU发货时间的备选项
func (r *TmallItemShiptimeUpdateRequest) SetOption(_option *UpdateItemShipTimeOption) error {
    r._option = _option
    r.Set("option", _option)
    return nil
}

// Option Getter
func (r TmallItemShiptimeUpdateRequest) GetOption() *UpdateItemShipTimeOption {
    return r._option
}
