package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallStoreOrderCreateAPIRequest 门店订单创建api API请求
// tmall.store.order.create
//
// 门店订单创建api
type TmallStoreOrderCreateAPIRequest struct {
	model.Params
	// 系统自动生成
	_appInfo *AppInfo
	// 创建订单请求
	_createOrderRequest *CreateOrderRequest
}

// NewTmallStoreOrderCreateRequest 初始化TmallStoreOrderCreateAPIRequest对象
func NewTmallStoreOrderCreateRequest() *TmallStoreOrderCreateAPIRequest {
	return &TmallStoreOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallStoreOrderCreateAPIRequest) GetApiMethodName() string {
	return "tmall.store.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallStoreOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
