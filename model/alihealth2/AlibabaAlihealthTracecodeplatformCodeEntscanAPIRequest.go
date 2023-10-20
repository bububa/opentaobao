package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodeplatformcodeentscanAPIRequest 药品商家扫码 API请求
// alibaba.alihealth.tracecodeplatform.code.entscan
//
// 药品商家扫描药品监管码，只有该商家的药才返回
type AlibabaalihealthtracecodeplatformcodeentscanAPIRequest struct {
	model.Params
	// 药监码
	_code string
	// 不同企业有不同的标识
	_serviceFlag string
}

// NewAlibabaalihealthtracecodeplatformcodeentscanRequest 初始化AlibabaalihealthtracecodeplatformcodeentscanAPIRequest对象
func NewAlibabaalihealthtracecodeplatformcodeentscanRequest() *AlibabaalihealthtracecodeplatformcodeentscanAPIRequest {
	return &AlibabaalihealthtracecodeplatformcodeentscanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodeplatformcodeentscanAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeplatform.code.entscan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodeplatformcodeentscanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodeplatformcodeentscanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 药监码
func (r *AlibabaalihealthtracecodeplatformcodeentscanAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthtracecodeplatformcodeentscanAPIRequest) GetCode() string {
	return r._code
}

// SetServiceFlag is ServiceFlag Setter
// 不同企业有不同的标识
func (r *AlibabaalihealthtracecodeplatformcodeentscanAPIRequest) SetServiceFlag(_serviceFlag string) error {
	r._serviceFlag = _serviceFlag
	r.Set("service_flag", _serviceFlag)
	return nil
}

// GetServiceFlag ServiceFlag Getter
func (r AlibabaalihealthtracecodeplatformcodeentscanAPIRequest) GetServiceFlag() string {
	return r._serviceFlag
}
