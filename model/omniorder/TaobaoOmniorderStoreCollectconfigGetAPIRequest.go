package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreCollectconfigGetAPIRequest 查询门店自提配置内容 API请求
// taobao.omniorder.store.collectconfig.get
//
// 查询门店自提配置内容
type TaobaoOmniorderStoreCollectconfigGetAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 是否是活动期
	_activity bool
}

// NewTaobaoOmniorderStoreCollectconfigGetRequest 初始化TaobaoOmniorderStoreCollectconfigGetAPIRequest对象
func NewTaobaoOmniorderStoreCollectconfigGetRequest() *TaobaoOmniorderStoreCollectconfigGetAPIRequest {
	return &TaobaoOmniorderStoreCollectconfigGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreCollectconfigGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.collectconfig.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreCollectconfigGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreCollectconfigGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoOmniorderStoreCollectconfigGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is Activity Setter
// 是否是活动期
func (r *TaobaoOmniorderStoreCollectconfigGetAPIRequest) SetActivity(_activity bool) error {
	r._activity = _activity
	r.Set("activity", _activity)
	return nil
}

// Get Activity Getter
func (r TaobaoOmniorderStoreCollectconfigGetAPIRequest) GetActivity() bool {
	return r._activity
}
