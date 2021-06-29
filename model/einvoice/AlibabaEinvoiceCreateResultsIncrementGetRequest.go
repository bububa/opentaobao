package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ERP增量开票结果获取 APIRequest
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

func NewAlibabaEinvoiceCreateResultsIncrementGetRequest() *AlibabaEinvoiceCreateResultsIncrementGetRequest{
    return &AlibabaEinvoiceCreateResultsIncrementGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.create.results.increment.get"
}

func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetStatus() string {
    return r.status
}

func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetStartModified(startModified string) error {
    r.startModified = startModified
    r.Set("start_modified", startModified)
    return nil
}

func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetStartModified() string {
    return r.startModified
}

func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetEndModified(endModified string) error {
    r.endModified = endModified
    r.Set("end_modified", endModified)
    return nil
}

func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetEndModified() string {
    return r.endModified
}

func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetPageNo() int64 {
    return r.pageNo
}

