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
    _bizType   string
    // 商品主图链接，最多5张，传入完整URL
    _images   []string
    // 商品类型。b:一口价  a:拍卖  默认值b一口价
    _itemType   string
    // 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
    _market   string
    // 商品类目ID
    _catId   int64
    // 产品ID，天猫市场(market=tmall)时必填
    _spuId   int64
    // 商品条码
    _barcode   string
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
func (r *AlibabaItemPublishSchemaGetRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r AlibabaItemPublishSchemaGetRequest) GetBizType() string {
    return r._bizType
}
// Images Setter
// 商品主图链接，最多5张，传入完整URL
func (r *AlibabaItemPublishSchemaGetRequest) SetImages(_images []string) error {
    r._images = _images
    r.Set("images", _images)
    return nil
}

// Images Getter
func (r AlibabaItemPublishSchemaGetRequest) GetImages() []string {
    return r._images
}
// ItemType Setter
// 商品类型。b:一口价  a:拍卖  默认值b一口价
func (r *AlibabaItemPublishSchemaGetRequest) SetItemType(_itemType string) error {
    r._itemType = _itemType
    r.Set("item_type", _itemType)
    return nil
}

// ItemType Getter
func (r AlibabaItemPublishSchemaGetRequest) GetItemType() string {
    return r._itemType
}
// Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaItemPublishSchemaGetRequest) SetMarket(_market string) error {
    r._market = _market
    r.Set("market", _market)
    return nil
}

// Market Getter
func (r AlibabaItemPublishSchemaGetRequest) GetMarket() string {
    return r._market
}
// CatId Setter
// 商品类目ID
func (r *AlibabaItemPublishSchemaGetRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlibabaItemPublishSchemaGetRequest) GetCatId() int64 {
    return r._catId
}
// SpuId Setter
// 产品ID，天猫市场(market=tmall)时必填
func (r *AlibabaItemPublishSchemaGetRequest) SetSpuId(_spuId int64) error {
    r._spuId = _spuId
    r.Set("spu_id", _spuId)
    return nil
}

// SpuId Getter
func (r AlibabaItemPublishSchemaGetRequest) GetSpuId() int64 {
    return r._spuId
}
// Barcode Setter
// 商品条码
func (r *AlibabaItemPublishSchemaGetRequest) SetBarcode(_barcode string) error {
    r._barcode = _barcode
    r.Set("barcode", _barcode)
    return nil
}

// Barcode Getter
func (r AlibabaItemPublishSchemaGetRequest) GetBarcode() string {
    return r._barcode
}
