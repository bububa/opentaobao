package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreDeleteAPIRequest
门店删除接口 API请求
taobao.qimen.store.delete

商家在ERP等系统中调用该接口，删除线下门店 */
type TaobaoQimenStoreDeleteAPIRequest struct {
	model.Params
	// 要删除的门店id
	_storeId int64
}

// NewTaobaoQimenStoreDeleteRequest 初始化TaobaoQimenStoreDeleteAPIRequest对象
func NewTaobaoQimenStoreDeleteRequest() *TaobaoQimenStoreDeleteAPIRequest {
	return &TaobaoQimenStoreDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.store.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 要删除的门店id
func (r *TaobaoQimenStoreDeleteAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoQimenStoreDeleteAPIRequest) GetStoreId() int64 {
	return r._storeId
}
