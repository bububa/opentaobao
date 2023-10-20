package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodelistcodegovAPIRequest 政府码查询接口 API请求
// alibaba.alihealth.drug.code.list.code.gov
//
// 政府码查询接口
type AlibabaalihealthdrugcodelistcodegovAPIRequest struct {
	model.Params
	// 追溯码
	_codeList []string
	// 证件编号
	_certIsvNo string
	// 调用方式：formal-正式、test-测试
	_invocation string
	// 终端类型 1：零售
	_terminalType string
	// 调用零售药店名称
	_terminalName string
	// 城市名称
	_bureauName string
	// 错误信息
	_errorMessage string
}

// NewAlibabaalihealthdrugcodelistcodegovRequest 初始化AlibabaalihealthdrugcodelistcodegovAPIRequest对象
func NewAlibabaalihealthdrugcodelistcodegovRequest() *AlibabaalihealthdrugcodelistcodegovAPIRequest {
	return &AlibabaalihealthdrugcodelistcodegovAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodelistcodegovAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.list.code.gov"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodelistcodegovAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodelistcodegovAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodeList is CodeList Setter
// 追溯码
func (r *AlibabaalihealthdrugcodelistcodegovAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaalihealthdrugcodelistcodegovAPIRequest) GetCodeList() []string {
	return r._codeList
}

// SetCertIsvNo is CertIsvNo Setter
// 证件编号
func (r *AlibabaalihealthdrugcodelistcodegovAPIRequest) SetCertIsvNo(_certIsvNo string) error {
	r._certIsvNo = _certIsvNo
	r.Set("cert_isv_no", _certIsvNo)
	return nil
}

// GetCertIsvNo CertIsvNo Getter
func (r AlibabaalihealthdrugcodelistcodegovAPIRequest) GetCertIsvNo() string {
	return r._certIsvNo
}

// SetInvocation is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaalihealthdrugcodelistcodegovAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// GetInvocation Invocation Getter
func (r AlibabaalihealthdrugcodelistcodegovAPIRequest) GetInvocation() string {
	return r._invocation
}

// SetTerminalType is TerminalType Setter
// 终端类型 1：零售
func (r *AlibabaalihealthdrugcodelistcodegovAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r AlibabaalihealthdrugcodelistcodegovAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetTerminalName is TerminalName Setter
// 调用零售药店名称
func (r *AlibabaalihealthdrugcodelistcodegovAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// GetTerminalName TerminalName Getter
func (r AlibabaalihealthdrugcodelistcodegovAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// SetBureauName is BureauName Setter
// 城市名称
func (r *AlibabaalihealthdrugcodelistcodegovAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// GetBureauName BureauName Getter
func (r AlibabaalihealthdrugcodelistcodegovAPIRequest) GetBureauName() string {
	return r._bureauName
}

// SetErrorMessage is ErrorMessage Setter
// 错误信息
func (r *AlibabaalihealthdrugcodelistcodegovAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// GetErrorMessage ErrorMessage Getter
func (r AlibabaalihealthdrugcodelistcodegovAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}
