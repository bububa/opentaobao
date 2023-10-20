package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemSkuOverrideAPIRequest 【API3.0】商品级别日历价格库存修改，全量覆盖 API请求
// taobao.alitrip.travel.item.sku.override
//
// 旅行度假新商品日历价格库存信息修改接口 第三版。提供商家通过TOP API方式修改商品sku信息。
type TaobaoAlitripTravelItemSkuOverrideAPIRequest struct {
	model.Params
	// 商品日历价格库存套餐
	_skus []PontusTravelItemSkuInfo
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
}

// NewTaobaoAlitripTravelItemSkuOverrideRequest 初始化TaobaoAlitripTravelItemSkuOverrideAPIRequest对象
func NewTaobaoAlitripTravelItemSkuOverrideRequest() *TaobaoAlitripTravelItemSkuOverrideAPIRequest {
	return &TaobaoAlitripTravelItemSkuOverrideAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelItemSkuOverrideAPIRequest) Reset() {
	r._skus = r._skus[:0]
	r._outProductId = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemSkuOverrideAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.sku.override"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemSkuOverrideAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelItemSkuOverrideAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkus is Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoAlitripTravelItemSkuOverrideAPIRequest) SetSkus(_skus []PontusTravelItemSkuInfo) error {
	r._skus = _skus
	r.Set("skus", _skus)
	return nil
}

// GetSkus Skus Getter
func (r TaobaoAlitripTravelItemSkuOverrideAPIRequest) GetSkus() []PontusTravelItemSkuInfo {
	return r._skus
}

// SetOutProductId is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuOverrideAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r TaobaoAlitripTravelItemSkuOverrideAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuOverrideAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoAlitripTravelItemSkuOverrideAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoAlitripTravelItemSkuOverrideAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelItemSkuOverrideRequest()
	},
}

// GetTaobaoAlitripTravelItemSkuOverrideRequest 从 sync.Pool 获取 TaobaoAlitripTravelItemSkuOverrideAPIRequest
func GetTaobaoAlitripTravelItemSkuOverrideAPIRequest() *TaobaoAlitripTravelItemSkuOverrideAPIRequest {
	return poolTaobaoAlitripTravelItemSkuOverrideAPIRequest.Get().(*TaobaoAlitripTravelItemSkuOverrideAPIRequest)
}

// ReleaseTaobaoAlitripTravelItemSkuOverrideAPIRequest 将 TaobaoAlitripTravelItemSkuOverrideAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelItemSkuOverrideAPIRequest(v *TaobaoAlitripTravelItemSkuOverrideAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelItemSkuOverrideAPIRequest.Put(v)
}
