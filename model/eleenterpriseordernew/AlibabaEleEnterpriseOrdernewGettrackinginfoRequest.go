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
type AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest struct {
    model.Params
    // 饿了么订单ID
    _orderId   string
    // 用户手机号
    _phone   string
}

// 初始化AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest对象
func NewAlibabaEleEnterpriseOrdernewGettrackinginfoRequest() *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest{
    return &AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.gettrackinginfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 饿了么订单ID
func (r *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) GetOrderId() string {
    return r._orderId
}
// Phone Setter
// 用户手机号
func (r *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) GetPhone() string {
    return r._phone
}
