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
    _orderId   string
    // 支付流水号
    _paySerialNumber   string
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
func (r *AlibabaEleEnterpriseOrdernewPaymentstatusRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEleEnterpriseOrdernewPaymentstatusRequest) GetOrderId() string {
    return r._orderId
}
// PaySerialNumber Setter
// 支付流水号
func (r *AlibabaEleEnterpriseOrdernewPaymentstatusRequest) SetPaySerialNumber(_paySerialNumber string) error {
    r._paySerialNumber = _paySerialNumber
    r.Set("pay_serial_number", _paySerialNumber)
    return nil
}

// PaySerialNumber Getter
func (r AlibabaEleEnterpriseOrdernewPaymentstatusRequest) GetPaySerialNumber() string {
    return r._paySerialNumber
}