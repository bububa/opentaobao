package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户投诉 API请求
alibaba.happytrip.taxi.order.complain

一个订单只能投诉一次，不可重复投诉

投诉选项
0	其他原因；
1	司机原因导致行程被取消；
2	服务态度恶劣；	
3	未坐车产生费用；	
4	额外收取不合理费用；
5	路不熟多产生费用；	
6	提前计费；
7	未及时结束计费；	
8	司机绕路；	
9	司机迟到；	
10	司机爽约或拒载；	
11	骚扰乘客；	
12	危险驾驶；	
13	不是订单显示车辆或司机；
*/
type AlibabaHappytripTaxiOrderComplainAPIRequest struct {
    model.Params
    // 订单id
    _orderId   string
    // 投诉选项 0	其他原因； 1	司机原因导致行程被取消； 2	服务态度恶劣；	 3	未坐车产生费用；	 4	额外收取不合理费用； 5	路不熟多产生费用；	 6	提前计费； 7	未及时结束计费；	 8	司机绕路；	 9	司机迟到；	 10	司机爽约或拒载；	 11	骚扰乘客；	 12	危险驾驶；	 13	不是订单显示车辆或司机；
    _type   int64
    // 投诉选项外的其他投诉内容,不能多于40个汉字
    _content   string
    // 该用户的真实号码
    _mobile   string
    // 投诉时间（毫秒）
    _time   int64
}

// 初始化AlibabaHappytripTaxiOrderComplainAPIRequest对象
func NewAlibabaHappytripTaxiOrderComplainRequest() *AlibabaHappytripTaxiOrderComplainAPIRequest{
    return &AlibabaHappytripTaxiOrderComplainAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderComplainAPIRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.complain"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderComplainAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderComplainAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderComplainAPIRequest) GetOrderId() string {
    return r._orderId
}
// Type Setter
// 投诉选项 0	其他原因； 1	司机原因导致行程被取消； 2	服务态度恶劣；	 3	未坐车产生费用；	 4	额外收取不合理费用； 5	路不熟多产生费用；	 6	提前计费； 7	未及时结束计费；	 8	司机绕路；	 9	司机迟到；	 10	司机爽约或拒载；	 11	骚扰乘客；	 12	危险驾驶；	 13	不是订单显示车辆或司机；
func (r *AlibabaHappytripTaxiOrderComplainAPIRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaHappytripTaxiOrderComplainAPIRequest) GetType() int64 {
    return r._type
}
// Content Setter
// 投诉选项外的其他投诉内容,不能多于40个汉字
func (r *AlibabaHappytripTaxiOrderComplainAPIRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r AlibabaHappytripTaxiOrderComplainAPIRequest) GetContent() string {
    return r._content
}
// Mobile Setter
// 该用户的真实号码
func (r *AlibabaHappytripTaxiOrderComplainAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHappytripTaxiOrderComplainAPIRequest) GetMobile() string {
    return r._mobile
}
// Time Setter
// 投诉时间（毫秒）
func (r *AlibabaHappytripTaxiOrderComplainAPIRequest) SetTime(_time int64) error {
    r._time = _time
    r.Set("time", _time)
    return nil
}

// Time Getter
func (r AlibabaHappytripTaxiOrderComplainAPIRequest) GetTime() int64 {
    return r._time
}
