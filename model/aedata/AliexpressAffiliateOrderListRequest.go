package aedata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE推广者订单批量获取接口 API请求
aliexpress.affiliate.order.list

AE联盟推广者订单分页查询接口
*/
type AliexpressAffiliateOrderListRequest struct {
    model.Params
    // 开始时间
    startTime   string
    // 结束时间
    endTime   string
    // 订单状态:Payment Completed,Buyer Confirmed Receipt
    status   string
    // 站点信息：global、ru_site、es_site、it_site
    localeSite   string
    // 页数
    pageNo   int64
    // 每页记录数
    pageSize   int64
    // 返回的字段信息
    fields   string
    // 安全签名
    appSignature   string
}

// 初始化AliexpressAffiliateOrderListRequest对象
func NewAliexpressAffiliateOrderListRequest() *AliexpressAffiliateOrderListRequest{
    return &AliexpressAffiliateOrderListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateOrderListRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.order.list"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateOrderListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartTime Setter
// 开始时间
func (r *AliexpressAffiliateOrderListRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AliexpressAffiliateOrderListRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间
func (r *AliexpressAffiliateOrderListRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AliexpressAffiliateOrderListRequest) GetEndTime() string {
    return r.endTime
}
// Status Setter
// 订单状态:Payment Completed,Buyer Confirmed Receipt
func (r *AliexpressAffiliateOrderListRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AliexpressAffiliateOrderListRequest) GetStatus() string {
    return r.status
}
// LocaleSite Setter
// 站点信息：global、ru_site、es_site、it_site
func (r *AliexpressAffiliateOrderListRequest) SetLocaleSite(localeSite string) error {
    r.localeSite = localeSite
    r.Set("locale_site", localeSite)
    return nil
}

// LocaleSite Getter
func (r AliexpressAffiliateOrderListRequest) GetLocaleSite() string {
    return r.localeSite
}
// PageNo Setter
// 页数
func (r *AliexpressAffiliateOrderListRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressAffiliateOrderListRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页记录数
func (r *AliexpressAffiliateOrderListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressAffiliateOrderListRequest) GetPageSize() int64 {
    return r.pageSize
}
// Fields Setter
// 返回的字段信息
func (r *AliexpressAffiliateOrderListRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateOrderListRequest) GetFields() string {
    return r.fields
}
// AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateOrderListRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateOrderListRequest) GetAppSignature() string {
    return r.appSignature
}
