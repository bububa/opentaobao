package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest 修改门店发货配置内容 API请求
// taobao.omniorder.store.deliverconfig.update
//
// 修改门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 卖家发货配置
	_storeDeliverConfig *StoreDeliverConfig
}

// NewTaobaoOmniorderStoreDeliverconfigUpdateRequest 初始化TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest对象
func NewTaobaoOmniorderStoreDeliverconfigUpdateRequest() *TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest {
	return &TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) Reset() {
	r._storeId = 0
	r._storeDeliverConfig = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.deliverconfig.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetStoreDeliverConfig is StoreDeliverConfig Setter
// 卖家发货配置
func (r *TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) SetStoreDeliverConfig(_storeDeliverConfig *StoreDeliverConfig) error {
	r._storeDeliverConfig = _storeDeliverConfig
	r.Set("store_deliver_config", _storeDeliverConfig)
	return nil
}

// GetStoreDeliverConfig StoreDeliverConfig Getter
func (r TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) GetStoreDeliverConfig() *StoreDeliverConfig {
	return r._storeDeliverConfig
}

var poolTaobaoOmniorderStoreDeliverconfigUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderStoreDeliverconfigUpdateRequest()
	},
}

// GetTaobaoOmniorderStoreDeliverconfigUpdateRequest 从 sync.Pool 获取 TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest
func GetTaobaoOmniorderStoreDeliverconfigUpdateAPIRequest() *TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest {
	return poolTaobaoOmniorderStoreDeliverconfigUpdateAPIRequest.Get().(*TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest)
}

// ReleaseTaobaoOmniorderStoreDeliverconfigUpdateAPIRequest 将 TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderStoreDeliverconfigUpdateAPIRequest(v *TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderStoreDeliverconfigUpdateAPIRequest.Put(v)
}
