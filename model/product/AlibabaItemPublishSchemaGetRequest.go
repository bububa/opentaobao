package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品发布规则信息 APIRequest
alibaba.item.publish.schema.get

新商品发布，获取商品发布规则信息
*/
type AlibabaItemPublishSchemaGetRequest struct {
    model.Params

    // 业务扩展参数，需与平台约定好
    bizType   string 

    // 商品主图链接，最多5张，传入完整URL
    images   []String 

    // 商品类型。b:一口价  a:拍卖  默认值b一口价
    itemType   string 

    // 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
    market   string 

    // 商品类目ID
    catId   int64 

    // 产品ID，天猫市场(market=tmall)时必填
    spuId   int64 

    // 商品条码
    barcode   string 

}

func NewAlibabaItemPublishSchemaGetRequest() *AlibabaItemPublishSchemaGetRequest{
    return &AlibabaItemPublishSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItemPublishSchemaGetRequest) GetApiMethodName() string {
    return "alibaba.item.publish.schema.get"
}

func (r AlibabaItemPublishSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItemPublishSchemaGetRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r AlibabaItemPublishSchemaGetRequest) GetBizType() string {
    return r.bizType
}

func (r *AlibabaItemPublishSchemaGetRequest) SetImages(images []String) error {
    r.images = images
    r.Set("images", images)
    return nil
}

func (r AlibabaItemPublishSchemaGetRequest) GetImages() []String {
    return r.images
}

func (r *AlibabaItemPublishSchemaGetRequest) SetItemType(itemType string) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

func (r AlibabaItemPublishSchemaGetRequest) GetItemType() string {
    return r.itemType
}

func (r *AlibabaItemPublishSchemaGetRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

func (r AlibabaItemPublishSchemaGetRequest) GetMarket() string {
    return r.market
}

func (r *AlibabaItemPublishSchemaGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r AlibabaItemPublishSchemaGetRequest) GetCatId() int64 {
    return r.catId
}

func (r *AlibabaItemPublishSchemaGetRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

func (r AlibabaItemPublishSchemaGetRequest) GetSpuId() int64 {
    return r.spuId
}

func (r *AlibabaItemPublishSchemaGetRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

func (r AlibabaItemPublishSchemaGetRequest) GetBarcode() string {
    return r.barcode
}

