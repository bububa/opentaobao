package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemPublishSubmitAPIRequest 商品发布 API请求
// alibaba.item.publish.submit
//
// 新商品发布，提交商品发布信息
type AlibabaItemPublishSubmitAPIRequest struct {
	model.Params
	// 业务扩展参数，需与平台约定好
	_bizType string
	// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
	_market string
	// 商品类目ID
	_catId int64
	// 产品ID，天猫市场(market=tmall)时必填
	_spuId int64
	// 商品条码
	_barcode string
	// 商品schema信息，通过alibaba.item.publish.props.get获取并补全后提交
	_schema string
}

// NewAlibabaItemPublishSubmitRequest 初始化AlibabaItemPublishSubmitAPIRequest对象
func NewAlibabaItemPublishSubmitRequest() *AlibabaItemPublishSubmitAPIRequest {
	return &AlibabaItemPublishSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemPublishSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.item.publish.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemPublishSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizType is BizType Setter
// 业务扩展参数，需与平台约定好
func (r *AlibabaItemPublishSubmitAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaItemPublishSubmitAPIRequest) GetBizType() string {
	return r._bizType
}

// SetMarket is Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaItemPublishSubmitAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// GetMarket Market Getter
func (r AlibabaItemPublishSubmitAPIRequest) GetMarket() string {
	return r._market
}

// SetCatId is CatId Setter
// 商品类目ID
func (r *AlibabaItemPublishSubmitAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaItemPublishSubmitAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetSpuId is SpuId Setter
// 产品ID，天猫市场(market=tmall)时必填
func (r *AlibabaItemPublishSubmitAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// GetSpuId SpuId Getter
func (r AlibabaItemPublishSubmitAPIRequest) GetSpuId() int64 {
	return r._spuId
}

// SetBarcode is Barcode Setter
// 商品条码
func (r *AlibabaItemPublishSubmitAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r AlibabaItemPublishSubmitAPIRequest) GetBarcode() string {
	return r._barcode
}

// SetSchema is Schema Setter
// 商品schema信息，通过alibaba.item.publish.props.get获取并补全后提交
func (r *AlibabaItemPublishSubmitAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r AlibabaItemPublishSubmitAPIRequest) GetSchema() string {
	return r._schema
}
