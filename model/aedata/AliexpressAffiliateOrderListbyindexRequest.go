package aedata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE联盟推广者订单查询接口-按游标索引查询 API请求
aliexpress.affiliate.order.listbyindex

AE联盟推广者订单按游标查询接口
*/
type AliexpressAffiliateOrderListbyindexRequest struct {
    model.Params
    // 开始时间
    _startTime   string
    // 查询索引开始值：若不传，则只能查第一页
    _startQueryIndexId   string
    // 结束时间
    _endTime   string
    // 订单状态:Payment Completed,Buyer Confirmed Receipt
    _status   string
    // 每页记录数
    _pageSize   int64
    // 返回的字段信息
    _fields   string
    // 安全签名
    _appSignature   string
}

// 初始化AliexpressAffiliateOrderListbyindexRequest对象
func NewAliexpressAffiliateOrderListbyindexRequest() *AliexpressAffiliateOrderListbyindexRequest{
    return &AliexpressAffiliateOrderListbyindexRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateOrderListbyindexRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.order.listbyindex"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateOrderListbyindexRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartTime Setter
// 开始时间
func (r *AliexpressAffiliateOrderListbyindexRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetStartTime() string {
    return r._startTime
}
// StartQueryIndexId Setter
// 查询索引开始值：若不传，则只能查第一页
func (r *AliexpressAffiliateOrderListbyindexRequest) SetStartQueryIndexId(_startQueryIndexId string) error {
    r._startQueryIndexId = _startQueryIndexId
    r.Set("start_query_index_id", _startQueryIndexId)
    return nil
}

// StartQueryIndexId Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetStartQueryIndexId() string {
    return r._startQueryIndexId
}
// EndTime Setter
// 结束时间
func (r *AliexpressAffiliateOrderListbyindexRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetEndTime() string {
    return r._endTime
}
// Status Setter
// 订单状态:Payment Completed,Buyer Confirmed Receipt
func (r *AliexpressAffiliateOrderListbyindexRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetStatus() string {
    return r._status
}
// PageSize Setter
// 每页记录数
func (r *AliexpressAffiliateOrderListbyindexRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetPageSize() int64 {
    return r._pageSize
}
// Fields Setter
// 返回的字段信息
func (r *AliexpressAffiliateOrderListbyindexRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetFields() string {
    return r._fields
}
// AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateOrderListbyindexRequest) SetAppSignature(_appSignature string) error {
    r._appSignature = _appSignature
    r.Set("app_signature", _appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetAppSignature() string {
    return r._appSignature
}
