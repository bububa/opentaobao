package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemPublishPropsGetAPIRequest 商品级联属性信息获取 API请求
// alibaba.item.publish.props.get
//
// 新商品发布，商品级联属性信息获取
type AlibabaItemPublishPropsGetAPIRequest struct {
	model.Params
	// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
	_market string
	// 商品条码
	_barcode string
	// 类目属性渲染schema
	_schema string
	// 商品类目ID
	_catId int64
	// 属性ID
	_propId int64
}

// NewAlibabaItemPublishPropsGetRequest 初始化AlibabaItemPublishPropsGetAPIRequest对象
func NewAlibabaItemPublishPropsGetRequest() *AlibabaItemPublishPropsGetAPIRequest {
	return &AlibabaItemPublishPropsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemPublishPropsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.item.publish.props.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemPublishPropsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItemPublishPropsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMarket is Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaItemPublishPropsGetAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// GetMarket Market Getter
func (r AlibabaItemPublishPropsGetAPIRequest) GetMarket() string {
	return r._market
}

// SetBarcode is Barcode Setter
// 商品条码
func (r *AlibabaItemPublishPropsGetAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r AlibabaItemPublishPropsGetAPIRequest) GetBarcode() string {
	return r._barcode
}

// SetSchema is Schema Setter
// 类目属性渲染schema
func (r *AlibabaItemPublishPropsGetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r AlibabaItemPublishPropsGetAPIRequest) GetSchema() string {
	return r._schema
}

// SetCatId is CatId Setter
// 商品类目ID
func (r *AlibabaItemPublishPropsGetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaItemPublishPropsGetAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetPropId is PropId Setter
// 属性ID
func (r *AlibabaItemPublishPropsGetAPIRequest) SetPropId(_propId int64) error {
	r._propId = _propId
	r.Set("prop_id", _propId)
	return nil
}

// GetPropId PropId Getter
func (r AlibabaItemPublishPropsGetAPIRequest) GetPropId() int64 {
	return r._propId
}
