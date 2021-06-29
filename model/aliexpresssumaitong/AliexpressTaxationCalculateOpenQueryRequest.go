package aliexpresssumaitong

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关务所需的申报清关字段 API请求
aliexpress.taxation.calculate.open.query

关务所需的申报清关字段
*/
type AliexpressTaxationCalculateOpenQueryRequest struct {
    model.Params
    // 主订单id
    _orderId   string
}

// 初始化AliexpressTaxationCalculateOpenQueryRequest对象
func NewAliexpressTaxationCalculateOpenQueryRequest() *AliexpressTaxationCalculateOpenQueryRequest{
    return &AliexpressTaxationCalculateOpenQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressTaxationCalculateOpenQueryRequest) GetApiMethodName() string {
    return "aliexpress.taxation.calculate.open.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressTaxationCalculateOpenQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 主订单id
func (r *AliexpressTaxationCalculateOpenQueryRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AliexpressTaxationCalculateOpenQueryRequest) GetOrderId() string {
    return r._orderId
}
