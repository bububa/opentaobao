package aliexpresssumaitong

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关务所需的申报清关字段 APIRequest
aliexpress.taxation.calculate.open.query

关务所需的申报清关字段
*/
type AliexpressTaxationCalculateOpenQueryRequest struct {
    model.Params

    // 主订单id
    orderId   string 

}

func NewAliexpressTaxationCalculateOpenQueryRequest() *AliexpressTaxationCalculateOpenQueryRequest{
    return &AliexpressTaxationCalculateOpenQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressTaxationCalculateOpenQueryRequest) GetApiMethodName() string {
    return "aliexpress.taxation.calculate.open.query"
}

func (r AliexpressTaxationCalculateOpenQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressTaxationCalculateOpenQueryRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AliexpressTaxationCalculateOpenQueryRequest) GetOrderId() string {
    return r.orderId
}

