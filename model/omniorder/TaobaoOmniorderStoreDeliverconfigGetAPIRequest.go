package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstoredeliverconfiggetAPIRequest 查询门店发货配置内容 API请求
// taobao.omniorder.store.deliverconfig.get
//
// 查询门店发货配置内容
type TaobaoomniorderstoredeliverconfiggetAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 是否是活动期
	_activity bool
}

// NewTaobaoomniorderstoredeliverconfiggetRequest 初始化TaobaoomniorderstoredeliverconfiggetAPIRequest对象
func NewTaobaoomniorderstoredeliverconfiggetRequest() *TaobaoomniorderstoredeliverconfiggetAPIRequest {
	return &TaobaoomniorderstoredeliverconfiggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderstoredeliverconfiggetAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.deliverconfig.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderstoredeliverconfiggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderstoredeliverconfiggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoomniorderstoredeliverconfiggetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoomniorderstoredeliverconfiggetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetActivity is Activity Setter
// 是否是活动期
func (r *TaobaoomniorderstoredeliverconfiggetAPIRequest) SetActivity(_activity bool) error {
	r._activity = _activity
	r.Set("activity", _activity)
	return nil
}

// GetActivity Activity Getter
func (r TaobaoomniorderstoredeliverconfiggetAPIRequest) GetActivity() bool {
	return r._activity
}
