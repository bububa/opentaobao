package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单详情 API请求
alibaba.ele.enterprise.ordernew.get

查询订单详情
*/
type AlibabaEleEnterpriseOrdernewGetRequest struct {
    model.Params
    // 饿了么订单ID
    orderId   string
    // 电话号码
    phone   string
}

// 初始化AlibabaEleEnterpriseOrdernewGetRequest对象
func NewAlibabaEleEnterpriseOrdernewGetRequest() *AlibabaEleEnterpriseOrdernewGetRequest{
    return &AlibabaEleEnterpriseOrdernewGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewGetRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 饿了么订单ID
func (r *AlibabaEleEnterpriseOrdernewGetRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEleEnterpriseOrdernewGetRequest) GetOrderId() string {
    return r.orderId
}
// Phone Setter
// 电话号码
func (r *AlibabaEleEnterpriseOrdernewGetRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseOrdernewGetRequest) GetPhone() string {
    return r.phone
}
