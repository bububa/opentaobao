package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退单和申诉 APIRequest
alibaba.ele.enterprise.ordernew.getrefundinfo

退单和申诉
*/
type AlibabaEleEnterpriseOrdernewGetrefundinfoRequest struct {
    model.Params

    // 饿了么订单ID
    orderId   string 

}

func NewAlibabaEleEnterpriseOrdernewGetrefundinfoRequest() *AlibabaEleEnterpriseOrdernewGetrefundinfoRequest{
    return &AlibabaEleEnterpriseOrdernewGetrefundinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseOrdernewGetrefundinfoRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.getrefundinfo"
}

func (r AlibabaEleEnterpriseOrdernewGetrefundinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseOrdernewGetrefundinfoRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewGetrefundinfoRequest) GetOrderId() string {
    return r.orderId
}

