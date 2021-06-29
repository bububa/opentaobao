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
    startTime   string
    // 查询索引开始值：若不传，则只能查第一页
    startQueryIndexId   string
    // 结束时间
    endTime   string
    // 订单状态:Payment Completed,Buyer Confirmed Receipt
    status   string
    // 每页记录数
    pageSize   int64
    // 返回的字段信息
    fields   string
    // 安全签名
    appSignature   string
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
func (r *AliexpressAffiliateOrderListbyindexRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetStartTime() string {
    return r.startTime
}
// StartQueryIndexId Setter
// 查询索引开始值：若不传，则只能查第一页
func (r *AliexpressAffiliateOrderListbyindexRequest) SetStartQueryIndexId(startQueryIndexId string) error {
    r.startQueryIndexId = startQueryIndexId
    r.Set("start_query_index_id", startQueryIndexId)
    return nil
}

// StartQueryIndexId Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetStartQueryIndexId() string {
    return r.startQueryIndexId
}
// EndTime Setter
// 结束时间
func (r *AliexpressAffiliateOrderListbyindexRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetEndTime() string {
    return r.endTime
}
// Status Setter
// 订单状态:Payment Completed,Buyer Confirmed Receipt
func (r *AliexpressAffiliateOrderListbyindexRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetStatus() string {
    return r.status
}
// PageSize Setter
// 每页记录数
func (r *AliexpressAffiliateOrderListbyindexRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetPageSize() int64 {
    return r.pageSize
}
// Fields Setter
// 返回的字段信息
func (r *AliexpressAffiliateOrderListbyindexRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetFields() string {
    return r.fields
}
// AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateOrderListbyindexRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateOrderListbyindexRequest) GetAppSignature() string {
    return r.appSignature
}
