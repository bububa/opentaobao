package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybillcloudprintnetprintverifycodeAPIRequest 打印验证码 API请求
// cainiao.waybill.cloudprint.netprint.verifycode
//
// 打印获取验证码
type CainiaowaybillcloudprintnetprintverifycodeAPIRequest struct {
	model.Params
	// 请求
	_printer *CloudPrinterVerifyCodeRequest
}

// NewCainiaowaybillcloudprintnetprintverifycodeRequest 初始化CainiaowaybillcloudprintnetprintverifycodeAPIRequest对象
func NewCainiaowaybillcloudprintnetprintverifycodeRequest() *CainiaowaybillcloudprintnetprintverifycodeAPIRequest {
	return &CainiaowaybillcloudprintnetprintverifycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybillcloudprintnetprintverifycodeAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.cloudprint.netprint.verifycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybillcloudprintnetprintverifycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybillcloudprintnetprintverifycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPrinter is Printer Setter
// 请求
func (r *CainiaowaybillcloudprintnetprintverifycodeAPIRequest) SetPrinter(_printer *CloudPrinterVerifyCodeRequest) error {
	r._printer = _printer
	r.Set("printer", _printer)
	return nil
}

// GetPrinter Printer Getter
func (r CainiaowaybillcloudprintnetprintverifycodeAPIRequest) GetPrinter() *CloudPrinterVerifyCodeRequest {
	return r._printer
}
