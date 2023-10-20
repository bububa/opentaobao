package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytrecordinfoAPIRequest 快易通健康检查 API请求
// alibaba.alihealth.drug.kyt.recordinfo
//
// 快易通健康检查
type AlibabaalihealthdrugkytrecordinfoAPIRequest struct {
	model.Params
	// 级别
	_logLevel string
	// 服务名
	_serviceName string
	// 类型
	_serviceType string
	// 输入参数
	_inputParam string
	// 其他参数
	_otherParam string
}

// NewAlibabaalihealthdrugkytrecordinfoRequest 初始化AlibabaalihealthdrugkytrecordinfoAPIRequest对象
func NewAlibabaalihealthdrugkytrecordinfoRequest() *AlibabaalihealthdrugkytrecordinfoAPIRequest {
	return &AlibabaalihealthdrugkytrecordinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytrecordinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.recordinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytrecordinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytrecordinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogLevel is LogLevel Setter
// 级别
func (r *AlibabaalihealthdrugkytrecordinfoAPIRequest) SetLogLevel(_logLevel string) error {
	r._logLevel = _logLevel
	r.Set("log_level", _logLevel)
	return nil
}

// GetLogLevel LogLevel Getter
func (r AlibabaalihealthdrugkytrecordinfoAPIRequest) GetLogLevel() string {
	return r._logLevel
}

// SetServiceName is ServiceName Setter
// 服务名
func (r *AlibabaalihealthdrugkytrecordinfoAPIRequest) SetServiceName(_serviceName string) error {
	r._serviceName = _serviceName
	r.Set("service_name", _serviceName)
	return nil
}

// GetServiceName ServiceName Getter
func (r AlibabaalihealthdrugkytrecordinfoAPIRequest) GetServiceName() string {
	return r._serviceName
}

// SetServiceType is ServiceType Setter
// 类型
func (r *AlibabaalihealthdrugkytrecordinfoAPIRequest) SetServiceType(_serviceType string) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// GetServiceType ServiceType Getter
func (r AlibabaalihealthdrugkytrecordinfoAPIRequest) GetServiceType() string {
	return r._serviceType
}

// SetInputParam is InputParam Setter
// 输入参数
func (r *AlibabaalihealthdrugkytrecordinfoAPIRequest) SetInputParam(_inputParam string) error {
	r._inputParam = _inputParam
	r.Set("input_param", _inputParam)
	return nil
}

// GetInputParam InputParam Getter
func (r AlibabaalihealthdrugkytrecordinfoAPIRequest) GetInputParam() string {
	return r._inputParam
}

// SetOtherParam is OtherParam Setter
// 其他参数
func (r *AlibabaalihealthdrugkytrecordinfoAPIRequest) SetOtherParam(_otherParam string) error {
	r._otherParam = _otherParam
	r.Set("other_param", _otherParam)
	return nil
}

// GetOtherParam OtherParam Getter
func (r AlibabaalihealthdrugkytrecordinfoAPIRequest) GetOtherParam() string {
	return r._otherParam
}
