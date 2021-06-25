package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取租赁订单信息 APIRequest
tmall.car.leaseorder.get

卖家通过供销平台获取代销商的订单信息，但是部分情况下网商银行订单号获取不到，需要提供接口或者工具给卖家
*/
type TmallCarLeaseorderGetRequest struct {
    model.Params

    // 订单号
    orderId   int64 

}

func NewTmallCarLeaseorderGetRequest() *TmallCarLeaseorderGetRequest{
    return &TmallCarLeaseorderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeaseorderGetRequest) GetApiMethodName() string {
    return "tmall.car.leaseorder.get"
}

func (r TmallCarLeaseorderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarLeaseorderGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TmallCarLeaseorderGetRequest) GetOrderId() int64 {
    return r.orderId
}

