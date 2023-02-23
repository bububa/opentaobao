package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreDeliverconfigGetAPIRequest 查询门店发货配置内容 API请求
// taobao.omniorder.store.deliverconfig.get
//
// 查询门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigGetAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 是否是活动期
	_activity bool
}

// NewTaobaoOmniorderStoreDeliverconfigGetRequest 初始化TaobaoOmniorderStoreDeliverconfigGetAPIRequest对象
func NewTaobaoOmniorderStoreDeliverconfigGetRequest() *TaobaoOmniorderStoreDeliverconfigGetAPIRequest {
	return &TaobaoOmniorderStoreDeliverconfigGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreDeliverconfigGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.deliverconfig.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreDeliverconfigGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderStoreDeliverconfigGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreDeliverconfigGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoOmniorderStoreDeliverconfigGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetActivity is Activity Setter
// 是否是活动期
func (r *TaobaoOmniorderStoreDeliverconfigGetAPIRequest) SetActivity(_activity bool) error {
	r._activity = _activity
	r.Set("activity", _activity)
	return nil
}

// GetActivity Activity Getter
func (r TaobaoOmniorderStoreDeliverconfigGetAPIRequest) GetActivity() bool {
	return r._activity
}
