package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstoredeliverconfigupdateAPIRequest 修改门店发货配置内容 API请求
// taobao.omniorder.store.deliverconfig.update
//
// 修改门店发货配置内容
type TaobaoomniorderstoredeliverconfigupdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 卖家发货配置
	_storeDeliverConfig *StoreDeliverConfig
}

// NewTaobaoomniorderstoredeliverconfigupdateRequest 初始化TaobaoomniorderstoredeliverconfigupdateAPIRequest对象
func NewTaobaoomniorderstoredeliverconfigupdateRequest() *TaobaoomniorderstoredeliverconfigupdateAPIRequest {
	return &TaobaoomniorderstoredeliverconfigupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstoredeliverconfigupdateAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.deliverconfig.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstoredeliverconfigupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstoredeliverconfigupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoomniorderstoredeliverconfigupdateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoomniorderstoredeliverconfigupdateAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetStoreDeliverConfig is StoreDeliverConfig Setter
// 卖家发货配置
func (r *TaobaoomniorderstoredeliverconfigupdateAPIRequest) SetStoreDeliverConfig(_storeDeliverConfig *StoreDeliverConfig) error {
	r._storeDeliverConfig = _storeDeliverConfig
	r.Set("store_deliver_config", _storeDeliverConfig)
	return nil
}

// GetStoreDeliverConfig StoreDeliverConfig Getter
func (r TaobaoomniorderstoredeliverconfigupdateAPIRequest) GetStoreDeliverConfig() *StoreDeliverConfig {
	return r._storeDeliverConfig
}
