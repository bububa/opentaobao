package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ERP增量开票结果获取 API请求
alibaba.einvoice.create.results.increment.get

增量开票结果获取
*/
type AlibabaEinvoiceCreateResultsIncrementGetRequest struct {
    model.Params
    // 开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败)
    status   string
    // 起始查询时间
    startModified   string
    // 收款方税务登记证号
    payeeRegisterNo   string
    // 终止查询时间
    endModified   string
    // 页面大小(不能超过200)
    pageSize   int64
    // 显示的页码
    pageNo   int64
}

// 初始化AlibabaEinvoiceCreateResultsIncrementGetRequest对象
func NewAlibabaEinvoiceCreateResultsIncrementGetRequest() *AlibabaEinvoiceCreateResultsIncrementGetRequest{
    return &AlibabaEinvoiceCreateResultsIncrementGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.create.results.increment.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败)
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetStatus() string {
    return r.status
}
// StartModified Setter
// 起始查询时间
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetStartModified(startModified string) error {
    r.startModified = startModified
    r.Set("start_modified", startModified)
    return nil
}

// StartModified Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetStartModified() string {
    return r.startModified
}
// PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// EndModified Setter
// 终止查询时间
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetEndModified(endModified string) error {
    r.endModified = endModified
    r.Set("end_modified", endModified)
    return nil
}

// EndModified Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetEndModified() string {
    return r.endModified
}
// PageSize Setter
// 页面大小(不能超过200)
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNo Setter
// 显示的页码
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetPageNo() int64 {
    return r.pageNo
}
