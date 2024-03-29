package retail

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallStoreOrderCreateAPIRequest 门店订单创建api API请求
// tmall.store.order.create
//
// 门店订单创建api
type TmallStoreOrderCreateAPIRequest struct {
	model.Params
	// 创建订单请求
	_createOrderRequest *CreateOrderRequest
	// 系统自动生成
	_appInfo *AppInfo
}

// NewTmallStoreOrderCreateRequest 初始化TmallStoreOrderCreateAPIRequest对象
func NewTmallStoreOrderCreateRequest() *TmallStoreOrderCreateAPIRequest {
	return &TmallStoreOrderCreateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallStoreOrderCreateAPIRequest) Reset() {
	r._createOrderRequest = nil
	r._appInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallStoreOrderCreateAPIRequest) GetApiMethodName() string {
	return "tmall.store.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallStoreOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallStoreOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateOrderRequest is CreateOrderRequest Setter
// 创建订单请求
func (r *TmallStoreOrderCreateAPIRequest) SetCreateOrderRequest(_createOrderRequest *CreateOrderRequest) error {
	r._createOrderRequest = _createOrderRequest
	r.Set("create_order_request", _createOrderRequest)
	return nil
}

// GetCreateOrderRequest CreateOrderRequest Getter
func (r TmallStoreOrderCreateAPIRequest) GetCreateOrderRequest() *CreateOrderRequest {
	return r._createOrderRequest
}

// SetAppInfo is AppInfo Setter
// 系统自动生成
func (r *TmallStoreOrderCreateAPIRequest) SetAppInfo(_appInfo *AppInfo) error {
	r._appInfo = _appInfo
	r.Set("app_info", _appInfo)
	return nil
}

// GetAppInfo AppInfo Getter
func (r TmallStoreOrderCreateAPIRequest) GetAppInfo() *AppInfo {
	return r._appInfo
}

var poolTmallStoreOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallStoreOrderCreateRequest()
	},
}

// GetTmallStoreOrderCreateRequest 从 sync.Pool 获取 TmallStoreOrderCreateAPIRequest
func GetTmallStoreOrderCreateAPIRequest() *TmallStoreOrderCreateAPIRequest {
	return poolTmallStoreOrderCreateAPIRequest.Get().(*TmallStoreOrderCreateAPIRequest)
}

// ReleaseTmallStoreOrderCreateAPIRequest 将 TmallStoreOrderCreateAPIRequest 放入 sync.Pool
func ReleaseTmallStoreOrderCreateAPIRequest(v *TmallStoreOrderCreateAPIRequest) {
	v.Reset()
	poolTmallStoreOrderCreateAPIRequest.Put(v)
}
