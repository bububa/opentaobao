package wlb

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillCloudprintNetprintPrintAPIRequest) Reset() {
	r._printerPrintData = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillCloudprintNetprintPrintAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.cloudprint.netprint.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillCloudprintNetprintPrintAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillCloudprintNetprintPrintAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolCainiaoWaybillCloudprintNetprintPrintAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillCloudprintNetprintPrintRequest()
	},
}

// GetCainiaoWaybillCloudprintNetprintPrintRequest 从 sync.Pool 获取 CainiaoWaybillCloudprintNetprintPrintAPIRequest
func GetCainiaoWaybillCloudprintNetprintPrintAPIRequest() *CainiaoWaybillCloudprintNetprintPrintAPIRequest {
	return poolCainiaoWaybillCloudprintNetprintPrintAPIRequest.Get().(*CainiaoWaybillCloudprintNetprintPrintAPIRequest)
}

// ReleaseCainiaoWaybillCloudprintNetprintPrintAPIRequest 将 CainiaoWaybillCloudprintNetprintPrintAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillCloudprintNetprintPrintAPIRequest(v *CainiaoWaybillCloudprintNetprintPrintAPIRequest) {
	v.Reset()
	poolCainiaoWaybillCloudprintNetprintPrintAPIRequest.Put(v)
}
