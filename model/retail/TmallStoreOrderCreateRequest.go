package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店订单创建api APIRequest
tmall.store.order.create

门店订单创建api
*/
type TmallStoreOrderCreateRequest struct {
    model.Params

    // 系统自动生成
    appInfo   *AppInfo 

    // 创建订单请求
    createOrderRequest   *CreateOrderRequest 

}

func NewTmallStoreOrderCreateRequest() *TmallStoreOrderCreateRequest{
    return &TmallStoreOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallStoreOrderCreateRequest) GetApiMethodName() string {
    return "tmall.store.order.create"
}

func (r TmallStoreOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallStoreOrderCreateRequest) SetAppInfo(appInfo *AppInfo) error {
    r.appInfo = appInfo
    r.Set("app_info", appInfo)
    return nil
}

func (r TmallStoreOrderCreateRequest) GetAppInfo() *AppInfo {
    return r.appInfo
}

func (r *TmallStoreOrderCreateRequest) SetCreateOrderRequest(createOrderRequest *CreateOrderRequest) error {
    r.createOrderRequest = createOrderRequest
    r.Set("create_order_request", createOrderRequest)
    return nil
}

func (r TmallStoreOrderCreateRequest) GetCreateOrderRequest() *CreateOrderRequest {
    return r.createOrderRequest
}

