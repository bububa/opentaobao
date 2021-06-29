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
    _barCode   string
    // 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
    _itemId   int64
    // 商品outerId
    _outerId   string
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
func (r *TaobaoOmniitemItemDeleteRequest) SetBarCode(_barCode string) error {
    r._barCode = _barCode
    r.Set("bar_code", _barCode)
    return nil
}

// BarCode Getter
func (r TaobaoOmniitemItemDeleteRequest) GetBarCode() string {
    return r._barCode
}
// ItemId Setter
// 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
func (r *TaobaoOmniitemItemDeleteRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOmniitemItemDeleteRequest) GetItemId() int64 {
    return r._itemId
}
// OuterId Setter
// 商品outerId
func (r *TaobaoOmniitemItemDeleteRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoOmniitemItemDeleteRequest) GetOuterId() string {
    return r._outerId
}
