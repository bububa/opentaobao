package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单取消 API请求
alibaba.ele.enterprise.ordernew.cancel

订单取消
*/
type AlibabaEleEnterpriseOrdernewCancelRequest struct {
    model.Params
    // 饿了么订单ID
    _orderId   string
    // 用户手机号
    _phone   string
    // 取消原因(取消时提供)
    _reason   string
}

// 初始化AlibabaEleEnterpriseOrdernewCancelRequest对象
func NewAlibabaEleEnterpriseOrdernewCancelRequest() *AlibabaEleEnterpriseOrdernewCancelRequest{
    return &AlibabaEleEnterpriseOrdernewCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewCancelRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 饿了么订单ID
func (r *AlibabaEleEnterpriseOrdernewCancelRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEleEnterpriseOrdernewCancelRequest) GetOrderId() string {
    return r._orderId
}
// Phone Setter
// 用户手机号
func (r *AlibabaEleEnterpriseOrdernewCancelRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseOrdernewCancelRequest) GetPhone() string {
    return r._phone
}
// Reason Setter
// 取消原因(取消时提供)
func (r *AlibabaEleEnterpriseOrdernewCancelRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r AlibabaEleEnterpriseOrdernewCancelRequest) GetReason() string {
    return r._reason
}
