package aedata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE联盟推广者订单查询接口-按游标索引查询 APIRequest
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

func NewAliexpressAffiliateOrderListbyindexRequest() *AliexpressAffiliateOrderListbyindexRequest{
    return &AliexpressAffiliateOrderListbyindexRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAffiliateOrderListbyindexRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.order.listbyindex"
}

func (r AliexpressAffiliateOrderListbyindexRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAffiliateOrderListbyindexRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r AliexpressAffiliateOrderListbyindexRequest) GetStartTime() string {
    return r.startTime
}

func (r *AliexpressAffiliateOrderListbyindexRequest) SetStartQueryIndexId(startQueryIndexId string) error {
    r.startQueryIndexId = startQueryIndexId
    r.Set("start_query_index_id", startQueryIndexId)
    return nil
}

func (r AliexpressAffiliateOrderListbyindexRequest) GetStartQueryIndexId() string {
    return r.startQueryIndexId
}

func (r *AliexpressAffiliateOrderListbyindexRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AliexpressAffiliateOrderListbyindexRequest) GetEndTime() string {
    return r.endTime
}

func (r *AliexpressAffiliateOrderListbyindexRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AliexpressAffiliateOrderListbyindexRequest) GetStatus() string {
    return r.status
}

func (r *AliexpressAffiliateOrderListbyindexRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AliexpressAffiliateOrderListbyindexRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AliexpressAffiliateOrderListbyindexRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r AliexpressAffiliateOrderListbyindexRequest) GetFields() string {
    return r.fields
}

func (r *AliexpressAffiliateOrderListbyindexRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

func (r AliexpressAffiliateOrderListbyindexRequest) GetAppSignature() string {
    return r.appSignature
}

