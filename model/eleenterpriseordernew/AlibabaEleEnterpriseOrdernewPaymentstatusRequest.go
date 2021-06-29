package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置订单支付 API请求
alibaba.ele.enterprise.ordernew.paymentstatus

设置订单支付成功
*/
type AlibabaEleEnterpriseOrdernewPaymentstatusRequest struct {
    model.Params
    // 订单id
    orderId   string
    // 支付流水号
    paySerialNumber   string
}

// 初始化AlibabaEleEnterpriseOrdernewPaymentstatusRequest对象
func NewAlibabaEleEnterpriseOrdernewPaymentstatusRequest() *AlibabaEleEnterpriseOrdernewPaymentstatusRequest{
    return &AlibabaEleEnterpriseOrdernewPaymentstatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewPaymentstatusRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.paymentstatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewPaymentstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaEleEnterpriseOrdernewPaymentstatusRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEleEnterpriseOrdernewPaymentstatusRequest) GetOrderId() string {
    return r.orderId
}
// PaySerialNumber Setter
// 支付流水号
func (r *AlibabaEleEnterpriseOrdernewPaymentstatusRequest) SetPaySerialNumber(paySerialNumber string) error {
    r.paySerialNumber = paySerialNumber
    r.Set("pay_serial_number", paySerialNumber)
    return nil
}

// PaySerialNumber Getter
func (r AlibabaEleEnterpriseOrdernewPaymentstatusRequest) GetPaySerialNumber() string {
    return r.paySerialNumber
}
