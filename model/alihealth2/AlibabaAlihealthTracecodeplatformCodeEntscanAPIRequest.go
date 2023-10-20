package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest 药品商家扫码 API请求
// alibaba.alihealth.tracecodeplatform.code.entscan
//
// 药品商家扫描药品监管码，只有该商家的药才返回
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) Reset() {
	r._code = ""
	r._serviceFlag = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeplatform.code.entscan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 药监码
func (r *AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) GetCode() string {
	return r._code
}

// SetServiceFlag is ServiceFlag Setter
// 不同企业有不同的标识
func (r *AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) SetServiceFlag(_serviceFlag string) error {
	r._serviceFlag = _serviceFlag
	r.Set("service_flag", _serviceFlag)
	return nil
}

// GetServiceFlag ServiceFlag Getter
func (r AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) GetServiceFlag() string {
	return r._serviceFlag
}

var poolAlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTracecodeplatformCodeEntscanRequest()
	},
}

// GetAlibabaAlihealthTracecodeplatformCodeEntscanRequest 从 sync.Pool 获取 AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest
func GetAlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest() *AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest {
	return poolAlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest.Get().(*AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest)
}

// ReleaseAlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest 将 AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest(v *AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest.Put(v)
}
