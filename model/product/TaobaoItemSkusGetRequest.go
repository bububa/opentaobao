package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID列表获取SKU信息 APIRequest
taobao.item.skus.get

* 获取多个商品下的所有sku
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemSkusGetRequest struct {
    model.Params

    // 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
    fields   []string 

    // sku所属商品数字id，必选。num_iid个数不能超过40个
    numIids   string 

}

func NewTaobaoItemSkusGetRequest() *TaobaoItemSkusGetRequest{
    return &TaobaoItemSkusGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemSkusGetRequest) GetApiMethodName() string {
    return "taobao.item.skus.get"
}

func (r TaobaoItemSkusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemSkusGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoItemSkusGetRequest) GetFields() []string {
    return r.fields
}

func (r *TaobaoItemSkusGetRequest) SetNumIids(numIids string) error {
    r.numIids = numIids
    r.Set("num_iids", numIids)
    return nil
}

func (r TaobaoItemSkusGetRequest) GetNumIids() string {
    return r.numIids
}

