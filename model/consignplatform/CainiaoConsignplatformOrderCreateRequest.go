package consignplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟发货工作台创建订单 APIRequest
cainiao.consignplatform.order.create

菜鸟发货工作台，商家或者isv通过api进行订单写入操作
*/
type CainiaoConsignplatformOrderCreateRequest struct {
    model.Params

    // 订单创建入参
    createRequest   *OrderCreateRequest 

}

func NewCainiaoConsignplatformOrderCreateRequest() *CainiaoConsignplatformOrderCreateRequest{
    return &CainiaoConsignplatformOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoConsignplatformOrderCreateRequest) GetApiMethodName() string {
    return "cainiao.consignplatform.order.create"
}

func (r CainiaoConsignplatformOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoConsignplatformOrderCreateRequest) SetCreateRequest(createRequest *OrderCreateRequest) error {
    r.createRequest = createRequest
    r.Set("create_request", createRequest)
    return nil
}

func (r CainiaoConsignplatformOrderCreateRequest) GetCreateRequest() *OrderCreateRequest {
    return r.createRequest
}

