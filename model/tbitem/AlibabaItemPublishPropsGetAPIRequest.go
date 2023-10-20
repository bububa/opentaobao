package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitempublishpropsgetAPIRequest 商品级联属性信息获取 API请求
// alibaba.item.publish.props.get
//
// 新商品发布，商品级联属性信息获取
type AlibabaitempublishpropsgetAPIRequest struct {
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

// NewAlibabaitempublishpropsgetRequest 初始化AlibabaitempublishpropsgetAPIRequest对象
func NewAlibabaitempublishpropsgetRequest() *AlibabaitempublishpropsgetAPIRequest {
	return &AlibabaitempublishpropsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitempublishpropsgetAPIRequest) GetApiMethodName() string {
	return "alibaba.item.publish.props.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitempublishpropsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitempublishpropsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMarket is Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaitempublishpropsgetAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// GetMarket Market Getter
func (r AlibabaitempublishpropsgetAPIRequest) GetMarket() string {
	return r._market
}

// SetBarcode is Barcode Setter
// 商品条码
func (r *AlibabaitempublishpropsgetAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r AlibabaitempublishpropsgetAPIRequest) GetBarcode() string {
	return r._barcode
}

// SetSchema is Schema Setter
// 类目属性渲染schema
func (r *AlibabaitempublishpropsgetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r AlibabaitempublishpropsgetAPIRequest) GetSchema() string {
	return r._schema
}

// SetCatId is CatId Setter
// 商品类目ID
func (r *AlibabaitempublishpropsgetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaitempublishpropsgetAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetPropId is PropId Setter
// 属性ID
func (r *AlibabaitempublishpropsgetAPIRequest) SetPropId(_propId int64) error {
	r._propId = _propId
	r.Set("prop_id", _propId)
	return nil
}

// GetPropId PropId Getter
func (r AlibabaitempublishpropsgetAPIRequest) GetPropId() int64 {
	return r._propId
}
