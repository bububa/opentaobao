package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest
药品商家扫码 API请求
alibaba.alihealth.tracecodeplatform.code.entscan

药品商家扫描药品监管码，只有该商家的药才返回 */
type AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest struct {
	model.Params
	// 药监码
	_code string
	// 不同企业有不同的标识
	_serviceFlag string
}

// NewAlibabaAlihealthTracecodeplatformCodeEntscanRequest 初始化AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest对象
func NewAlibabaAlihealthTracecodeplatformCodeEntscanRequest() *AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest {
	return &AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeplatform.code.entscan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Code Setter
// 药监码
func (r *AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) GetCode() string {
	return r._code
}

// Set is ServiceFlag Setter
// 不同企业有不同的标识
func (r *AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) SetServiceFlag(_serviceFlag string) error {
	r._serviceFlag = _serviceFlag
	r.Set("service_flag", _serviceFlag)
	return nil
}

// Get ServiceFlag Getter
func (r AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) GetServiceFlag() string {
	return r._serviceFlag
}
