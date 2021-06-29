package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品级联属性信息获取 API请求
alibaba.item.publish.props.get

新商品发布，商品级联属性信息获取
*/
type AlibabaItemPublishPropsGetRequest struct {
    model.Params
    // 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
    _market   string
    // 商品类目ID
    _catId   int64
    // 商品条码
    _barcode   string
    // 类目属性渲染schema
    _schema   string
    // 属性ID
    _propId   int64
}

// 初始化AlibabaItemPublishPropsGetRequest对象
func NewAlibabaItemPublishPropsGetRequest() *AlibabaItemPublishPropsGetRequest{
    return &AlibabaItemPublishPropsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemPublishPropsGetRequest) GetApiMethodName() string {
    return "alibaba.item.publish.props.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemPublishPropsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaItemPublishPropsGetRequest) SetMarket(_market string) error {
    r._market = _market
    r.Set("market", _market)
    return nil
}

// Market Getter
func (r AlibabaItemPublishPropsGetRequest) GetMarket() string {
    return r._market
}
// CatId Setter
// 商品类目ID
func (r *AlibabaItemPublishPropsGetRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlibabaItemPublishPropsGetRequest) GetCatId() int64 {
    return r._catId
}
// Barcode Setter
// 商品条码
func (r *AlibabaItemPublishPropsGetRequest) SetBarcode(_barcode string) error {
    r._barcode = _barcode
    r.Set("barcode", _barcode)
    return nil
}

// Barcode Getter
func (r AlibabaItemPublishPropsGetRequest) GetBarcode() string {
    return r._barcode
}
// Schema Setter
// 类目属性渲染schema
func (r *AlibabaItemPublishPropsGetRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r AlibabaItemPublishPropsGetRequest) GetSchema() string {
    return r._schema
}
// PropId Setter
// 属性ID
func (r *AlibabaItemPublishPropsGetRequest) SetPropId(_propId int64) error {
    r._propId = _propId
    r.Set("prop_id", _propId)
    return nil
}

// PropId Getter
func (r AlibabaItemPublishPropsGetRequest) GetPropId() int64 {
    return r._propId
}
