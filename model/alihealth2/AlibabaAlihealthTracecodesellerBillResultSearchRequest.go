package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询出入库单处理结果api API请求
alibaba.alihealth.tracecodeseller.bill.result.search

查询出入库单处理结果api
*/
type AlibabaAlihealthTracecodesellerBillResultSearchRequest struct {
    model.Params
    // top身份认证
    skeyCode   string
    // 商家id
    entInfoId   int64
    // 单据编号
    billCode   string
    // 查询开始日期
    beginDate   string
    // 查询结束日期
    endDate   string
    // 不需要
    sellerName   string
    // 每页条数
    pageSize   int64
    // 当前页
    page   int64
}

// 初始化AlibabaAlihealthTracecodesellerBillResultSearchRequest对象
func NewAlibabaAlihealthTracecodesellerBillResultSearchRequest() *AlibabaAlihealthTracecodesellerBillResultSearchRequest{
    return &AlibabaAlihealthTracecodesellerBillResultSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.bill.result.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkeyCode Setter
// top身份认证
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetSkeyCode(skeyCode string) error {
    r.skeyCode = skeyCode
    r.Set("skey_code", skeyCode)
    return nil
}

// SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetSkeyCode() string {
    return r.skeyCode
}
// EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetEntInfoId(entInfoId int64) error {
    r.entInfoId = entInfoId
    r.Set("ent_info_id", entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetEntInfoId() int64 {
    return r.entInfoId
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetBillCode() string {
    return r.billCode
}
// BeginDate Setter
// 查询开始日期
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetBeginDate() string {
    return r.beginDate
}
// EndDate Setter
// 查询结束日期
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetEndDate() string {
    return r.endDate
}
// SellerName Setter
// 不需要
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetSellerName(sellerName string) error {
    r.sellerName = sellerName
    r.Set("seller_name", sellerName)
    return nil
}

// SellerName Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetSellerName() string {
    return r.sellerName
}
// PageSize Setter
// 每页条数
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetPageSize() int64 {
    return r.pageSize
}
// Page Setter
// 当前页
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetPage() int64 {
    return r.page
}
