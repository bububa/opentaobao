package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询出入库单处理结果api APIRequest
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

func NewAlibabaAlihealthTracecodesellerBillResultSearchRequest() *AlibabaAlihealthTracecodesellerBillResultSearchRequest{
    return &AlibabaAlihealthTracecodesellerBillResultSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.bill.result.search"
}

func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetSkeyCode(skeyCode string) error {
    r.skeyCode = skeyCode
    r.Set("skey_code", skeyCode)
    return nil
}

func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetSkeyCode() string {
    return r.skeyCode
}

func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetEntInfoId(entInfoId int64) error {
    r.entInfoId = entInfoId
    r.Set("ent_info_id", entInfoId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetEntInfoId() int64 {
    return r.entInfoId
}

func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetBeginDate() string {
    return r.beginDate
}

func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetSellerName(sellerName string) error {
    r.sellerName = sellerName
    r.Set("seller_name", sellerName)
    return nil
}

func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetSellerName() string {
    return r.sellerName
}

func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthTracecodesellerBillResultSearchRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthTracecodesellerBillResultSearchRequest) GetPage() int64 {
    return r.page
}

