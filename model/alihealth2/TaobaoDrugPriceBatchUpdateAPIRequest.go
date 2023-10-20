package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDrugPriceBatchUpdateAPIRequest 商家批量更新宝贝价格 API请求
// taobao.drug.price.batch.update
//
// 商家批量更新宝贝价格
type TaobaoDrugPriceBatchUpdateAPIRequest struct {
	model.Params
	// 外部店铺ID
	_outStoreId string
	// 商品ID和价格
	_outItemIdPriceMap string
}

// NewTaobaoDrugPriceBatchUpdateRequest 初始化TaobaoDrugPriceBatchUpdateAPIRequest对象
func NewTaobaoDrugPriceBatchUpdateRequest() *TaobaoDrugPriceBatchUpdateAPIRequest {
	return &TaobaoDrugPriceBatchUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDrugPriceBatchUpdateAPIRequest) Reset() {
	r._outStoreId = ""
	r._outItemIdPriceMap = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDrugPriceBatchUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.drug.price.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDrugPriceBatchUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDrugPriceBatchUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutStoreId is OutStoreId Setter
// 外部店铺ID
func (r *TaobaoDrugPriceBatchUpdateAPIRequest) SetOutStoreId(_outStoreId string) error {
	r._outStoreId = _outStoreId
	r.Set("out_store_id", _outStoreId)
	return nil
}

// GetOutStoreId OutStoreId Getter
func (r TaobaoDrugPriceBatchUpdateAPIRequest) GetOutStoreId() string {
	return r._outStoreId
}

// SetOutItemIdPriceMap is OutItemIdPriceMap Setter
// 商品ID和价格
func (r *TaobaoDrugPriceBatchUpdateAPIRequest) SetOutItemIdPriceMap(_outItemIdPriceMap string) error {
	r._outItemIdPriceMap = _outItemIdPriceMap
	r.Set("out_item_id_price_map", _outItemIdPriceMap)
	return nil
}

// GetOutItemIdPriceMap OutItemIdPriceMap Getter
func (r TaobaoDrugPriceBatchUpdateAPIRequest) GetOutItemIdPriceMap() string {
	return r._outItemIdPriceMap
}

var poolTaobaoDrugPriceBatchUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDrugPriceBatchUpdateRequest()
	},
}

// GetTaobaoDrugPriceBatchUpdateRequest 从 sync.Pool 获取 TaobaoDrugPriceBatchUpdateAPIRequest
func GetTaobaoDrugPriceBatchUpdateAPIRequest() *TaobaoDrugPriceBatchUpdateAPIRequest {
	return poolTaobaoDrugPriceBatchUpdateAPIRequest.Get().(*TaobaoDrugPriceBatchUpdateAPIRequest)
}

// ReleaseTaobaoDrugPriceBatchUpdateAPIRequest 将 TaobaoDrugPriceBatchUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoDrugPriceBatchUpdateAPIRequest(v *TaobaoDrugPriceBatchUpdateAPIRequest) {
	v.Reset()
	poolTaobaoDrugPriceBatchUpdateAPIRequest.Put(v)
}
