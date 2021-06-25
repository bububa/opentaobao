package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个商品详细信息 APIRequest
taobao.item.seller.get

获取单个商品的全部信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemSellerGetRequest struct {
    model.Params

    // 需要返回的商品字段列表。可选值：Item商品结构体中所有字段均可返回，多个字段用“,”分隔。
    fields   string 

    // 商品数字ID
    numIid   int64 

}

func NewTaobaoItemSellerGetRequest() *TaobaoItemSellerGetRequest{
    return &TaobaoItemSellerGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemSellerGetRequest) GetApiMethodName() string {
    return "taobao.item.seller.get"
}

func (r TaobaoItemSellerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemSellerGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoItemSellerGetRequest) GetFields() string {
    return r.fields
}

func (r *TaobaoItemSellerGetRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemSellerGetRequest) GetNumIid() int64 {
    return r.numIid
}

