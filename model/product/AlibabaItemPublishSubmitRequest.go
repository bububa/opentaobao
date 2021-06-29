package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品发布 API请求
alibaba.item.publish.submit

新商品发布，提交商品发布信息
*/
type AlibabaItemPublishSubmitRequest struct {
    model.Params
    // 业务扩展参数，需与平台约定好
    bizType   string
    // 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
    market   string
    // 商品类目ID
    catId   int64
    // 产品ID，天猫市场(market=tmall)时必填
    spuId   int64
    // 商品条码
    barcode   string
    // 商品schema信息，通过alibaba.item.publish.props.get获取并补全后提交
    schema   string
}

// 初始化AlibabaItemPublishSubmitRequest对象
func NewAlibabaItemPublishSubmitRequest() *AlibabaItemPublishSubmitRequest{
    return &AlibabaItemPublishSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemPublishSubmitRequest) GetApiMethodName() string {
    return "alibaba.item.publish.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemPublishSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务扩展参数，需与平台约定好
func (r *AlibabaItemPublishSubmitRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r AlibabaItemPublishSubmitRequest) GetBizType() string {
    return r.bizType
}
// Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaItemPublishSubmitRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

// Market Getter
func (r AlibabaItemPublishSubmitRequest) GetMarket() string {
    return r.market
}
// CatId Setter
// 商品类目ID
func (r *AlibabaItemPublishSubmitRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r AlibabaItemPublishSubmitRequest) GetCatId() int64 {
    return r.catId
}
// SpuId Setter
// 产品ID，天猫市场(market=tmall)时必填
func (r *AlibabaItemPublishSubmitRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

// SpuId Getter
func (r AlibabaItemPublishSubmitRequest) GetSpuId() int64 {
    return r.spuId
}
// Barcode Setter
// 商品条码
func (r *AlibabaItemPublishSubmitRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

// Barcode Getter
func (r AlibabaItemPublishSubmitRequest) GetBarcode() string {
    return r.barcode
}
// Schema Setter
// 商品schema信息，通过alibaba.item.publish.props.get获取并补全后提交
func (r *AlibabaItemPublishSubmitRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r AlibabaItemPublishSubmitRequest) GetSchema() string {
    return r.schema
}
