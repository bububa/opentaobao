package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintCmdprintRenderAPIRequest 生成打印机渲染命令（通过打印机命令打印） API请求
// cainiao.cloudprint.cmdprint.render
//
// isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。
type CainiaoCloudprintCmdprintRenderAPIRequest struct {
	model.Params
	// 参数对象
	_params *CmdRenderParams
}

// NewCainiaoCloudprintCmdprintRenderRequest 初始化CainiaoCloudprintCmdprintRenderAPIRequest对象
func NewCainiaoCloudprintCmdprintRenderRequest() *CainiaoCloudprintCmdprintRenderAPIRequest {
	return &CainiaoCloudprintCmdprintRenderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintCmdprintRenderAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.cmdprint.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintCmdprintRenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCloudprintCmdprintRenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParams is Params Setter
// 参数对象
func (r *CainiaoCloudprintCmdprintRenderAPIRequest) SetParams(_params *CmdRenderParams) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r CainiaoCloudprintCmdprintRenderAPIRequest) GetParams() *CmdRenderParams {
	return r._params
}
