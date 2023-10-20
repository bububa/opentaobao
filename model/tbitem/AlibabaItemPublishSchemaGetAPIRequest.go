package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitempublishschemagetAPIRequest 获取商品发布规则信息 API请求
// alibaba.item.publish.schema.get
//
// 新商品发布，获取商品发布规则信息
type AlibabaitempublishschemagetAPIRequest struct {
	model.Params
	// 商品主图链接，最多5张，传入完整URL
	_images []string
	// 商品类型。b:一口价  a:拍卖  默认值b一口价
	_itemType string
	// 业务扩展参数，需与平台约定好
	_bizType string
	// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
	_market string
	// 商品条码
	_barcode string
	// 商品类目ID
	_catId int64
	// 产品ID，天猫市场(market=tmall)时必填
	_spuId int64
}

// NewAlibabaitempublishschemagetRequest 初始化AlibabaitempublishschemagetAPIRequest对象
func NewAlibabaitempublishschemagetRequest() *AlibabaitempublishschemagetAPIRequest {
	return &AlibabaitempublishschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitempublishschemagetAPIRequest) GetApiMethodName() string {
	return "alibaba.item.publish.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitempublishschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitempublishschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImages is Images Setter
// 商品主图链接，最多5张，传入完整URL
func (r *AlibabaitempublishschemagetAPIRequest) SetImages(_images []string) error {
	r._images = _images
	r.Set("images", _images)
	return nil
}

// GetImages Images Getter
func (r AlibabaitempublishschemagetAPIRequest) GetImages() []string {
	return r._images
}

// SetItemType is ItemType Setter
// 商品类型。b:一口价  a:拍卖  默认值b一口价
func (r *AlibabaitempublishschemagetAPIRequest) SetItemType(_itemType string) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// GetItemType ItemType Getter
func (r AlibabaitempublishschemagetAPIRequest) GetItemType() string {
	return r._itemType
}

// SetBizType is BizType Setter
// 业务扩展参数，需与平台约定好
func (r *AlibabaitempublishschemagetAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaitempublishschemagetAPIRequest) GetBizType() string {
	return r._bizType
}

// SetMarket is Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaitempublishschemagetAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// GetMarket Market Getter
func (r AlibabaitempublishschemagetAPIRequest) GetMarket() string {
	return r._market
}

// SetBarcode is Barcode Setter
// 商品条码
func (r *AlibabaitempublishschemagetAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r AlibabaitempublishschemagetAPIRequest) GetBarcode() string {
	return r._barcode
}

// SetCatId is CatId Setter
// 商品类目ID
func (r *AlibabaitempublishschemagetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaitempublishschemagetAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetSpuId is SpuId Setter
// 产品ID，天猫市场(market=tmall)时必填
func (r *AlibabaitempublishschemagetAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// GetSpuId SpuId Getter
func (r AlibabaitempublishschemagetAPIRequest) GetSpuId() int64 {
	return r._spuId
}
