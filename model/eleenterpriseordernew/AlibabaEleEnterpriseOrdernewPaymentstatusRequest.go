package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置订单支付 APIRequest
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

func NewAlibabaEleEnterpriseOrdernewPaymentstatusRequest() *AlibabaEleEnterpriseOrdernewPaymentstatusRequest{
    return &AlibabaEleEnterpriseOrdernewPaymentstatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseOrdernewPaymentstatusRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.paymentstatus"
}

func (r AlibabaEleEnterpriseOrdernewPaymentstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseOrdernewPaymentstatusRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewPaymentstatusRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaEleEnterpriseOrdernewPaymentstatusRequest) SetPaySerialNumber(paySerialNumber string) error {
    r.paySerialNumber = paySerialNumber
    r.Set("pay_serial_number", paySerialNumber)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewPaymentstatusRequest) GetPaySerialNumber() string {
    return r.paySerialNumber
}

