package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店订单创建api API请求
tmall.store.order.create

门店订单创建api
*/
type TmallStoreOrderCreateRequest struct {
    model.Params
    // 系统自动生成
    _appInfo   *AppInfo
    // 创建订单请求
    _createOrderRequest   *CreateOrderRequest
}

// 初始化TmallStoreOrderCreateRequest对象
func NewTmallStoreOrderCreateRequest() *TmallStoreOrderCreateRequest{
    return &TmallStoreOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallStoreOrderCreateRequest) GetApiMethodName() string {
    return "tmall.store.order.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallStoreOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppInfo Setter
// 系统自动生成
func (r *TmallStoreOrderCreateRequest) SetAppInfo(_appInfo *AppInfo) error {
    r._appInfo = _appInfo
    r.Set("app_info", _appInfo)
    return nil
}

// AppInfo Getter
func (r TmallStoreOrderCreateRequest) GetAppInfo() *AppInfo {
    return r._appInfo
}
// CreateOrderRequest Setter
// 创建订单请求
func (r *TmallStoreOrderCreateRequest) SetCreateOrderRequest(_createOrderRequest *CreateOrderRequest) error {
    r._createOrderRequest = _createOrderRequest
    r.Set("create_order_request", _createOrderRequest)
    return nil
}

// CreateOrderRequest Getter
func (r TmallStoreOrderCreateRequest) GetCreateOrderRequest() *CreateOrderRequest {
    return r._createOrderRequest
}