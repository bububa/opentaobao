package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionSaleQueryAPIRequest 查询商品销售区域 API请求
// taobao.region.sale.query
//
// 查询商品销售区域
type TaobaoRegionSaleQueryAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 无sku传0
	_skuId int64
	// 1:国家;2:省;3: 市;4:区县
	_saleRegionLevel int64
}

// NewTaobaoRegionSaleQueryRequest 初始化TaobaoRegionSaleQueryAPIRequest对象
func NewTaobaoRegionSaleQueryRequest() *TaobaoRegionSaleQueryAPIRequest {
	return &TaobaoRegionSaleQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRegionSaleQueryAPIRequest) GetApiMethodName() string {
	return "taobao.region.sale.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRegionSaleQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoRegionSaleQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoRegionSaleQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// 无sku传0
func (r *TaobaoRegionSaleQueryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoRegionSaleQueryAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetSaleRegionLevel is SaleRegionLevel Setter
// 1:国家;2:省;3: 市;4:区县
func (r *TaobaoRegionSaleQueryAPIRequest) SetSaleRegionLevel(_saleRegionLevel int64) error {
	r._saleRegionLevel = _saleRegionLevel
	r.Set("sale_region_level", _saleRegionLevel)
	return nil
}

// GetSaleRegionLevel SaleRegionLevel Getter
func (r TaobaoRegionSaleQueryAPIRequest) GetSaleRegionLevel() int64 {
	return r._saleRegionLevel
}
