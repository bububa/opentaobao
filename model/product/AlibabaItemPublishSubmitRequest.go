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
    _bizType   string
    // 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
    _market   string
    // 商品类目ID
    _catId   int64
    // 产品ID，天猫市场(market=tmall)时必填
    _spuId   int64
    // 商品条码
    _barcode   string
    // 商品schema信息，通过alibaba.item.publish.props.get获取并补全后提交
    _schema   string
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
func (r *AlibabaItemPublishSubmitRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r AlibabaItemPublishSubmitRequest) GetBizType() string {
    return r._bizType
}
// Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaItemPublishSubmitRequest) SetMarket(_market string) error {
    r._market = _market
    r.Set("market", _market)
    return nil
}

// Market Getter
func (r AlibabaItemPublishSubmitRequest) GetMarket() string {
    return r._market
}
// CatId Setter
// 商品类目ID
func (r *AlibabaItemPublishSubmitRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlibabaItemPublishSubmitRequest) GetCatId() int64 {
    return r._catId
}
// SpuId Setter
// 产品ID，天猫市场(market=tmall)时必填
func (r *AlibabaItemPublishSubmitRequest) SetSpuId(_spuId int64) error {
    r._spuId = _spuId
    r.Set("spu_id", _spuId)
    return nil
}

// SpuId Getter
func (r AlibabaItemPublishSubmitRequest) GetSpuId() int64 {
    return r._spuId
}
// Barcode Setter
// 商品条码
func (r *AlibabaItemPublishSubmitRequest) SetBarcode(_barcode string) error {
    r._barcode = _barcode
    r.Set("barcode", _barcode)
    return nil
}

// Barcode Getter
func (r AlibabaItemPublishSubmitRequest) GetBarcode() string {
    return r._barcode
}
// Schema Setter
// 商品schema信息，通过alibaba.item.publish.props.get获取并补全后提交
func (r *AlibabaItemPublishSubmitRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r AlibabaItemPublishSubmitRequest) GetSchema() string {
    return r._schema
}
