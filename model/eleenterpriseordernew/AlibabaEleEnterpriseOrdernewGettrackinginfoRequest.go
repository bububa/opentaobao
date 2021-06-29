package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单配送信息跟踪 APIRequest
alibaba.ele.enterprise.ordernew.gettrackinginfo

订单配送信息跟踪
*/
type AlibabaEleEnterpriseOrdernewGettrackinginfoRequest struct {
    model.Params

    // 饿了么订单ID
    orderId   string 

    // 用户手机号
    phone   string 

}

func NewAlibabaEleEnterpriseOrdernewGettrackinginfoRequest() *AlibabaEleEnterpriseOrdernewGettrackinginfoRequest{
    return &AlibabaEleEnterpriseOrdernewGettrackinginfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.gettrackinginfo"
}

func (r AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) GetPhone() string {
    return r.phone
}

