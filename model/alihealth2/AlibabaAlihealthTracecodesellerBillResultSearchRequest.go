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
    _skeyCode   string
    // 商家id
    _entInfoId   int64
    // 单据编号
    _billCode   string
    // 查询开始日期
    _beginDate   string
    // 查询结束日期
    _endDate   string
    // 不需要
    _sellerName   string
    // 每页条数
    _pageSize   int64
    // 当前页
    _page   int64
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
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetSkeyCode(_skeyCode string) error {
    r._skeyCode = _skeyCode
    r.Set("skey_code", _skeyCode)
    return nil
}

// SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetSkeyCode() string {
    return r._skeyCode
}
// EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetEntInfoId(_entInfoId int64) error {
    r._entInfoId = _entInfoId
    r.Set("ent_info_id", _entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetEntInfoId() int64 {
    return r._entInfoId
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetBillCode() string {
    return r._billCode
}
// BeginDate Setter
// 查询开始日期
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetBeginDate(_beginDate string) error {
    r._beginDate = _beginDate
    r.Set("begin_date", _beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetBeginDate() string {
    return r._beginDate
}
// EndDate Setter
// 查询结束日期
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetEndDate() string {
    return r._endDate
}
// SellerName Setter
// 不需要
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetSellerName(_sellerName string) error {
    r._sellerName = _sellerName
    r.Set("seller_name", _sellerName)
    return nil
}

// SellerName Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetSellerName() string {
    return r._sellerName
}
// PageSize Setter
// 每页条数
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
// Page Setter
// 当前页
func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetPage() int64 {
    return r._page
}
