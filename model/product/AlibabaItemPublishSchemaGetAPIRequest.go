package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemPublishSchemaGetAPIRequest 获取商品发布规则信息 API请求
// alibaba.item.publish.schema.get
//
// 新商品发布，获取商品发布规则信息
type AlibabaItemPublishSchemaGetAPIRequest struct {
	model.Params
	// 业务扩展参数，需与平台约定好
	_bizType string
	// 商品主图链接，最多5张，传入完整URL
	_images []string
	// 商品类型。b:一口价  a:拍卖  默认值b一口价
	_itemType string
	// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
	_market string
	// 商品类目ID
	_catId int64
	// 产品ID，天猫市场(market=tmall)时必填
	_spuId int64
	// 商品条码
	_barcode string
}

// NewAlibabaItemPublishSchemaGetRequest 初始化AlibabaItemPublishSchemaGetAPIRequest对象
func NewAlibabaItemPublishSchemaGetRequest() *AlibabaItemPublishSchemaGetAPIRequest {
	return &AlibabaItemPublishSchemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemPublishSchemaGetAPIRequest) GetApiMethodName() string {
	return "alibaba.item.publish.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemPublishSchemaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizType Setter
// 业务扩展参数，需与平台约定好
func (r *AlibabaItemPublishSchemaGetAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// Get BizType Getter
func (r AlibabaItemPublishSchemaGetAPIRequest) GetBizType() string {
	return r._bizType
}

// Set is Images Setter
// 商品主图链接，最多5张，传入完整URL
func (r *AlibabaItemPublishSchemaGetAPIRequest) SetImages(_images []string) error {
	r._images = _images
	r.Set("images", _images)
	return nil
}

// Get Images Getter
func (r AlibabaItemPublishSchemaGetAPIRequest) GetImages() []string {
	return r._images
}

// Set is ItemType Setter
// 商品类型。b:一口价  a:拍卖  默认值b一口价
func (r *AlibabaItemPublishSchemaGetAPIRequest) SetItemType(_itemType string) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// Get ItemType Getter
func (r AlibabaItemPublishSchemaGetAPIRequest) GetItemType() string {
	return r._itemType
}

// Set is Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaItemPublishSchemaGetAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// Get Market Getter
func (r AlibabaItemPublishSchemaGetAPIRequest) GetMarket() string {
	return r._market
}

// Set is CatId Setter
// 商品类目ID
func (r *AlibabaItemPublishSchemaGetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// Get CatId Getter
func (r AlibabaItemPublishSchemaGetAPIRequest) GetCatId() int64 {
	return r._catId
}

// Set is SpuId Setter
// 产品ID，天猫市场(market=tmall)时必填
func (r *AlibabaItemPublishSchemaGetAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// Get SpuId Getter
func (r AlibabaItemPublishSchemaGetAPIRequest) GetSpuId() int64 {
	return r._spuId
}

// Set is Barcode Setter
// 商品条码
func (r *AlibabaItemPublishSchemaGetAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// Get Barcode Getter
func (r AlibabaItemPublishSchemaGetAPIRequest) GetBarcode() string {
	return r._barcode
}
