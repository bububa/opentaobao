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
type AlibabaWttOfflineRecordQueryagentinfoAPIRequest struct {
    model.Params
    // 淘宝订单号
    _orderId   int64
    // 业务号码
    _phone   string
}

// 初始化AlibabaWttOfflineRecordQueryagentinfoAPIRequest对象
func NewAlibabaWttOfflineRecordQueryagentinfoRequest() *AlibabaWttOfflineRecordQueryagentinfoAPIRequest{
    return &AlibabaWttOfflineRecordQueryagentinfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWttOfflineRecordQueryagentinfoAPIRequest) GetApiMethodName() string {
    return "alibaba.wtt.offline.record.queryagentinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWttOfflineRecordQueryagentinfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 淘宝订单号
func (r *AlibabaWttOfflineRecordQueryagentinfoAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaWttOfflineRecordQueryagentinfoAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// Phone Setter
// 业务号码
func (r *AlibabaWttOfflineRecordQueryagentinfoAPIRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaWttOfflineRecordQueryagentinfoAPIRequest) GetPhone() string {
    return r._phone
}
