package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytybgetcoderelationAPIRequest 医保-查询码的所有子码 API请求
// alibaba.alihealth.drug.kyt.yb.getcoderelation
//
// 应用于药店或医院入库环节，通过扫码获取下级码进行入库；
// 通过码查询所有子码以及包装比例
type AlibabaalihealthdrugkytybgetcoderelationAPIRequest struct {
	model.Params
	// 社保局(所属地市名称)
	_bureauName string
	// 请求终端名称
	_terminalName string
	// 终端类型：1005100-零售，1005200-医疗
	_terminalType string
	// 调用方式：formal-正式、test-测试
	_invocation string
	// 追溯码
	_code string
}

// NewAlibabaalihealthdrugkytybgetcoderelationRequest 初始化AlibabaalihealthdrugkytybgetcoderelationAPIRequest对象
func NewAlibabaalihealthdrugkytybgetcoderelationRequest() *AlibabaalihealthdrugkytybgetcoderelationAPIRequest {
	return &AlibabaalihealthdrugkytybgetcoderelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytybgetcoderelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yb.getcoderelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytybgetcoderelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytybgetcoderelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBureauName is BureauName Setter
// 社保局(所属地市名称)
func (r *AlibabaalihealthdrugkytybgetcoderelationAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// GetBureauName BureauName Getter
func (r AlibabaalihealthdrugkytybgetcoderelationAPIRequest) GetBureauName() string {
	return r._bureauName
}

// SetTerminalName is TerminalName Setter
// 请求终端名称
func (r *AlibabaalihealthdrugkytybgetcoderelationAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// GetTerminalName TerminalName Getter
func (r AlibabaalihealthdrugkytybgetcoderelationAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// SetTerminalType is TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaalihealthdrugkytybgetcoderelationAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r AlibabaalihealthdrugkytybgetcoderelationAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetInvocation is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaalihealthdrugkytybgetcoderelationAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// GetInvocation Invocation Getter
func (r AlibabaalihealthdrugkytybgetcoderelationAPIRequest) GetInvocation() string {
	return r._invocation
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugkytybgetcoderelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugkytybgetcoderelationAPIRequest) GetCode() string {
	return r._code
}
