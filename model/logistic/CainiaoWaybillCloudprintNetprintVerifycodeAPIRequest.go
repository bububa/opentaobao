package logistic

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.cloudprint.netprint.verifycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
