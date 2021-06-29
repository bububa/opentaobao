package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下推广充送等业务订单来源 API请求
alibaba.wtt.offline.record.queryagentinfo

线下推广充送等业务订单来源的查询接口
*/
type AlibabaWttOfflineRecordQueryagentinfoRequest struct {
    model.Params
    // 淘宝订单号
    orderId   int64
    // 业务号码
    phone   string
}

// 初始化AlibabaWttOfflineRecordQueryagentinfoRequest对象
func NewAlibabaWttOfflineRecordQueryagentinfoRequest() *AlibabaWttOfflineRecordQueryagentinfoRequest{
    return &AlibabaWttOfflineRecordQueryagentinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWttOfflineRecordQueryagentinfoRequest) GetApiMethodName() string {
    return "alibaba.wtt.offline.record.queryagentinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWttOfflineRecordQueryagentinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 淘宝订单号
func (r *AlibabaWttOfflineRecordQueryagentinfoRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaWttOfflineRecordQueryagentinfoRequest) GetOrderId() int64 {
    return r.orderId
}
// Phone Setter
// 业务号码
func (r *AlibabaWttOfflineRecordQueryagentinfoRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r AlibabaWttOfflineRecordQueryagentinfoRequest) GetPhone() string {
    return r.phone
}
