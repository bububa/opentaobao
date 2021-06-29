package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据外部ID取商品SKU API请求
taobao.skus.custom.get

跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku <br/>这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
*/
type TaobaoSkusCustomGetRequest struct {
    model.Params
    // Sku的外部商家ID
    _outerId   string
    // 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开
    _fields   string
}

// 初始化TaobaoSkusCustomGetRequest对象
func NewTaobaoSkusCustomGetRequest() *TaobaoSkusCustomGetRequest{
    return &TaobaoSkusCustomGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSkusCustomGetRequest) GetApiMethodName() string {
    return "taobao.skus.custom.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSkusCustomGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// Sku的外部商家ID
func (r *TaobaoSkusCustomGetRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoSkusCustomGetRequest) GetOuterId() string {
    return r._outerId
}
// Fields Setter
// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开
func (r *TaobaoSkusCustomGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoSkusCustomGetRequest) GetFields() string {
    return r._fields
}
