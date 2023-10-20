package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybillcloudprintnetprintprintAPIRequest 网络打印机打印接口 API请求
// cainiao.waybill.cloudprint.netprint.print
//
// 打印接口
type CainiaowaybillcloudprintnetprintprintAPIRequest struct {
	model.Params
	// 请求
	_printerPrintData *CloudPrinterPrintRequest
}

// NewCainiaowaybillcloudprintnetprintprintRequest 初始化CainiaowaybillcloudprintnetprintprintAPIRequest对象
func NewCainiaowaybillcloudprintnetprintprintRequest() *CainiaowaybillcloudprintnetprintprintAPIRequest {
	return &CainiaowaybillcloudprintnetprintprintAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybillcloudprintnetprintprintAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.cloudprint.netprint.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybillcloudprintnetprintprintAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybillcloudprintnetprintprintAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPrinterPrintData is PrinterPrintData Setter
// 请求
func (r *CainiaowaybillcloudprintnetprintprintAPIRequest) SetPrinterPrintData(_printerPrintData *CloudPrinterPrintRequest) error {
	r._printerPrintData = _printerPrintData
	r.Set("printer_print_data", _printerPrintData)
	return nil
}

// GetPrinterPrintData PrinterPrintData Getter
func (r CainiaowaybillcloudprintnetprintprintAPIRequest) GetPrinterPrintData() *CloudPrinterPrintRequest {
	return r._printerPrintData
}
