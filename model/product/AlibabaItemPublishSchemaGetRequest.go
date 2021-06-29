package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品发布规则信息 API请求
alibaba.item.publish.schema.get

新商品发布，获取商品发布规则信息
*/
type AlibabaItemPublishSchemaGetRequest struct {
    model.Params
    // 业务扩展参数，需与平台约定好
    bizType   string
    // 商品主图链接，最多5张，传入完整URL
    images   []string
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

// 初始化AlibabaItemPublishSchemaGetRequest对象
func NewAlibabaItemPublishSchemaGetRequest() *AlibabaItemPublishSchemaGetRequest{
    return &AlibabaItemPublishSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemPublishSchemaGetRequest) GetApiMethodName() string {
    return "alibaba.item.publish.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemPublishSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务扩展参数，需与平台约定好
func (r *AlibabaItemPublishSchemaGetRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r AlibabaItemPublishSchemaGetRequest) GetBizType() string {
    return r.bizType
}
// Images Setter
// 商品主图链接，最多5张，传入完整URL
func (r *AlibabaItemPublishSchemaGetRequest) SetImages(images []string) error {
    r.images = images
    r.Set("images", images)
    return nil
}

// Images Getter
func (r AlibabaItemPublishSchemaGetRequest) GetImages() []string {
    return r.images
}
// ItemType Setter
// 商品类型。b:一口价  a:拍卖  默认值b一口价
func (r *AlibabaItemPublishSchemaGetRequest) SetItemType(itemType string) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

// ItemType Getter
func (r AlibabaItemPublishSchemaGetRequest) GetItemType() string {
    return r.itemType
}
// Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaItemPublishSchemaGetRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

// Market Getter
func (r AlibabaItemPublishSchemaGetRequest) GetMarket() string {
    return r.market
}
// CatId Setter
// 商品类目ID
func (r *AlibabaItemPublishSchemaGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r AlibabaItemPublishSchemaGetRequest) GetCatId() int64 {
    return r.catId
}
// SpuId Setter
// 产品ID，天猫市场(market=tmall)时必填
func (r *AlibabaItemPublishSchemaGetRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

// SpuId Getter
func (r AlibabaItemPublishSchemaGetRequest) GetSpuId() int64 {
    return r.spuId
}
// Barcode Setter
// 商品条码
func (r *AlibabaItemPublishSchemaGetRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

// Barcode Getter
func (r AlibabaItemPublishSchemaGetRequest) GetBarcode() string {
    return r.barcode
}
