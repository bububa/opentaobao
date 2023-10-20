package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstorecollectconfiggetAPIRequest 查询门店自提配置内容 API请求
// taobao.omniorder.store.collectconfig.get
//
// 查询门店自提配置内容
type TaobaoomniorderstorecollectconfiggetAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 是否是活动期
	_activity bool
}

// NewTaobaoomniorderstorecollectconfiggetRequest 初始化TaobaoomniorderstorecollectconfiggetAPIRequest对象
func NewTaobaoomniorderstorecollectconfiggetRequest() *TaobaoomniorderstorecollectconfiggetAPIRequest {
	return &TaobaoomniorderstorecollectconfiggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstorecollectconfiggetAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.collectconfig.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstorecollectconfiggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstorecollectconfiggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoomniorderstorecollectconfiggetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoomniorderstorecollectconfiggetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetActivity is Activity Setter
// 是否是活动期
func (r *TaobaoomniorderstorecollectconfiggetAPIRequest) SetActivity(_activity bool) error {
	r._activity = _activity
	r.Set("activity", _activity)
	return nil
}

// GetActivity Activity Getter
func (r TaobaoomniorderstorecollectconfiggetAPIRequest) GetActivity() bool {
	return r._activity
}
