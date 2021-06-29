package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单详情 APIRequest
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

func NewAlibabaEleEnterpriseOrdernewGetRequest() *AlibabaEleEnterpriseOrdernewGetRequest{
    return &AlibabaEleEnterpriseOrdernewGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseOrdernewGetRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.get"
}

func (r AlibabaEleEnterpriseOrdernewGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseOrdernewGetRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewGetRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaEleEnterpriseOrdernewGetRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewGetRequest) GetPhone() string {
    return r.phone
}

