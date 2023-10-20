package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) Reset() {
	r._bureauName = ""
	r._terminalName = ""
	r._terminalType = ""
	r._invocation = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yb.getcoderelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBureauName is BureauName Setter
// 社保局(所属地市名称)
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// GetBureauName BureauName Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetBureauName() string {
	return r._bureauName
}

// SetTerminalName is TerminalName Setter
// 请求终端名称
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// GetTerminalName TerminalName Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// SetTerminalType is TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetInvocation is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// GetInvocation Invocation Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetInvocation() string {
	return r._invocation
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaAlihealthDrugKytYbGetcoderelationAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytYbGetcoderelationRequest()
	},
}

// GetAlibabaAlihealthDrugKytYbGetcoderelationRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest
func GetAlibabaAlihealthDrugKytYbGetcoderelationAPIRequest() *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest {
	return poolAlibabaAlihealthDrugKytYbGetcoderelationAPIRequest.Get().(*AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytYbGetcoderelationAPIRequest 将 AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytYbGetcoderelationAPIRequest(v *AlibabaAlihealthDrugKytYbGetcoderelationAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytYbGetcoderelationAPIRequest.Put(v)
}
