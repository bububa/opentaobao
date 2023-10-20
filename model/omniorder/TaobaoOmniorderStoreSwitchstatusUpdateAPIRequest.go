package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstoreswitchstatusupdateAPIRequest switchstatus.update API请求
// taobao.omniorder.store.switchstatus.update
//
// 变更门店发货、门店自提状态
type TaobaoomniorderstoreswitchstatusupdateAPIRequest struct {
	model.Params
	// 门店发货自提状态
	_status string
	// 门店ID
	_storeId int64
}

// NewTaobaoomniorderstoreswitchstatusupdateRequest 初始化TaobaoomniorderstoreswitchstatusupdateAPIRequest对象
func NewTaobaoomniorderstoreswitchstatusupdateRequest() *TaobaoomniorderstoreswitchstatusupdateAPIRequest {
	return &TaobaoomniorderstoreswitchstatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstoreswitchstatusupdateAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.switchstatus.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstoreswitchstatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstoreswitchstatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 门店发货自提状态
func (r *TaobaoomniorderstoreswitchstatusupdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoomniorderstoreswitchstatusupdateAPIRequest) GetStatus() string {
	return r._status
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoomniorderstoreswitchstatusupdateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoomniorderstoreswitchstatusupdateAPIRequest) GetStoreId() int64 {
	return r._storeId
}
