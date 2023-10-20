package openmall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallItemSaleareaGetAPIRequest 查询商品可售区域 API请求
// taobao.openmall.item.salearea.get
//
// 获取openmall商品的可售区域
type TaobaoOpenmallItemSaleareaGetAPIRequest struct {
	model.Params
	// 商品SKU
	_skuIds string
	// 商品ID
	_itemId int64
}

// NewTaobaoOpenmallItemSaleareaGetRequest 初始化TaobaoOpenmallItemSaleareaGetAPIRequest对象
func NewTaobaoOpenmallItemSaleareaGetRequest() *TaobaoOpenmallItemSaleareaGetAPIRequest {
	return &TaobaoOpenmallItemSaleareaGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallItemSaleareaGetAPIRequest) Reset() {
	r._skuIds = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallItemSaleareaGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.item.salearea.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallItemSaleareaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallItemSaleareaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuIds is SkuIds Setter
// 商品SKU
func (r *TaobaoOpenmallItemSaleareaGetAPIRequest) SetSkuIds(_skuIds string) error {
	r._skuIds = _skuIds
	r.Set("sku_ids", _skuIds)
	return nil
}

// GetSkuIds SkuIds Getter
func (r TaobaoOpenmallItemSaleareaGetAPIRequest) GetSkuIds() string {
	return r._skuIds
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoOpenmallItemSaleareaGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoOpenmallItemSaleareaGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoOpenmallItemSaleareaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallItemSaleareaGetRequest()
	},
}

// GetTaobaoOpenmallItemSaleareaGetRequest 从 sync.Pool 获取 TaobaoOpenmallItemSaleareaGetAPIRequest
func GetTaobaoOpenmallItemSaleareaGetAPIRequest() *TaobaoOpenmallItemSaleareaGetAPIRequest {
	return poolTaobaoOpenmallItemSaleareaGetAPIRequest.Get().(*TaobaoOpenmallItemSaleareaGetAPIRequest)
}

// ReleaseTaobaoOpenmallItemSaleareaGetAPIRequest 将 TaobaoOpenmallItemSaleareaGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallItemSaleareaGetAPIRequest(v *TaobaoOpenmallItemSaleareaGetAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallItemSaleareaGetAPIRequest.Put(v)
}
