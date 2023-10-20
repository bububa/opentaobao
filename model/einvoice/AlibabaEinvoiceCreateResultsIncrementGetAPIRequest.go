package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicecreateresultsincrementgetAPIRequest ERP增量开票结果获取 API请求
// alibaba.einvoice.create.results.increment.get
//
// 增量开票结果获取
type AlibabaeinvoicecreateresultsincrementgetAPIRequest struct {
	model.Params
	// 终止查询时间
	_endModified string
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 起始查询时间
	_startModified string
	// 开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败)
	_status string
	// 页面大小(不能超过200)
	_pageSize int64
	// 显示的页码
	_pageNo int64
}

// NewAlibabaeinvoicecreateresultsincrementgetRequest 初始化AlibabaeinvoicecreateresultsincrementgetAPIRequest对象
func NewAlibabaeinvoicecreateresultsincrementgetRequest() *AlibabaeinvoicecreateresultsincrementgetAPIRequest {
	return &AlibabaeinvoicecreateresultsincrementgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicecreateresultsincrementgetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.create.results.increment.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicecreateresultsincrementgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicecreateresultsincrementgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndModified is EndModified Setter
// 终止查询时间
func (r *AlibabaeinvoicecreateresultsincrementgetAPIRequest) SetEndModified(_endModified string) error {
	r._endModified = _endModified
	r.Set("end_modified", _endModified)
	return nil
}

// GetEndModified EndModified Getter
func (r AlibabaeinvoicecreateresultsincrementgetAPIRequest) GetEndModified() string {
	return r._endModified
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaeinvoicecreateresultsincrementgetAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoicecreateresultsincrementgetAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetStartModified is StartModified Setter
// 起始查询时间
func (r *AlibabaeinvoicecreateresultsincrementgetAPIRequest) SetStartModified(_startModified string) error {
	r._startModified = _startModified
	r.Set("start_modified", _startModified)
	return nil
}

// GetStartModified StartModified Getter
func (r AlibabaeinvoicecreateresultsincrementgetAPIRequest) GetStartModified() string {
	return r._startModified
}

// SetStatus is Status Setter
// 开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败)
func (r *AlibabaeinvoicecreateresultsincrementgetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaeinvoicecreateresultsincrementgetAPIRequest) GetStatus() string {
	return r._status
}

// SetPageSize is PageSize Setter
// 页面大小(不能超过200)
func (r *AlibabaeinvoicecreateresultsincrementgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaeinvoicecreateresultsincrementgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 显示的页码
func (r *AlibabaeinvoicecreateresultsincrementgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AlibabaeinvoicecreateresultsincrementgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
