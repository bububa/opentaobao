package consignplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟发货工作台创建订单 API请求
cainiao.consignplatform.order.create

菜鸟发货工作台，商家或者isv通过api进行订单写入操作
*/
type CainiaoConsignplatformOrderCreateRequest struct {
    model.Params
    // 订单创建入参
    createRequest   *OrderCreateRequest
}

// 初始化CainiaoConsignplatformOrderCreateRequest对象
func NewCainiaoConsignplatformOrderCreateRequest() *CainiaoConsignplatformOrderCreateRequest{
    return &CainiaoConsignplatformOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoConsignplatformOrderCreateRequest) GetApiMethodName() string {
    return "cainiao.consignplatform.order.create"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoConsignplatformOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateRequest Setter
// 订单创建入参
func (r *CainiaoConsignplatformOrderCreateRequest) SetCreateRequest(createRequest *OrderCreateRequest) error {
    r.createRequest = createRequest
    r.Set("create_request", createRequest)
    return nil
}

// CreateRequest Getter
func (r CainiaoConsignplatformOrderCreateRequest) GetCreateRequest() *OrderCreateRequest {
    return r.createRequest
}
