package omniorder

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderStoreCollectconfigGetAPIRequest) Reset() {
	r._storeId = 0
	r._activity = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreCollectconfigGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.collectconfig.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreCollectconfigGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderStoreCollectconfigGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreCollectconfigGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoOmniorderStoreCollectconfigGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetActivity is Activity Setter
// 是否是活动期
func (r *TaobaoOmniorderStoreCollectconfigGetAPIRequest) SetActivity(_activity bool) error {
	r._activity = _activity
	r.Set("activity", _activity)
	return nil
}

// GetActivity Activity Getter
func (r TaobaoOmniorderStoreCollectconfigGetAPIRequest) GetActivity() bool {
	return r._activity
}

var poolTaobaoOmniorderStoreCollectconfigGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderStoreCollectconfigGetRequest()
	},
}

// GetTaobaoOmniorderStoreCollectconfigGetRequest 从 sync.Pool 获取 TaobaoOmniorderStoreCollectconfigGetAPIRequest
func GetTaobaoOmniorderStoreCollectconfigGetAPIRequest() *TaobaoOmniorderStoreCollectconfigGetAPIRequest {
	return poolTaobaoOmniorderStoreCollectconfigGetAPIRequest.Get().(*TaobaoOmniorderStoreCollectconfigGetAPIRequest)
}

// ReleaseTaobaoOmniorderStoreCollectconfigGetAPIRequest 将 TaobaoOmniorderStoreCollectconfigGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderStoreCollectconfigGetAPIRequest(v *TaobaoOmniorderStoreCollectconfigGetAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderStoreCollectconfigGetAPIRequest.Put(v)
}
