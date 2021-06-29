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
    _status   string
    // 起始查询时间
    _startModified   string
    // 收款方税务登记证号
    _payeeRegisterNo   string
    // 终止查询时间
    _endModified   string
    // 页面大小(不能超过200)
    _pageSize   int64
    // 显示的页码
    _pageNo   int64
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
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetStatus() string {
    return r._status
}
// StartModified Setter
// 起始查询时间
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetStartModified(_startModified string) error {
    r._startModified = _startModified
    r.Set("start_modified", _startModified)
    return nil
}

// StartModified Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetStartModified() string {
    return r._startModified
}
// PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// EndModified Setter
// 终止查询时间
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetEndModified(_endModified string) error {
    r._endModified = _endModified
    r.Set("end_modified", _endModified)
    return nil
}

// EndModified Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetEndModified() string {
    return r._endModified
}
// PageSize Setter
// 页面大小(不能超过200)
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 显示的页码
func (r *AlibabaEinvoiceCreateResultsIncrementGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaEinvoiceCreateResultsIncrementGetRequest) GetPageNo() int64 {
    return r._pageNo
}
