package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退单和申诉 API请求
alibaba.ele.enterprise.ordernew.getrefundinfo

退单和申诉
*/
type AlibabaEleEnterpriseOrdernewGetrefundinfoRequest struct {
    model.Params
    // 饿了么订单ID
    orderId   string
}

// 初始化AlibabaEleEnterpriseOrdernewGetrefundinfoRequest对象
func NewAlibabaEleEnterpriseOrdernewGetrefundinfoRequest() *AlibabaEleEnterpriseOrdernewGetrefundinfoRequest{
    return &AlibabaEleEnterpriseOrdernewGetrefundinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewGetrefundinfoRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.getrefundinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewGetrefundinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 饿了么订单ID
func (r *AlibabaEleEnterpriseOrdernewGetrefundinfoRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEleEnterpriseOrdernewGetrefundinfoRequest) GetOrderId() string {
    return r.orderId
}
