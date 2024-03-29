package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDrugQuantityBatchUpdateAPIRequest 批量同步库存接口 API请求
// taobao.drug.quantity.batch.update
//
// 商家通过top接口可以批量修改商品库存
type TaobaoDrugQuantityBatchUpdateAPIRequest struct {
	model.Params
	// 外部店铺ID
	_outStoreId string
	// 商品ID和库存
	_outItemIdQuantityMap string
}

// NewTaobaoDrugQuantityBatchUpdateRequest 初始化TaobaoDrugQuantityBatchUpdateAPIRequest对象
func NewTaobaoDrugQuantityBatchUpdateRequest() *TaobaoDrugQuantityBatchUpdateAPIRequest {
	return &TaobaoDrugQuantityBatchUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDrugQuantityBatchUpdateAPIRequest) Reset() {
	r._outStoreId = ""
	r._outItemIdQuantityMap = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDrugQuantityBatchUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.drug.quantity.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDrugQuantityBatchUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDrugQuantityBatchUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutStoreId is OutStoreId Setter
// 外部店铺ID
func (r *TaobaoDrugQuantityBatchUpdateAPIRequest) SetOutStoreId(_outStoreId string) error {
	r._outStoreId = _outStoreId
	r.Set("out_store_id", _outStoreId)
	return nil
}

// GetOutStoreId OutStoreId Getter
func (r TaobaoDrugQuantityBatchUpdateAPIRequest) GetOutStoreId() string {
	return r._outStoreId
}

// SetOutItemIdQuantityMap is OutItemIdQuantityMap Setter
// 商品ID和库存
func (r *TaobaoDrugQuantityBatchUpdateAPIRequest) SetOutItemIdQuantityMap(_outItemIdQuantityMap string) error {
	r._outItemIdQuantityMap = _outItemIdQuantityMap
	r.Set("out_item_id_quantity_map", _outItemIdQuantityMap)
	return nil
}

// GetOutItemIdQuantityMap OutItemIdQuantityMap Getter
func (r TaobaoDrugQuantityBatchUpdateAPIRequest) GetOutItemIdQuantityMap() string {
	return r._outItemIdQuantityMap
}

var poolTaobaoDrugQuantityBatchUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDrugQuantityBatchUpdateRequest()
	},
}

// GetTaobaoDrugQuantityBatchUpdateRequest 从 sync.Pool 获取 TaobaoDrugQuantityBatchUpdateAPIRequest
func GetTaobaoDrugQuantityBatchUpdateAPIRequest() *TaobaoDrugQuantityBatchUpdateAPIRequest {
	return poolTaobaoDrugQuantityBatchUpdateAPIRequest.Get().(*TaobaoDrugQuantityBatchUpdateAPIRequest)
}

// ReleaseTaobaoDrugQuantityBatchUpdateAPIRequest 将 TaobaoDrugQuantityBatchUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoDrugQuantityBatchUpdateAPIRequest(v *TaobaoDrugQuantityBatchUpdateAPIRequest) {
	v.Reset()
	poolTaobaoDrugQuantityBatchUpdateAPIRequest.Put(v)
}
