package aedata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE推广者订单批量获取接口 APIRequest
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

func NewAliexpressAffiliateOrderListRequest() *AliexpressAffiliateOrderListRequest{
    return &AliexpressAffiliateOrderListRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAffiliateOrderListRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.order.list"
}

func (r AliexpressAffiliateOrderListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAffiliateOrderListRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r AliexpressAffiliateOrderListRequest) GetStartTime() string {
    return r.startTime
}

func (r *AliexpressAffiliateOrderListRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AliexpressAffiliateOrderListRequest) GetEndTime() string {
    return r.endTime
}

func (r *AliexpressAffiliateOrderListRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AliexpressAffiliateOrderListRequest) GetStatus() string {
    return r.status
}

func (r *AliexpressAffiliateOrderListRequest) SetLocaleSite(localeSite string) error {
    r.localeSite = localeSite
    r.Set("locale_site", localeSite)
    return nil
}

func (r AliexpressAffiliateOrderListRequest) GetLocaleSite() string {
    return r.localeSite
}

func (r *AliexpressAffiliateOrderListRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AliexpressAffiliateOrderListRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *AliexpressAffiliateOrderListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AliexpressAffiliateOrderListRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AliexpressAffiliateOrderListRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r AliexpressAffiliateOrderListRequest) GetFields() string {
    return r.fields
}

func (r *AliexpressAffiliateOrderListRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

func (r AliexpressAffiliateOrderListRequest) GetAppSignature() string {
    return r.appSignature
}

