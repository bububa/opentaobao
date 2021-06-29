package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单配送信息跟踪 API请求
alibaba.ele.enterprise.ordernew.gettrackinginfo

订单配送信息跟踪
*/
type AlibabaEleEnterpriseOrdernewGettrackinginfoRequest struct {
    model.Params
    // 饿了么订单ID
    _orderId   string
    // 用户手机号
    _phone   string
}

// 初始化AlibabaEleEnterpriseOrdernewGettrackinginfoRequest对象
func NewAlibabaEleEnterpriseOrdernewGettrackinginfoRequest() *AlibabaEleEnterpriseOrdernewGettrackinginfoRequest{
    return &AlibabaEleEnterpriseOrdernewGettrackinginfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.gettrackinginfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 饿了么订单ID
func (r *AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) GetOrderId() string {
    return r._orderId
}
// Phone Setter
// 用户手机号
func (r *AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoRequest) GetPhone() string {
    return r._phone
}
