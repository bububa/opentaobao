package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillCloudprintNetprintPrintAPIRequest 网络打印机打印接口 API请求
// cainiao.waybill.cloudprint.netprint.print
//
// 打印接口
type CainiaoWaybillCloudprintNetprintPrintAPIRequest struct {
	model.Params
	// 请求
	_printerPrintData *CloudPrinterPrintRequest
}

// NewCainiaoWaybillCloudprintNetprintPrintRequest 初始化CainiaoWaybillCloudprintNetprintPrintAPIRequest对象
func NewCainiaoWaybillCloudprintNetprintPrintRequest() *CainiaoWaybillCloudprintNetprintPrintAPIRequest {
	return &CainiaoWaybillCloudprintNetprintPrintAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillCloudprintNetprintPrintAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.cloudprint.netprint.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillCloudprintNetprintPrintAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPrinterPrintData is PrinterPrintData Setter
// 请求
func (r *CainiaoWaybillCloudprintNetprintPrintAPIRequest) SetPrinterPrintData(_printerPrintData *CloudPrinterPrintRequest) error {
	r._printerPrintData = _printerPrintData
	r.Set("printer_print_data", _printerPrintData)
	return nil
}

// GetPrinterPrintData PrinterPrintData Getter
func (r CainiaoWaybillCloudprintNetprintPrintAPIRequest) GetPrinterPrintData() *CloudPrinterPrintRequest {
	return r._printerPrintData
}
