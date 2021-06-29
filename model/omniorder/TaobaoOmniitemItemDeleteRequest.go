package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品删除 API请求
taobao.omniitem.item.delete

全渠道商品删除，能够对门店商品库商品进行删除动作
*/
type TaobaoOmniitemItemDeleteRequest struct {
    model.Params
    // 条形码
    barCode   string
    // 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
    itemId   int64
    // 商品outerId
    outerId   string
}

// 初始化TaobaoOmniitemItemDeleteRequest对象
func NewTaobaoOmniitemItemDeleteRequest() *TaobaoOmniitemItemDeleteRequest{
    return &TaobaoOmniitemItemDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemDeleteRequest) GetApiMethodName() string {
    return "taobao.omniitem.item.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BarCode Setter
// 条形码
func (r *TaobaoOmniitemItemDeleteRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

// BarCode Getter
func (r TaobaoOmniitemItemDeleteRequest) GetBarCode() string {
    return r.barCode
}
// ItemId Setter
// 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
func (r *TaobaoOmniitemItemDeleteRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOmniitemItemDeleteRequest) GetItemId() int64 {
    return r.itemId
}
// OuterId Setter
// 商品outerId
func (r *TaobaoOmniitemItemDeleteRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoOmniitemItemDeleteRequest) GetOuterId() string {
    return r.outerId
}
