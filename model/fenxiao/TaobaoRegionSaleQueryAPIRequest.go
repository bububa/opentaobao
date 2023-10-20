package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoregionsalequeryAPIRequest 查询商品销售区域 API请求
// taobao.region.sale.query
//
// 查询商品销售区域
type TaobaoregionsalequeryAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 无sku传0
	_skuId int64
	// 1:国家;2:省;3: 市;4:区县
	_saleRegionLevel int64
}

// NewTaobaoregionsalequeryRequest 初始化TaobaoregionsalequeryAPIRequest对象
func NewTaobaoregionsalequeryRequest() *TaobaoregionsalequeryAPIRequest {
	return &TaobaoregionsalequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoregionsalequeryAPIRequest) GetApiMethodName() string {
	return "taobao.region.sale.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoregionsalequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoregionsalequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoregionsalequeryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoregionsalequeryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// 无sku传0
func (r *TaobaoregionsalequeryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoregionsalequeryAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetSaleRegionLevel is SaleRegionLevel Setter
// 1:国家;2:省;3: 市;4:区县
func (r *TaobaoregionsalequeryAPIRequest) SetSaleRegionLevel(_saleRegionLevel int64) error {
	r._saleRegionLevel = _saleRegionLevel
	r.Set("sale_region_level", _saleRegionLevel)
	return nil
}

// GetSaleRegionLevel SaleRegionLevel Getter
func (r TaobaoregionsalequeryAPIRequest) GetSaleRegionLevel() int64 {
	return r._saleRegionLevel
}
