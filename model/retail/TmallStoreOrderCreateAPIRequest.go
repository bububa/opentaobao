package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallstoreordercreateAPIRequest 门店订单创建api API请求
// tmall.store.order.create
//
// 门店订单创建api
type TmallstoreordercreateAPIRequest struct {
	model.Params
	// 创建订单请求
	_createOrderRequest *CreateOrderRequest
	// 系统自动生成
	_appInfo *AppInfo
}

// NewTmallstoreordercreateRequest 初始化TmallstoreordercreateAPIRequest对象
func NewTmallstoreordercreateRequest() *TmallstoreordercreateAPIRequest {
	return &TmallstoreordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallstoreordercreateAPIRequest) GetApiMethodName() string {
	return "tmall.store.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallstoreordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallstoreordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateOrderRequest is CreateOrderRequest Setter
// 创建订单请求
func (r *TmallstoreordercreateAPIRequest) SetCreateOrderRequest(_createOrderRequest *CreateOrderRequest) error {
	r._createOrderRequest = _createOrderRequest
	r.Set("create_order_request", _createOrderRequest)
	return nil
}

// GetCreateOrderRequest CreateOrderRequest Getter
func (r TmallstoreordercreateAPIRequest) GetCreateOrderRequest() *CreateOrderRequest {
	return r._createOrderRequest
}

// SetAppInfo is AppInfo Setter
// 系统自动生成
func (r *TmallstoreordercreateAPIRequest) SetAppInfo(_appInfo *AppInfo) error {
	r._appInfo = _appInfo
	r.Set("app_info", _appInfo)
	return nil
}

// GetAppInfo AppInfo Getter
func (r TmallstoreordercreateAPIRequest) GetAppInfo() *AppInfo {
	return r._appInfo
}
