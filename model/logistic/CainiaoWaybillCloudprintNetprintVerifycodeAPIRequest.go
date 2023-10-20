package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest 打印验证码 API请求
// cainiao.waybill.cloudprint.netprint.verifycode
//
// 打印获取验证码
type CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest struct {
	model.Params
	// 请求
	_printer *CloudPrinterVerifyCodeRequest
}

// NewCainiaoWaybillCloudprintNetprintVerifycodeRequest 初始化CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest对象
func NewCainiaoWaybillCloudprintNetprintVerifycodeRequest() *CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest {
	return &CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest) Reset() {
	r._printer = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.cloudprint.netprint.verifycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPrinter is Printer Setter
// 请求
func (r *CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest) SetPrinter(_printer *CloudPrinterVerifyCodeRequest) error {
	r._printer = _printer
	r.Set("printer", _printer)
	return nil
}

// GetPrinter Printer Getter
func (r CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest) GetPrinter() *CloudPrinterVerifyCodeRequest {
	return r._printer
}

var poolCainiaoWaybillCloudprintNetprintVerifycodeAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillCloudprintNetprintVerifycodeRequest()
	},
}

// GetCainiaoWaybillCloudprintNetprintVerifycodeRequest 从 sync.Pool 获取 CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest
func GetCainiaoWaybillCloudprintNetprintVerifycodeAPIRequest() *CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest {
	return poolCainiaoWaybillCloudprintNetprintVerifycodeAPIRequest.Get().(*CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest)
}

// ReleaseCainiaoWaybillCloudprintNetprintVerifycodeAPIRequest 将 CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillCloudprintNetprintVerifycodeAPIRequest(v *CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest) {
	v.Reset()
	poolCainiaoWaybillCloudprintNetprintVerifycodeAPIRequest.Put(v)
}
