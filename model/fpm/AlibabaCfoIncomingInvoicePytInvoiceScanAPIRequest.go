package fpm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacfoincominginvoicepytinvoicescanAPIRequest 票易通发票ocr信息同步 API请求
// alibaba.cfo.incoming.invoice.pyt.invoice.scan
//
// 票易通发票ocr信息同步
type AlibabacfoincominginvoicepytinvoicescanAPIRequest struct {
	model.Params
	// ocr录入请求体
	_scanRequest *InvoiceScanRequest
}

// NewAlibabacfoincominginvoicepytinvoicescanRequest 初始化AlibabacfoincominginvoicepytinvoicescanAPIRequest对象
func NewAlibabacfoincominginvoicepytinvoicescanRequest() *AlibabacfoincominginvoicepytinvoicescanAPIRequest {
	return &AlibabacfoincominginvoicepytinvoicescanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacfoincominginvoicepytinvoicescanAPIRequest) GetApiMethodName() string {
	return "alibaba.cfo.incoming.invoice.pyt.invoice.scan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacfoincominginvoicepytinvoicescanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacfoincominginvoicepytinvoicescanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScanRequest is ScanRequest Setter
// ocr录入请求体
func (r *AlibabacfoincominginvoicepytinvoicescanAPIRequest) SetScanRequest(_scanRequest *InvoiceScanRequest) error {
	r._scanRequest = _scanRequest
	r.Set("scan_request", _scanRequest)
	return nil
}

// GetScanRequest ScanRequest Getter
func (r AlibabacfoincominginvoicepytinvoicescanAPIRequest) GetScanRequest() *InvoiceScanRequest {
	return r._scanRequest
}
