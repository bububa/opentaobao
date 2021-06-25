package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据外部ID取商品SKU APIRequest
taobao.skus.custom.get

跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku <br/>这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
*/
type TaobaoSkusCustomGetRequest struct {
    model.Params

    // Sku的外部商家ID
    outerId   string 

    // 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开
    fields   string 

}

func NewTaobaoSkusCustomGetRequest() *TaobaoSkusCustomGetRequest{
    return &TaobaoSkusCustomGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSkusCustomGetRequest) GetApiMethodName() string {
    return "taobao.skus.custom.get"
}

func (r TaobaoSkusCustomGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSkusCustomGetRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoSkusCustomGetRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoSkusCustomGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoSkusCustomGetRequest) GetFields() string {
    return r.fields
}

