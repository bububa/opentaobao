package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDrugPriceUpdateAPIRequest 商家更新宝贝价格 API请求
// taobao.drug.price.update
//
// 商家更新价格
type TaobaoDrugPriceUpdateAPIRequest struct {
	model.Params
	// 对应的外部店铺ID
	_outStoreId string
	// 对应的外部商品编码
	_outItemId string
	// 商品价格
	_price float64
}

// NewTaobaoDrugPriceUpdateRequest 初始化TaobaoDrugPriceUpdateAPIRequest对象
func NewTaobaoDrugPriceUpdateRequest() *TaobaoDrugPriceUpdateAPIRequest {
	return &TaobaoDrugPriceUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDrugPriceUpdateAPIRequest) Reset() {
	r._outStoreId = ""
	r._outItemId = ""
	r._price = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDrugPriceUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.drug.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDrugPriceUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDrugPriceUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutStoreId is OutStoreId Setter
// 对应的外部店铺ID
func (r *TaobaoDrugPriceUpdateAPIRequest) SetOutStoreId(_outStoreId string) error {
	r._outStoreId = _outStoreId
	r.Set("out_store_id", _outStoreId)
	return nil
}

// GetOutStoreId OutStoreId Getter
func (r TaobaoDrugPriceUpdateAPIRequest) GetOutStoreId() string {
	return r._outStoreId
}

// SetOutItemId is OutItemId Setter
// 对应的外部商品编码
func (r *TaobaoDrugPriceUpdateAPIRequest) SetOutItemId(_outItemId string) error {
	r._outItemId = _outItemId
	r.Set("out_item_id", _outItemId)
	return nil
}

// GetOutItemId OutItemId Getter
func (r TaobaoDrugPriceUpdateAPIRequest) GetOutItemId() string {
	return r._outItemId
}

// SetPrice is Price Setter
// 商品价格
func (r *TaobaoDrugPriceUpdateAPIRequest) SetPrice(_price float64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TaobaoDrugPriceUpdateAPIRequest) GetPrice() float64 {
	return r._price
}

var poolTaobaoDrugPriceUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDrugPriceUpdateRequest()
	},
}

// GetTaobaoDrugPriceUpdateRequest 从 sync.Pool 获取 TaobaoDrugPriceUpdateAPIRequest
func GetTaobaoDrugPriceUpdateAPIRequest() *TaobaoDrugPriceUpdateAPIRequest {
	return poolTaobaoDrugPriceUpdateAPIRequest.Get().(*TaobaoDrugPriceUpdateAPIRequest)
}

// ReleaseTaobaoDrugPriceUpdateAPIRequest 将 TaobaoDrugPriceUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoDrugPriceUpdateAPIRequest(v *TaobaoDrugPriceUpdateAPIRequest) {
	v.Reset()
	poolTaobaoDrugPriceUpdateAPIRequest.Put(v)
}
