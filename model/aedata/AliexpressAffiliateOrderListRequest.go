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
    _startTime   string
    // 结束时间
    _endTime   string
    // 订单状态:Payment Completed,Buyer Confirmed Receipt
    _status   string
    // 站点信息：global、ru_site、es_site、it_site
    _localeSite   string
    // 页数
    _pageNo   int64
    // 每页记录数
    _pageSize   int64
    // 返回的字段信息
    _fields   string
    // 安全签名
    _appSignature   string
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
func (r *AliexpressAffiliateOrderListRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r AliexpressAffiliateOrderListRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间
func (r *AliexpressAffiliateOrderListRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r AliexpressAffiliateOrderListRequest) GetEndTime() string {
    return r._endTime
}
// Status Setter
// 订单状态:Payment Completed,Buyer Confirmed Receipt
func (r *AliexpressAffiliateOrderListRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AliexpressAffiliateOrderListRequest) GetStatus() string {
    return r._status
}
// LocaleSite Setter
// 站点信息：global、ru_site、es_site、it_site
func (r *AliexpressAffiliateOrderListRequest) SetLocaleSite(_localeSite string) error {
    r._localeSite = _localeSite
    r.Set("locale_site", _localeSite)
    return nil
}

// LocaleSite Getter
func (r AliexpressAffiliateOrderListRequest) GetLocaleSite() string {
    return r._localeSite
}
// PageNo Setter
// 页数
func (r *AliexpressAffiliateOrderListRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressAffiliateOrderListRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页记录数
func (r *AliexpressAffiliateOrderListRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressAffiliateOrderListRequest) GetPageSize() int64 {
    return r._pageSize
}
// Fields Setter
// 返回的字段信息
func (r *AliexpressAffiliateOrderListRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateOrderListRequest) GetFields() string {
    return r._fields
}
// AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateOrderListRequest) SetAppSignature(_appSignature string) error {
    r._appSignature = _appSignature
    r.Set("app_signature", _appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateOrderListRequest) GetAppSignature() string {
    return r._appSignature
}
