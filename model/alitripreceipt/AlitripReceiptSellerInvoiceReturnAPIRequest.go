package alitripreceipt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripreceiptsellerinvoicereturnAPIRequest 飞猪发票商家回调接口 API请求
// alitrip.receipt.seller.invoice.return
//
// 飞猪发票回调接口
type AlitripreceiptsellerinvoicereturnAPIRequest struct {
	model.Params
	// 入参对象
	_receiptDo *ReceiptDo
}

// NewAlitripreceiptsellerinvoicereturnRequest 初始化AlitripreceiptsellerinvoicereturnAPIRequest对象
func NewAlitripreceiptsellerinvoicereturnRequest() *AlitripreceiptsellerinvoicereturnAPIRequest {
	return &AlitripreceiptsellerinvoicereturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripreceiptsellerinvoicereturnAPIRequest) GetApiMethodName() string {
	return "alitrip.receipt.seller.invoice.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripreceiptsellerinvoicereturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripreceiptsellerinvoicereturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiptDo is ReceiptDo Setter
// 入参对象
func (r *AlitripreceiptsellerinvoicereturnAPIRequest) SetReceiptDo(_receiptDo *ReceiptDo) error {
	r._receiptDo = _receiptDo
	r.Set("receipt_do", _receiptDo)
	return nil
}

// GetReceiptDo ReceiptDo Getter
func (r AlitripreceiptsellerinvoicereturnAPIRequest) GetReceiptDo() *ReceiptDo {
	return r._receiptDo
}
