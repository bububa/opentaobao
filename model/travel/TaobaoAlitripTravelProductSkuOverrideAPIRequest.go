package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelProductSkuOverrideAPIRequest （供销）产品级别日历价格库存修改，全量覆盖 API请求
// taobao.alitrip.travel.product.sku.override
//
// （供销）产品级别日历价格库存修改，全量覆盖
type TaobaoAlitripTravelProductSkuOverrideAPIRequest struct {
	model.Params
	// 商品日历价格库存套餐
	_skus []ItemSkuInfo
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
}

// NewTaobaoAlitripTravelProductSkuOverrideRequest 初始化TaobaoAlitripTravelProductSkuOverrideAPIRequest对象
func NewTaobaoAlitripTravelProductSkuOverrideRequest() *TaobaoAlitripTravelProductSkuOverrideAPIRequest {
	return &TaobaoAlitripTravelProductSkuOverrideAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelProductSkuOverrideAPIRequest) Reset() {
	r._skus = r._skus[:0]
	r._outProductId = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelProductSkuOverrideAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.product.sku.override"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelProductSkuOverrideAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelProductSkuOverrideAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkus is Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoAlitripTravelProductSkuOverrideAPIRequest) SetSkus(_skus []ItemSkuInfo) error {
	r._skus = _skus
	r.Set("skus", _skus)
	return nil
}

// GetSkus Skus Getter
func (r TaobaoAlitripTravelProductSkuOverrideAPIRequest) GetSkus() []ItemSkuInfo {
	return r._skus
}

// SetOutProductId is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelProductSkuOverrideAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r TaobaoAlitripTravelProductSkuOverrideAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelProductSkuOverrideAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoAlitripTravelProductSkuOverrideAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoAlitripTravelProductSkuOverrideAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelProductSkuOverrideRequest()
	},
}

// GetTaobaoAlitripTravelProductSkuOverrideRequest 从 sync.Pool 获取 TaobaoAlitripTravelProductSkuOverrideAPIRequest
func GetTaobaoAlitripTravelProductSkuOverrideAPIRequest() *TaobaoAlitripTravelProductSkuOverrideAPIRequest {
	return poolTaobaoAlitripTravelProductSkuOverrideAPIRequest.Get().(*TaobaoAlitripTravelProductSkuOverrideAPIRequest)
}

// ReleaseTaobaoAlitripTravelProductSkuOverrideAPIRequest 将 TaobaoAlitripTravelProductSkuOverrideAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelProductSkuOverrideAPIRequest(v *TaobaoAlitripTravelProductSkuOverrideAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelProductSkuOverrideAPIRequest.Put(v)
}
