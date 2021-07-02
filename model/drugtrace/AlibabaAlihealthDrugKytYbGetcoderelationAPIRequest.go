package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest 医保-查询码的所有子码 API请求
// alibaba.alihealth.drug.kyt.yb.getcoderelation
//
// 应用于药店或医院入库环节，通过扫码获取下级码进行入库；
// 通过码查询所有子码以及包装比例
type AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest struct {
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

// NewAlibabaAlihealthDrugKytYbGetcoderelationRequest 初始化AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest对象
func NewAlibabaAlihealthDrugKytYbGetcoderelationRequest() *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest {
	return &AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yb.getcoderelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BureauName Setter
// 社保局(所属地市名称)
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// Get BureauName Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetBureauName() string {
	return r._bureauName
}

// Set is TerminalName Setter
// 请求终端名称
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// Get TerminalName Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// Set is TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// Get TerminalType Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// Set is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// Get Invocation Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetInvocation() string {
	return r._invocation
}

// Set is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetCode() string {
	return r._code
}
