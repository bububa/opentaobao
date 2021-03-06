package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest switchstatus.update API请求
// taobao.omniorder.store.switchstatus.update
//
// 变更门店发货、门店自提状态
type TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest struct {
	model.Params
	// 门店发货自提状态
	_status string
	// 门店ID
	_storeId int64
}

// NewTaobaoOmniorderStoreSwitchstatusUpdateRequest 初始化TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest对象
func NewTaobaoOmniorderStoreSwitchstatusUpdateRequest() *TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest {
	return &TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.switchstatus.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStatus is Status Setter
// 门店发货自提状态
func (r *TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest) GetStatus() string {
	return r._status
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest) GetStoreId() int64 {
	return r._storeId
}
