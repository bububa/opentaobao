package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreQueryAPIRequest
门店信息查询接口 API请求
taobao.qimen.store.query

商家在ERP等系统中调用该接口，查询门店相关信息 */
type TaobaoQimenStoreQueryAPIRequest struct {
	model.Params
	// 已分配的线上门店ID
	_storeId int64
}

// NewTaobaoQimenStoreQueryRequest 初始化TaobaoQimenStoreQueryAPIRequest对象
func NewTaobaoQimenStoreQueryRequest() *TaobaoQimenStoreQueryAPIRequest {
	return &TaobaoQimenStoreQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.store.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 已分配的线上门店ID
func (r *TaobaoQimenStoreQueryAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoQimenStoreQueryAPIRequest) GetStoreId() int64 {
	return r._storeId
}
