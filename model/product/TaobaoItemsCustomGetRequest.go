package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据外部ID取商品 APIRequest
taobao.items.custom.get

跟据卖家设定的商品外部id获取商品，只能获取授权卖家的商品
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemsCustomGetRequest struct {
    model.Params

    // 商品的外部商品ID，支持批量，最多不超过40个。
    outerId   string 

    // 需返回的字段列表，参考：Item商品结构体说明，其中barcode、sku.barcode等条形码字段暂不支持；多个字段之间用“,”分隔。
    fields   string 

}

func NewTaobaoItemsCustomGetRequest() *TaobaoItemsCustomGetRequest{
    return &TaobaoItemsCustomGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemsCustomGetRequest) GetApiMethodName() string {
    return "taobao.items.custom.get"
}

func (r TaobaoItemsCustomGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemsCustomGetRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoItemsCustomGetRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoItemsCustomGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoItemsCustomGetRequest) GetFields() string {
    return r.fields
}

