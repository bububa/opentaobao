package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstorecollectconfigupdateAPIRequest 门店自提配置修改 API请求
// taobao.omniorder.store.collectconfig.update
//
// 修改门店自提配置内容
type TaobaoomniorderstorecollectconfigupdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 门店自提配置
	_storeCollectConfig *StoreCollectConfig
}

// NewTaobaoomniorderstorecollectconfigupdateRequest 初始化TaobaoomniorderstorecollectconfigupdateAPIRequest对象
func NewTaobaoomniorderstorecollectconfigupdateRequest() *TaobaoomniorderstorecollectconfigupdateAPIRequest {
	return &TaobaoomniorderstorecollectconfigupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstorecollectconfigupdateAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.collectconfig.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstorecollectconfigupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstorecollectconfigupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoomniorderstorecollectconfigupdateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoomniorderstorecollectconfigupdateAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetStoreCollectConfig is StoreCollectConfig Setter
// 门店自提配置
func (r *TaobaoomniorderstorecollectconfigupdateAPIRequest) SetStoreCollectConfig(_storeCollectConfig *StoreCollectConfig) error {
	r._storeCollectConfig = _storeCollectConfig
	r.Set("store_collect_config", _storeCollectConfig)
	return nil
}

// GetStoreCollectConfig StoreCollectConfig Getter
func (r TaobaoomniorderstorecollectconfigupdateAPIRequest) GetStoreCollectConfig() *StoreCollectConfig {
	return r._storeCollectConfig
}
