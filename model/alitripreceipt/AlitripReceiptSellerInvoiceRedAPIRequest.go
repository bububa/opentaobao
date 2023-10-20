package alitripreceipt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripreceiptsellerinvoiceredAPIRequest 飞猪发票冲红 API请求
// alitrip.receipt.seller.invoice.red
//
// 飞猪发票创建
type AlitripreceiptsellerinvoiceredAPIRequest struct {
	model.Params
	// 入参对象
	_redReceiptParam *RedReceiptParam
}

// NewAlitripreceiptsellerinvoiceredRequest 初始化AlitripreceiptsellerinvoiceredAPIRequest对象
func NewAlitripreceiptsellerinvoiceredRequest() *AlitripreceiptsellerinvoiceredAPIRequest {
	return &AlitripreceiptsellerinvoiceredAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripreceiptsellerinvoiceredAPIRequest) GetApiMethodName() string {
	return "alitrip.receipt.seller.invoice.red"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripreceiptsellerinvoiceredAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripreceiptsellerinvoiceredAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRedReceiptParam is RedReceiptParam Setter
// 入参对象
func (r *AlitripreceiptsellerinvoiceredAPIRequest) SetRedReceiptParam(_redReceiptParam *RedReceiptParam) error {
	r._redReceiptParam = _redReceiptParam
	r.Set("red_receipt_param", _redReceiptParam)
	return nil
}

// GetRedReceiptParam RedReceiptParam Getter
func (r AlitripreceiptsellerinvoiceredAPIRequest) GetRedReceiptParam() *RedReceiptParam {
	return r._redReceiptParam
}
