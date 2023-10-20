package fenxiao

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRegionSaleQueryAPIRequest) Reset() {
	r._itemId = 0
	r._skuId = 0
	r._saleRegionLevel = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRegionSaleQueryAPIRequest) GetApiMethodName() string {
	return "taobao.region.sale.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRegionSaleQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRegionSaleQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoRegionSaleQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRegionSaleQueryRequest()
	},
}

// GetTaobaoRegionSaleQueryRequest 从 sync.Pool 获取 TaobaoRegionSaleQueryAPIRequest
func GetTaobaoRegionSaleQueryAPIRequest() *TaobaoRegionSaleQueryAPIRequest {
	return poolTaobaoRegionSaleQueryAPIRequest.Get().(*TaobaoRegionSaleQueryAPIRequest)
}

// ReleaseTaobaoRegionSaleQueryAPIRequest 将 TaobaoRegionSaleQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoRegionSaleQueryAPIRequest(v *TaobaoRegionSaleQueryAPIRequest) {
	v.Reset()
	poolTaobaoRegionSaleQueryAPIRequest.Put(v)
}
