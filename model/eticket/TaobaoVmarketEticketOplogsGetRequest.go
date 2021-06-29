package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证操作日志查询 API请求
taobao.vmarket.eticket.oplogs.get

电子凭证核销日志查询
*/
type TaobaoVmarketEticketOplogsGetRequest struct {
    model.Params
    // 0:全部 1:核销 2:冲正
    type   int64
    // 开始时间
    startTime   string
    // 结束时间
    endTime   string
    // 核销码
    code   string
    // 手机号后四位
    mobile   string
    // 当前页码
    pageNo   int64
    // 每页显示的记录数，最大为40，默认为40
    pageSize   int64
    // 排序方式
    sort   string
    // 核销身份
    posid   string
    // 码商ID
    codemerchantId   int64
}

// 初始化TaobaoVmarketEticketOplogsGetRequest对象
func NewTaobaoVmarketEticketOplogsGetRequest() *TaobaoVmarketEticketOplogsGetRequest{
    return &TaobaoVmarketEticketOplogsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketOplogsGetRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.oplogs.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketOplogsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 0:全部 1:核销 2:冲正
func (r *TaobaoVmarketEticketOplogsGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoVmarketEticketOplogsGetRequest) GetType() int64 {
    return r.type
}
// StartTime Setter
// 开始时间
func (r *TaobaoVmarketEticketOplogsGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoVmarketEticketOplogsGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间
func (r *TaobaoVmarketEticketOplogsGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoVmarketEticketOplogsGetRequest) GetEndTime() string {
    return r.endTime
}
// Code Setter
// 核销码
func (r *TaobaoVmarketEticketOplogsGetRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r TaobaoVmarketEticketOplogsGetRequest) GetCode() string {
    return r.code
}
// Mobile Setter
// 手机号后四位
func (r *TaobaoVmarketEticketOplogsGetRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r TaobaoVmarketEticketOplogsGetRequest) GetMobile() string {
    return r.mobile
}
// PageNo Setter
// 当前页码
func (r *TaobaoVmarketEticketOplogsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoVmarketEticketOplogsGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页显示的记录数，最大为40，默认为40
func (r *TaobaoVmarketEticketOplogsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoVmarketEticketOplogsGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// Sort Setter
// 排序方式
func (r *TaobaoVmarketEticketOplogsGetRequest) SetSort(sort string) error {
    r.sort = sort
    r.Set("sort", sort)
    return nil
}

// Sort Getter
func (r TaobaoVmarketEticketOplogsGetRequest) GetSort() string {
    return r.sort
}
// Posid Setter
// 核销身份
func (r *TaobaoVmarketEticketOplogsGetRequest) SetPosid(posid string) error {
    r.posid = posid
    r.Set("posid", posid)
    return nil
}

// Posid Getter
func (r TaobaoVmarketEticketOplogsGetRequest) GetPosid() string {
    return r.posid
}
// CodemerchantId Setter
// 码商ID
func (r *TaobaoVmarketEticketOplogsGetRequest) SetCodemerchantId(codemerchantId int64) error {
    r.codemerchantId = codemerchantId
    r.Set("codemerchant_id", codemerchantId)
    return nil
}

// CodemerchantId Getter
func (r TaobaoVmarketEticketOplogsGetRequest) GetCodemerchantId() int64 {
    return r.codemerchantId
}
