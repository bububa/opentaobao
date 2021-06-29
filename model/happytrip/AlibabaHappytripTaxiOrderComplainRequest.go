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
type AlibabaHappytripTaxiOrderComplainRequest struct {
    model.Params
    // 订单id
    orderId   string
    // 投诉选项 0	其他原因； 1	司机原因导致行程被取消； 2	服务态度恶劣；	 3	未坐车产生费用；	 4	额外收取不合理费用； 5	路不熟多产生费用；	 6	提前计费； 7	未及时结束计费；	 8	司机绕路；	 9	司机迟到；	 10	司机爽约或拒载；	 11	骚扰乘客；	 12	危险驾驶；	 13	不是订单显示车辆或司机；
    type   int64
    // 投诉选项外的其他投诉内容,不能多于40个汉字
    content   string
    // 该用户的真实号码
    mobile   string
    // 投诉时间（毫秒）
    time   int64
}

// 初始化AlibabaHappytripTaxiOrderComplainRequest对象
func NewAlibabaHappytripTaxiOrderComplainRequest() *AlibabaHappytripTaxiOrderComplainRequest{
    return &AlibabaHappytripTaxiOrderComplainRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderComplainRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.complain"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderComplainRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderComplainRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderComplainRequest) GetOrderId() string {
    return r.orderId
}
// Type Setter
// 投诉选项 0	其他原因； 1	司机原因导致行程被取消； 2	服务态度恶劣；	 3	未坐车产生费用；	 4	额外收取不合理费用； 5	路不熟多产生费用；	 6	提前计费； 7	未及时结束计费；	 8	司机绕路；	 9	司机迟到；	 10	司机爽约或拒载；	 11	骚扰乘客；	 12	危险驾驶；	 13	不是订单显示车辆或司机；
func (r *AlibabaHappytripTaxiOrderComplainRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaHappytripTaxiOrderComplainRequest) GetType() int64 {
    return r.type
}
// Content Setter
// 投诉选项外的其他投诉内容,不能多于40个汉字
func (r *AlibabaHappytripTaxiOrderComplainRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r AlibabaHappytripTaxiOrderComplainRequest) GetContent() string {
    return r.content
}
// Mobile Setter
// 该用户的真实号码
func (r *AlibabaHappytripTaxiOrderComplainRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHappytripTaxiOrderComplainRequest) GetMobile() string {
    return r.mobile
}
// Time Setter
// 投诉时间（毫秒）
func (r *AlibabaHappytripTaxiOrderComplainRequest) SetTime(time int64) error {
    r.time = time
    r.Set("time", time)
    return nil
}

// Time Getter
func (r AlibabaHappytripTaxiOrderComplainRequest) GetTime() int64 {
    return r.time
}
