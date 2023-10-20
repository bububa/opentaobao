package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybillcloudprintnetprintbindAPIRequest 网络打印机绑定 API请求
// cainiao.waybill.cloudprint.netprint.bind
//
// 绑定打印机接口
type CainiaowaybillcloudprintnetprintbindAPIRequest struct {
	model.Params
	// req
	_params *CloudPrinterBindRequest
}

// NewCainiaowaybillcloudprintnetprintbindRequest 初始化CainiaowaybillcloudprintnetprintbindAPIRequest对象
func NewCainiaowaybillcloudprintnetprintbindRequest() *CainiaowaybillcloudprintnetprintbindAPIRequest {
	return &CainiaowaybillcloudprintnetprintbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybillcloudprintnetprintbindAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.cloudprint.netprint.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybillcloudprintnetprintbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybillcloudprintnetprintbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParams is Params Setter
// req
func (r *CainiaowaybillcloudprintnetprintbindAPIRequest) SetParams(_params *CloudPrinterBindRequest) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r CainiaowaybillcloudprintnetprintbindAPIRequest) GetParams() *CloudPrinterBindRequest {
	return r._params
}
