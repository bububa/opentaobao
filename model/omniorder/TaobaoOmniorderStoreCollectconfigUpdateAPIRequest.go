package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreCollectconfigUpdateAPIRequest 门店自提配置修改 API请求
// taobao.omniorder.store.collectconfig.update
//
// 修改门店自提配置内容
type TaobaoOmniorderStoreCollectconfigUpdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 门店自提配置
	_storeCollectConfig *StoreCollectConfig
}

// NewTaobaoOmniorderStoreCollectconfigUpdateRequest 初始化TaobaoOmniorderStoreCollectconfigUpdateAPIRequest对象
func NewTaobaoOmniorderStoreCollectconfigUpdateRequest() *TaobaoOmniorderStoreCollectconfigUpdateAPIRequest {
	return &TaobaoOmniorderStoreCollectconfigUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderStoreCollectconfigUpdateAPIRequest) Reset() {
	r._storeId = 0
	r._storeCollectConfig = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreCollectconfigUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.collectconfig.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreCollectconfigUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderStoreCollectconfigUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreCollectconfigUpdateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoOmniorderStoreCollectconfigUpdateAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetStoreCollectConfig is StoreCollectConfig Setter
// 门店自提配置
func (r *TaobaoOmniorderStoreCollectconfigUpdateAPIRequest) SetStoreCollectConfig(_storeCollectConfig *StoreCollectConfig) error {
	r._storeCollectConfig = _storeCollectConfig
	r.Set("store_collect_config", _storeCollectConfig)
	return nil
}

// GetStoreCollectConfig StoreCollectConfig Getter
func (r TaobaoOmniorderStoreCollectconfigUpdateAPIRequest) GetStoreCollectConfig() *StoreCollectConfig {
	return r._storeCollectConfig
}

var poolTaobaoOmniorderStoreCollectconfigUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderStoreCollectconfigUpdateRequest()
	},
}

// GetTaobaoOmniorderStoreCollectconfigUpdateRequest 从 sync.Pool 获取 TaobaoOmniorderStoreCollectconfigUpdateAPIRequest
func GetTaobaoOmniorderStoreCollectconfigUpdateAPIRequest() *TaobaoOmniorderStoreCollectconfigUpdateAPIRequest {
	return poolTaobaoOmniorderStoreCollectconfigUpdateAPIRequest.Get().(*TaobaoOmniorderStoreCollectconfigUpdateAPIRequest)
}

// ReleaseTaobaoOmniorderStoreCollectconfigUpdateAPIRequest 将 TaobaoOmniorderStoreCollectconfigUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderStoreCollectconfigUpdateAPIRequest(v *TaobaoOmniorderStoreCollectconfigUpdateAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderStoreCollectconfigUpdateAPIRequest.Put(v)
}
