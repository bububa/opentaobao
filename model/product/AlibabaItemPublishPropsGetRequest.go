package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品级联属性信息获取 APIRequest
alibaba.item.publish.props.get

新商品发布，商品级联属性信息获取
*/
type AlibabaItemPublishPropsGetRequest struct {
    model.Params

    // 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
    market   string 

    // 商品类目ID
    catId   int64 

    // 商品条码
    barcode   string 

    // 类目属性渲染schema
    schema   string 

    // 属性ID
    propId   int64 

}

func NewAlibabaItemPublishPropsGetRequest() *AlibabaItemPublishPropsGetRequest{
    return &AlibabaItemPublishPropsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItemPublishPropsGetRequest) GetApiMethodName() string {
    return "alibaba.item.publish.props.get"
}

func (r AlibabaItemPublishPropsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItemPublishPropsGetRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

func (r AlibabaItemPublishPropsGetRequest) GetMarket() string {
    return r.market
}

func (r *AlibabaItemPublishPropsGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r AlibabaItemPublishPropsGetRequest) GetCatId() int64 {
    return r.catId
}

func (r *AlibabaItemPublishPropsGetRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

func (r AlibabaItemPublishPropsGetRequest) GetBarcode() string {
    return r.barcode
}

func (r *AlibabaItemPublishPropsGetRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r AlibabaItemPublishPropsGetRequest) GetSchema() string {
    return r.schema
}

func (r *AlibabaItemPublishPropsGetRequest) SetPropId(propId int64) error {
    r.propId = propId
    r.Set("prop_id", propId)
    return nil
}

func (r AlibabaItemPublishPropsGetRequest) GetPropId() int64 {
    return r.propId
}

