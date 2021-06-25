package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退货入库单创建接口 APIRequest
taobao.qimen.returnorder.create

ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作
*/
type TaobaoQimenReturnorderCreateRequest struct {
    model.Params

    // 
    request   *ReturnOrderCreateRequest 

}

func NewTaobaoQimenReturnorderCreateRequest() *TaobaoQimenReturnorderCreateRequest{
    return &TaobaoQimenReturnorderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenReturnorderCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.returnorder.create"
}

func (r TaobaoQimenReturnorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenReturnorderCreateRequest) SetRequest(request *ReturnOrderCreateRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenReturnorderCreateRequest) GetRequest() *ReturnOrderCreateRequest {
    return r.request
}

