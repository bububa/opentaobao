package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillCloudprintNetprintBindAPIRequest 网络打印机绑定 API请求
// cainiao.waybill.cloudprint.netprint.bind
//
// 绑定打印机接口
type CainiaoWaybillCloudprintNetprintBindAPIRequest struct {
	model.Params
	// req
	_params *CloudPrinterBindRequest
}

// NewCainiaoWaybillCloudprintNetprintBindRequest 初始化CainiaoWaybillCloudprintNetprintBindAPIRequest对象
func NewCainiaoWaybillCloudprintNetprintBindRequest() *CainiaoWaybillCloudprintNetprintBindAPIRequest {
	return &CainiaoWaybillCloudprintNetprintBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillCloudprintNetprintBindAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.cloudprint.netprint.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillCloudprintNetprintBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParams is Params Setter
// req
func (r *CainiaoWaybillCloudprintNetprintBindAPIRequest) SetParams(_params *CloudPrinterBindRequest) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r CainiaoWaybillCloudprintNetprintBindAPIRequest) GetParams() *CloudPrinterBindRequest {
	return r._params
}
