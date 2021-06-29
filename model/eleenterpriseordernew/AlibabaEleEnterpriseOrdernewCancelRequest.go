package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单取消 APIRequest
alibaba.ele.enterprise.ordernew.cancel

订单取消
*/
type AlibabaEleEnterpriseOrdernewCancelRequest struct {
    model.Params

    // 饿了么订单ID
    orderId   string 

    // 用户手机号
    phone   string 

    // 取消原因(取消时提供)
    reason   string 

}

func NewAlibabaEleEnterpriseOrdernewCancelRequest() *AlibabaEleEnterpriseOrdernewCancelRequest{
    return &AlibabaEleEnterpriseOrdernewCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseOrdernewCancelRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.cancel"
}

func (r AlibabaEleEnterpriseOrdernewCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseOrdernewCancelRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCancelRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaEleEnterpriseOrdernewCancelRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCancelRequest) GetPhone() string {
    return r.phone
}

func (r *AlibabaEleEnterpriseOrdernewCancelRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCancelRequest) GetReason() string {
    return r.reason
}

