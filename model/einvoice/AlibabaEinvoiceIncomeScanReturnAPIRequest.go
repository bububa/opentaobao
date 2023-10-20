package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceincomescanreturnAPIRequest 进项扫描状态回传 API请求
// alibaba.einvoice.income.scan.return
//
// 回传进项扫描每个阶段的状态，比如ocr开始，ocr结束，查验开始，查验结束等
type AlibabaeinvoiceincomescanreturnAPIRequest struct {
	model.Params
	// 扫描的批次号
	_batchNo string
	// 错误码，success=false时填入
	_errorCode string
	// 错误信息，success=false时必填
	_errorMessage string
	// 扫描状态，0=开始ocr，1=ocr结束，2=开始查验，3=查验结束
	_status int64
	// 该批次对应的发票数量，扫描结束和查验结束status=1，3时必填
	_invoiceCount int64
	// 驱动是否成功，true=成功，false=失败
	_success bool
}

// NewAlibabaeinvoiceincomescanreturnRequest 初始化AlibabaeinvoiceincomescanreturnAPIRequest对象
func NewAlibabaeinvoiceincomescanreturnRequest() *AlibabaeinvoiceincomescanreturnAPIRequest {
	return &AlibabaeinvoiceincomescanreturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceincomescanreturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.scan.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceincomescanreturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceincomescanreturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchNo is BatchNo Setter
// 扫描的批次号
func (r *AlibabaeinvoiceincomescanreturnAPIRequest) SetBatchNo(_batchNo string) error {
	r._batchNo = _batchNo
	r.Set("batch_no", _batchNo)
	return nil
}

// GetBatchNo BatchNo Getter
func (r AlibabaeinvoiceincomescanreturnAPIRequest) GetBatchNo() string {
	return r._batchNo
}

// SetErrorCode is ErrorCode Setter
// 错误码，success=false时填入
func (r *AlibabaeinvoiceincomescanreturnAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r AlibabaeinvoiceincomescanreturnAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// SetErrorMessage is ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaeinvoiceincomescanreturnAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// GetErrorMessage ErrorMessage Getter
func (r AlibabaeinvoiceincomescanreturnAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// SetStatus is Status Setter
// 扫描状态，0=开始ocr，1=ocr结束，2=开始查验，3=查验结束
func (r *AlibabaeinvoiceincomescanreturnAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaeinvoiceincomescanreturnAPIRequest) GetStatus() int64 {
	return r._status
}

// SetInvoiceCount is InvoiceCount Setter
// 该批次对应的发票数量，扫描结束和查验结束status=1，3时必填
func (r *AlibabaeinvoiceincomescanreturnAPIRequest) SetInvoiceCount(_invoiceCount int64) error {
	r._invoiceCount = _invoiceCount
	r.Set("invoice_count", _invoiceCount)
	return nil
}

// GetInvoiceCount InvoiceCount Getter
func (r AlibabaeinvoiceincomescanreturnAPIRequest) GetInvoiceCount() int64 {
	return r._invoiceCount
}

// SetSuccess is Success Setter
// 驱动是否成功，true=成功，false=失败
func (r *AlibabaeinvoiceincomescanreturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r AlibabaeinvoiceincomescanreturnAPIRequest) GetSuccess() bool {
	return r._success
}
