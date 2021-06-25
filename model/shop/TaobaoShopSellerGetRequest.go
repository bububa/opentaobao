package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家店铺基础信息查询 APIRequest
taobao.shop.seller.get

获取卖家店铺的基本信息
*/
type TaobaoShopSellerGetRequest struct {
    model.Params

    // 需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔
    fields   string 

}

func NewTaobaoShopSellerGetRequest() *TaobaoShopSellerGetRequest{
    return &TaobaoShopSellerGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoShopSellerGetRequest) GetApiMethodName() string {
    return "taobao.shop.seller.get"
}

func (r TaobaoShopSellerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoShopSellerGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoShopSellerGetRequest) GetFields() string {
    return r.fields
}

