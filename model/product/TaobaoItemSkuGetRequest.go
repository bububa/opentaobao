package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取SKU APIRequest
taobao.item.sku.get

获取sku_id所对应的sku数据 
sku_id对应的sku要属于传入的nick对应的卖家
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemSkuGetRequest struct {
    model.Params

    // 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
    fields   string 

    // Sku的id。可以通过taobao.item.seller.get得到
    skuId   int64 

    // 商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。
    numIid   int64 

}

func NewTaobaoItemSkuGetRequest() *TaobaoItemSkuGetRequest{
    return &TaobaoItemSkuGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemSkuGetRequest) GetApiMethodName() string {
    return "taobao.item.sku.get"
}

func (r TaobaoItemSkuGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemSkuGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoItemSkuGetRequest) GetFields() string {
    return r.fields
}

func (r *TaobaoItemSkuGetRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TaobaoItemSkuGetRequest) GetSkuId() int64 {
    return r.skuId
}

func (r *TaobaoItemSkuGetRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemSkuGetRequest) GetNumIid() int64 {
    return r.numIid
}

