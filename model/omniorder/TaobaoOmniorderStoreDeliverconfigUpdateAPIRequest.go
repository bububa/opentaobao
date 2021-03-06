package omniorder

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.deliverconfig.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
